package main

import (
	"context"
	"lo/internal/app"
	"log"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	a, err := app.New(
		ctx,
		app.WithLogger(),
	)

	if err != nil {
		log.Printf("Error creating new app %s", err.Error())
		os.Exit(1)
	}

	if err := a.Run(ctx, cancel); err != nil {
		log.Printf("Error running app %s", err.Error())
		os.Exit(1)
	}
}
