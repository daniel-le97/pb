// main.go
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	routers "github.com/daniel-le97/pb/routers"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	// uncomment once you have at least one .go migration file in the "migrations" directory
	_ "github.com/daniel-le97/pb/migrations"
	// _ "leploy/routers"
)

func isDockerInstalled() bool {
	cmd := exec.Command("docker", "--version")
	err := cmd.Run()
	return err == nil
}
func isNixpacksInstalled() bool {
	cmd := exec.Command("nixpacks", "--version")
	err := cmd.Run()
	return err == nil
}

func installNixPacks() {
	// Run the Docker installation script
	cmd := exec.Command("sh", "-c", "curl -sSL https://nixpacks.com/install.sh | bash")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error installing nixpacks: %s\n", err)
		return
	}

	fmt.Println("nixpacks installed successfully!")
}
func installDocker() {
    
	// Run the Docker installation script
	cmd := exec.Command("sh", "-c", "curl -fsSL https://get.docker.com -o get-docker.sh && sh get-docker.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error installing Docker: %s\n", err)
		return
	}

	fmt.Println("Docker installed successfully!")
}

func main() {
	if isDockerInstalled() {
		fmt.Println("Docker is already installed.")
	} else {
		installDocker()
	}
	if isNixpacksInstalled() {
		fmt.Println("Nixpacks is already installed.")
	} else {
		installNixPacks()
	}
	app := pocketbase.New()
	routers.Initialize(app)
	// loosely check if it was executed using "go run"
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// serves static files from the provided dir (if exists)
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("pb_public/public"), false))
	
		return nil
	})
	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Admin UI
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
