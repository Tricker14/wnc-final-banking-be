package startup

import (
	"github.com/VuKhoa23/advanced-web-be/internal"
	"github.com/VuKhoa23/advanced-web-be/internal/controller"
	"github.com/VuKhoa23/advanced-web-be/internal/database"
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
