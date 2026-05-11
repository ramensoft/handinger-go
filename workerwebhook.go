// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package handinger

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/ramensoft/handinger-go/internal/apijson"
	"github.com/ramensoft/handinger-go/internal/apiquery"
	shimjson "github.com/ramensoft/handinger-go/internal/encoding/json"
	"github.com/ramensoft/handinger-go/internal/requestconfig"
	"github.com/ramensoft/handinger-go/option"
	"github.com/ramensoft/handinger-go/packages/param"
	"github.com/ramensoft/handinger-go/packages/respjson"
)

// Configure outbound webhooks delivered when a worker's tasks complete.
//
// WorkerWebhookService contains methods and other services that help with
// interacting with the handinger API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkerWebhookService] method instead.
type WorkerWebhookService struct {
	options []option.RequestOption
}

// NewWorkerWebhookService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkerWebhookService(opts ...option.RequestOption) (r WorkerWebhookService) {
	r = WorkerWebhookService{}
	r.options = opts
	return
}

// Retrieve the webhook URL and shared token configured for a worker. Both fields
// are `null` when no webhook is configured. Only the worker creator can read the
// webhook configuration.
func (r *WorkerWebhookService) Get(ctx context.Context, workerID string, opts ...option.RequestOption) (res *Webhook, err error) {
	opts = slices.Concat(r.options, opts)
	if workerID == "" {
		err = errors.New("missing required workerId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/workers/%s/webhook", url.PathEscape(workerID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Set or replace the webhook URL for a worker. A fresh token is generated the
// first time a URL is set; subsequent updates keep the existing token. Pass
// `url: null` to clear the webhook (use the dedicated DELETE for the same effect).
// Only the worker creator can update the webhook.
func (r *WorkerWebhookService) Update(ctx context.Context, workerID string, body WorkerWebhookUpdateParams, opts ...option.RequestOption) (res *Webhook, err error) {
	opts = slices.Concat(r.options, opts)
	if workerID == "" {
		err = errors.New("missing required workerId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/workers/%s/webhook", url.PathEscape(workerID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Remove the webhook from a worker. Both `url` and `token` are cleared and no
// further deliveries are attempted. Only the worker creator can delete the
// webhook.
func (r *WorkerWebhookService) Delete(ctx context.Context, workerID string, opts ...option.RequestOption) (res *Webhook, err error) {
	opts = slices.Concat(r.options, opts)
	if workerID == "" {
		err = errors.New("missing required workerId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/workers/%s/webhook", url.PathEscape(workerID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// List recent webhook delivery attempts for a worker, newest first, paginated 50
// per page. Only the worker creator can read execution history.
func (r *WorkerWebhookService) ListExecutions(ctx context.Context, workerID string, query WorkerWebhookListExecutionsParams, opts ...option.RequestOption) (res *WebhookExecutionList, err error) {
	opts = slices.Concat(r.options, opts)
	if workerID == "" {
		err = errors.New("missing required workerId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/workers/%s/webhook/executions", url.PathEscape(workerID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Issue a new shared token for the webhook, invalidating the previous one. The
// webhook URL is preserved. Only the worker creator can regenerate the token.
func (r *WorkerWebhookService) RegenerateToken(ctx context.Context, workerID string, opts ...option.RequestOption) (res *Webhook, err error) {
	opts = slices.Concat(r.options, opts)
	if workerID == "" {
		err = errors.New("missing required workerId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/workers/%s/webhook/regenerate-token", url.PathEscape(workerID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// The property URL is required.
type UpdateWebhookParam struct {
	// HTTPS endpoint Handinger should POST to when a task finishes. Pass `null` to
	// remove the webhook and clear its token.
	URL param.Opt[string] `json:"url,omitzero" api:"required" format:"uri"`
	paramObj
}

func (r UpdateWebhookParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateWebhookParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateWebhookParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Webhook struct {
	// Shared secret sent in the `X-Handinger-Token` header on each delivery. `null`
	// when no webhook is configured.
	Token string `json:"token" api:"required"`
	// HTTPS endpoint that receives webhook deliveries when a task completes. `null`
	// when no webhook is configured.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Webhook) RawJSON() string { return r.JSON.raw }
func (r *Webhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookExecution struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Wall-clock time spent on the delivery attempt.
	DurationMs int64 `json:"durationMs" api:"required"`
	// Failure reason when `requestStatus` is `error`.
	ErrorMessage string `json:"errorMessage" api:"required"`
	// `success` when the endpoint returned a 2xx response, `error` otherwise.
	//
	// Any of "success", "error".
	RequestStatus WebhookExecutionRequestStatus `json:"requestStatus" api:"required"`
	// HTTP status returned by the endpoint, when reachable.
	ResponseStatus int64 `json:"responseStatus" api:"required"`
	// Task that triggered the delivery, when available.
	TaskID string `json:"taskId" api:"required"`
	// Title of the originating task, when available.
	TaskTitle string `json:"taskTitle" api:"required"`
	// Endpoint Handinger attempted to deliver to.
	URL      string `json:"url" api:"required"`
	WorkerID string `json:"workerId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		DurationMs     respjson.Field
		ErrorMessage   respjson.Field
		RequestStatus  respjson.Field
		ResponseStatus respjson.Field
		TaskID         respjson.Field
		TaskTitle      respjson.Field
		URL            respjson.Field
		WorkerID       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookExecution) RawJSON() string { return r.JSON.raw }
func (r *WebhookExecution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `success` when the endpoint returned a 2xx response, `error` otherwise.
type WebhookExecutionRequestStatus string

const (
	WebhookExecutionRequestStatusSuccess WebhookExecutionRequestStatus = "success"
	WebhookExecutionRequestStatusError   WebhookExecutionRequestStatus = "error"
)

type WebhookExecutionList struct {
	Logs []WebhookExecution `json:"logs" api:"required"`
	// Current page number.
	Page int64 `json:"page" api:"required"`
	// Total number of pages available.
	PageCount int64 `json:"pageCount" api:"required"`
	// Total number of executions recorded.
	TotalCount int64 `json:"totalCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Logs        respjson.Field
		Page        respjson.Field
		PageCount   respjson.Field
		TotalCount  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookExecutionList) RawJSON() string { return r.JSON.raw }
func (r *WebhookExecutionList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerWebhookUpdateParams struct {
	UpdateWebhook UpdateWebhookParam
	paramObj
}

func (r WorkerWebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateWebhook)
}
func (r *WorkerWebhookUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerWebhookListExecutionsParams struct {
	// Page number (1-indexed). Defaults to 1.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkerWebhookListExecutionsParams]'s query parameters as
// `url.Values`.
func (r WorkerWebhookListExecutionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
