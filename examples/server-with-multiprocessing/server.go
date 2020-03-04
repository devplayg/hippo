package main

import (
	"context"
	"fmt"
	"github.com/devplayg/hippo/v2"
	"math/rand"
	"strings"
	"sync/atomic"
	"time"
)

var (
	taskId         int
	currentWorking int32
)

func main() {
	// Random seed
	rand.Seed(time.Now().UnixNano())

	// Config
	config := &hippo.Config{
		Debug: true,
	}

	// Server
	company := &Server{
		workerCount: 8,
	}

	// Start engine after loading the server into the engine
	engine := hippo.NewEngine(company, config)
	if err := engine.Start(); err != nil {
		panic(err)
	}
}

type Server struct {
	hippo.Launcher
	workerCount int
}

func (s *Server) Start() error {
	s.Log.Debugf("%s has been started", s.Engine.Config.Name)
	if err := s.run(); err != nil {
		s.Log.Error(err)
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	s.Log.Debugf("%s has been stopped", s.Engine.Config.Name)
	return nil
}

func (s *Server) run() error {
	s.Log.Debugf("%d workers are ready", s.workerCount)
	ch := make(chan bool, s.workerCount)

	for {
		tasks := getTasks()
		s.Log.Debugf("got %d tasks", len(tasks))

		err := s.distributeTasks(tasks, ch)
		if err != nil {
			// is error from context canceled?
			if strings.Contains(err.Error(), context.Canceled.Error()) {
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

		if s.Ctx.Err() == nil {

			// waits for start time
			ch <- true

			// Handle task
			go func(myTask string) {
				defer func() {
					s.Log.Debugf("done %s", myTask)

					// Mark an worker as waiting
					atomic.AddInt32(&currentWorking, -1)

					<-ch
				}()

				// Mark an worker as working
				atomic.AddInt32(&currentWorking, 1)

				// Handle task
				if err := s.handleTask(myTask); err != nil {
					s.Log.Error(err)
				}
			}(task)
		}

		select {
		case <-s.Ctx.Done(): // for gracefully shutdown
			// Wait for working workers
			s.Log.Debugf("waiting for working %d workers", currentWorking)
			for currentWorking > 0 {
				time.Sleep(100 * time.Millisecond)
			}
			s.Log.Debug("works of all workers are over")
			return s.Ctx.Err()
		case <-time.After(1 * time.Millisecond):
		}
	}

	return nil
}

func (s *Server) handleTask(task string) error {
	s.Log.Debugf("working %s", task)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	return nil
}

// Task generator
func getTasks() []string {
	tasks := make([]string, 0)
	for i := 0; i < rand.Intn(10); i++ {
		tasks = append(tasks, fmt.Sprintf("task-%d", taskId))
		taskId++
	}
	return tasks
}
