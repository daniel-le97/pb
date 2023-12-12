package router

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/spf13/cobra"
	// uncomment once you have at least one .go migration file in the "migrations" directory
)

func Initialize(app *pocketbase.PocketBase) {
	fmt.Println("Initialization logic in init function")
    cmd := exec.Command("bun", "smee.ts")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		fmt.Printf("Error running Node.js script: %s\n", err)
		return
	}

	fmt.Println("Node.js script executed successfully")
	app.RootCmd.AddCommand(&cobra.Command{
		Use: "hello",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Hello world!")
		},
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.POST("/webhooks/:name", func(c echo.Context) error {
			name := c.PathParam("name")
            // log.Println("Hello world!")
            headers := c.Request()
			// if(err != nil){
			// 	return c.JSON(http.StatusBadRequest, map[string]string{"message": "Hello " + name})
				
			// }
			log.Println(headers)
			return c.JSON(http.StatusOK, map[string]string{"message": "Hello " + name})
		}, apis.ActivityLogger(app))
		// e.Router.GET("/webhooks/:name", func(c echo.Context) error {
		// 	name := c.PathParam("name")
        //     log.Println("Hello world!")
		// 	return c.JSON(http.StatusOK, map[string]string{"message": "Hello " + name})
		// } /* optional middlewares */)

		return nil
	})
}