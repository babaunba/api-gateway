package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/babaunba/project-management/api-gateway/internal/domain"
)

const (
	taskQueue = "labels-tasks"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalf("failed to connect to temporal: %v", err)
	}

	w := worker.New(c, taskQueue, worker.Options{})
	w.RegisterWorkflow(domain.New().GetLabelsWF)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalf("failed to start worker: %v", err)
	}
}
