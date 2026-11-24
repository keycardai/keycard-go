// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycard

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/keycardai/keycard-go/internal/apijson"
	"github.com/keycardai/keycard-go/internal/apiquery"
	"github.com/keycardai/keycard-go/internal/requestconfig"
	"github.com/keycardai/keycard-go/option"
	"github.com/keycardai/keycard-go/packages/param"
	"github.com/keycardai/keycard-go/packages/respjson"
)

// Declarative read/apply of a policy set's full content
//
// ZonePolicySetManifestService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZonePolicySetManifestService] method instead.
type ZonePolicySetManifestService struct {
	Options []option.RequestOption
}

// NewZonePolicySetManifestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZonePolicySetManifestService(opts ...option.RequestOption) (r ZonePolicySetManifestService) {
	r = ZonePolicySetManifestService{}
	r.Options = opts
	return
}

// Returns the resolved manifest of the **latest non-archived version** by default.
// Pass `?version=` with a policy set version public ID to read a specific
// version's manifest (e.g. for rollback: GET an old manifest, PUT it back — which
// mints a new version at the tip with the old content; history is never
// re-pointed).
//
// There is deliberately no `?version=active` keyword: a binding slot is
// `(zone, target_type, target_id, mode)` and a set may have zero or several active
// bindings, so "the active version" is not a function of the set alone. Resolve
// the version ID first — from the enriched policy set representation
// (`active_version_id`) or `GET /policy-bindings` — then pass it here.
//
// The default (latest) read anchors the apply flow: the `ETag` response header is
// the JCS SHA-256 digest of the latest manifest and is the value to send as
// `If-Match` on `PUT`. **`ETag` is emitted only on the default latest read** —
// versioned reads return `manifest_sha` in the body but no `ETag`, so a stale
// token cannot be round-tripped as a precondition.
//
// A policy set with no versions yet returns `200` with empty `entries`, null
// `policy_set_version_id`/`version`/`schema_version`/`manifest_sha`, and the
// digest of the empty manifest as `ETag`, so even the first apply can be
// conditional.
func (r *ZonePolicySetManifestService) Get(ctx context.Context, policySetID string, params ZonePolicySetManifestGetParams, opts ...option.RequestOption) (res *ZonePolicySetManifestGetResponse, err error) {
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
	path := fmt.Sprintf("zones/%s/policy-sets/%s/manifest", url.PathEscape(params.ZoneID), url.PathEscape(policySetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Declarative full-replace apply. `entries` is the complete desired content of the
// policy set: **any policy omitted from `entries` is dropped from the resulting
// version** (its rows are retained and a later apply may reference it again).
// Entry order is preserved for newly added policies; retained policies keep their
// current manifest order, so an unchanged apply is byte-identical and dedups to
// the existing version.
//
// Per entry, the server diffs by content SHA against the version pinned in the
// latest manifest: unchanged content under the same schema version reuses the
// pinned policy version; changed content mints a new policy version; `new_policy`
// entries mint a customer-owned policy with the requested name. All of this, plus
// the policy set version mint and the optional activation, happens in **one
// transaction** (row lock on the policy set). Cedar validation runs against the
// resolved schema before the transaction opens — an invalid policy rejects the
// whole apply with zero writes.
//
// **Schema resolution:** `schema_version` omitted → the latest version's schema if
// still writable, else the zone default. Provided → authoritative
// (`schema_version_unsupported` if unknown, unattested, or archived). All entries
// validate against the single resolved version (uniform-schema-per-version is
// retained); a pinned policy version created under a different schema is
// re-versioned even if its content is unchanged.
//
// **Activation:** `activate: true` requires a zone-targeted set and rebinds the
// `(zone, target_type=zone, mode=active)` slot to the resulting version in the
// same transaction. It requires `policy_set_bindings:activate` permission and
// preserves the previous binding in binding history. Unchanged content returns
// `created: false` and still activates the resulting version. An already active
// version keeps its binding unchanged.
//
// **Concurrency:** `If-Match` is optional (omitted = last-writer-wins). When
// provided it is compared, under the row lock, against the digest of the latest
// version's manifest; mismatch returns `412` (`etag_mismatch`) carrying the
// current `ETag`. The token is content-addressed, not a version identity: it
// guards "the content I read is still the content in place", so two versions with
// identical manifests are interchangeable for the precondition, and archiving the
// tip does not invalidate a token whose digest the new latest version still
// carries. Do not re-key it on `policy_set_version_id`.
//
// **Identity and reverts:** a desired state identical to the **latest** version is
// a no-op (`created: false`, latest returned, nothing minted). A desired state
// matching an **older** version mints a new version at the tip — history is
// linear, never forked. Rolling back deployment without a content edit is not an
// apply at all: bind the older version via `PUT /policy-bindings`.
//
// **Authorization:** the route gate is `update` on the policy set. When
// `activate: true`, the handler additionally requires the `activate` action on
// `policy_set_bindings`.
//
// Domain error codes: `policy_set_not_found`, `policy_not_found`,
// `policy_version_not_found`, `version_archived`, `schema_version_mismatch`,
// `manifest_duplicate_policy`, `missing_cedar_content`, `invalid_cedar`,
// `schema_version_unsupported`, `policy_name_conflict`, `policy_archived`,
// `set_archived`, `etag_mismatch`, `activate_requires_zone_target`,
// `shadow_equals_active`.
func (r *ZonePolicySetManifestService) Apply(ctx context.Context, policySetID string, params ZonePolicySetManifestApplyParams, opts ...option.RequestOption) (res *ZonePolicySetManifestApplyResponse, err error) {
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
	path := fmt.Sprintf("zones/%s/policy-sets/%s/manifest", url.PathEscape(params.ZoneID), url.PathEscape(policySetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Readback shape of a policy set version's manifest, richer than the write shape:
// each entry carries the resolved policy version ID and content SHA.
type ZonePolicySetManifestGetResponse struct {
	Entries     []PolicySetManifestEntry `json:"entries" api:"required"`
	PolicySetID string                   `json:"policy_set_id" api:"required"`
	// JCS SHA-256 digest of the stored manifest of the returned version. Null when no
	// versions exist yet. On the default latest read this is the same value the `ETag`
	// header carries.
	ManifestSha string `json:"manifest_sha" api:"nullable"`
	// Null only for a policy set with no versions yet.
	PolicySetVersionID string `json:"policy_set_version_id" api:"nullable"`
	// Null when no versions exist yet.
	SchemaVersion string `json:"schema_version" api:"nullable"`
	// Monotonic version number. Null when no versions exist yet.
	Version int64 `json:"version" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entries            respjson.Field
		PolicySetID        respjson.Field
		ManifestSha        respjson.Field
		PolicySetVersionID respjson.Field
		SchemaVersion      respjson.Field
		Version            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicySetManifestGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicySetManifestGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicySetManifestApplyResponse struct {
	// Per-policy outcome of the diff.
	Changes []ZonePolicySetManifestApplyResponseChange `json:"changes" api:"required"`
	// True when the server mints a new policy set version. False when the content
	// matches the latest version, even when activate changes the active binding.
	// Reverts to older content mint a new version at the tip and return true.
	Created bool `json:"created" api:"required"`
	// Present when activate was true.
	Binding PolicySetBinding `json:"binding"`
	// The resulting version. Omitted on dry_run.
	PolicySetVersion PolicySetVersion `json:"policy_set_version"`
	// Non-fatal findings — e.g. Cedar actions unknown to the resolved schema. Present
	// only when non-empty.
	Warnings []ZonePolicySetManifestApplyResponseWarning `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes          respjson.Field
		Created          respjson.Field
		Binding          respjson.Field
		PolicySetVersion respjson.Field
		Warnings         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicySetManifestApplyResponse) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicySetManifestApplyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicySetManifestApplyResponseChange struct {
	// `repinned`: an explicit `policy_version_id` replaced a different version the set
	// already pinned for that policy; no version minted.
	//
	// Any of "created_policy", "created_version", "reused", "repinned", "dropped".
	Action string `json:"action" api:"required"`
	// The policy's name. Lets a caller correlate a `created_policy` row with its
	// `new_policy` request entry without a re-list.
	Name     string `json:"name" api:"required"`
	PolicyID string `json:"policy_id" api:"required"`
	// Absent when action is dropped.
	PolicyVersionID string `json:"policy_version_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action          respjson.Field
		Name            respjson.Field
		PolicyID        respjson.Field
		PolicyVersionID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicySetManifestApplyResponseChange) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicySetManifestApplyResponseChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicySetManifestApplyResponseWarning struct {
	// Machine-readable warning code, e.g. unknown_actions.
	Code    string `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicySetManifestApplyResponseWarning) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicySetManifestApplyResponseWarning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicySetManifestGetParams struct {
	ZoneID string `path:"zone_id" api:"required" json:"-"`
	// Policy set version public ID. Default: latest non-archived version.
	Version          param.Opt[string] `query:"version,omitzero" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ZonePolicySetManifestGetParams]'s query parameters as
// `url.Values`.
func (r ZonePolicySetManifestGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZonePolicySetManifestApplyParams struct {
	ZoneID string `path:"zone_id" api:"required" json:"-"`
	// Complete desired content. Order is preserved for newly added policies; retained
	// policies keep their current manifest order (so an unchanged apply reproduces the
	// identical manifest and dedups).
	Entries []ZonePolicySetManifestApplyParamsEntryUnion `json:"entries,omitzero" api:"required"`
	// Not yet implemented — `true` returns 501 until ACC-1002. When true, performs
	// validation, schema resolution, and the content diff, and returns the would-be
	// `changes` and `warnings` — but persists nothing. `policy_set_version` is omitted
	// from the response and no ETag is returned.
	DryRun param.Opt[bool] `query:"dry_run,omitzero" json:"-"`
	// Also rebind the zone/active slot to the resulting version in the same
	// transaction. Requires the `activate` permission on policy set bindings. Targets
	// only the zone/active slot; other slots go through PUT /policy-bindings.
	Activate param.Opt[bool] `json:"activate,omitzero"`
	// Cedar schema version to validate and pin the resulting version against. Omitted:
	// the server resolves it (latest version's schema if still writable, else the zone
	// default).
	SchemaVersion    param.Opt[string] `json:"schema_version,omitzero"`
	IfMatch          param.Opt[string] `header:"If-Match,omitzero" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ZonePolicySetManifestApplyParams) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicySetManifestApplyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicySetManifestApplyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ZonePolicySetManifestApplyParams]'s query parameters as
// `url.Values`.
func (r ZonePolicySetManifestApplyParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZonePolicySetManifestApplyParamsEntryUnion struct {
	OfZonePolicySetManifestApplysEntryPdpExistingPolicyEntry *ZonePolicySetManifestApplyParamsEntryPdpExistingPolicyEntry `json:",omitzero,inline"`
	OfZonePolicySetManifestApplysEntryPdpNewPolicyEntry      *ZonePolicySetManifestApplyParamsEntryPdpNewPolicyEntry      `json:",omitzero,inline"`
	paramUnion
}

func (u ZonePolicySetManifestApplyParamsEntryUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfZonePolicySetManifestApplysEntryPdpExistingPolicyEntry, u.OfZonePolicySetManifestApplysEntryPdpNewPolicyEntry)
}
func (u *ZonePolicySetManifestApplyParamsEntryUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Reference to an existing (non-archived) policy in the zone — not limited to
// policies already in this set. With `cedar_raw`/`cedar_json` (mutually
// exclusive): the server diffs by content SHA; unchanged content under the
// resolved schema reuses the pinned policy version, changed content mints a new
// one. Without content ("pin as-is"): reuses the version pinned in the latest
// manifest, or the policy's latest version when the policy is newly added to this
// set. Bare pins are re-versioned when the resolved schema differs from the pinned
// version's schema. With `policy_version_id`: pins exactly that existing version
// and mints nothing. Mutually exclusive with `cedar_raw`/`cedar_json` (a version
// is content; 400 when both are supplied). The version must belong to `policy_id`,
// must not be archived (`version_archived`), and must have been validated against
// the resolved schema (`schema_version_mismatch`; no re-versioning). Reported as
// `repinned` when the set already pins a different version of the policy,
// otherwise `reused`. Platform-owned policies accept bare pins and
// `policy_version_id` (customers cannot mint versions of those).
//
// The property PolicyID is required.
type ZonePolicySetManifestApplyParamsEntryPdpExistingPolicyEntry struct {
	// Public ID of an existing policy in the zone.
	PolicyID string `json:"policy_id" api:"required"`
	// Cedar policy text. Mutually exclusive with cedar_json.
	CedarRaw param.Opt[string] `json:"cedar_raw,omitzero"`
	// Public ID of an existing version of `policy_id` to pin. Mutually exclusive with
	// cedar_raw and cedar_json.
	PolicyVersionID param.Opt[string] `json:"policy_version_id,omitzero"`
	// Cedar policy JSON. Mutually exclusive with cedar_raw.
	CedarJson any `json:"cedar_json,omitzero"`
	paramObj
}

func (r ZonePolicySetManifestApplyParamsEntryPdpExistingPolicyEntry) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicySetManifestApplyParamsEntryPdpExistingPolicyEntry
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicySetManifestApplyParamsEntryPdpExistingPolicyEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Mints a new customer-owned policy with the requested name (409
// `policy_name_conflict` on collision) plus its first version from the supplied
// content. Exactly one of `cedar_raw`/`cedar_json` is required.
//
// The property NewPolicy is required.
type ZonePolicySetManifestApplyParamsEntryPdpNewPolicyEntry struct {
	NewPolicy ZonePolicySetManifestApplyParamsEntryPdpNewPolicyEntryNewPolicy `json:"new_policy,omitzero" api:"required"`
	// Cedar policy text. Mutually exclusive with cedar_json.
	CedarRaw param.Opt[string] `json:"cedar_raw,omitzero"`
	// Cedar policy JSON. Mutually exclusive with cedar_raw.
	CedarJson any `json:"cedar_json,omitzero"`
	paramObj
}

func (r ZonePolicySetManifestApplyParamsEntryPdpNewPolicyEntry) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicySetManifestApplyParamsEntryPdpNewPolicyEntry
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicySetManifestApplyParamsEntryPdpNewPolicyEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type ZonePolicySetManifestApplyParamsEntryPdpNewPolicyEntryNewPolicy struct {
	Name        string            `json:"name" api:"required"`
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r ZonePolicySetManifestApplyParamsEntryPdpNewPolicyEntryNewPolicy) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicySetManifestApplyParamsEntryPdpNewPolicyEntryNewPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicySetManifestApplyParamsEntryPdpNewPolicyEntryNewPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
