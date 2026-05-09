// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package handinger_test

import (
	"context"
	"os"
	"testing"

	"github.com/Ramensoft/handinger-go"
	"github.com/Ramensoft/handinger-go/internal/testutil"
	"github.com/Ramensoft/handinger-go/option"
)

func TestUsage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := handinger.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	worker, err := client.Tasks.New(context.TODO(), handinger.TaskNewParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", worker.ID)
}
