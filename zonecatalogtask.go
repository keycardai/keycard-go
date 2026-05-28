// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycard

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/keycardai/keycard-go/internal/apijson"
	"github.com/keycardai/keycard-go/internal/requestconfig"
	"github.com/keycardai/keycard-go/option"
	"github.com/keycardai/keycard-go/packages/param"
	"github.com/keycardai/keycard-go/packages/respjson"
)

// ZoneCatalogTaskService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneCatalogTaskService] method instead.
type ZoneCatalogTaskService struct {
	Options []option.RequestOption
}

// NewZoneCatalogTaskService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneCatalogTaskService(opts ...option.RequestOption) (r ZoneCatalogTaskService) {
	r = ZoneCatalogTaskService{}
	r.Options = opts
	return
}

// Returns 200 with task details when pending, running, or failed. Returns 303
// redirect to the install when completed.
func (r *ZoneCatalogTaskService) Get(ctx context.Context, taskID string, params ZoneCatalogTaskGetParams, opts ...option.RequestOption) (res *Task, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("/")}, opts...)
	if params.ZoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/catalog_tasks/%s", url.PathEscape(params.ZoneID), url.PathEscape(taskID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type Task struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Any of "create", "delete".
	Operation TaskOperation `json:"operation" api:"required"`
	// Any of "pending", "running", "completed", "failed".
	Status         TaskStatus `json:"status" api:"required"`
	UpdatedAt      time.Time  `json:"updated_at" api:"required" format:"date-time"`
	ErrorMessage   string     `json:"error_message"`
	InstallID      string     `json:"install_id"`
	Links          []TaskLink `json:"links"`
	PackageID      string     `json:"package_id"`
	PackageSlug    string     `json:"package_slug"`
	PackageVersion int64      `json:"package_version"`
	// Informational warnings about the task outcome. For delete tasks, warns when
	// adopted entities (pre-existing resources not created by the catalog) will be
	// preserved rather than deleted.
	Warnings []TaskWarning `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		Operation      respjson.Field
		Status         respjson.Field
		UpdatedAt      respjson.Field
		ErrorMessage   respjson.Field
		InstallID      respjson.Field
		Links          respjson.Field
		PackageID      respjson.Field
		PackageSlug    respjson.Field
		PackageVersion respjson.Field
		Warnings       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Task) RawJSON() string { return r.JSON.raw }
func (r *Task) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A directed, typed relationship from one entity (the subject) to another (the
// target).
//
// Follows the structure of RFC 7033 JRD link objects, adapted for intra-graph
// entity references. The subject is the entity whose `links` array contains this
// link.
type TaskLink struct {
	// Target reference.
	//
	// Fragment URIs (`#name`) reference other entities in the same graph by their
	// local name (the key in the entity map). Absolute paths and URLs reference
	// external resources outside the graph.
	Href string `json:"href" api:"required"`
	// Link relation type.
	Rel string `json:"rel" api:"required"`
	// Additional metadata keyed by property name.
	Properties map[string]any `json:"properties"`
	// Human-readable titles keyed by BCP 47 language tag.
	Titles map[string]string `json:"titles"`
	// Media type of the target resource (per RFC 7033 section 4.4.4.3). Applies to
	// external `href`s; typically omitted for intra-graph references.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Href        respjson.Field
		Rel         respjson.Field
		Properties  respjson.Field
		Titles      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskLink) RawJSON() string { return r.JSON.raw }
func (r *TaskLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an error that has occurred in the Keycard system.
type TaskWarning struct {
	// Any of "validation_error", "bad_request", "unauthorized", "forbidden",
	// "not_found", "conflict", "rate_limit_exceeded", "internal_error",
	// "service_unavailable".
	Code    string              `json:"code" api:"required"`
	Details []TaskWarningDetail `json:"details" api:"required"`
	// summary of the error
	Message   string `json:"message" api:"required"`
	Path      string `json:"path" api:"required" format:"url"`
	RequestID string `json:"request_id" api:"required" format:"uuid"`
	// HTTP Status Code
	Status    int64     `json:"status" api:"required"`
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Details     respjson.Field
		Message     respjson.Field
		Path        respjson.Field
		RequestID   respjson.Field
		Status      respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskWarning) RawJSON() string { return r.JSON.raw }
func (r *TaskWarning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskWarningDetail struct {
	// Any of "validation_error", "bad_request", "unauthorized", "forbidden",
	// "not_found", "conflict", "rate_limit_exceeded", "internal_error",
	// "service_unavailable".
	Code string `json:"code" api:"required"`
	// valid json path for request body
	Field string `json:"field" api:"required"`
	// error message for specific error
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Field       respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskWarningDetail) RawJSON() string { return r.JSON.raw }
func (r *TaskWarningDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskOperation string

const (
	TaskOperationCreate TaskOperation = "create"
	TaskOperationDelete TaskOperation = "delete"
)

type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusRunning   TaskStatus = "running"
	TaskStatusCompleted TaskStatus = "completed"
	TaskStatusFailed    TaskStatus = "failed"
)

type ZoneCatalogTaskGetParams struct {
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}
