package routers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"


	// "io/ioutil"

	"net/http"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/labstack/echo/v5"

	// router "github.com/daniel-le97/pb/routers"
	"github.com/google/go-github/v39/github"
	"github.com/joho/godotenv"

	// "github.com/labstack/echo"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/spf13/cobra"
	// uncomment once you have at least one .go migration file in the "migrations" directory
)
type Config struct {
	GitHubClient         *github.Client
	GitHubWebhookSecret  string
	GitHubPrivateKey  string
	ServerPort           string
}

var config Config
type Event string

const (
	Install     Event = "installation"
	Ping        Event = "ping"
	Push        Event = "push"
	PullRequest Event = "pull_request"
)

var Events = []Event{
	Install,
	Ping,
	Push,
	PullRequest,
}



var Consumers = map[string]func(EventPayload) error{
	string(Install):     consumeInstallEvent,
	string(Ping):        consumePingEvent,
	string(Push):        consumePushEvent,
	string(PullRequest): consumePullRequestEvent,
}

func consumeInstallEvent(payload EventPayload) error {
	// Handle installation event...
	return nil
}

func consumePingEvent(payload EventPayload) error {
	// Handle ping event...
	return nil
}

func consumePushEvent(payload EventPayload) error {
	// Process push event...
	log.Printf("Received push from %s, by user %s, on branch %s",
		payload.Repository.FullName,
		payload.Pusher.Name,
		payload.Ref)

		fmt.Println(payload.Repository.HTMLURL)
	// Enumerating commits
	var commits []string
	for _, commit := range payload.Commits {
		commits = append(commits, commit.ID)
	}
	log.Printf("Pushed commits: %v", commits)

	return nil
}

func consumePullRequestEvent(payload EventPayload) error {
	// Handle pull request event...
	return nil
}

func InitGitHubClient() {
	if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
	config.GitHubPrivateKey = os.Getenv("PRIVATE_KEY")
	tr := http.DefaultTransport
	itr, err := ghinstallation.NewKeyFromFile(tr, 12345, 123456789, config.GitHubPrivateKey)

	if err != nil {
		log.Fatal(err)
	}

	config.GitHubClient = github.NewClient(&http.Client{Transport: itr})
}

func VerifySignature(payload []byte, signature string) bool {
	if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
	config.GitHubWebhookSecret = os.Getenv("WEBHOOK_SECRET")
	key := hmac.New(sha256.New, []byte(config.GitHubWebhookSecret))
	key.Write([]byte(string(payload)))
	computedSignature := "sha256=" + hex.EncodeToString(key.Sum(nil))
	log.Printf("computed signature: %s", computedSignature)

	return computedSignature == signature
}

func ConsumeEvent(c echo.Context) error {
	payload, err := io.ReadAll(c.Request().Body)
	// fmt.Println()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if !VerifySignature(payload, c.Request().Header.Get("X-Hub-Signature-256")) {
		return c.String(http.StatusUnauthorized, "signatures don't match")
	}

	event := c.Request().Header.Get("X-GitHub-Event")

	for _, e := range Events {
		if string(e) == event {
			log.Printf("consuming event: %s", e)
			var p EventPayload
			json.Unmarshal(payload, &p)
			if err := Consumers[string(e)](p); err != nil {
				log.Printf("couldn't consume event %s, error: %+v", string(e), err)
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"reason": err})
			}

			log.Printf("consumed event: %s", e)

			// TODO probably want to find all the projects that have this repoURL




			log.Printf("consumed payload: %+v", p)
			return c.NoContent(http.StatusNoContent)
		}
	}

	log.Printf("Unsupported event: %s", event)
	return c.JSON(http.StatusNotImplemented, map[string]interface{}{"reason": "Unsupported event: " + event})
}

func Initialize(app *pocketbase.PocketBase) {
	if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

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
	// 	e.Router.POST("/webhooks/:name", func(c echo.Context) error {
	// 		name := c.PathParam("name")
	// 		body, err := io.ReadAll(c.Request().Body)
	// 		if err != nil {
	// 			log.Println("Error reading request body:", err)
	// 			return c.String(http.StatusInternalServerError, "Internal Server Error")
	// 		}
	// 		var myRequest EventPayload
	// if err := json.Unmarshal(body, &myRequest); err != nil {
	// 	log.Println("Error decoding JSON:", err)
	// 	return c.String(http.StatusBadRequest, "Bad Request")
	// }
	// log.Printf("Received request: %+v", myRequest)
	
	// 		return c.JSON(http.StatusOK, map[string]string{"message": "Hello " + name})
	// 	}, /* optional middlewares */)
		e.Router.POST("/webhooks/:name", ConsumeEvent /* optional middlewares */)

		return nil
	})
}