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

// OrganizationServiceAccountService contains methods and other services that help
// with interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationServiceAccountService] method instead.
type OrganizationServiceAccountService struct {
	Options     []option.RequestOption
	Credentials OrganizationServiceAccountCredentialService
}

// NewOrganizationServiceAccountService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOrganizationServiceAccountService(opts ...option.RequestOption) (r OrganizationServiceAccountService) {
	r = OrganizationServiceAccountService{}
	r.Options = opts
	r.Credentials = NewOrganizationServiceAccountCredentialService(opts...)
	return
}

// Create a new service account for an organization
func (r *OrganizationServiceAccountService) New(ctx context.Context, organizationID string, params OrganizationServiceAccountNewParams, opts ...option.RequestOption) (res *ServiceAccount, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%s", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/service-accounts", url.PathEscape(organizationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get a specific service account
func (r *OrganizationServiceAccountService) Get(ctx context.Context, serviceAccountID string, params OrganizationServiceAccountGetParams, opts ...option.RequestOption) (res *ServiceAccount, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%s", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if params.OrganizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	if serviceAccountID == "" {
		err = errors.New("missing required service_account_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/service-accounts/%s", url.PathEscape(params.OrganizationID), url.PathEscape(serviceAccountID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Update a service account
func (r *OrganizationServiceAccountService) Update(ctx context.Context, serviceAccountID string, params OrganizationServiceAccountUpdateParams, opts ...option.RequestOption) (res *ServiceAccount, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%s", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if params.OrganizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	if serviceAccountID == "" {
		err = errors.New("missing required service_account_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/service-accounts/%s", url.PathEscape(params.OrganizationID), url.PathEscape(serviceAccountID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List service accounts for an organization
func (r *OrganizationServiceAccountService) List(ctx context.Context, organizationID string, params OrganizationServiceAccountListParams, opts ...option.RequestOption) (res *OrganizationServiceAccountListResponse, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%s", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/service-accounts", url.PathEscape(organizationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Delete a service account
func (r *OrganizationServiceAccountService) Delete(ctx context.Context, serviceAccountID string, params OrganizationServiceAccountDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%s", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.OrganizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	if serviceAccountID == "" {
		err = errors.New("missing required service_account_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/service-accounts/%s", url.PathEscape(params.OrganizationID), url.PathEscape(serviceAccountID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type ServiceAccount struct {
	// Identifier for API resources. A 26-char nanoid (URL/DNS safe).
	ID string `json:"id,required"`
	// The time the entity was created in utc
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A name for the entity to be displayed in UI
	Name string `json:"name,required"`
	// The time the entity was mostly recently updated in utc
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Optional description of the service account
	Description string `json:"description"`
	// Permissions granted to the authenticated principal for this resource. Only
	// populated when the 'expand[]=permissions' query parameter is provided. Keys are
	// resource types (e.g., "organizations"), values are objects mapping permission
	// names to boolean values indicating if the permission is granted.
	Permissions map[string]map[string]bool `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServiceAccount) RawJSON() string { return r.JSON.raw }
func (r *ServiceAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationServiceAccountListResponse struct {
	Items []ServiceAccount `json:"items,required"`
	// Pagination information using cursor-based pagination
	PageInfo PageInfoCursor `json:"page_info,required"`
	// Permissions granted to the authenticated principal for this resource. Only
	// populated when the 'expand[]=permissions' query parameter is provided. Keys are
	// resource types (e.g., "organizations"), values are objects mapping permission
	// names to boolean values indicating if the permission is granted.
	Permissions map[string]map[string]bool `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		PageInfo    respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationServiceAccountListResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationServiceAccountListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationServiceAccountNewParams struct {
	// Service account name
	Name string `json:"name,required"`
	// Optional description of the service account
	Description      param.Opt[string] `json:"description,omitzero"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r OrganizationServiceAccountNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationServiceAccountNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationServiceAccountNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationServiceAccountGetParams struct {
	// Organization ID or label identifier
	OrganizationID   string            `path:"organization_id,required" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Fields to expand in the response. Currently supports "permissions" to include
	// the permissions field with the caller's permissions for the resource.
	//
	// Any of "permissions".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationServiceAccountGetParams]'s query parameters as
// `url.Values`.
func (r OrganizationServiceAccountGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationServiceAccountUpdateParams struct {
	// Organization ID or label identifier
	OrganizationID string `path:"organization_id,required" json:"-"`
	// Optional description of the service account
	Description param.Opt[string] `json:"description,omitzero"`
	// Service account name
	Name             param.Opt[string] `json:"name,omitzero"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r OrganizationServiceAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationServiceAccountUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationServiceAccountUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationServiceAccountListParams struct {
	// Cursor for forward pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of service accounts to return
	Limit            param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Fields to expand in the response. Currently supports "permissions" to include
	// the permissions field with the caller's permissions for the resource.
	//
	// Any of "permissions".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationServiceAccountListParams]'s query parameters as
// `url.Values`.
func (r OrganizationServiceAccountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationServiceAccountDeleteParams struct {
	// Organization ID or label identifier
	OrganizationID   string            `path:"organization_id,required" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}
