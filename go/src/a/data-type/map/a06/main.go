package main

import "fmt"

// Profile struct.
type Profile struct {
	Name    string
	Age     int
	married bool
}

type classicQueryKey struct {
	Name string
	Age  int
}

type multiQueryKey struct {
	Name string
	Age  int
}

func (c *classicQueryKey) hash() int {
	return simpleHash(c.Name) + c.Age*10000000
}

func simpleHash(s string) (ret int) {
	for i := 0; i < len(s); i++ {
		c := s[i]
		ret += int(c)
	}
	return
}

var mapper = make(map[int][]*Profile)
var mapper2 = make(map[multiQueryKey]*Profile)

func main() {
	list := []*Profile{
		{"张三", 30, true},
		{"李四", 38, false},
		{"王五", 28, true},
	}
	builder(list)
	query("张三", 30)
	query("李四", 30)

	fmt.Println()

	builder2(list)
	query2("张三", 30)
	query2("李四", 30)
}

func builder(list []*Profile) {
	for _, profile := range list {
		key := classicQueryKey{profile.Name, profile.Age}
		value := mapper[key.hash()]
		value = append(value, profile)
		mapper[key.hash()] = value
	}
}

func builder2(list []*Profile) {
	for _, profile := range list {
		key := multiQueryKey{Name: profile.Name, Age: profile.Age}
		mapper2[key] = profile
	}
}

func query(name string, age int) {
	keyToQuery := classicQueryKey{Name: name, Age: age}
	list := mapper[keyToQuery.hash()]
	for _, result := range list {
		if result.Name == name && result.Age == age {
			fmt.Printf("🎉 %T, %v\n", result, result)
			return
		}
	}
	fmt.Println("🧨 Query data not exist")
}

func query2(name string, age int) {
	keyToQuery := multiQueryKey{name, age}
	result, ok := mapper2[keyToQuery]
	if ok {
		fmt.Printf("👌 %T, %v\n", result, result)
	} else {
		fmt.Println("📛 Query data not exist")
	}
}
