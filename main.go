package main

import (
	"context"
	"log"
	"os"

	"github.com/MoonMoon1919/gignore"
	"github.com/MoonMoon1919/gignore-cli/internal/builder"
)

func run(args []string, svc gignore.Service) {
	cmd := builder.AppBuilder(svc)

	if err := cmd.Run(context.Background(), args); err != nil {
		log.Fatal(err)
	}
}

func main() {
	repo := gignore.NewFileRepository(gignore.RenderOptions{
		TrailingNewLine: true,
		HeaderComment:   "This file was automatically generated, do not make manual edits",
	})
	svc := gignore.NewService(&repo)

	run(os.Args, svc)
}
