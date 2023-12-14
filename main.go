// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strings"

	routers "github.com/daniel-le97/pb/routers"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"

	// "github.com/pocketbase/pocketbase/db"

	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	// "github.com/pocketbase/pocketbase/models"
	// uncomment once you have at least one .go migration file in the "migrations" directory
	_ "github.com/daniel-le97/pb/migrations"
	model "github.com/daniel-le97/pb/models"
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

func ConvertRecordToStruct(record *models.Record, targetStruct interface{}) error {
	recordJSON, err := json.Marshal(record)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(recordJSON, targetStruct); err != nil {
		return err
	}

	return nil
}

func printStructFields(data interface{}) {
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i).Interface()
		fmt.Printf("%s: %v\n", field.Name, fieldValue)
	}
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
	app.OnModelAfterCreate("queue").Add(func(e *core.ModelEvent) error {

	
		
		
		QueueRecord, err := app.Dao().FindRecordById("queue", e.Model.GetId())
	if err != nil {
		return err
	}

	// Create an instance of Users to hold the data
	var queue model.Queue

	// Use the helper function to convert the *models.Record to Users
	if err := ConvertRecordToStruct(QueueRecord, &queue); err != nil {
		fmt.Println("Error:", err)
		return err
	}
	printStructFields(queue)

	ProjectRecord, err := app.Dao().FindRecordById("queue", e.Model.GetId())
	if err != nil {
		return err
	}
	var project model.Project
	if err := ConvertRecordToStruct(ProjectRecord, &project); err != nil {
		fmt.Println("Error:", err)
		return err
	}
	printStructFields(project)
	// Now 'users' holds the data from the *models.Record in Users struct
        return nil
    })

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
