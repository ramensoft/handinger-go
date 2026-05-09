// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package handinger_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Ramensoft/handinger-go"
	"github.com/Ramensoft/handinger-go/internal/testutil"
	"github.com/Ramensoft/handinger-go/option"
)

func TestTaskNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Tasks.New(context.TODO(), handinger.TaskNewParams{
		CreateTask: handinger.CreateTaskParam{
			CreateWorkerParam: handinger.CreateWorkerParam{
				Instructions: handinger.String("instructions"),
				OutputSchema: map[string]any{
					"foo": "bar",
				},
				Prompt:     handinger.String("prompt"),
				Summary:    handinger.String("summary"),
				Title:      handinger.String("Brand voice analyzer"),
				Visibility: handinger.CreateWorkerVisibilityPublic,
			},
			WorkerID: "wrk_vk81XUHKHG-qr4",
			TaskID:   handinger.String("tsk_2Z-YWz3hFq6VlW"),
		},
	})
	if err != nil {
		var apierr *handinger.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTaskGet(t *testing.T) {
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
	_, err := client.Tasks.Get(context.TODO(), "tsk_01HZY31W2SZJ8MJ2FQTR3M1K9D")
	if err != nil {
		var apierr *handinger.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
