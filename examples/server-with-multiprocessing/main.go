package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/devplayg/hippo"
	"math/rand"
	"sync/atomic"
	"time"
)

var (
	taskId int
)

func main() {
	// Random seed
	rand.Seed(time.Now().UnixNano())

	// Server
	company := &Server{
		workerCount: 8,
	}

	// Start hippo after loading the server into the hippo
	hippo := hippo.NewHippo(company, nil)
	if err := hippo.Start(); err != nil {
		panic(err)
	}
}

type Server struct {
	hippo.Launcher
	workerCount    int
	currentWorking int32
}

func (s *Server) Start() error {
	s.Log.Print("server has been started")
	if err := s.run(); err != nil {
		s.Log.Print(err)
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	s.Log.Print("server has been stopped")
	return nil
}

func (s *Server) run() error {
	s.Log.Printf("%d workers are ready", s.workerCount)
	ch := make(chan bool, s.workerCount)

	for {
		if s.currentWorking == 0 && s.Ctx.Err() != nil {
			if errors.Is(s.Ctx.Err(), context.Canceled) {
				return nil
			}
			return s.Ctx.Err()
		}

		tasks, _ := s.getTasks()
		s.Log.Printf("found %d tasks", len(tasks))

		err := s.distributeTasks(tasks, ch)
		if err != nil {
			// is error from context canceled?
			if errors.Is(err, context.Canceled) {
				return nil
			}
			return err
		}

		select {
		case <-time.After(1 * time.Second):
		}
	}
}

func (s *Server) distributeTasks(tasks []string, ch chan bool) error {
	for _, task := range tasks {

		// waits for start time
		if s.Ctx.Err() == nil {

			// waits for start time
			ch <- true

			if s.Ctx.Err() == nil { // Recheck if context was canceled while waiting
				// Handle task
				go func(myTask string) {
					defer func() {
						s.Log.Printf("done %s", myTask)

						// Mark an worker as waiting
						atomic.AddInt32(&s.currentWorking, -1)

						<-ch
					}()

					// Mark an worker as working
					atomic.AddInt32(&s.currentWorking, 1)

					// Handle task
					if err := s.handleTask(myTask); err != nil {
						s.Log.Print(err)
					}
				}(task)
			}
		}

		select {
		case <-s.Ctx.Done(): // for gracefully shutdown
			// Wait for working workers
			s.Log.Printf("waiting for working %d workers", s.currentWorking)
			for s.currentWorking > 0 {
				time.Sleep(100 * time.Millisecond)
			}
			s.Log.Print("works of all workers are over")
			return s.Ctx.Err()
		case <-time.After(1 * time.Millisecond):
		}
	}

	return nil
}

func (s *Server) handleTask(task string) error {
	s.Log.Printf("working %s", task)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	return nil
}

func (s *Server) getTasks() ([]string, error) {
	tasks := make([]string, 0)
	for i := 0; i < rand.Intn(10); i++ {
		tasks = append(tasks, fmt.Sprintf("task-%d", taskId))
		taskId++
	}
	return tasks, nil
}
