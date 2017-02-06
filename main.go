package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/drone/drone-plugin-go/plugin"
	_ "github.com/gemnasium/migrate/driver/postgres" // postgres driver
	"github.com/gemnasium/migrate/migrate"
)

type Config struct {
	DatabaseURL    string `json:"database_url"`
	MigrationsPath string `json:"migrations_path"`
}

var (
	buildDate string
)

func main() {
	fmt.Printf("Drone Migrate Plugin built at %s\n", buildDate)

	workspace := plugin.Workspace{}
	vargs := Config{}

	plugin.Param("workspace", &workspace)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	migrationsPath := filepath.Join(workspace.Path, vargs.MigrationsPath)
	errs, ok := migrate.UpSync(vargs.DatabaseURL, migrationsPath)
	if !ok {
		fmt.Println("Error migrating database:")
		for _, e := range errs {
			fmt.Println(e.Error())
		}
		os.Exit(1)
	}

	version, err := migrate.Version(vargs.DatabaseURL, migrationsPath)
	if err != nil {
		fmt.Printf("Error getting database version: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Database succesfully migrated to version %v\n", version)
}
