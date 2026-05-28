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
	"github.com/keycardai/keycard-go/internal/apiquery"
	"github.com/keycardai/keycard-go/internal/requestconfig"
	"github.com/keycardai/keycard-go/option"
	"github.com/keycardai/keycard-go/packages/param"
	"github.com/keycardai/keycard-go/packages/respjson"
)

// Install packages and manage package installations.
//
// ZoneInstallService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneInstallService] method instead.
type ZoneInstallService struct {
	Options []option.RequestOption
}

// NewZoneInstallService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneInstallService(opts ...option.RequestOption) (r ZoneInstallService) {
	r = ZoneInstallService{}
	r.Options = opts
	return
}

// Create an install of a package
func (r *ZoneInstallService) New(ctx context.Context, zoneID string, params ZoneInstallNewParams, opts ...option.RequestOption) (res *Task, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("/")}, opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/installs", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get a specific zone install
func (r *ZoneInstallService) Get(ctx context.Context, installID string, params ZoneInstallGetParams, opts ...option.RequestOption) (res *Install, err error) {
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
	if installID == "" {
		err = errors.New("missing required install_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/installs/%s", url.PathEscape(params.ZoneID), url.PathEscape(installID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List installs in a zone
func (r *ZoneInstallService) List(ctx context.Context, zoneID string, params ZoneInstallListParams, opts ...option.RequestOption) (res *InstallList, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("/")}, opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/installs", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Delete a zone install
func (r *ZoneInstallService) Delete(ctx context.Context, installID string, params ZoneInstallDeleteParams, opts ...option.RequestOption) (res *Task, err error) {
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
	if installID == "" {
		err = errors.New("missing required install_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/installs/%s", url.PathEscape(params.ZoneID), url.PathEscape(installID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type Install struct {
	ID          string    `json:"id" api:"required"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	PackageID   string    `json:"package_id" api:"required"`
	PackageSlug string    `json:"package_slug" api:"required"`
	// Any of "pending", "active", "deleting", "failed", "deleted".
	Status    InstallStatus `json:"status" api:"required"`
	UpdatedAt time.Time     `json:"updated_at" api:"required" format:"date-time"`
	// Install-specific input values that supplement the package's inputs. Merged with
	// the package's input values to form the complete `entities.inputs` for entity
	// binding evaluation.
	Inputs map[string]any `json:"inputs"`
	Links  []InstallLink  `json:"links"`
	OrgID  string         `json:"org_id"`
	// Resolved output values produced by the provisioner, conforming to the package's
	// `Package.outputs.schema`. Flat — the provisioner evaluates
	// `Package.outputs.bindings` against the resolved entity graph.
	Outputs        map[string]any `json:"outputs"`
	PackageVersion int64          `json:"package_version"`
	ZoneID         string         `json:"zone_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		PackageID      respjson.Field
		PackageSlug    respjson.Field
		Status         respjson.Field
		UpdatedAt      respjson.Field
		Inputs         respjson.Field
		Links          respjson.Field
		OrgID          respjson.Field
		Outputs        respjson.Field
		PackageVersion respjson.Field
		ZoneID         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Install) RawJSON() string { return r.JSON.raw }
func (r *Install) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A directed, typed relationship from one entity (the subject) to another (the
// target).
//
// Follows the structure of RFC 7033 JRD link objects, adapted for intra-graph
// entity references. The subject is the entity whose `links` array contains this
// link.
type InstallLink struct {
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
func (r InstallLink) RawJSON() string { return r.JSON.raw }
func (r *InstallLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstallList struct {
	Items []Install `json:"items" api:"required"`
	// Cursor-based pagination metadata returned alongside a list of results
	Pagination InstallListPagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstallList) RawJSON() string { return r.JSON.raw }
func (r *InstallList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination metadata returned alongside a list of results
type InstallListPagination struct {
	// An opaque cursor used for paginating through a list of results
	AfterCursor string `json:"after_cursor" api:"required"`
	// An opaque cursor used for paginating through a list of results
	BeforeCursor string `json:"before_cursor" api:"required"`
	// Total number of items across all pages. Only present when the request includes
	// ?expand[]=total_count.
	TotalCount int64 `json:"total_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AfterCursor  respjson.Field
		BeforeCursor respjson.Field
		TotalCount   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstallListPagination) RawJSON() string { return r.JSON.raw }
func (r *InstallListPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstallStatus string

const (
	InstallStatusPending  InstallStatus = "pending"
	InstallStatusActive   InstallStatus = "active"
	InstallStatusDeleting InstallStatus = "deleting"
	InstallStatusFailed   InstallStatus = "failed"
	InstallStatusDeleted  InstallStatus = "deleted"
)

type ZoneInstallNewParams struct {
	// Public ID of the package to install.
	PackageID string `json:"package_id" api:"required"`
	// Specific package version to install. Defaults to latest.
	Version          param.Opt[int64]  `json:"version,omitzero"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Parametric inputs required by the package.
	Inputs map[string]any `json:"inputs,omitzero"`
	paramObj
}

func (r ZoneInstallNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ZoneInstallNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneInstallNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneInstallGetParams struct {
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ZoneInstallListParams struct {
	// Cursor for forward pagination. Returned in `Pagination.after_cursor`. Mutually
	// exclusive with `before`.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination. Returned in `Pagination.before_cursor`. Mutually
	// exclusive with `after`.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of items to return per page.
	Limit            param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneInstallListParams]'s query parameters as `url.Values`.
func (r ZoneInstallListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneInstallDeleteParams struct {
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}
