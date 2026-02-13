// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycardapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/keycardlabs/keycard-go/internal/apijson"
	"github.com/keycardlabs/keycard-go/internal/apiquery"
	"github.com/keycardlabs/keycard-go/internal/requestconfig"
	"github.com/keycardlabs/keycard-go/option"
	"github.com/keycardlabs/keycard-go/packages/param"
	"github.com/keycardlabs/keycard-go/packages/respjson"
)

// ZoneApplicationDependencyService contains methods and other services that help
// with interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneApplicationDependencyService] method instead.
type ZoneApplicationDependencyService struct {
	Options []option.RequestOption
}

// NewZoneApplicationDependencyService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneApplicationDependencyService(opts ...option.RequestOption) (r ZoneApplicationDependencyService) {
	r = ZoneApplicationDependencyService{}
	r.Options = opts
	return
}

// Retrieves a specific resource dependency of an application
func (r *ZoneApplicationDependencyService) Get(ctx context.Context, dependencyID string, query ZoneApplicationDependencyGetParams, opts ...option.RequestOption) (res *Resource, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if query.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	if query.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if dependencyID == "" {
		err = errors.New("missing required dependencyId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/applications/%s/dependencies/%s", url.PathEscape(query.ZoneID), url.PathEscape(query.ID), url.PathEscape(dependencyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns resource dependencies for an application
func (r *ZoneApplicationDependencyService) List(ctx context.Context, id string, params ZoneApplicationDependencyListParams, opts ...option.RequestOption) (res *ZoneApplicationDependencyListResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/applications/%s/dependencies", url.PathEscape(params.ZoneID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Adds a resource dependency to an application
func (r *ZoneApplicationDependencyService) Add(ctx context.Context, dependencyID string, params ZoneApplicationDependencyAddParams, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if dependencyID == "" {
		err = errors.New("missing required dependencyId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/applications/%s/dependencies/%s", url.PathEscape(params.ZoneID), url.PathEscape(params.ID), url.PathEscape(dependencyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return
}

// Removes a resource dependency from an application
func (r *ZoneApplicationDependencyService) Remove(ctx context.Context, dependencyID string, body ZoneApplicationDependencyRemoveParams, opts ...option.RequestOption) (err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if dependencyID == "" {
		err = errors.New("missing required dependencyId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/applications/%s/dependencies/%s", url.PathEscape(body.ZoneID), url.PathEscape(body.ID), url.PathEscape(dependencyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// A Resource is a system that exposes protected information or functionality. It
// requires authentication of the requesting actor, which may be a user or
// application, before allowing access.
type Resource struct {
	// Unique identifier of the resource
	ID string `json:"id,required"`
	// Entity creation timestamp
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// User specified identifier, unique within the zone
	Identifier string `json:"identifier,required"`
	// Human-readable name
	Name string `json:"name,required"`
	// Organization that owns this resource
	OrganizationID string `json:"organization_id,required"`
	// URL-safe identifier, unique within the zone
	Slug string `json:"slug,required"`
	// Entity update timestamp
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Zone this resource belongs to
	ZoneID string `json:"zone_id,required"`
	// An Application is a software system with an associated identity that can access
	// Resources. It may act on its own behalf (machine-to-machine) or on behalf of a
	// user (delegated access).
	//
	// Deprecated: deprecated
	Application Application `json:"application"`
	// ID of the application that provides this resource
	ApplicationID string `json:"application_id"`
	// A Provider is a system that supplies access to Resources and allows actors
	// (Users or Applications) to authenticate.
	//
	// Deprecated: deprecated
	CredentialProvider Provider `json:"credential_provider"`
	// ID of the credential provider for this resource
	CredentialProviderID string `json:"credential_provider_id"`
	// Human-readable description
	Description string `json:"description,nullable"`
	// Entity metadata
	Metadata Metadata `json:"metadata"`
	// Scopes supported by the resource
	Scopes []string `json:"scopes,nullable"`
	// List of resource IDs that, when accessed, make this dependency available. Only
	// present when this resource is returned as a dependency.
	WhenAccessing []string `json:"when_accessing"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		CreatedAt            respjson.Field
		Identifier           respjson.Field
		Name                 respjson.Field
		OrganizationID       respjson.Field
		Slug                 respjson.Field
		UpdatedAt            respjson.Field
		ZoneID               respjson.Field
		Application          respjson.Field
		ApplicationID        respjson.Field
		CredentialProvider   respjson.Field
		CredentialProviderID respjson.Field
		Description          respjson.Field
		Metadata             respjson.Field
		Scopes               respjson.Field
		WhenAccessing        respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Resource) RawJSON() string { return r.JSON.raw }
func (r *Resource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneApplicationDependencyListResponse struct {
	Items []Resource `json:"items,required"`
	// Pagination information
	PageInfo PageInfoPagination `json:"page_info,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		PageInfo    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneApplicationDependencyListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneApplicationDependencyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneApplicationDependencyGetParams struct {
	ZoneID string `path:"zoneId,required" json:"-"`
	ID     string `path:"id,required" json:"-"`
	paramObj
}

type ZoneApplicationDependencyListParams struct {
	ZoneID        string            `path:"zoneId,required" json:"-"`
	Cursor        param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit         param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	WhenAccessing param.Opt[string] `query:"when_accessing,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneApplicationDependencyListParams]'s query parameters as
// `url.Values`.
func (r ZoneApplicationDependencyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneApplicationDependencyAddParams struct {
	ZoneID        string   `path:"zoneId,required" json:"-"`
	ID            string   `path:"id,required" json:"-"`
	WhenAccessing []string `query:"when_accessing,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneApplicationDependencyAddParams]'s query parameters as
// `url.Values`.
func (r ZoneApplicationDependencyAddParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneApplicationDependencyRemoveParams struct {
	ZoneID string `path:"zoneId,required" json:"-"`
	ID     string `path:"id,required" json:"-"`
	paramObj
}
