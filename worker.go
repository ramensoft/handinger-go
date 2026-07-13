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
	"github.com/ramensoft/handinger-go/internal/apiquery"
	shimjson "github.com/ramensoft/handinger-go/internal/encoding/json"
	"github.com/ramensoft/handinger-go/internal/requestconfig"
	"github.com/ramensoft/handinger-go/option"
	"github.com/ramensoft/handinger-go/packages/param"
	"github.com/ramensoft/handinger-go/packages/respjson"
	"github.com/ramensoft/handinger-go/shared/constant"
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
	// Configure outbound webhooks delivered when a worker's tasks complete.
	Webhooks WorkerWebhookService
}

// NewWorkerService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWorkerService(opts ...option.RequestOption) (r WorkerService) {
	r = WorkerService{}
	r.options = opts
	r.Schedules = NewWorkerScheduleService(opts...)
	r.Webhooks = NewWorkerWebhookService(opts...)
	return
}

// Create a new worker. The worker is a reusable agent template; tasks are runs
// against this template. Use `POST /tasks` to actually run the agent.
func (r *WorkerService) New(ctx context.Context, body WorkerNewParams, opts ...option.RequestOption) (res *WorkerTemplate, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/workers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve the current worker state and messages from its most recent task (or a
// specific task via `taskId`).
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

// Update a worker's instructions, title, summary, visibility, or output schema.
// Only the fields you send are changed; omitted fields keep their current values.
// Only the worker creator can update a worker.
func (r *WorkerService) Update(ctx context.Context, workerID string, body WorkerUpdateParams, opts ...option.RequestOption) (res *WorkerTemplate, err error) {
	opts = slices.Concat(r.options, opts)
	if workerID == "" {
		err = errors.New("missing required workerId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/workers/%s", url.PathEscape(workerID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Soft-delete a worker template so it no longer appears in list or retrieve
// endpoints. Tasks, turns, files, schedules, and integrations remain in the
// database for analytics. Only the worker creator can delete a worker.
func (r *WorkerService) Delete(ctx context.Context, workerID string, opts ...option.RequestOption) (res *DeleteWorkerResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if workerID == "" {
		err = errors.New("missing required workerId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/workers/%s", url.PathEscape(workerID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
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
	// Short one-line description of the worker's purpose. Auto-generated when omitted
	// and a `prompt` is provided.
	Summary param.Opt[string] `json:"summary,omitzero"`
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

type DeleteWorkerResponse struct {
	Deleted bool `json:"deleted" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deleted     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeleteWorkerResponse) RawJSON() string { return r.JSON.raw }
func (r *DeleteWorkerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UpdateWorkerParam struct {
	// Replaces the persistent system prompt. Subsequent tasks pick up the new
	// instructions immediately; in-flight tasks keep using the previous version.
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	// Replaces the worker's short one-line summary.
	Summary param.Opt[string] `json:"summary,omitzero"`
	// New display name for the worker.
	Title param.Opt[string] `json:"title,omitzero"`
	// Replace the worker's structured output schema. Pass `null` to clear it and
	// return to free-form text responses.
	OutputSchema map[string]any `json:"outputSchema,omitzero"`
	// Change visibility between `public` (any org member can run tasks) and `private`
	// (only invited members).
	//
	// Any of "public", "private".
	Visibility UpdateWorkerVisibility `json:"visibility,omitzero"`
	paramObj
}

func (r UpdateWorkerParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateWorkerParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateWorkerParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Change visibility between `public` (any org member can run tasks) and `private`
// (only invited members).
type UpdateWorkerVisibility string

const (
	UpdateWorkerVisibilityPublic  UpdateWorkerVisibility = "public"
	UpdateWorkerVisibilityPrivate UpdateWorkerVisibility = "private"
)

type Worker struct {
	ID                string          `json:"id" api:"required"`
	CreatedAt         int64           `json:"created_at" api:"required"`
	Error             any             `json:"error" api:"required"`
	Files             []WorkerFile    `json:"files" api:"required"`
	IncompleteDetails any             `json:"incomplete_details" api:"required"`
	Messages          []any           `json:"messages" api:"required"`
	Metadata          map[string]any  `json:"metadata" api:"required"`
	Object            constant.Worker `json:"object" default:"worker"`
	Output            []WorkerOutput  `json:"output" api:"required"`
	OutputText        string          `json:"output_text" api:"required"`
	Running           bool            `json:"running" api:"required"`
	Sources           []WorkerSource  `json:"sources" api:"required"`
	// Any of "running", "completed", "pending".
	Status           WorkerStatus   `json:"status" api:"required"`
	StructuredOutput map[string]any `json:"structured_output" api:"required"`
	// Web URL of the worker in the Handinger dashboard.
	URL   string      `json:"url" api:"required"`
	Usage WorkerUsage `json:"usage"`
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
		URL               respjson.Field
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
	URL       string `json:"url" api:"required" format:"uri"`
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
func (r WorkerFile) RawJSON() string { return r.JSON.raw }
func (r *WorkerFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerOutput struct {
	ID      string                `json:"id" api:"required"`
	Content []WorkerOutputContent `json:"content" api:"required"`
	Role    constant.Assistant    `json:"role" default:"assistant"`
	Status  constant.Completed    `json:"status" default:"completed"`
	Type    constant.Message      `json:"type" default:"message"`
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
	Text string              `json:"text" api:"required"`
	Type constant.OutputText `json:"type" default:"output_text"`
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
	ID    string       `json:"id" api:"required"`
	Title string       `json:"title" api:"required"`
	Type  constant.URL `json:"type" default:"url"`
	URL   string       `json:"url" api:"required"`
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
	DurationMs int64 `json:"durationMs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
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

type WorkerTemplate struct {
	ID             string         `json:"id" api:"required"`
	CreatedAt      string         `json:"createdAt" api:"required"`
	Instructions   string         `json:"instructions" api:"required"`
	OrganizationID string         `json:"organizationId" api:"required"`
	OutputSchema   map[string]any `json:"outputSchema" api:"required"`
	Summary        string         `json:"summary" api:"required"`
	Title          string         `json:"title" api:"required"`
	UpdatedAt      string         `json:"updatedAt" api:"required"`
	// Web URL of the worker in the Handinger dashboard.
	URL    string `json:"url" api:"required"`
	UserID string `json:"userId" api:"required"`
	// Any of "public", "private".
	Visibility WorkerTemplateVisibility `json:"visibility" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		Instructions   respjson.Field
		OrganizationID respjson.Field
		OutputSchema   respjson.Field
		Summary        respjson.Field
		Title          respjson.Field
		UpdatedAt      respjson.Field
		URL            respjson.Field
		UserID         respjson.Field
		Visibility     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerTemplate) RawJSON() string { return r.JSON.raw }
func (r *WorkerTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerTemplateVisibility string

const (
	WorkerTemplateVisibilityPublic  WorkerTemplateVisibility = "public"
	WorkerTemplateVisibilityPrivate WorkerTemplateVisibility = "private"
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
	// Return the worker state and messages for a specific task instead of the most
	// recent one.
	TaskID param.Opt[string] `query:"taskId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkerGetParams]'s query parameters as `url.Values`.
func (r WorkerGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkerUpdateParams struct {
	UpdateWorker UpdateWorkerParam
	paramObj
}

func (r WorkerUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateWorker)
}
func (r *WorkerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
