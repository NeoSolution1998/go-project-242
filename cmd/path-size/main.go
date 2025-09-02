package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {

	cmd := &cli.Command{
		Name:  "path-size - print size of a file or directory;",
		Usage: "path-size [global options]",
		Action: func(context.Context, *cli.Command) error {

			Hello("World")
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func Hello(name string) string {
	message := fmt.Sprintf("Hello, %v!", name)
	fmt.Println(message)
	return message
}
