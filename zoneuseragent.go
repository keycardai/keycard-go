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

// ZoneUserAgentService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneUserAgentService] method instead.
type ZoneUserAgentService struct {
	Options []option.RequestOption
}

// NewZoneUserAgentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneUserAgentService(opts ...option.RequestOption) (r ZoneUserAgentService) {
	r = ZoneUserAgentService{}
	r.Options = opts
	return
}

// Returns details of a specific user agent by user agent ID
func (r *ZoneUserAgentService) Get(ctx context.Context, id string, query ZoneUserAgentGetParams, opts ...option.RequestOption) (res *UserAgent, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if query.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/user-agents/%s", url.PathEscape(query.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of user agents in the specified zone. User agents represent
// client software (browsers, desktop apps, CLI tools) registered via OAuth 2.0
// Dynamic Client Registration.
func (r *ZoneUserAgentService) List(ctx context.Context, zoneID string, query ZoneUserAgentListParams, opts ...option.RequestOption) (res *ZoneUserAgentListResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/user-agents", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// A User Agent represents a user agent (browser, desktop app, CLI tool) that can
// initiate user sessions via OAuth 2.0 Dynamic Client Registration.
type UserAgent struct {
	// Unique identifier of the user agent
	ID string `json:"id" api:"required"`
	// Entity creation timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// User agent identifier (serves as OAuth client_id). Format: ua:{sha256_hash}
	Identifier string `json:"identifier" api:"required"`
	// Human-readable name
	Name string `json:"name" api:"required"`
	// Organization that owns this user agent
	OrganizationID string `json:"organization_id" api:"required"`
	// URL-safe identifier, unique within the zone
	Slug string `json:"slug" api:"required"`
	// Entity update timestamp
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Zone this user agent belongs to
	ZoneID string `json:"zone_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		Identifier     respjson.Field
		Name           respjson.Field
		OrganizationID respjson.Field
		Slug           respjson.Field
		UpdatedAt      respjson.Field
		ZoneID         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserAgent) RawJSON() string { return r.JSON.raw }
func (r *UserAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneUserAgentListResponse struct {
	Items []UserAgent `json:"items" api:"required"`
	// Cursor-based pagination metadata
	Pagination ZoneUserAgentListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneUserAgentListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneUserAgentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination metadata
type ZoneUserAgentListResponsePagination struct {
	// An opaque cursor used for paginating through a list of results
	AfterCursor string `json:"after_cursor" api:"required"`
	// An opaque cursor used for paginating through a list of results
	BeforeCursor string `json:"before_cursor" api:"required"`
	// Total number of items matching the query. Only included when
	// expand[]=total_count is requested.
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
func (r ZoneUserAgentListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ZoneUserAgentListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneUserAgentGetParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	paramObj
}

type ZoneUserAgentListParams struct {
	// Cursor for forward pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of items to return
	Limit  param.Opt[int64]                   `query:"limit,omitzero" json:"-"`
	Expand ZoneUserAgentListParamsExpandUnion `query:"expand[],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneUserAgentListParams]'s query parameters as
// `url.Values`.
func (r ZoneUserAgentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZoneUserAgentListParamsExpandUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfZoneUserAgentListsExpandString)
	OfZoneUserAgentListsExpandString         param.Opt[string] `query:",omitzero,inline"`
	OfZoneUserAgentListsExpandArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

type ZoneUserAgentListParamsExpandString string

const (
	ZoneUserAgentListParamsExpandStringTotalCount ZoneUserAgentListParamsExpandString = "total_count"
)
