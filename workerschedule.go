// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package handinger

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Ramensoft/handinger-go/internal/apijson"
	"github.com/Ramensoft/handinger-go/internal/requestconfig"
	"github.com/Ramensoft/handinger-go/option"
	"github.com/Ramensoft/handinger-go/packages/param"
	"github.com/Ramensoft/handinger-go/packages/respjson"
)

// Manage future and recurring worker tasks.
//
// WorkerScheduleService contains methods and other services that help with
// interacting with the handinger API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkerScheduleService] method instead.
type WorkerScheduleService struct {
	options []option.RequestOption
}

// NewWorkerScheduleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkerScheduleService(opts ...option.RequestOption) (r WorkerScheduleService) {
	r = WorkerScheduleService{}
	r.options = opts
	return
}

// Schedule a worker instruction for a future or recurring run.
func (r *WorkerScheduleService) New(ctx context.Context, workerID string, body WorkerScheduleNewParams, opts ...option.RequestOption) (res *WorkerScheduleUnion, err error) {
	opts = slices.Concat(r.options, opts)
	if workerID == "" {
		err = errors.New("missing required workerId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/workers/%s/schedules", url.PathEscape(workerID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List scheduled tasks for a worker.
func (r *WorkerScheduleService) List(ctx context.Context, workerID string, opts ...option.RequestOption) (res *WorkerScheduleListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if workerID == "" {
		err = errors.New("missing required workerId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/workers/%s/schedules", url.PathEscape(workerID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Cancel a scheduled task for a worker.
func (r *WorkerScheduleService) Cancel(ctx context.Context, scheduleID string, body WorkerScheduleCancelParams, opts ...option.RequestOption) (res *WorkerScheduleCancelResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.WorkerID == "" {
		err = errors.New("missing required workerId parameter")
		return nil, err
	}
	if scheduleID == "" {
		err = errors.New("missing required scheduleId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/workers/%s/schedules/%s", url.PathEscape(body.WorkerID), url.PathEscape(scheduleID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// WorkerScheduleUnion contains all possible properties and values from
// [WorkerScheduleObject], [WorkerScheduleObject2], [WorkerScheduleObject3],
// [WorkerScheduleObject4].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WorkerScheduleUnion struct {
	ID        string    `json:"id"`
	Budget    string    `json:"budget"`
	Input     string    `json:"input"`
	NextRunAt time.Time `json:"nextRunAt"`
	Type      string    `json:"type"`
	// This field is from variant [WorkerScheduleObject2].
	DelayInSeconds int64 `json:"delayInSeconds"`
	// This field is from variant [WorkerScheduleObject3].
	Cron string `json:"cron"`
	// This field is from variant [WorkerScheduleObject4].
	IntervalSeconds int64 `json:"intervalSeconds"`
	JSON            struct {
		ID              respjson.Field
		Budget          respjson.Field
		Input           respjson.Field
		NextRunAt       respjson.Field
		Type            respjson.Field
		DelayInSeconds  respjson.Field
		Cron            respjson.Field
		IntervalSeconds respjson.Field
		raw             string
	} `json:"-"`
}

func (u WorkerScheduleUnion) AsWorkerScheduleObject() (v WorkerScheduleObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkerScheduleUnion) AsWorkerScheduleObject2() (v WorkerScheduleObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkerScheduleUnion) AsWorkerScheduleObject3() (v WorkerScheduleObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkerScheduleUnion) AsWorkerScheduleObject4() (v WorkerScheduleObject4) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WorkerScheduleUnion) RawJSON() string { return u.JSON.raw }

func (r *WorkerScheduleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScheduleObject struct {
	ID string `json:"id" api:"required"`
	// Any of "low", "standard", "high", "unlimited".
	Budget    string    `json:"budget" api:"required"`
	Input     string    `json:"input" api:"required"`
	NextRunAt time.Time `json:"nextRunAt" api:"required" format:"date-time"`
	// Any of "scheduled".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Budget      respjson.Field
		Input       respjson.Field
		NextRunAt   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerScheduleObject) RawJSON() string { return r.JSON.raw }
func (r *WorkerScheduleObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScheduleObject2 struct {
	ID string `json:"id" api:"required"`
	// Any of "low", "standard", "high", "unlimited".
	Budget         string    `json:"budget" api:"required"`
	DelayInSeconds int64     `json:"delayInSeconds" api:"required"`
	Input          string    `json:"input" api:"required"`
	NextRunAt      time.Time `json:"nextRunAt" api:"required" format:"date-time"`
	// Any of "delayed".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Budget         respjson.Field
		DelayInSeconds respjson.Field
		Input          respjson.Field
		NextRunAt      respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerScheduleObject2) RawJSON() string { return r.JSON.raw }
func (r *WorkerScheduleObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScheduleObject3 struct {
	ID string `json:"id" api:"required"`
	// Any of "low", "standard", "high", "unlimited".
	Budget    string    `json:"budget" api:"required"`
	Cron      string    `json:"cron" api:"required"`
	Input     string    `json:"input" api:"required"`
	NextRunAt time.Time `json:"nextRunAt" api:"required" format:"date-time"`
	// Any of "cron".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Budget      respjson.Field
		Cron        respjson.Field
		Input       respjson.Field
		NextRunAt   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerScheduleObject3) RawJSON() string { return r.JSON.raw }
func (r *WorkerScheduleObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScheduleObject4 struct {
	ID string `json:"id" api:"required"`
	// Any of "low", "standard", "high", "unlimited".
	Budget          string    `json:"budget" api:"required"`
	Input           string    `json:"input" api:"required"`
	IntervalSeconds int64     `json:"intervalSeconds" api:"required"`
	NextRunAt       time.Time `json:"nextRunAt" api:"required" format:"date-time"`
	// Any of "interval".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Budget          respjson.Field
		Input           respjson.Field
		IntervalSeconds respjson.Field
		NextRunAt       respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerScheduleObject4) RawJSON() string { return r.JSON.raw }
func (r *WorkerScheduleObject4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScheduleListResponse struct {
	Schedules []WorkerScheduleUnion `json:"schedules" api:"required"`
	WorkerID  string                `json:"workerId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Schedules   respjson.Field
		WorkerID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerScheduleListResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkerScheduleListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScheduleCancelResponse struct {
	Cancelled bool `json:"cancelled" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cancelled   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkerScheduleCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkerScheduleCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScheduleNewParams struct {
	Input string                           `json:"input" api:"required"`
	When  WorkerScheduleNewParamsWhenUnion `json:"when,omitzero" api:"required"`
	// Any of "low", "standard", "high", "unlimited".
	Budget WorkerScheduleNewParamsBudget `json:"budget,omitzero"`
	paramObj
}

func (r WorkerScheduleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkerScheduleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkerScheduleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkerScheduleNewParamsWhenUnion struct {
	OfWorkerScheduleNewsWhenObject  *WorkerScheduleNewParamsWhenObject  `json:",omitzero,inline"`
	OfWorkerScheduleNewsWhenObject2 *WorkerScheduleNewParamsWhenObject2 `json:",omitzero,inline"`
	OfWorkerScheduleNewsWhenObject3 *WorkerScheduleNewParamsWhenObject3 `json:",omitzero,inline"`
	OfWorkerScheduleNewsWhenObject4 *WorkerScheduleNewParamsWhenObject4 `json:",omitzero,inline"`
	paramUnion
}

func (u WorkerScheduleNewParamsWhenUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWorkerScheduleNewsWhenObject, u.OfWorkerScheduleNewsWhenObject2, u.OfWorkerScheduleNewsWhenObject3, u.OfWorkerScheduleNewsWhenObject4)
}
func (u *WorkerScheduleNewParamsWhenUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Date, Type are required.
type WorkerScheduleNewParamsWhenObject struct {
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// Any of "scheduled".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r WorkerScheduleNewParamsWhenObject) MarshalJSON() (data []byte, err error) {
	type shadow WorkerScheduleNewParamsWhenObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkerScheduleNewParamsWhenObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkerScheduleNewParamsWhenObject](
		"type", "scheduled",
	)
}

// The properties DelayInSeconds, Type are required.
type WorkerScheduleNewParamsWhenObject2 struct {
	DelayInSeconds int64 `json:"delayInSeconds" api:"required"`
	// Any of "delayed".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r WorkerScheduleNewParamsWhenObject2) MarshalJSON() (data []byte, err error) {
	type shadow WorkerScheduleNewParamsWhenObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkerScheduleNewParamsWhenObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkerScheduleNewParamsWhenObject2](
		"type", "delayed",
	)
}

// The properties Cron, Type are required.
type WorkerScheduleNewParamsWhenObject3 struct {
	Cron string `json:"cron" api:"required"`
	// Any of "cron".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r WorkerScheduleNewParamsWhenObject3) MarshalJSON() (data []byte, err error) {
	type shadow WorkerScheduleNewParamsWhenObject3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkerScheduleNewParamsWhenObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkerScheduleNewParamsWhenObject3](
		"type", "cron",
	)
}

// The properties IntervalSeconds, Type are required.
type WorkerScheduleNewParamsWhenObject4 struct {
	IntervalSeconds int64 `json:"intervalSeconds" api:"required"`
	// Any of "interval".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r WorkerScheduleNewParamsWhenObject4) MarshalJSON() (data []byte, err error) {
	type shadow WorkerScheduleNewParamsWhenObject4
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkerScheduleNewParamsWhenObject4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkerScheduleNewParamsWhenObject4](
		"type", "interval",
	)
}

type WorkerScheduleNewParamsBudget string

const (
	WorkerScheduleNewParamsBudgetLow       WorkerScheduleNewParamsBudget = "low"
	WorkerScheduleNewParamsBudgetStandard  WorkerScheduleNewParamsBudget = "standard"
	WorkerScheduleNewParamsBudgetHigh      WorkerScheduleNewParamsBudget = "high"
	WorkerScheduleNewParamsBudgetUnlimited WorkerScheduleNewParamsBudget = "unlimited"
)

type WorkerScheduleCancelParams struct {
	WorkerID string `path:"workerId" api:"required" json:"-"`
	paramObj
}
