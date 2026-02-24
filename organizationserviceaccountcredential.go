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

// OrganizationServiceAccountCredentialService contains methods and other services
// that help with interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationServiceAccountCredentialService] method instead.
type OrganizationServiceAccountCredentialService struct {
	Options []option.RequestOption
}

// NewOrganizationServiceAccountCredentialService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewOrganizationServiceAccountCredentialService(opts ...option.RequestOption) (r OrganizationServiceAccountCredentialService) {
	r = OrganizationServiceAccountCredentialService{}
	r.Options = opts
	return
}

// Create a new credential for a service account
func (r *OrganizationServiceAccountCredentialService) New(ctx context.Context, serviceAccountID string, params OrganizationServiceAccountCredentialNewParams, opts ...option.RequestOption) (res *OrganizationServiceAccountCredentialNewResponse, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
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
	path := fmt.Sprintf("organizations/%s/service-accounts/%s/credentials", url.PathEscape(params.OrganizationID), url.PathEscape(serviceAccountID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get a specific service account credential
func (r *OrganizationServiceAccountCredentialService) Get(ctx context.Context, credentialID string, params OrganizationServiceAccountCredentialGetParams, opts ...option.RequestOption) (res *ServiceAccountCredential, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if params.OrganizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	if params.ServiceAccountID == "" {
		err = errors.New("missing required service_account_id parameter")
		return
	}
	if credentialID == "" {
		err = errors.New("missing required credential_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/service-accounts/%s/credentials/%s", url.PathEscape(params.OrganizationID), url.PathEscape(params.ServiceAccountID), url.PathEscape(credentialID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Update a service account credential
func (r *OrganizationServiceAccountCredentialService) Update(ctx context.Context, credentialID string, params OrganizationServiceAccountCredentialUpdateParams, opts ...option.RequestOption) (res *ServiceAccountCredential, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if params.OrganizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	if params.ServiceAccountID == "" {
		err = errors.New("missing required service_account_id parameter")
		return
	}
	if credentialID == "" {
		err = errors.New("missing required credential_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/service-accounts/%s/credentials/%s", url.PathEscape(params.OrganizationID), url.PathEscape(params.ServiceAccountID), url.PathEscape(credentialID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List credentials for a service account
func (r *OrganizationServiceAccountCredentialService) List(ctx context.Context, serviceAccountID string, params OrganizationServiceAccountCredentialListParams, opts ...option.RequestOption) (res *OrganizationServiceAccountCredentialListResponse, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
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
	path := fmt.Sprintf("organizations/%s/service-accounts/%s/credentials", url.PathEscape(params.OrganizationID), url.PathEscape(serviceAccountID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Delete a service account credential
func (r *OrganizationServiceAccountCredentialService) Delete(ctx context.Context, credentialID string, params OrganizationServiceAccountCredentialDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.OrganizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	if params.ServiceAccountID == "" {
		err = errors.New("missing required service_account_id parameter")
		return
	}
	if credentialID == "" {
		err = errors.New("missing required credential_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/service-accounts/%s/credentials/%s", url.PathEscape(params.OrganizationID), url.PathEscape(params.ServiceAccountID), url.PathEscape(credentialID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Service account credential (without secret)
type ServiceAccountCredential struct {
	// Identifier for API resources. A 26-char nanoid (URL/DNS safe).
	ID string `json:"id,required"`
	// The client ID for authentication
	ClientID string `json:"client_id,required"`
	// The time the entity was created in utc
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A name for the entity to be displayed in UI
	Name string `json:"name,required"`
	// Optional description of the credential
	Description string `json:"description"`
	// When the credential was last used
	LastUsedAt time.Time `json:"last_used_at" format:"date-time"`
	// Permissions granted to the authenticated principal for this resource. Only
	// populated when the 'expand[]=permissions' query parameter is provided. Keys are
	// resource types (e.g., "organizations"), values are objects mapping permission
	// names to boolean values indicating if the permission is granted.
	Permissions map[string]map[string]bool `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ClientID    respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Description respjson.Field
		LastUsedAt  respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServiceAccountCredential) RawJSON() string { return r.JSON.raw }
func (r *ServiceAccountCredential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Service account credential with plaintext secret (only returned on creation)
type OrganizationServiceAccountCredentialNewResponse struct {
	// Identifier for API resources. A 26-char nanoid (URL/DNS safe).
	ID string `json:"id,required"`
	// The client ID for authentication
	ClientID string `json:"client_id,required"`
	// The client secret
	ClientSecret string `json:"client_secret,required"`
	// The time the entity was created in utc
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A name for the entity to be displayed in UI
	Name string `json:"name,required"`
	// Optional description of the credential
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		ClientID     respjson.Field
		ClientSecret respjson.Field
		CreatedAt    respjson.Field
		Name         respjson.Field
		Description  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationServiceAccountCredentialNewResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationServiceAccountCredentialNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationServiceAccountCredentialListResponse struct {
	Items []ServiceAccountCredential `json:"items,required"`
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
func (r OrganizationServiceAccountCredentialListResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationServiceAccountCredentialListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationServiceAccountCredentialNewParams struct {
	// Organization ID or label identifier
	OrganizationID string `path:"organization_id,required" json:"-"`
	// Credential name
	Name string `json:"name,required"`
	// Optional description of the credential
	Description      param.Opt[string] `json:"description,omitzero"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r OrganizationServiceAccountCredentialNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationServiceAccountCredentialNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationServiceAccountCredentialNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationServiceAccountCredentialGetParams struct {
	// Organization ID or label identifier
	OrganizationID string `path:"organization_id,required" json:"-"`
	// Identifier for API resources. A 26-char nanoid (URL/DNS safe).
	ServiceAccountID string            `path:"service_account_id,required" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Fields to expand in the response. Currently supports "permissions" to include
	// the permissions field with the caller's permissions for the resource.
	//
	// Any of "permissions".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationServiceAccountCredentialGetParams]'s query
// parameters as `url.Values`.
func (r OrganizationServiceAccountCredentialGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationServiceAccountCredentialUpdateParams struct {
	// Organization ID or label identifier
	OrganizationID string `path:"organization_id,required" json:"-"`
	// Identifier for API resources. A 26-char nanoid (URL/DNS safe).
	ServiceAccountID string `path:"service_account_id,required" json:"-"`
	// Optional description of the credential
	Description param.Opt[string] `json:"description,omitzero"`
	// Credential name
	Name             param.Opt[string] `json:"name,omitzero"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r OrganizationServiceAccountCredentialUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationServiceAccountCredentialUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationServiceAccountCredentialUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationServiceAccountCredentialListParams struct {
	// Organization ID or label identifier
	OrganizationID string `path:"organization_id,required" json:"-"`
	// Cursor for forward pagination
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of credentials to return
	Limit            param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Fields to expand in the response. Currently supports "permissions" to include
	// the permissions field with the caller's permissions for the resource.
	//
	// Any of "permissions".
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationServiceAccountCredentialListParams]'s query
// parameters as `url.Values`.
func (r OrganizationServiceAccountCredentialListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationServiceAccountCredentialDeleteParams struct {
	// Organization ID or label identifier
	OrganizationID string `path:"organization_id,required" json:"-"`
	// Identifier for API resources. A 26-char nanoid (URL/DNS safe).
	ServiceAccountID string            `path:"service_account_id,required" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}
