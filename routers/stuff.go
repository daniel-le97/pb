package router

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/v39/github"
	"github.com/labstack/echo"
)

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

type EventPayload struct {
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
	Pusher struct {
		Name string `json:"name"`
	} `json:"pusher"`
	Ref     string        `json:"ref"`
	Commits []CommitEntry `json:"commits"`
}

type CommitEntry struct {
	ID string `json:"id"`
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
	tr := http.DefaultTransport
	itr, err := ghinstallation.NewKeyFromFile(tr, 12345, 123456789, "/config/github-app.pem")

	if err != nil {
		log.Fatal(err)
	}

	config.GitHubClient = github.NewClient(&http.Client{Transport: itr})
}

func VerifySignature(payload []byte, signature string) bool {
	key := hmac.New(sha256.New, []byte(config.GitHubWebhookSecret))
	key.Write([]byte(string(payload)))
	computedSignature := "sha256=" + hex.EncodeToString(key.Sum(nil))
	log.Printf("computed signature: %s", computedSignature)

	return computedSignature == signature
}

func ConsumeEvent(c echo.Context) error {
	payload, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
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
			return c.NoContent(http.StatusNoContent)
		}
	}

	log.Printf("Unsupported event: %s", event)
	return c.JSON(http.StatusNotImplemented, map[string]interface{}{"reason": "Unsupported event: " + event})
}

