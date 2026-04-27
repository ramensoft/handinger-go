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

// Create, retrieve, and continue agent workers.
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

// Create a new agent worker and start it with the supplied instruction. Send
// `multipart/form-data` to attach files alongside the instruction; the bytes are
// bootstrapped into the worker's workspace before the first turn.
func (r *WorkerService) New(ctx context.Context, body WorkerNewParams, opts ...option.RequestOption) (res *Worker, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/workers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve the current worker state and messages. Returns a JSON worker object by
// default, or a server-sent event stream when `stream=true`.
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

// Send another instruction to an existing worker. Send `multipart/form-data` to
// attach additional files; the bytes are bootstrapped into the worker's workspace
// before the next turn.
func (r *WorkerService) Continue(ctx context.Context, workerID string, body WorkerContinueParams, opts ...option.RequestOption) (res *Worker, err error) {
	opts = slices.Concat(r.options, opts)
	if workerID == "" {
		err = errors.New("missing required workerId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/workers/%s", url.PathEscape(workerID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve the inbound email address for a worker.
func (r *WorkerService) GetEmail(ctx context.Context, workerID string, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.options, opts)
	if workerID == "" {
		err = errors.New("missing required workerId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/workers/%s/email", url.PathEscape(workerID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The property Input is required.
type CreateWorkerParam struct {
	Input  string          `json:"input" api:"required"`
	Stream param.Opt[bool] `json:"stream,omitzero"`
	// Any of "low", "standard", "high", "unlimited".
	Budget CreateWorkerBudget `json:"budget,omitzero"`
	paramObj
}

func (r CreateWorkerParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateWorkerParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateWorkerParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreateWorkerBudget string

const (
	CreateWorkerBudgetLow       CreateWorkerBudget = "low"
	CreateWorkerBudgetStandard  CreateWorkerBudget = "standard"
	CreateWorkerBudgetHigh      CreateWorkerBudget = "high"
	CreateWorkerBudgetUnlimited CreateWorkerBudget = "unlimited"
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
	Status WorkerStatus `json:"status" api:"required"`
	Usage  WorkerUsage  `json:"usage"`
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
	// messages and then continues with live chunks from the active turn (if any)
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
// messages and then continues with live chunks from the active turn (if any)
// before closing.
type WorkerGetParamsStream string

const (
	WorkerGetParamsStreamTrue  WorkerGetParamsStream = "true"
	WorkerGetParamsStreamFalse WorkerGetParamsStream = "false"
)

type WorkerContinueParams struct {
	CreateWorker CreateWorkerParam
	paramObj
}

func (r WorkerContinueParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateWorker)
}
func (r *WorkerContinueParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
