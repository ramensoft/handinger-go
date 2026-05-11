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

	"github.com/ramensoft/handinger-go/internal/apijson"
	"github.com/ramensoft/handinger-go/internal/requestconfig"
	"github.com/ramensoft/handinger-go/option"
	"github.com/ramensoft/handinger-go/packages/param"
	"github.com/ramensoft/handinger-go/packages/respjson"
	"github.com/ramensoft/handinger-go/shared/constant"
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
// [WorkerScheduleScheduled], [WorkerScheduleDelayed], [WorkerScheduleCron],
// [WorkerScheduleInterval].
//
// Use the [WorkerScheduleUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WorkerScheduleUnion struct {
	ID        string    `json:"id"`
	Budget    string    `json:"budget"`
	Input     string    `json:"input"`
	NextRunAt time.Time `json:"nextRunAt"`
	// Any of "scheduled", "delayed", "cron", "interval".
	Type string `json:"type"`
	// This field is from variant [WorkerScheduleDelayed].
	DelayInSeconds int64 `json:"delayInSeconds"`
	// This field is from variant [WorkerScheduleCron].
	Cron string `json:"cron"`
	// This field is from variant [WorkerScheduleInterval].
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

// anyWorkerSchedule is implemented by each variant of [WorkerScheduleUnion] to add
// type safety for the return type of [WorkerScheduleUnion.AsAny]
type anyWorkerSchedule interface {
	implWorkerScheduleUnion()
}

func (WorkerScheduleScheduled) implWorkerScheduleUnion() {}
func (WorkerScheduleDelayed) implWorkerScheduleUnion()   {}
func (WorkerScheduleCron) implWorkerScheduleUnion()      {}
func (WorkerScheduleInterval) implWorkerScheduleUnion()  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := WorkerScheduleUnion.AsAny().(type) {
//	case handinger.WorkerScheduleScheduled:
//	case handinger.WorkerScheduleDelayed:
//	case handinger.WorkerScheduleCron:
//	case handinger.WorkerScheduleInterval:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u WorkerScheduleUnion) AsAny() anyWorkerSchedule {
	switch u.Type {
	case "scheduled":
		return u.AsScheduled()
	case "delayed":
		return u.AsDelayed()
	case "cron":
		return u.AsCron()
	case "interval":
		return u.AsInterval()
	}
	return nil
}

func (u WorkerScheduleUnion) AsScheduled() (v WorkerScheduleScheduled) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkerScheduleUnion) AsDelayed() (v WorkerScheduleDelayed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkerScheduleUnion) AsCron() (v WorkerScheduleCron) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkerScheduleUnion) AsInterval() (v WorkerScheduleInterval) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WorkerScheduleUnion) RawJSON() string { return u.JSON.raw }

func (r *WorkerScheduleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScheduleScheduled struct {
	ID string `json:"id" api:"required"`
	// Any of "low", "standard", "high", "unlimited".
	Budget    string             `json:"budget" api:"required"`
	Input     string             `json:"input" api:"required"`
	NextRunAt time.Time          `json:"nextRunAt" api:"required" format:"date-time"`
	Type      constant.Scheduled `json:"type" default:"scheduled"`
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
func (r WorkerScheduleScheduled) RawJSON() string { return r.JSON.raw }
func (r *WorkerScheduleScheduled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScheduleDelayed struct {
	ID string `json:"id" api:"required"`
	// Any of "low", "standard", "high", "unlimited".
	Budget         string           `json:"budget" api:"required"`
	DelayInSeconds int64            `json:"delayInSeconds" api:"required"`
	Input          string           `json:"input" api:"required"`
	NextRunAt      time.Time        `json:"nextRunAt" api:"required" format:"date-time"`
	Type           constant.Delayed `json:"type" default:"delayed"`
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
func (r WorkerScheduleDelayed) RawJSON() string { return r.JSON.raw }
func (r *WorkerScheduleDelayed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScheduleCron struct {
	ID string `json:"id" api:"required"`
	// Any of "low", "standard", "high", "unlimited".
	Budget    string        `json:"budget" api:"required"`
	Cron      string        `json:"cron" api:"required"`
	Input     string        `json:"input" api:"required"`
	NextRunAt time.Time     `json:"nextRunAt" api:"required" format:"date-time"`
	Type      constant.Cron `json:"type" default:"cron"`
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
func (r WorkerScheduleCron) RawJSON() string { return r.JSON.raw }
func (r *WorkerScheduleCron) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScheduleInterval struct {
	ID string `json:"id" api:"required"`
	// Any of "low", "standard", "high", "unlimited".
	Budget          string            `json:"budget" api:"required"`
	Input           string            `json:"input" api:"required"`
	IntervalSeconds int64             `json:"intervalSeconds" api:"required"`
	NextRunAt       time.Time         `json:"nextRunAt" api:"required" format:"date-time"`
	Type            constant.Interval `json:"type" default:"interval"`
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
func (r WorkerScheduleInterval) RawJSON() string { return r.JSON.raw }
func (r *WorkerScheduleInterval) UnmarshalJSON(data []byte) error {
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
	OfScheduled *WorkerScheduleNewParamsWhenScheduled `json:",omitzero,inline"`
	OfDelayed   *WorkerScheduleNewParamsWhenDelayed   `json:",omitzero,inline"`
	OfCron      *WorkerScheduleNewParamsWhenCron      `json:",omitzero,inline"`
	OfInterval  *WorkerScheduleNewParamsWhenInterval  `json:",omitzero,inline"`
	paramUnion
}

func (u WorkerScheduleNewParamsWhenUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfScheduled, u.OfDelayed, u.OfCron, u.OfInterval)
}
func (u *WorkerScheduleNewParamsWhenUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[WorkerScheduleNewParamsWhenUnion](
		"type",
		apijson.Discriminator[WorkerScheduleNewParamsWhenScheduled]("scheduled"),
		apijson.Discriminator[WorkerScheduleNewParamsWhenDelayed]("delayed"),
		apijson.Discriminator[WorkerScheduleNewParamsWhenCron]("cron"),
		apijson.Discriminator[WorkerScheduleNewParamsWhenInterval]("interval"),
	)
}

// The properties Date, Type are required.
type WorkerScheduleNewParamsWhenScheduled struct {
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// This field can be elided, and will marshal its zero value as "scheduled".
	Type constant.Scheduled `json:"type" default:"scheduled"`
	paramObj
}

func (r WorkerScheduleNewParamsWhenScheduled) MarshalJSON() (data []byte, err error) {
	type shadow WorkerScheduleNewParamsWhenScheduled
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkerScheduleNewParamsWhenScheduled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DelayInSeconds, Type are required.
type WorkerScheduleNewParamsWhenDelayed struct {
	DelayInSeconds int64 `json:"delayInSeconds" api:"required"`
	// This field can be elided, and will marshal its zero value as "delayed".
	Type constant.Delayed `json:"type" default:"delayed"`
	paramObj
}

func (r WorkerScheduleNewParamsWhenDelayed) MarshalJSON() (data []byte, err error) {
	type shadow WorkerScheduleNewParamsWhenDelayed
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkerScheduleNewParamsWhenDelayed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Cron, Type are required.
type WorkerScheduleNewParamsWhenCron struct {
	Cron string `json:"cron" api:"required"`
	// This field can be elided, and will marshal its zero value as "cron".
	Type constant.Cron `json:"type" default:"cron"`
	paramObj
}

func (r WorkerScheduleNewParamsWhenCron) MarshalJSON() (data []byte, err error) {
	type shadow WorkerScheduleNewParamsWhenCron
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkerScheduleNewParamsWhenCron) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties IntervalSeconds, Type are required.
type WorkerScheduleNewParamsWhenInterval struct {
	IntervalSeconds int64 `json:"intervalSeconds" api:"required"`
	// This field can be elided, and will marshal its zero value as "interval".
	Type constant.Interval `json:"type" default:"interval"`
	paramObj
}

func (r WorkerScheduleNewParamsWhenInterval) MarshalJSON() (data []byte, err error) {
	type shadow WorkerScheduleNewParamsWhenInterval
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkerScheduleNewParamsWhenInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
