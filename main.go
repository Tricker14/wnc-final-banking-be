package main

import (
	"os"

	_ "github.com/VuKhoa23/advanced-web-be/docs"
	"github.com/VuKhoa23/advanced-web-be/startup"
)

func hasCommand(args []string) bool {
	return len(args) > 1
}

func main() {
	if hasCommand(os.Args) && os.Args[1] == "migrate-up" {
		startup.Migrate()
		return
	}

	startup.Execute()
}