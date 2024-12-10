package startup

import (
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/controller"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/database"
)

func Migrate() {
	// Open the database connection
	db := database.Open()

	database.MigrateUp(db)
}

func registerDependencies() *controller.ApiContainer {
	// Open the database connection
	db := database.Open()

	return internal.InitializeContainer(db)
}

func Execute() {
	container := registerDependencies()
	container.HttpServer.Run()
}
