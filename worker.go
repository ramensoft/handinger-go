// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package handinger

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/Ramensoft/handinger-go/internal/apijson"
	"github.com/Ramensoft/handinger-go/internal/apiquery"
	shimjson "github.com/Ramensoft/handinger-go/internal/encoding/json"
	"github.com/Ramensoft/handinger-go/internal/requestconfig"
	"github.com/Ramensoft/handinger-go/option"
	"github.com/Ramensoft/handinger-go/packages/param"
	"github.com/Ramensoft/handinger-go/packages/respjson"
)

// Create, retrieve, and manage agent worker templates.
//
// WorkerService contains methods and other services that help with interacting
// with the handinger API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkerService] method instead.
type WorkerService struct {
	options []option.RequestOption
	// Manage future and recurring worker tasks.
	Schedules WorkerScheduleService
}

// NewWorkerService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWorkerService(opts ...option.RequestOption) (r WorkerService) {
	r = WorkerService{}
	r.options = opts
	r.Schedules = NewWorkerScheduleService(opts...)
	return
}

// Create a new worker. The worker is a reusable agent template; tasks are runs
// against this template. Use `POST /tasks` to actually run the agent.
func (r *WorkerService) New(ctx context.Context, body WorkerNewParams, opts ...option.RequestOption) (res *WorkerNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/workers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve the current worker state and messages from its most recent task.
// Returns a JSON worker object by default, or a server-sent event stream when
// `stream=true`.
func (r *WorkerService) Get(ctx context.Context, workerID string, query WorkerGetParams, opts ...option.RequestOption) (res *Worker, err error) {
	opts = slices.Concat(r.options, opts)
	if workerID == "" {
		err = errors.New("missing required workerId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/workers/%s", url.PathEscape(workerID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the inbound email address for a worker.
func (r *WorkerService) GetEmail(ctx context.Context, workerID string, opts ...option.RequestOption) (res *WorkerGetEmailResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if workerID == "" {
		err = errors.New("missing required workerId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/workers/%s/email", url.PathEscape(workerID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type CreateWorkerParam struct {
	// Persistent system prompt the worker uses for every task it runs.
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	// Natural-language description of the worker to use for AI-generated instructions
	// when `instructions` is omitted.
	Prompt param.Opt[string] `json:"prompt,omitzero"`
	// Optional display name. When omitted, Handinger assigns a random dog-themed name.
	Title param.Opt[string] `json:"title,omitzero"`
	// Optional JSON Schema (Draft-07) describing the structured object the worker must
	// produce. When set, every task response is validated against the schema and
	// exposed as `structuredOutput`.
	OutputSchema map[string]any `json:"outputSchema,omitzero"`
	// `public` (default) is visible to all org members. `private` is only visible to
	// invited members.
	//
	// Any of "public", "private".
	Visibility CreateWorkerVisibility `json:"visibility,omitzero"`
	paramObj
}

func (r CreateWorkerParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateWorkerParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateWorkerParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `public` (default) is visible to all org members. `private` is only visible to
// invited members.
type CreateWorkerVisibility string

const (
	CreateWorkerVisibilityPublic  CreateWorkerVisibility = "public"
	CreateWorkerVisibilityPrivate CreateWorkerVisibility = "private"
)

type Worker struct {
	ID                string         `json:"id" api:"required"`
	CreatedAt         int64          `json:"created_at" api:"required"`
	Error             any            `json:"error" api:"required"`
	Files             []WorkerFile   `json:"files" api:"required"`
	IncompleteDetails any            `json:"incomplete_details" api:"required"`
	Messages          []any          `json:"messages" api:"required"`
	Metadata          map[string]any `json:"metadata" api:"required"`
	// Any of "worker".
	Object     WorkerObject   `json:"object" api:"required"`
	Output     []WorkerOutput `json:"output" api:"required"`
	OutputText string         `json:"output_text" api:"required"`
	Running    bool           `json:"running" api:"required"`
	Sources    []WorkerSource `json:"sources" api:"required"`
	// Any of "running", "completed", "pending".
	Status           WorkerStatus   `json:"status" api:"required"`
	StructuredOutput map[string]any `json:"structured_output" api:"required"`
	Usage            WorkerUsage    `json:"usage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		Error             respjson.Field
		Files             respjson.Field
		IncompleteDetails respjson.Field
		Messages          respjson.Field
		Metadata          respjson.Field
		Object            respjson.Field
		Output            respjson.Field
		OutputText        respjson.Field
		Running           respjson.Field
		Sources           respjson.Field
		Status            respjson.Field
		StructuredOutput  respjson.Field
		Usage             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Worker) RawJSON() string { return r.JSON.raw }
func (r *Worker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerFile struct {
	Filename  string `json:"filename" api:"required"`
	MediaType string `json:"mediaType" api:"required"`
	URL       string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filename    respjson.Field
		MediaType   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerFile) RawJSON() string { return r.JSON.raw }
func (r *WorkerFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerObject string

const (
	WorkerObjectWorker WorkerObject = "worker"
)

type WorkerOutput struct {
	ID      string                `json:"id" api:"required"`
	Content []WorkerOutputContent `json:"content" api:"required"`
	// Any of "assistant".
	Role string `json:"role" api:"required"`
	// Any of "completed".
	Status string `json:"status" api:"required"`
	// Any of "message".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Content     respjson.Field
		Role        respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerOutput) RawJSON() string { return r.JSON.raw }
func (r *WorkerOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerOutputContent struct {
	Text string `json:"text" api:"required"`
	// Any of "output_text".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerOutputContent) RawJSON() string { return r.JSON.raw }
func (r *WorkerOutputContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerSource struct {
	ID    string `json:"id" api:"required"`
	Title string `json:"title" api:"required"`
	// Any of "url".
	Type string `json:"type" api:"required"`
	URL  string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Title       respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerSource) RawJSON() string { return r.JSON.raw }
func (r *WorkerSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerStatus string

const (
	WorkerStatusRunning   WorkerStatus = "running"
	WorkerStatusCompleted WorkerStatus = "completed"
	WorkerStatusPending   WorkerStatus = "pending"
)

type WorkerUsage struct {
	Credits    int64 `json:"credits"`
	DurationMs int64 `json:"durationMs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Credits     respjson.Field
		DurationMs  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerUsage) RawJSON() string { return r.JSON.raw }
func (r *WorkerUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerNewResponse struct {
	ID             string         `json:"id" api:"required"`
	CreatedAt      string         `json:"createdAt" api:"required"`
	Instructions   string         `json:"instructions" api:"required"`
	OrganizationID string         `json:"organizationId" api:"required"`
	OutputSchema   map[string]any `json:"outputSchema" api:"required"`
	Title          string         `json:"title" api:"required"`
	UpdatedAt      string         `json:"updatedAt" api:"required"`
	UserID         string         `json:"userId" api:"required"`
	// Any of "public", "private".
	Visibility WorkerNewResponseVisibility `json:"visibility" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		Instructions   respjson.Field
		OrganizationID respjson.Field
		OutputSchema   respjson.Field
		Title          respjson.Field
		UpdatedAt      respjson.Field
		UserID         respjson.Field
		Visibility     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkerNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerNewResponseVisibility string

const (
	WorkerNewResponseVisibilityPublic  WorkerNewResponseVisibility = "public"
	WorkerNewResponseVisibilityPrivate WorkerNewResponseVisibility = "private"
)

type WorkerGetEmailResponse struct {
	Email string `json:"email" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerGetEmailResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkerGetEmailResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerNewParams struct {
	CreateWorker CreateWorkerParam
	paramObj
}

func (r WorkerNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateWorker)
}
func (r *WorkerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerGetParams struct {
	// Set to "true" to receive a server-sent event stream that replays all stored
	// messages and then continues with live chunks from the active task (if any)
	// before closing.
	//
	// Any of "true", "false".
	Stream WorkerGetParamsStream `query:"stream,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkerGetParams]'s query parameters as `url.Values`.
func (r WorkerGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Set to "true" to receive a server-sent event stream that replays all stored
// messages and then continues with live chunks from the active task (if any)
// before closing.
type WorkerGetParamsStream string

const (
	WorkerGetParamsStreamTrue  WorkerGetParamsStream = "true"
	WorkerGetParamsStreamFalse WorkerGetParamsStream = "false"
)
