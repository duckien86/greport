package main

import (
	"2ndbrand-api/component/asyncjob"
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {

	job1 := asyncjob.NewJob(func(ctx context.Context) error {
		jNum := 1
		randNumb := rand.Intn(100)
		log.Printf("Do the job %v --> %v", jNum, randNumb)
		if randNumb > 0 && randNumb%2 == 0 {
			time.Sleep(time.Second * 1)
			log.Printf(" job %v --> DONE", jNum)
			return nil
		}
		return fmt.Errorf("error at %v", jNum)
	})
	job2 := asyncjob.NewJob(func(ctx context.Context) error {
		jNum := 2
		randNumb := rand.Intn(100)
		log.Printf("Do the job %v --> %v", jNum, randNumb)
		if randNumb > 0 && randNumb%2 == 0 {
			// time.Sleep(time.Second * 1)
			log.Printf(" job %v --> DONE", jNum)
			return nil
		}
		return fmt.Errorf("error at %v", jNum)
	})

	job3 := asyncjob.NewJob(func(ctx context.Context) error {
		jNum := 3
		randNumb := rand.Intn(100)
		log.Printf("Do the job %v --> %v", jNum, randNumb)
		if randNumb > 0 && randNumb%2 == 0 {
			// time.Sleep(time.Second * 1)
			log.Printf(" job %v --> DONE", jNum)
			return nil
		}
		return fmt.Errorf("error at %v", jNum)
	})

	group := asyncjob.NewGroup(true, job1, job2, job3)
	if err := group.Run(context.Background()); err != nil {
		log.Println("Last error : ", err)
	}
}
