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

func TestWorkerScheduleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Workers.Schedules.New(
		context.TODO(),
		"t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM",
		handinger.WorkerScheduleNewParams{
			Input: "x",
			When: handinger.WorkerScheduleNewParamsWhenUnion{
				OfScheduled: &handinger.WorkerScheduleNewParamsWhenScheduled{
					Date: "x",
				},
			},
			Budget: handinger.WorkerScheduleNewParamsBudgetLow,
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

func TestWorkerScheduleList(t *testing.T) {
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
	_, err := client.Workers.Schedules.List(context.TODO(), "t_org_123_w_01HZY2ZJQ8G7K42W2D7WF6V4GM")
	if err != nil {
		var apierr *handinger.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkerScheduleCancel(t *testing.T) {
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
	_, err := client.Workers.Schedules.Cancel(
		context.TODO(),
		"sch_01HZY31W2SZJ8MJ2FQTR3M1K9D",
		handinger.WorkerScheduleCancelParams{
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
}
