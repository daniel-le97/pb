// main.go
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	// uncomment once you have at least one .go migration file in the "migrations" directory
	_ "leploy/migrations"
	_ "leploy/routers"
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

// func initialize(app *pocketbase.PocketBase) {
// 	fmt.Println("Initialization logic in init function")
// 	cmd := exec.Command("bun", "smee.ts")
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr

// 	err := cmd.Start()
// 	if err != nil {
// 		fmt.Printf("Error running Node.js script: %s\n", err)
// 		return
// 	}

// 	fmt.Println("Node.js script executed successfully")
// 	app.RootCmd.AddCommand(&cobra.Command{
// 		Use: "hello",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			log.Println("Hello world!")
// 		},
// 	})

// 	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
// 		e.Router.POST("/webhooks/:name", func(c echo.Context) error {
// 			name := c.PathParam("name")
// 			log.Println("Hello world!")
// 			headers := c.Request().Header.Get("x-github-event")
// 			log.Println(headers)
// 			return c.JSON(http.StatusOK, map[string]string{"message": "Hello " + name})
// 		} /* optional middlewares */)
// 		e.Router.GET("/webhooks/:name", func(c echo.Context) error {
// 			name := c.PathParam("name")
// 			log.Println("Hello world!")
// 			return c.JSON(http.StatusOK, map[string]string{"message": "Hello " + name})
// 		} /* optional middlewares */)

// 		return nil
// 	})
// }

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
	initialize(app)
	// loosely check if it was executed using "go run"
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Admin UI
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
