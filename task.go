// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package handinger

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/ramensoft/handinger-go/internal/apijson"
	shimjson "github.com/ramensoft/handinger-go/internal/encoding/json"
	"github.com/ramensoft/handinger-go/internal/requestconfig"
	"github.com/ramensoft/handinger-go/option"
	"github.com/ramensoft/handinger-go/packages/param"
	"github.com/ramensoft/handinger-go/packages/respjson"
)

// Run and inspect tasks against a worker.
//
// TaskService contains methods and other services that help with interacting with
// the handinger API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTaskService] method instead.
type TaskService struct {
	options []option.RequestOption
}

// NewTaskService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTaskService(opts ...option.RequestOption) (r TaskService) {
	r = TaskService{}
	r.options = opts
	return
}

// Run a new task against an existing worker and wait for the result. Send a
// `taskId` of a prior task to add a follow-up turn instead of starting a fresh
// task. Send `multipart/form-data` to attach files; the bytes are bootstrapped
// into the worker's workspace before the task starts. The task runs to completion
// on the server even if the connection drops; subscribe to task webhooks for
// long-running tasks.
func (r *TaskService) New(ctx context.Context, body TaskNewParams, opts ...option.RequestOption) (res *Worker, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/tasks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a single task and its individual turns.
func (r *TaskService) Get(ctx context.Context, taskID string, opts ...option.RequestOption) (res *TaskWithTurns, err error) {
	opts = slices.Concat(r.options, opts)
	if taskID == "" {
		err = errors.New("missing required taskId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/tasks/%s", url.PathEscape(taskID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Archive a task so it stops appearing in `GET /tasks` results. Turns and files
// are retained for audit purposes. Only the worker creator can archive a task.
func (r *TaskService) Delete(ctx context.Context, taskID string, opts ...option.RequestOption) (res *DeleteTaskResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if taskID == "" {
		err = errors.New("missing required taskId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/tasks/%s", url.PathEscape(taskID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// The property Input is required.
type CreateTaskParam struct {
	Input string `json:"input" api:"required"`
	// Optional client-provided task id. Reuse this id to add turns to an existing
	// task.
	TaskID param.Opt[string] `json:"taskId,omitzero"`
	// Worker id the task belongs to. If omitted, a new worker is created on-the-fly
	// using the input as instructions.
	WorkerID param.Opt[string] `json:"workerId,omitzero"`
	// Compute budget the worker is allowed to spend on the task. Defaults to
	// `standard`.
	//
	// Any of "low", "standard", "high", "unlimited".
	Budget CreateTaskBudget `json:"budget,omitzero"`
	paramObj
}

func (r CreateTaskParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateTaskParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateTaskParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compute budget the worker is allowed to spend on the task. Defaults to
// `standard`.
type CreateTaskBudget string

const (
	CreateTaskBudgetLow       CreateTaskBudget = "low"
	CreateTaskBudgetStandard  CreateTaskBudget = "standard"
	CreateTaskBudgetHigh      CreateTaskBudget = "high"
	CreateTaskBudgetUnlimited CreateTaskBudget = "unlimited"
)

type DeleteTaskResponse struct {
	Archived bool `json:"archived" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Archived    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeleteTaskResponse) RawJSON() string { return r.JSON.raw }
func (r *DeleteTaskResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Task struct {
	ID              string `json:"id" api:"required"`
	CompletedAt     string `json:"completedAt" api:"required"`
	CreatedAt       string `json:"createdAt" api:"required"`
	CreatedByUserID string `json:"createdByUserId" api:"required"`
	OrganizationID  string `json:"organizationId" api:"required"`
	// Any of "pending", "running", "completed", "error", "aborted".
	Status TaskStatus `json:"status" api:"required"`
	Title  string     `json:"title" api:"required"`
	// Aggregate credit spend, elapsed wall-clock, and number of turns across the task.
	Totals TaskTotals `json:"totals" api:"required"`
	// Any of "api", "email", "schedule", "ui".
	TriggeredBy TaskTriggeredBy `json:"triggeredBy" api:"required"`
	// Web URL of the task in the Handinger dashboard.
	URL      string `json:"url" api:"required"`
	WorkerID string `json:"workerId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CompletedAt     respjson.Field
		CreatedAt       respjson.Field
		CreatedByUserID respjson.Field
		OrganizationID  respjson.Field
		Status          respjson.Field
		Title           respjson.Field
		Totals          respjson.Field
		TriggeredBy     respjson.Field
		URL             respjson.Field
		WorkerID        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Task) RawJSON() string { return r.JSON.raw }
func (r *Task) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusRunning   TaskStatus = "running"
	TaskStatusCompleted TaskStatus = "completed"
	TaskStatusError     TaskStatus = "error"
	TaskStatusAborted   TaskStatus = "aborted"
)

// Aggregate credit spend, elapsed wall-clock, and number of turns across the task.
type TaskTotals struct {
	Credits    int64 `json:"credits" api:"required"`
	DurationMs int64 `json:"durationMs" api:"required"`
	TurnCount  int64 `json:"turnCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Credits     respjson.Field
		DurationMs  respjson.Field
		TurnCount   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskTotals) RawJSON() string { return r.JSON.raw }
func (r *TaskTotals) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskTriggeredBy string

const (
	TaskTriggeredByAPI      TaskTriggeredBy = "api"
	TaskTriggeredByEmail    TaskTriggeredBy = "email"
	TaskTriggeredBySchedule TaskTriggeredBy = "schedule"
	TaskTriggeredByUi       TaskTriggeredBy = "ui"
)

type TaskWithTurns struct {
	Task  Task                `json:"task" api:"required"`
	Turns []TaskWithTurnsTurn `json:"turns" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Task        respjson.Field
		Turns       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskWithTurns) RawJSON() string { return r.JSON.raw }
func (r *TaskWithTurns) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskWithTurnsTurn struct {
	ID          string `json:"id" api:"required"`
	CompletedAt string `json:"completedAt" api:"required"`
	Credits     int64  `json:"credits" api:"required"`
	DurationMs  int64  `json:"durationMs" api:"required"`
	// Files published by this turn.
	Files        []TaskWithTurnsTurnFile `json:"files" api:"required"`
	Input        string                  `json:"input" api:"required"`
	InputTokens  int64                   `json:"inputTokens" api:"required"`
	OutputText   string                  `json:"outputText" api:"required"`
	OutputTokens int64                   `json:"outputTokens" api:"required"`
	Role         string                  `json:"role" api:"required"`
	Seq          int64                   `json:"seq" api:"required"`
	StartedAt    string                  `json:"startedAt" api:"required"`
	Status       string                  `json:"status" api:"required"`
	// Structured JSON payload when the worker is configured with an output schema.
	// `null` otherwise.
	StructuredOutput map[string]any `json:"structuredOutput" api:"required"`
	TaskID           string         `json:"taskId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CompletedAt      respjson.Field
		Credits          respjson.Field
		DurationMs       respjson.Field
		Files            respjson.Field
		Input            respjson.Field
		InputTokens      respjson.Field
		OutputText       respjson.Field
		OutputTokens     respjson.Field
		Role             respjson.Field
		Seq              respjson.Field
		StartedAt        respjson.Field
		Status           respjson.Field
		StructuredOutput respjson.Field
		TaskID           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskWithTurnsTurn) RawJSON() string { return r.JSON.raw }
func (r *TaskWithTurnsTurn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskWithTurnsTurnFile struct {
	Filename  string `json:"filename" api:"required"`
	MediaType string `json:"mediaType" api:"required"`
	URL       string `json:"url" api:"required"`
	Size      int64  `json:"size"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filename    respjson.Field
		MediaType   respjson.Field
		URL         respjson.Field
		Size        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskWithTurnsTurnFile) RawJSON() string { return r.JSON.raw }
func (r *TaskWithTurnsTurnFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskNewParams struct {
	CreateTask CreateTaskParam
	paramObj
}

func (r TaskNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateTask)
}
func (r *TaskNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
