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

func TestWorkerWebhookGet(t *testing.T) {
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
	_, err := client.Workers.Webhooks.Get(context.TODO(), "t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM")
	if err != nil {
		var apierr *handinger.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkerWebhookUpdate(t *testing.T) {
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
	_, err := client.Workers.Webhooks.Update(
		context.TODO(),
		"t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM",
		handinger.WorkerWebhookUpdateParams{
			UpdateWebhook: handinger.UpdateWebhookParam{
				URL: handinger.String("https://example.com/handinger-webhook"),
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

func TestWorkerWebhookDelete(t *testing.T) {
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
	_, err := client.Workers.Webhooks.Delete(context.TODO(), "t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM")
	if err != nil {
		var apierr *handinger.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkerWebhookListExecutionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Webhooks.ListExecutions(
		context.TODO(),
		"t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM",
		handinger.WorkerWebhookListExecutionsParams{
			Page: handinger.Int(1),
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

func TestWorkerWebhookRegenerateToken(t *testing.T) {
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
	_, err := client.Workers.Webhooks.RegenerateToken(context.TODO(), "t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM")
	if err != nil {
		var apierr *handinger.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
