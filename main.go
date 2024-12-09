package main

import (
	"os"

	_ "github.com/VuKhoa23/advanced-web-be/docs"
	"github.com/VuKhoa23/advanced-web-be/startup"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		startup.Migrate()
		return
	}

	startup.Execute()
}