package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"sync/atomic"
	"time"

	"golang.org/x/sync/errgroup"
)

const nWorkers = 10

var users = []string{"Tom", "Bob", "Frank", "Jack", "Alice"}

// User struct.
type User struct {
	ID   int64
	Name string
}

// FriendIterator mocks out an iterator to a remote datastore query.
type FriendIterator struct {
	cursor int
}

// Next receiver.
func (it *FriendIterator) Next(ctx context.Context) (int64, error) {
	if it.cursor >= len(users) {
		return 0, io.EOF
	}

	// Pretend each item takes 10ms to come back over the network.
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case <-time.After(10 * time.Millisecond):
		r := int64(it.cursor)
		it.cursor++
		log.Printf("Found friend id: %d", r)
		return r, nil
	}
}

// GetFriendIds mocks out a remote friend list lookup.
func GetFriendIds(user int64) *FriendIterator {
	return &FriendIterator{}
}

// GetUserProfile func.
func GetUserProfile(ctx context.Context, id int64) (*User, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(100 * time.Microsecond):
		if id < 0 || id > int64(len(users)) {
			return nil, fmt.Errorf("Unknown user: %d", id)
		}
		log.Printf("Found user profile: %d", id)
		return &User{ID: id, Name: users[id]}, nil
	}
}

// GetFriendsSerial func.
func GetFriendsSerial(ctx context.Context, user int64) (map[string]*User, error) {
	var friendIds []int64
	for it := GetFriendIds(user); ; {
		id, err := it.Next(ctx)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("GetFriendIds %d: %s", user, err)
		}
		friendIds = append(friendIds, id)
	}
	ret := map[string]*User{}
	for _, friendID := range friendIds {
		friend, err := GetUserProfile(ctx, friendID)
		if err != nil {
			return nil, fmt.Errorf("GetUserProfile %d: %s", friendID, err)
		}
		ret[friend.Name] = friend
	}
	return ret, nil
}

// GetFriendsParallel func.
func GetFriendsParallel(ctx context.Context, user int64) (map[string]*User, error) {
	friendIds := make(chan int64)
	go func() {
		defer close(friendIds)
		for it := GetFriendIds(user); ; {
			if id, err := it.Next(ctx); err != nil {
				if err == io.EOF {
					break
				}
				log.Fatalf("GetFriendIds %d: %s", user, err)
			} else {
				friendIds <- id
			}
		}
	}()

	friends := make(chan *User)
	workers := int32(nWorkers)
	for i := 0; i < nWorkers; i++ {
		go func() {
			defer func() {
				if atomic.AddInt32(&workers, -1) == 0 {
					close(friends)
				}
			}()
			for id := range friendIds {
				if friend, err := GetUserProfile(ctx, id); err != nil {
					log.Fatalf("GerUserProfile %d: %s", user, err)
				} else {
					friends <- friend
				}
			}
		}()
	}
	ret := map[string]*User{}
	for friend := range friends {
		ret[friend.Name] = friend
	}
	return ret, nil
}

// GetFriendsErrGroup func.
func GetFriendsErrGroup(ctx context.Context, user int64) (map[string]*User, error) {
	g, ctx := errgroup.WithContext(ctx)
	friendIds := make(chan int64)

	g.Go(func() error {
		defer close(friendIds)
		for it := GetFriendIds(user); ; {
			id, err := it.Next(ctx)
			if err != nil {
				if err == io.EOF {
					return nil
				}
				return fmt.Errorf("GetFriendIds %d: %s", user, err)
			}
			friendIds <- id
		}
	})

	friends := make(chan *User)
	workers := int32(nWorkers)
	for i := 0; i < nWorkers; i++ {
		g.Go(func() error {
			defer func() {
				if atomic.AddInt32(&workers, -1) == 0 {
					close(friends)
				}
			}()

			for id := range friendIds {
				friend, err := GetUserProfile(ctx, id)
				if err != nil {
					return fmt.Errorf("GetUserProfile %d: %s", user, err)
				}
				friends <- friend
			}
			return nil
		})
	}

	ret := map[string]*User{}
	g.Go(func() error {
		for friend := range friends {
			ret[friend.Name] = friend
		}
		return nil
	})
	return ret, g.Wait()
}

// GetFriendsSelects func.
func GetFriendsSelects(ctx context.Context, user int64) (map[string]*User, error) {
	g, ctx := errgroup.WithContext(ctx)
	friendIds := make(chan int64)

	// Produce
	g.Go(func() error {
		defer close(friendIds)
		for it := GetFriendIds(user); ; {
			if id, err := it.Next(ctx); err != nil {
				if err == io.EOF {
					return nil
				}
			} else {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case friendIds <- id:
				}
			}
		}
	})

	friends := make(chan *User)
	workers := int32(nWorkers)
	for i := 0; i < nWorkers; i++ {
		g.Go(func() error {
			defer func() {
				if atomic.AddInt32(&workers, -1) == 0 {
					close(friends)
				}
			}()
			for id := range friendIds {
				friend, err := GetUserProfile(ctx, id)
				if err != nil {
					return fmt.Errorf("GetUserProfile %d: %s", user, err)
				}
				select {
				case <-ctx.Done():
					return ctx.Err()
				case friends <- friend:
				}
			}
			return nil
		})
	}

	ret := map[string]*User{}
	g.Go(func() error {
		for friend := range friends {
			ret[friend.Name] = friend
		}
		return nil
	})
	return ret, g.Wait()
}

func main() {
	type getFriendFunc func(ctx context.Context, user int64) (map[string]*User, error)
	type tc struct {
		name string
		fn   getFriendFunc
	}

	for _, tc := range []tc{
		{"GetFriendsSerial", GetFriendsSerial},
		{"GetFriendsParallel", GetFriendsParallel},
		{"GetFriendsErrGroup", GetFriendsErrGroup},
		{"GetFriendsSelects", GetFriendsSelects},
	} {
		time.Sleep(time.Second)
		log.Printf("%s\n", tc.name)

		ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
		start := time.Now()
		ret, err := tc.fn(ctx, 0)
		if err != nil {
			log.Fatalf("error: %s", err)
		} else {
			log.Printf("finished in %s: %s", time.Since(start), joinString(ret))
		}
		log.Println()
		cancel()
	}
}

func joinString(v interface{}) string {
	ret, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}
	return string(ret)
}
