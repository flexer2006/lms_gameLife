package main

import (
	"context"
	"os"

	"github.com/flexer2006/lms_gameLife/internal/application"
)

func main() {
	ctx := context.Background()
	cfg := application.Config{}
	app := application.New(cfg)
	os.Exit(app.Run(ctx))
}
