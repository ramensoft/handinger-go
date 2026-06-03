// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package handinger_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/ramensoft/handinger-go"
	"github.com/ramensoft/handinger-go/internal/testutil"
	"github.com/ramensoft/handinger-go/option"
)

func TestWorkerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.New(context.TODO(), handinger.WorkerNewParams{
		CreateWorker: handinger.CreateWorkerParam{
			Instructions: handinger.String("You are a brand voice analyzer. Read the input text and report whether it matches Acme's playful, plain-spoken house style. Quote specific phrases."),
			OutputSchema: map[string]any{
				"type":       "bar",
				"required":   "bar",
				"properties": "bar",
			},
			Prompt:     handinger.String("A worker that fact-checks short claims and returns a verdict with citations."),
			Summary:    handinger.String("Audits copy against the Acme brand voice guide."),
			Title:      handinger.String("Brand voice analyzer"),
			Visibility: handinger.CreateWorkerVisibilityPublic,
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

func TestWorkerGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Get(
		context.TODO(),
		"t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM",
		handinger.WorkerGetParams{
			TaskID: handinger.String("x"),
		},
	)
	if err != nil {
		var apierr *handinger.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Update(
		context.TODO(),
		"t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM",
		handinger.WorkerUpdateParams{
			UpdateWorker: handinger.UpdateWorkerParam{
				Instructions: handinger.String("You are a brand voice analyzer. Read the input text and report whether it matches Acme's playful, plain-spoken house style. Quote specific phrases."),
				OutputSchema: map[string]any{
					"type":       "bar",
					"required":   "bar",
					"properties": "bar",
				},
				Summary:    handinger.String("Audits copy against the Acme brand voice guide."),
				Title:      handinger.String("Claim verdict v2"),
				Visibility: handinger.UpdateWorkerVisibilityPrivate,
			},
		},
	)
	if err != nil {
		var apierr *handinger.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkerDelete(t *testing.T) {
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
	_, err := client.Workers.Delete(context.TODO(), "t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM")
	if err != nil {
		var apierr *handinger.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkerGetEmail(t *testing.T) {
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
	_, err := client.Workers.GetEmail(context.TODO(), "t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM")
	if err != nil {
		var apierr *handinger.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
