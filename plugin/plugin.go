package plugin

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

// Args provides plugin execution arguments.
type Args struct {
	Pipeline

	// Level defines the plugin log level.
	Level string `envconfig:"PLUGIN_LOG_LEVEL"`

	ProjectID string `envconfig:"PLUGIN_PROJECT_ID"`
}

// Exec executes the plugin.
func Exec(ctx context.Context, args Args) error {
	projectURL := DatoNotificationURL(args.ProjectID)
	if args.Pipeline.Build.Status != "success" {
		var resultMessage = []byte(`{ "status": "success" }`)
		http.Post(projectURL, "application/json", bytes.NewBuffer(resultMessage))
		fmt.Printf("Sent success webhook to DatoCMS for project %s\n", args.ProjectID)
	} else {
		var resultMessage = []byte(`{ "status": "error" }`)
		http.Post(projectURL, "application/json", bytes.NewBuffer(resultMessage))
		fmt.Printf("Send failure webhook to DatoCMS for project %s\n", args.ProjectID)
	}

	return nil
}

func DatoNotificationURL(projectId string) string {
	return fmt.Sprintf("https://webhooks.datocms.com/%s/deploy-results", projectId)
}

func stepFailed(step string, failedSteps []string) bool {
	for _, failedStep := range failedSteps {
		if step == failedStep {
			return true
		}
	}
	return false
}
