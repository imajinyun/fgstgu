package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

func task() {
	fmt.Println("I am running task...")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func main() {
	// defines a new scheduler that schedules and runs jobs.
	s1 := gocron.NewScheduler(time.UTC)
	s1.Every(3).Seconds().Do(task)
	s1.StartAsync()

	// do jobs without params.
	s2 := gocron.NewScheduler(time.UTC)
	s2.Every(1).Second().Do(task)
	s2.Every(2).Seconds().Do(task)
	s2.Every(1).Minute().Do(task)
	s2.Every(2).Minutes().Do(task)

	// check for errors.
	if _, err := s2.Every(1).Day().At("bad-time").Do(task); err != nil {
		log.Fatalf("🔥 error creating job: %v", err)
	}

	// do jobs with params.
	s2.Every(1).Second().Do(taskWithParams, 1, "hello world")

	// do jobs with tags.
	tag1 := []string{"tag1"}
	tag2 := []string{"tag2"}
	s2.Every(1).Week().SetTag(tag1).Do(task)
	s2.Every(1).Week().SetTag(tag2).Do(task)

	// removing job based on Tag.
	s2.RemoveJobByTag("tag1")

	// remove a job after its last execution.
	j, _ := s2.Every(1).StartAt(time.Now().Add(30 * time.Second)).Do(task)
	j.LimitRunsTo(1)
	j.RemoveAfterLastRun()

	// do jobs on specific weekday.
	s2.Every(1).Monday().Do(task)
	s2.Every(1).Thursday().Do(task)

	// do a job at a specific time - 'hour:min:sec' - seconds optional.
	s2.Every(1).Day().At("10:30").Do(task)
	s2.Every(1).Friday().At("17:00").Do(task)

	// begin job at a specific date/time.
	t := time.Date(2021, time.January, 1, 17, 0, 0, 0, time.UTC)
	s2.Every(1).Hour().StartAt(t).Do(task)

	// delay start of job.
	s2.Every(1).Hour().StartAt(time.Now().Add(time.Duration(1 * time.Hour))).Do(task)

	// NextRun gets the next running time.
	_, time := s2.NextRun()
	fmt.Println(time)

	// remove a specific job.
	s2.Remove(task)

	// clear all scheduled jobs.
	s2.Clear()

	// stop our first scheduler (it will exists but doesn't run anymore).
	s1.Stop()

	// execute the scheduler and blocks current thread.
	s2.StartBlocking()

	fmt.Println("task already finished...")
}
