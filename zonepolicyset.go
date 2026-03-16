// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycard

import (
	"context"
	"encoding/json"
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

// Policy set CRUD and binding management
//
// ZonePolicySetService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZonePolicySetService] method instead.
type ZonePolicySetService struct {
	Options []option.RequestOption
	// Immutable policy set manifest snapshots
	Versions ZonePolicySetVersionService
}

// NewZonePolicySetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZonePolicySetService(opts ...option.RequestOption) (r ZonePolicySetService) {
	r = ZonePolicySetService{}
	r.Options = opts
	r.Versions = NewZonePolicySetVersionService(opts...)
	return
}

// Creates an unbound policy set. Use updatePolicySet to bind after creating a
// version.
func (r *ZonePolicySetService) New(ctx context.Context, zoneID string, params ZonePolicySetNewParams, opts ...option.RequestOption) (res *PolicySetWithBinding, err error) {
	if !param.IsOmitted(params.XAPIVersion) {
		opts = append(opts, option.WithHeader("X-API-Version", fmt.Sprintf("%v", params.XAPIVersion.Value)))
	}
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policy-sets", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns the policy set with current binding information.
func (r *ZonePolicySetService) Get(ctx context.Context, policySetID string, params ZonePolicySetGetParams, opts ...option.RequestOption) (res *PolicySetWithBinding, err error) {
	if !param.IsOmitted(params.XAPIVersion) {
		opts = append(opts, option.WithHeader("X-API-Version", fmt.Sprintf("%v", params.XAPIVersion.Value)))
	}
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if policySetID == "" {
		err = errors.New("missing required policy_set_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policy-sets/%s", url.PathEscape(params.ZoneID), url.PathEscape(policySetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update metadata or manage binding. Set active=true to bind, active=false to
// unbind.
func (r *ZonePolicySetService) Update(ctx context.Context, policySetID string, params ZonePolicySetUpdateParams, opts ...option.RequestOption) (res *PolicySetWithBinding, err error) {
	if !param.IsOmitted(params.IfMatch) {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", params.IfMatch.Value)))
	}
	if !param.IsOmitted(params.XAPIVersion) {
		opts = append(opts, option.WithHeader("X-API-Version", fmt.Sprintf("%v", params.XAPIVersion.Value)))
	}
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if policySetID == "" {
		err = errors.New("missing required policy_set_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policy-sets/%s", url.PathEscape(params.ZoneID), url.PathEscape(policySetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// List policy sets in a zone
func (r *ZonePolicySetService) List(ctx context.Context, zoneID string, params ZonePolicySetListParams, opts ...option.RequestOption) (res *ZonePolicySetListResponse, err error) {
	if !param.IsOmitted(params.XAPIVersion) {
		opts = append(opts, option.WithHeader("X-API-Version", fmt.Sprintf("%v", params.XAPIVersion.Value)))
	}
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policy-sets", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Archive a policy set
func (r *ZonePolicySetService) Archive(ctx context.Context, policySetID string, params ZonePolicySetArchiveParams, opts ...option.RequestOption) (res *PolicySetWithBinding, err error) {
	if !param.IsOmitted(params.IfMatch) {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", params.IfMatch.Value)))
	}
	if !param.IsOmitted(params.XAPIVersion) {
		opts = append(opts, option.WithHeader("X-API-Version", fmt.Sprintf("%v", params.XAPIVersion.Value)))
	}
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if policySetID == "" {
		err = errors.New("missing required policy_set_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/policy-sets/%s", url.PathEscape(params.ZoneID), url.PathEscape(policySetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// JWS Flattened JSON Serialization (RFC 7515 §7.2.2) of a policy set attestation.
// The protected header carries the signing algorithm and key identifier; the
// payload is a base64url-encoded AttestationStatement canonicalized per RFC 8785
// (JCS). Verify using the zone JWKS endpoint (RFC 7517). Currently signed with
// RS256; future zone key types (e.g. EdDSA) will be indicated by the "alg" header
// — no envelope changes required.
type Attestation struct {
	// Base64url-encoded AttestationStatement (RFC 7515 §3). Decode to inspect
	// attestation content. The RFC 8785 canonical form of the decoded JSON is the JWS
	// Signing Input alongside the protected header.
	Payload string `json:"payload" api:"required"`
	// Base64url-encoded JWS protected header (RFC 7515 §4). Contains at minimum "alg"
	// (signing algorithm — currently RS256, will migrate to EdDSA) and "kid" (signing
	// key identifier resolvable via the zone JWKS endpoint).
	Protected string `json:"protected" api:"required"`
	// Base64url-encoded digital signature computed over the JWS Signing Input
	// (ASCII(protected) || '.' || payload) per RFC 7515 §5.1.
	Signature string `json:"signature" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Payload     respjson.Field
		Protected   respjson.Field
		Signature   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Attestation) RawJSON() string { return r.JSON.raw }
func (r *Attestation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicySet struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	CreatedBy string    `json:"created_by" api:"required"`
	Name      string    `json:"name" api:"required"`
	// Any of "platform", "customer".
	OwnerType PolicySetOwnerType `json:"owner_type" api:"required"`
	// Any of "zone", "resource", "user", "session".
	ScopeType  PolicySetScopeType `json:"scope_type" api:"required"`
	UpdatedAt  time.Time          `json:"updated_at" api:"required" format:"date-time"`
	ZoneID     string             `json:"zone_id" api:"required"`
	ArchivedAt time.Time          `json:"archived_at" api:"nullable" format:"date-time"`
	// Human-readable version number of the latest version (e.g., 1, 2, 3)
	LatestVersion   int64  `json:"latest_version" api:"nullable"`
	LatestVersionID string `json:"latest_version_id" api:"nullable"`
	UpdatedBy       string `json:"updated_by" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		CreatedBy       respjson.Field
		Name            respjson.Field
		OwnerType       respjson.Field
		ScopeType       respjson.Field
		UpdatedAt       respjson.Field
		ZoneID          respjson.Field
		ArchivedAt      respjson.Field
		LatestVersion   respjson.Field
		LatestVersionID respjson.Field
		UpdatedBy       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicySet) RawJSON() string { return r.JSON.raw }
func (r *PolicySet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicySetOwnerType string

const (
	PolicySetOwnerTypePlatform PolicySetOwnerType = "platform"
	PolicySetOwnerTypeCustomer PolicySetOwnerType = "customer"
)

type PolicySetScopeType string

const (
	PolicySetScopeTypeZone     PolicySetScopeType = "zone"
	PolicySetScopeTypeResource PolicySetScopeType = "resource"
	PolicySetScopeTypeUser     PolicySetScopeType = "user"
	PolicySetScopeTypeSession  PolicySetScopeType = "session"
)

type PolicySetManifest struct {
	Entries []PolicySetManifestEntry `json:"entries" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entries     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicySetManifest) RawJSON() string { return r.JSON.raw }
func (r *PolicySetManifest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PolicySetManifest to a PolicySetManifestParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PolicySetManifestParam.Overrides()
func (r PolicySetManifest) ToParam() PolicySetManifestParam {
	return param.Override[PolicySetManifestParam](json.RawMessage(r.RawJSON()))
}

// The property Entries is required.
type PolicySetManifestParam struct {
	Entries []PolicySetManifestEntryParam `json:"entries,omitzero" api:"required"`
	paramObj
}

func (r PolicySetManifestParam) MarshalJSON() (data []byte, err error) {
	type shadow PolicySetManifestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicySetManifestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicySetManifestEntry struct {
	PolicyID        string `json:"policy_id" api:"required"`
	PolicyVersionID string `json:"policy_version_id" api:"required"`
	// SHA-256 of the policy version content, populated by the server
	Sha string `json:"sha"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PolicyID        respjson.Field
		PolicyVersionID respjson.Field
		Sha             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicySetManifestEntry) RawJSON() string { return r.JSON.raw }
func (r *PolicySetManifestEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PolicySetManifestEntry to a PolicySetManifestEntryParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PolicySetManifestEntryParam.Overrides()
func (r PolicySetManifestEntry) ToParam() PolicySetManifestEntryParam {
	return param.Override[PolicySetManifestEntryParam](json.RawMessage(r.RawJSON()))
}

// The properties PolicyID, PolicyVersionID are required.
type PolicySetManifestEntryParam struct {
	PolicyID        string `json:"policy_id" api:"required"`
	PolicyVersionID string `json:"policy_version_id" api:"required"`
	// SHA-256 of the policy version content, populated by the server
	Sha param.Opt[string] `json:"sha,omitzero"`
	paramObj
}

func (r PolicySetManifestEntryParam) MarshalJSON() (data []byte, err error) {
	type shadow PolicySetManifestEntryParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicySetManifestEntryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PolicySetWithBinding struct {
	// Whether this policy set is currently bound to a scope
	Active bool `json:"active"`
	// Human-readable version number of the active version (e.g., 1, 2, 3)
	ActiveVersion int64 `json:"active_version" api:"nullable"`
	// Public ID of the currently active (bound) version
	ActiveVersionID string `json:"active_version_id" api:"nullable"`
	// Any of "active", "shadow".
	Mode          string `json:"mode" api:"nullable"`
	ScopeTargetID string `json:"scope_target_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active          respjson.Field
		ActiveVersion   respjson.Field
		ActiveVersionID respjson.Field
		Mode            respjson.Field
		ScopeTargetID   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
	PolicySet
}

// Returns the unmodified JSON received from the API
func (r PolicySetWithBinding) RawJSON() string { return r.JSON.raw }
func (r *PolicySetWithBinding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicySetListResponse struct {
	Items      []PolicySetWithBinding              `json:"items" api:"required"`
	Pagination ZonePolicySetListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicySetListResponse) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicySetListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicySetListResponsePagination struct {
	// Cursor of the last item on the current page. Pass to after for the next page.
	// Null when there is no next page.
	AfterCursor string `json:"after_cursor" api:"required"`
	// Cursor of the first item on the current page. Pass to before for the previous
	// page. Null when there is no previous page.
	BeforeCursor string `json:"before_cursor" api:"required"`
	// Total number of items matching the current filters. Only included when
	// expand=total_count is requested.
	TotalCount int64 `json:"total_count" api:"nullable"`
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
func (r ZonePolicySetListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicySetListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicySetNewParams struct {
	Name             string            `json:"name" api:"required"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Any of "zone", "resource", "user", "session".
	ScopeType ZonePolicySetNewParamsScopeType `json:"scope_type,omitzero"`
	paramObj
}

func (r ZonePolicySetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicySetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicySetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicySetNewParamsScopeType string

const (
	ZonePolicySetNewParamsScopeTypeZone     ZonePolicySetNewParamsScopeType = "zone"
	ZonePolicySetNewParamsScopeTypeResource ZonePolicySetNewParamsScopeType = "resource"
	ZonePolicySetNewParamsScopeTypeUser     ZonePolicySetNewParamsScopeType = "user"
	ZonePolicySetNewParamsScopeTypeSession  ZonePolicySetNewParamsScopeType = "session"
)

type ZonePolicySetGetParams struct {
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ZonePolicySetUpdateParams struct {
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	Name             param.Opt[string] `json:"name,omitzero"`
	IfMatch          param.Opt[string] `header:"If-Match,omitzero" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ZonePolicySetUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicySetUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicySetUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicySetListParams struct {
	// Return items after this cursor (forward pagination). Use after_cursor from a
	// previous response. Mutually exclusive with before.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Return items before this cursor (backward pagination). Use before_cursor from a
	// previous response. Mutually exclusive with after.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Maximum number of items to return
	Limit            param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Opt-in to additional response fields
	//
	// Any of "total_count".
	Expand []string `query:"expand,omitzero" json:"-"`
	// Sort direction. Default is desc (newest first).
	//
	// Any of "asc", "desc".
	Order ZonePolicySetListParamsOrder `query:"order,omitzero" json:"-"`
	// Field to sort by.
	//
	// Any of "created_at".
	Sort ZonePolicySetListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZonePolicySetListParams]'s query parameters as
// `url.Values`.
func (r ZonePolicySetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort direction. Default is desc (newest first).
type ZonePolicySetListParamsOrder string

const (
	ZonePolicySetListParamsOrderAsc  ZonePolicySetListParamsOrder = "asc"
	ZonePolicySetListParamsOrderDesc ZonePolicySetListParamsOrder = "desc"
)

// Field to sort by.
type ZonePolicySetListParamsSort string

const (
	ZonePolicySetListParamsSortCreatedAt ZonePolicySetListParamsSort = "created_at"
)

type ZonePolicySetArchiveParams struct {
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	IfMatch          param.Opt[string] `header:"If-Match,omitzero" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}
