// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package handinger_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stainless-sdks/handinger-go"
	"github.com/stainless-sdks/handinger-go/internal/testutil"
	"github.com/stainless-sdks/handinger-go/option"
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
			Input:  "x",
			Budget: handinger.CreateWorkerBudgetLow,
			Stream: handinger.Bool(true),
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
			Stream: handinger.Bool(true),
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

func TestWorkerContinueWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Continue(
		context.TODO(),
		"t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM",
		handinger.WorkerContinueParams{
			CreateWorker: handinger.CreateWorkerParam{
				Input:  "x",
				Budget: handinger.CreateWorkerBudgetLow,
				Stream: handinger.Bool(true),
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

func TestWorkerGetFile(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := handinger.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	resp, err := client.Workers.GetFile(
		context.TODO(),
		"scratchpad/plan.md",
		handinger.WorkerGetFileParams{
			WorkerID: "t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM",
		},
	)
	if err != nil {
		var apierr *handinger.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *handinger.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}
