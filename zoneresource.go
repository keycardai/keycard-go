// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycardapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/keycard-api-go/internal/apijson"
	"github.com/stainless-sdks/keycard-api-go/internal/apiquery"
	"github.com/stainless-sdks/keycard-api-go/internal/requestconfig"
	"github.com/stainless-sdks/keycard-api-go/option"
	"github.com/stainless-sdks/keycard-api-go/packages/param"
	"github.com/stainless-sdks/keycard-api-go/packages/respjson"
)

// ZoneResourceService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneResourceService] method instead.
type ZoneResourceService struct {
	Options []option.RequestOption
}

// NewZoneResourceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneResourceService(opts ...option.RequestOption) (r ZoneResourceService) {
	r = ZoneResourceService{}
	r.Options = opts
	return
}

// Creates a new Resource - a system that exposes protected information or
// functionality requiring authentication
func (r *ZoneResourceService) New(ctx context.Context, zoneID string, body ZoneResourceNewParams, opts ...option.RequestOption) (res *Resource, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/resources", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns details of a specific Resource by ID
func (r *ZoneResourceService) Get(ctx context.Context, id string, query ZoneResourceGetParams, opts ...option.RequestOption) (res *Resource, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/resources/%s", query.ZoneID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a Resource's configuration and metadata
func (r *ZoneResourceService) Update(ctx context.Context, id string, params ZoneResourceUpdateParams, opts ...option.RequestOption) (res *Resource, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/resources/%s", params.ZoneID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Returns a list of resources in the specified zone
func (r *ZoneResourceService) List(ctx context.Context, zoneID string, query ZoneResourceListParams, opts ...option.RequestOption) (res *ZoneResourceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/resources", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Permanently deletes a Resource
func (r *ZoneResourceService) Delete(ctx context.Context, id string, body ZoneResourceDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/resources/%s", body.ZoneID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type ZoneResourceListResponse struct {
	Items []Resource `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneResourceListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneResourceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneResourceNewParams struct {
	// User specified identifier, unique within the zone
	Identifier string `json:"identifier,required"`
	// Human-readable name
	Name string `json:"name,required"`
	// Human-readable description
	Description param.Opt[string] `json:"description,omitzero"`
	// ID of the application that provides this resource
	ApplicationID param.Opt[string] `json:"application_id,omitzero"`
	// ID of the credential provider to associate with the resource
	CredentialProviderID param.Opt[string] `json:"credential_provider_id,omitzero"`
	// Entity metadata
	Metadata MetadataParam `json:"metadata,omitzero"`
	// Scopes supported by the resource
	Scopes []string `json:"scopes,omitzero"`
	paramObj
}

func (r ZoneResourceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ZoneResourceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneResourceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneResourceGetParams struct {
	ZoneID string `path:"zoneId,required" json:"-"`
	paramObj
}

type ZoneResourceUpdateParams struct {
	ZoneID string `path:"zoneId,required" json:"-"`
	// ID of the application that provides this resource (set to null to unset)
	ApplicationID param.Opt[string] `json:"application_id,omitzero"`
	// ID of the credential provider to associate with the resource (set to null to
	// unset)
	CredentialProviderID param.Opt[string] `json:"credential_provider_id,omitzero"`
	// Human-readable description
	Description param.Opt[string] `json:"description,omitzero"`
	// User specified identifier, unique within the zone
	Identifier param.Opt[string] `json:"identifier,omitzero"`
	// Human-readable name
	Name param.Opt[string] `json:"name,omitzero"`
	// Entity metadata (set to null or {} to remove metadata)
	Metadata MetadataUpdateParam `json:"metadata,omitzero"`
	// Scopes supported by the resource (set to null to unset)
	Scopes []string `json:"scopes,omitzero"`
	paramObj
}

func (r ZoneResourceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ZoneResourceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneResourceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneResourceListParams struct {
	// Filter resources by credential provider ID
	CredentialProviderID param.Opt[string] `query:"credentialProviderId,omitzero" json:"-"`
	// Filter resources by identifier
	Identifier param.Opt[string] `query:"identifier,omitzero" json:"-"`
	Slug       param.Opt[string] `query:"slug,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZoneResourceListParams]'s query parameters as `url.Values`.
func (r ZoneResourceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneResourceDeleteParams struct {
	ZoneID string `path:"zoneId,required" json:"-"`
	paramObj
}
