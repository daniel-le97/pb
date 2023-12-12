package router

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
	"time"

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

type EventPayload struct {
	Repository struct {
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
		Owner    struct {
			Name              string `json:"name"`
			Email             string `json:"email"`
			Login             string `json:"login"`
			ID                int    `json:"id"`
			NodeID            string `json:"node_id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"owner"`
		HTMLURL                  string    `json:"html_url"`
		Description              any       `json:"description"`
		Fork                     bool      `json:"fork"`
		URL                      string    `json:"url"`
		ForksURL                 string    `json:"forks_url"`
		KeysURL                  string    `json:"keys_url"`
		CollaboratorsURL         string    `json:"collaborators_url"`
		TeamsURL                 string    `json:"teams_url"`
		HooksURL                 string    `json:"hooks_url"`
		IssueEventsURL           string    `json:"issue_events_url"`
		EventsURL                string    `json:"events_url"`
		AssigneesURL             string    `json:"assignees_url"`
		BranchesURL              string    `json:"branches_url"`
		TagsURL                  string    `json:"tags_url"`
		BlobsURL                 string    `json:"blobs_url"`
		GitTagsURL               string    `json:"git_tags_url"`
		GitRefsURL               string    `json:"git_refs_url"`
		TreesURL                 string    `json:"trees_url"`
		StatusesURL              string    `json:"statuses_url"`
		LanguagesURL             string    `json:"languages_url"`
		StargazersURL            string    `json:"stargazers_url"`
		ContributorsURL          string    `json:"contributors_url"`
		SubscribersURL           string    `json:"subscribers_url"`
		SubscriptionURL          string    `json:"subscription_url"`
		CommitsURL               string    `json:"commits_url"`
		GitCommitsURL            string    `json:"git_commits_url"`
		CommentsURL              string    `json:"comments_url"`
		IssueCommentURL          string    `json:"issue_comment_url"`
		ContentsURL              string    `json:"contents_url"`
		CompareURL               string    `json:"compare_url"`
		MergesURL                string    `json:"merges_url"`
		ArchiveURL               string    `json:"archive_url"`
		DownloadsURL             string    `json:"downloads_url"`
		IssuesURL                string    `json:"issues_url"`
		PullsURL                 string    `json:"pulls_url"`
		MilestonesURL            string    `json:"milestones_url"`
		NotificationsURL         string    `json:"notifications_url"`
		LabelsURL                string    `json:"labels_url"`
		ReleasesURL              string    `json:"releases_url"`
		DeploymentsURL           string    `json:"deployments_url"`
		CreatedAt                int       `json:"created_at"`
		UpdatedAt                time.Time `json:"updated_at"`
		PushedAt                 int       `json:"pushed_at"`
		GitURL                   string    `json:"git_url"`
		SSHURL                   string    `json:"ssh_url"`
		CloneURL                 string    `json:"clone_url"`
		SvnURL                   string    `json:"svn_url"`
		Homepage                 any       `json:"homepage"`
		Size                     int       `json:"size"`
		StargazersCount          int       `json:"stargazers_count"`
		WatchersCount            int       `json:"watchers_count"`
		Language                 string    `json:"language"`
		HasIssues                bool      `json:"has_issues"`
		HasProjects              bool      `json:"has_projects"`
		HasDownloads             bool      `json:"has_downloads"`
		HasWiki                  bool      `json:"has_wiki"`
		HasPages                 bool      `json:"has_pages"`
		HasDiscussions           bool      `json:"has_discussions"`
		ForksCount               int       `json:"forks_count"`
		MirrorURL                any       `json:"mirror_url"`
		Archived                 bool      `json:"archived"`
		Disabled                 bool      `json:"disabled"`
		OpenIssuesCount          int       `json:"open_issues_count"`
		License                  any       `json:"license"`
		AllowForking             bool      `json:"allow_forking"`
		IsTemplate               bool      `json:"is_template"`
		WebCommitSignoffRequired bool      `json:"web_commit_signoff_required"`
		Topics                   []any     `json:"topics"`
		Visibility               string    `json:"visibility"`
		Forks                    int       `json:"forks"`
		OpenIssues               int       `json:"open_issues"`
		Watchers                 int       `json:"watchers"`
		DefaultBranch            string    `json:"default_branch"`
		Stargazers               int       `json:"stargazers"`
		MasterBranch             string    `json:"master_branch"`
	} `json:"repository"`
	Pusher struct {
		Name  string `json:"name"`
		Email any    `json:"email"`
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