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

// Stateless policy evaluation testing
//
// ZonePolicyEvaluationService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZonePolicyEvaluationService] method instead.
type ZonePolicyEvaluationService struct {
	Options []option.RequestOption
}

// NewZonePolicyEvaluationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZonePolicyEvaluationService(opts ...option.RequestOption) (r ZonePolicyEvaluationService) {
	r = ZonePolicyEvaluationService{}
	r.Options = opts
	return
}

// Stateless evaluation endpoint for testing Cedar policy evaluation against real
// entities before deployment. Supports evaluating a single policy, a full policy
// set, or raw Cedar text.
func (r *ZonePolicyEvaluationService) New(ctx context.Context, zoneID string, params ZonePolicyEvaluationNewParams, opts ...option.RequestOption) (res *ZonePolicyEvaluationNewResponse, err error) {
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
	path := fmt.Sprintf("zones/%s/policy-evaluations", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of evaluation request entries captured from live
// credentials API traffic for the specified zone. Sorted by created_at descending
// (newest first). Use with the test evaluation endpoint for replay-based impact
// analysis.
func (r *ZonePolicyEvaluationService) ListRequests(ctx context.Context, zoneID string, params ZonePolicyEvaluationListRequestsParams, opts ...option.RequestOption) (res *ZonePolicyEvaluationListRequestsResponse, err error) {
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
	path := fmt.Sprintf("zones/%s/policy-evaluations/requests", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type ZonePolicyEvaluationNewResponse struct {
	// Any of "allow", "deny".
	Decision    ZonePolicyEvaluationNewResponseDecision `json:"decision" api:"required"`
	EvaluatedAt time.Time                               `json:"evaluated_at" api:"required" format:"date-time"`
	// Per-sub-check results, each with the policies that determined it. A simple
	// request has one actor check; an on-behalf-of / impersonation request has a
	// subject check and an actor check. The overall decision is deny-wins across
	// these. Determining-policy names and policy-set version populate only when
	// evaluating a minted policy set/version; raw text and drafts are structural.
	Checks []ZonePolicyEvaluationNewResponseCheck `json:"checks"`
	// Raw Cedar policy IDs that determined the decision, of the form
	// "<policy_version_id>::<cedar_key>", unioned across sub-checks. See `checks` for
	// the resolved, human-readable form with per-check attribution.
	DeterminingPolicies []string                                      `json:"determining_policies"`
	Diagnostics         []ZonePolicyEvaluationNewResponseDiagnostic   `json:"diagnostics"`
	EntitySnapshot      ZonePolicyEvaluationNewResponseEntitySnapshot `json:"entity_snapshot"`
	// Any of "complete", "partial".
	EvaluationStatus ZonePolicyEvaluationNewResponseEvaluationStatus `json:"evaluation_status"`
	// Human-readable explanation of the decision, suitable for presenting to end
	// users.
	Reason string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Decision            respjson.Field
		EvaluatedAt         respjson.Field
		Checks              respjson.Field
		DeterminingPolicies respjson.Field
		Diagnostics         respjson.Field
		EntitySnapshot      respjson.Field
		EvaluationStatus    respjson.Field
		Reason              respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyEvaluationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyEvaluationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicyEvaluationNewResponseDecision string

const (
	ZonePolicyEvaluationNewResponseDecisionAllow ZonePolicyEvaluationNewResponseDecision = "allow"
	ZonePolicyEvaluationNewResponseDecisionDeny  ZonePolicyEvaluationNewResponseDecision = "deny"
)

// Result of a single Cedar sub-check within a compound evaluation. Each check's
// determining policies and their effect are attributed to that check's own
// decision, not the deny-wins aggregate.
type ZonePolicyEvaluationNewResponseCheck struct {
	// Any of "allow", "deny".
	Decision            string                                                  `json:"decision"`
	DeterminingPolicies []ZonePolicyEvaluationNewResponseCheckDeterminingPolicy `json:"determining_policies"`
	// actor = the requesting principal's access (or its right to act for the subject
	// in an on-behalf-of flow); subject = the impersonated / delegated principal's own
	// access.
	//
	// Any of "actor", "subject".
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Decision            respjson.Field
		DeterminingPolicies respjson.Field
		Role                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyEvaluationNewResponseCheck) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyEvaluationNewResponseCheck) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A policy that contributed to the decision, enriched with human-readable names
// and version. For an allow these are permit policies; for an explicit deny these
// are forbid policies. An implicit (fall-through) deny has none. On the management
// test-evaluate endpoint, names and policy-set version populate only when a minted
// policy version is being evaluated; raw text and draft sources stay structural.
type ZonePolicyEvaluationNewResponseCheckDeterminingPolicy struct {
	// Any of "permit", "forbid".
	Effect string `json:"effect"`
	// Public ID of the policy.
	PolicyID string `json:"policy_id"`
	// Human-readable policy name.
	PolicyName string `json:"policy_name"`
	// Public ID of the policy set the policy belongs to.
	PolicySetID string `json:"policy_set_id"`
	// Human-readable policy set name.
	PolicySetName string `json:"policy_set_name"`
	// Version number of the policy set.
	PolicySetVersion int64 `json:"policy_set_version"`
	// Public ID of the policy version that was evaluated.
	PolicyVersionID string `json:"policy_version_id"`
	// Cedar per-rule `@id` annotation that determined the decision (e.g.
	// `limit-linear-scopes`). Falls back to the generated `policyN` key when the Cedar
	// statement carries no `@id`.
	RuleID string `json:"rule_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Effect           respjson.Field
		PolicyID         respjson.Field
		PolicyName       respjson.Field
		PolicySetID      respjson.Field
		PolicySetName    respjson.Field
		PolicySetVersion respjson.Field
		PolicyVersionID  respjson.Field
		RuleID           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyEvaluationNewResponseCheckDeterminingPolicy) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyEvaluationNewResponseCheckDeterminingPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicyEvaluationNewResponseDiagnostic struct {
	Message  string `json:"message"`
	PolicyID string `json:"policy_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		PolicyID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyEvaluationNewResponseDiagnostic) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyEvaluationNewResponseDiagnostic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicyEvaluationNewResponseEntitySnapshot struct {
	// The assembled Cedar context record.
	Context any `json:"context"`
	// Subset of the IAM User entity relevant to Cedar evaluation. All fields match the
	// IAM management API User schema.
	Principal ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalUnion `json:"principal"`
	// Subset of the IAM Resource entity relevant to Cedar evaluation. All fields match
	// the IAM management API Resource schema.
	Resource ZonePolicyEvaluationNewResponseEntitySnapshotResource `json:"resource"`
	// Present when subject_type/subject_id were provided.
	Subject ZonePolicyEvaluationNewResponseEntitySnapshotSubjectUnion `json:"subject"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Context     respjson.Field
		Principal   respjson.Field
		Resource    respjson.Field
		Subject     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyEvaluationNewResponseEntitySnapshot) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyEvaluationNewResponseEntitySnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalUnion contains all
// possible properties and values from
// [ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalPdpUserSnapshot],
// [ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalPdpApplicationSnapshot].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalUnion struct {
	ID string `json:"id"`
	// This field is from variant
	// [ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalPdpUserSnapshot].
	Email string `json:"email"`
	Type  string `json:"type"`
	// This field is from variant
	// [ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalPdpUserSnapshot].
	EmailVerified bool `json:"email_verified"`
	// This field is from variant
	// [ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalPdpApplicationSnapshot].
	Name string `json:"name"`
	// This field is from variant
	// [ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalPdpApplicationSnapshot].
	RegistrationMethod string `json:"registration_method"`
	// This field is from variant
	// [ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalPdpApplicationSnapshot].
	CredentialType string `json:"credential_type"`
	// This field is from variant
	// [ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalPdpApplicationSnapshot].
	Dependencies []ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalPdpApplicationSnapshotDependency `json:"dependencies"`
	JSON         struct {
		ID                 respjson.Field
		Email              respjson.Field
		Type               respjson.Field
		EmailVerified      respjson.Field
		Name               respjson.Field
		RegistrationMethod respjson.Field
		CredentialType     respjson.Field
		Dependencies       respjson.Field
		raw                string
	} `json:"-"`
}

func (u ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalUnion) AsZonePolicyEvaluationNewResponseEntitySnapshotPrincipalPdpUserSnapshot() (v ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalPdpUserSnapshot) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalUnion) AsZonePolicyEvaluationNewResponseEntitySnapshotPrincipalPdpApplicationSnapshot() (v ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalPdpApplicationSnapshot) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subset of the IAM User entity relevant to Cedar evaluation. All fields match the
// IAM management API User schema.
type ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalPdpUserSnapshot struct {
	// User public ID (same as IAM management API).
	ID    string `json:"id" api:"required"`
	Email string `json:"email" api:"required"`
	// Any of "user".
	Type          string `json:"type" api:"required"`
	EmailVerified bool   `json:"email_verified"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Email         respjson.Field
		Type          respjson.Field
		EmailVerified respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalPdpUserSnapshot) RawJSON() string {
	return r.JSON.raw
}
func (r *ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalPdpUserSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subset of the IAM Application entity relevant to Cedar evaluation. All fields
// match the IAM management API Application schema.
type ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalPdpApplicationSnapshot struct {
	// Application public ID (same as IAM management API).
	ID   string `json:"id" api:"required"`
	Name string `json:"name" api:"required"`
	// Any of "managed", "dcr".
	RegistrationMethod string `json:"registration_method" api:"required"`
	// Any of "application".
	Type string `json:"type" api:"required"`
	// Present only when resolved via credential lookup.
	CredentialType string `json:"credential_type"`
	// Resources this application depends on.
	Dependencies []ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalPdpApplicationSnapshotDependency `json:"dependencies"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Name               respjson.Field
		RegistrationMethod respjson.Field
		Type               respjson.Field
		CredentialType     respjson.Field
		Dependencies       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalPdpApplicationSnapshot) RawJSON() string {
	return r.JSON.raw
}
func (r *ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalPdpApplicationSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subset of the IAM Resource entity relevant to Cedar evaluation. All fields match
// the IAM management API Resource schema.
type ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalPdpApplicationSnapshotDependency struct {
	// Resource public ID (same as IAM management API).
	ID         string   `json:"id" api:"required"`
	Identifier string   `json:"identifier" api:"required"`
	Name       string   `json:"name" api:"required"`
	Scopes     []string `json:"scopes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Identifier  respjson.Field
		Name        respjson.Field
		Scopes      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalPdpApplicationSnapshotDependency) RawJSON() string {
	return r.JSON.raw
}
func (r *ZonePolicyEvaluationNewResponseEntitySnapshotPrincipalPdpApplicationSnapshotDependency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subset of the IAM Resource entity relevant to Cedar evaluation. All fields match
// the IAM management API Resource schema.
type ZonePolicyEvaluationNewResponseEntitySnapshotResource struct {
	// Resource public ID (same as IAM management API).
	ID         string   `json:"id" api:"required"`
	Identifier string   `json:"identifier" api:"required"`
	Name       string   `json:"name" api:"required"`
	Scopes     []string `json:"scopes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Identifier  respjson.Field
		Name        respjson.Field
		Scopes      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyEvaluationNewResponseEntitySnapshotResource) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyEvaluationNewResponseEntitySnapshotResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ZonePolicyEvaluationNewResponseEntitySnapshotSubjectUnion contains all possible
// properties and values from
// [ZonePolicyEvaluationNewResponseEntitySnapshotSubjectPdpUserSnapshot],
// [ZonePolicyEvaluationNewResponseEntitySnapshotSubjectPdpApplicationSnapshot].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ZonePolicyEvaluationNewResponseEntitySnapshotSubjectUnion struct {
	ID string `json:"id"`
	// This field is from variant
	// [ZonePolicyEvaluationNewResponseEntitySnapshotSubjectPdpUserSnapshot].
	Email string `json:"email"`
	Type  string `json:"type"`
	// This field is from variant
	// [ZonePolicyEvaluationNewResponseEntitySnapshotSubjectPdpUserSnapshot].
	EmailVerified bool `json:"email_verified"`
	// This field is from variant
	// [ZonePolicyEvaluationNewResponseEntitySnapshotSubjectPdpApplicationSnapshot].
	Name string `json:"name"`
	// This field is from variant
	// [ZonePolicyEvaluationNewResponseEntitySnapshotSubjectPdpApplicationSnapshot].
	RegistrationMethod string `json:"registration_method"`
	// This field is from variant
	// [ZonePolicyEvaluationNewResponseEntitySnapshotSubjectPdpApplicationSnapshot].
	CredentialType string `json:"credential_type"`
	// This field is from variant
	// [ZonePolicyEvaluationNewResponseEntitySnapshotSubjectPdpApplicationSnapshot].
	Dependencies []ZonePolicyEvaluationNewResponseEntitySnapshotSubjectPdpApplicationSnapshotDependency `json:"dependencies"`
	JSON         struct {
		ID                 respjson.Field
		Email              respjson.Field
		Type               respjson.Field
		EmailVerified      respjson.Field
		Name               respjson.Field
		RegistrationMethod respjson.Field
		CredentialType     respjson.Field
		Dependencies       respjson.Field
		raw                string
	} `json:"-"`
}

func (u ZonePolicyEvaluationNewResponseEntitySnapshotSubjectUnion) AsZonePolicyEvaluationNewResponseEntitySnapshotSubjectPdpUserSnapshot() (v ZonePolicyEvaluationNewResponseEntitySnapshotSubjectPdpUserSnapshot) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ZonePolicyEvaluationNewResponseEntitySnapshotSubjectUnion) AsZonePolicyEvaluationNewResponseEntitySnapshotSubjectPdpApplicationSnapshot() (v ZonePolicyEvaluationNewResponseEntitySnapshotSubjectPdpApplicationSnapshot) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ZonePolicyEvaluationNewResponseEntitySnapshotSubjectUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ZonePolicyEvaluationNewResponseEntitySnapshotSubjectUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subset of the IAM User entity relevant to Cedar evaluation. All fields match the
// IAM management API User schema.
type ZonePolicyEvaluationNewResponseEntitySnapshotSubjectPdpUserSnapshot struct {
	// User public ID (same as IAM management API).
	ID    string `json:"id" api:"required"`
	Email string `json:"email" api:"required"`
	// Any of "user".
	Type          string `json:"type" api:"required"`
	EmailVerified bool   `json:"email_verified"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Email         respjson.Field
		Type          respjson.Field
		EmailVerified respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyEvaluationNewResponseEntitySnapshotSubjectPdpUserSnapshot) RawJSON() string {
	return r.JSON.raw
}
func (r *ZonePolicyEvaluationNewResponseEntitySnapshotSubjectPdpUserSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subset of the IAM Application entity relevant to Cedar evaluation. All fields
// match the IAM management API Application schema.
type ZonePolicyEvaluationNewResponseEntitySnapshotSubjectPdpApplicationSnapshot struct {
	// Application public ID (same as IAM management API).
	ID   string `json:"id" api:"required"`
	Name string `json:"name" api:"required"`
	// Any of "managed", "dcr".
	RegistrationMethod string `json:"registration_method" api:"required"`
	// Any of "application".
	Type string `json:"type" api:"required"`
	// Present only when resolved via credential lookup.
	CredentialType string `json:"credential_type"`
	// Resources this application depends on.
	Dependencies []ZonePolicyEvaluationNewResponseEntitySnapshotSubjectPdpApplicationSnapshotDependency `json:"dependencies"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Name               respjson.Field
		RegistrationMethod respjson.Field
		Type               respjson.Field
		CredentialType     respjson.Field
		Dependencies       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyEvaluationNewResponseEntitySnapshotSubjectPdpApplicationSnapshot) RawJSON() string {
	return r.JSON.raw
}
func (r *ZonePolicyEvaluationNewResponseEntitySnapshotSubjectPdpApplicationSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Subset of the IAM Resource entity relevant to Cedar evaluation. All fields match
// the IAM management API Resource schema.
type ZonePolicyEvaluationNewResponseEntitySnapshotSubjectPdpApplicationSnapshotDependency struct {
	// Resource public ID (same as IAM management API).
	ID         string   `json:"id" api:"required"`
	Identifier string   `json:"identifier" api:"required"`
	Name       string   `json:"name" api:"required"`
	Scopes     []string `json:"scopes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Identifier  respjson.Field
		Name        respjson.Field
		Scopes      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyEvaluationNewResponseEntitySnapshotSubjectPdpApplicationSnapshotDependency) RawJSON() string {
	return r.JSON.raw
}
func (r *ZonePolicyEvaluationNewResponseEntitySnapshotSubjectPdpApplicationSnapshotDependency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicyEvaluationNewResponseEvaluationStatus string

const (
	ZonePolicyEvaluationNewResponseEvaluationStatusComplete ZonePolicyEvaluationNewResponseEvaluationStatus = "complete"
	ZonePolicyEvaluationNewResponseEvaluationStatusPartial  ZonePolicyEvaluationNewResponseEvaluationStatus = "partial"
)

type ZonePolicyEvaluationListRequestsResponse struct {
	Items []ZonePolicyEvaluationListRequestsResponseItem `json:"items" api:"required"`
	// Cursor-based pagination metadata returned alongside a list of results
	Pagination ZonePolicyEvaluationListRequestsResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyEvaluationListRequestsResponse) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyEvaluationListRequestsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicyEvaluationListRequestsResponseItem struct {
	// Unique identifier for this evaluation request entry.
	ID     string `json:"id" api:"required" format:"uuid"`
	Action string `json:"action" api:"required"`
	// Any of "allow", "deny".
	ActiveDecision string `json:"active_decision" api:"required"`
	// Evaluation duration in milliseconds.
	ActiveDurationMs int64 `json:"active_duration_ms" api:"required"`
	// ID of the active policy set version used for evaluation.
	ActivePsvID string `json:"active_psv_id" api:"required" format:"uuid"`
	// Cedar schema version of the active policy set version.
	ActiveSchemaVersion string    `json:"active_schema_version" api:"required"`
	CreatedAt           time.Time `json:"created_at" api:"required" format:"date-time"`
	PrincipalID         string    `json:"principal_id" api:"required"`
	// Any of "user", "application".
	PrincipalType string `json:"principal_type" api:"required"`
	// Correlation ID from the original credentials API request.
	RequestID  string `json:"request_id" api:"required"`
	ResourceID string `json:"resource_id" api:"required"`
	// The verified RFC 8693 acting principal (nominated via actor_token) from the
	// original check, when one was present.
	ActingPrincipalID string `json:"acting_principal_id" api:"nullable"`
	// Any of "delegation", "impersonation".
	AuthorityType string `json:"authority_type" api:"nullable"`
	// Additional context for Cedar authorization evaluation. Fields here are mapped
	// directly into the Cedar `context` record.
	Context ZonePolicyEvaluationListRequestsResponseItemContext `json:"context"`
	// ID of the shadow policy set version, if a shadow binding existed.
	ShadowPsvID string `json:"shadow_psv_id" api:"nullable" format:"uuid"`
	SubjectID   string `json:"subject_id" api:"nullable"`
	// Any of "user", "application".
	SubjectType string `json:"subject_type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Action              respjson.Field
		ActiveDecision      respjson.Field
		ActiveDurationMs    respjson.Field
		ActivePsvID         respjson.Field
		ActiveSchemaVersion respjson.Field
		CreatedAt           respjson.Field
		PrincipalID         respjson.Field
		PrincipalType       respjson.Field
		RequestID           respjson.Field
		ResourceID          respjson.Field
		ActingPrincipalID   respjson.Field
		AuthorityType       respjson.Field
		Context             respjson.Field
		ShadowPsvID         respjson.Field
		SubjectID           respjson.Field
		SubjectType         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyEvaluationListRequestsResponseItem) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyEvaluationListRequestsResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional context for Cedar authorization evaluation. Fields here are mapped
// directly into the Cedar `context` record.
type ZonePolicyEvaluationListRequestsResponseItemContext struct {
	// Token claims for the actor, subject, and/or acting principal. Mapped to
	// `context.actor_claims`, `context.subject_claims`, and
	// `context.acting_principal_claims` in Cedar.
	Claims ZonePolicyEvaluationListRequestsResponseItemContextClaims `json:"claims"`
	// OAuth scopes associated with the request. Mapped to `context.scopes` in the
	// Cedar evaluation as a `Set<String>`.
	Scopes []string `json:"scopes"`
	// Session identifier for the current authorization context. Mapped to
	// `context.session_id` in Cedar.
	SessionID string `json:"session_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Claims      respjson.Field
		Scopes      respjson.Field
		SessionID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyEvaluationListRequestsResponseItemContext) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyEvaluationListRequestsResponseItemContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token claims for the actor, subject, and/or acting principal. Mapped to
// `context.actor_claims`, `context.subject_claims`, and
// `context.acting_principal_claims` in Cedar.
type ZonePolicyEvaluationListRequestsResponseItemContextClaims struct {
	// Represents token claims passed into the Cedar context. Mapped to the Cedar
	// `Keycard::Claims` type. Known fields (`email`, `groups`, `issuer_claims`) are
	// typed; additional claims are passed through as untyped key-value pairs via
	// `additionalProperties`.
	ActingPrincipal ZonePolicyEvaluationListRequestsResponseItemContextClaimsActingPrincipal `json:"acting_principal"`
	// Represents token claims passed into the Cedar context. Mapped to the Cedar
	// `Keycard::Claims` type. Known fields (`email`, `groups`, `issuer_claims`) are
	// typed; additional claims are passed through as untyped key-value pairs via
	// `additionalProperties`.
	Actor ZonePolicyEvaluationListRequestsResponseItemContextClaimsActor `json:"actor"`
	// Represents token claims passed into the Cedar context. Mapped to the Cedar
	// `Keycard::Claims` type. Known fields (`email`, `groups`, `issuer_claims`) are
	// typed; additional claims are passed through as untyped key-value pairs via
	// `additionalProperties`.
	Subject ZonePolicyEvaluationListRequestsResponseItemContextClaimsSubject `json:"subject"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActingPrincipal respjson.Field
		Actor           respjson.Field
		Subject         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyEvaluationListRequestsResponseItemContextClaims) RawJSON() string {
	return r.JSON.raw
}
func (r *ZonePolicyEvaluationListRequestsResponseItemContextClaims) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents token claims passed into the Cedar context. Mapped to the Cedar
// `Keycard::Claims` type. Known fields (`email`, `groups`, `issuer_claims`) are
// typed; additional claims are passed through as untyped key-value pairs via
// `additionalProperties`.
type ZonePolicyEvaluationListRequestsResponseItemContextClaimsActingPrincipal struct {
	// Email claim from the authentication token.
	Email string `json:"email"`
	// Group membership claims from the authentication token.
	Groups []string `json:"groups"`
	// Issuer-specific claims forwarded opaquely from STS (e.g. GHA OIDC assertion
	// claims). Mapped to `context.actor_claims.issuer_claims` in Cedar. All fields are
	// optional and issuer-dependent. See issuer documentation for available claim
	// fields.
	IssuerClaims map[string]any `json:"issuer_claims"`
	ExtraFields  map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email        respjson.Field
		Groups       respjson.Field
		IssuerClaims respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyEvaluationListRequestsResponseItemContextClaimsActingPrincipal) RawJSON() string {
	return r.JSON.raw
}
func (r *ZonePolicyEvaluationListRequestsResponseItemContextClaimsActingPrincipal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents token claims passed into the Cedar context. Mapped to the Cedar
// `Keycard::Claims` type. Known fields (`email`, `groups`, `issuer_claims`) are
// typed; additional claims are passed through as untyped key-value pairs via
// `additionalProperties`.
type ZonePolicyEvaluationListRequestsResponseItemContextClaimsActor struct {
	// Email claim from the authentication token.
	Email string `json:"email"`
	// Group membership claims from the authentication token.
	Groups []string `json:"groups"`
	// Issuer-specific claims forwarded opaquely from STS (e.g. GHA OIDC assertion
	// claims). Mapped to `context.actor_claims.issuer_claims` in Cedar. All fields are
	// optional and issuer-dependent. See issuer documentation for available claim
	// fields.
	IssuerClaims map[string]any `json:"issuer_claims"`
	ExtraFields  map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email        respjson.Field
		Groups       respjson.Field
		IssuerClaims respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyEvaluationListRequestsResponseItemContextClaimsActor) RawJSON() string {
	return r.JSON.raw
}
func (r *ZonePolicyEvaluationListRequestsResponseItemContextClaimsActor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents token claims passed into the Cedar context. Mapped to the Cedar
// `Keycard::Claims` type. Known fields (`email`, `groups`, `issuer_claims`) are
// typed; additional claims are passed through as untyped key-value pairs via
// `additionalProperties`.
type ZonePolicyEvaluationListRequestsResponseItemContextClaimsSubject struct {
	// Email claim from the authentication token.
	Email string `json:"email"`
	// Group membership claims from the authentication token.
	Groups []string `json:"groups"`
	// Issuer-specific claims forwarded opaquely from STS (e.g. GHA OIDC assertion
	// claims). Mapped to `context.actor_claims.issuer_claims` in Cedar. All fields are
	// optional and issuer-dependent. See issuer documentation for available claim
	// fields.
	IssuerClaims map[string]any `json:"issuer_claims"`
	ExtraFields  map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email        respjson.Field
		Groups       respjson.Field
		IssuerClaims respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZonePolicyEvaluationListRequestsResponseItemContextClaimsSubject) RawJSON() string {
	return r.JSON.raw
}
func (r *ZonePolicyEvaluationListRequestsResponseItemContextClaimsSubject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination metadata returned alongside a list of results
type ZonePolicyEvaluationListRequestsResponsePagination struct {
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
func (r ZonePolicyEvaluationListRequestsResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ZonePolicyEvaluationListRequestsResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicyEvaluationNewParams struct {
	// Cedar action to evaluate (e.g. `any`). The engine namespaces actions
	// automatically as `Keycard::Action::"<name>"`.
	Action string `json:"action" api:"required"`
	// Specifies which policies to evaluate. Exactly one of policy, policy_set, raw, or
	// manifest must be provided.
	PolicySource ZonePolicyEvaluationNewParamsPolicySource `json:"policy_source,omitzero" api:"required"`
	// Public ID of the user or application.
	PrincipalID string `json:"principal_id" api:"required"`
	// Any of "user", "application".
	PrincipalType ZonePolicyEvaluationNewParamsPrincipalType `json:"principal_type,omitzero" api:"required"`
	// Public ID of the resource.
	ResourceID string `json:"resource_id" api:"required"`
	// ID of the acting application from a verified RFC 8693 `actor_token`, as sent on
	// `/policy/check`. Resolved with the same entity lookup as an application
	// `principal_id` and exposed to Cedar as `context.acting_principal.application` on
	// both checks. Lets a captured evaluation request be replayed faithfully.
	ActingPrincipalID param.Opt[string] `json:"acting_principal_id,omitzero"`
	// Public ID of the subject (required when subject_type is set).
	SubjectID        param.Opt[string] `json:"subject_id,omitzero"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Discriminator for the subject authority model. Determines how the subject
	// relationship is expressed in the Cedar context:
	//
	//   - `"delegation"` — sets `context.on_behalf = true` on the actor check. The actor
	//     is acting on behalf of the subject.
	//   - `"impersonation"` — sets `context.impersonate = true` on the actor check. The
	//     actor is impersonating the subject. When null and a subject is present,
	//     defaults to `"delegation"` for backward compatibility.
	//
	// Any of "delegation", "impersonation".
	AuthorityType ZonePolicyEvaluationNewParamsAuthorityType `json:"authority_type,omitzero"`
	// Additional context for Cedar authorization evaluation. Fields here are mapped
	// directly into the Cedar `context` record.
	Context ZonePolicyEvaluationNewParamsContext `json:"context,omitzero"`
	// Any of "user", "application".
	SubjectType ZonePolicyEvaluationNewParamsSubjectType `json:"subject_type,omitzero"`
	paramObj
}

func (r ZonePolicyEvaluationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicyEvaluationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicyEvaluationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies which policies to evaluate. Exactly one of policy, policy_set, raw, or
// manifest must be provided.
type ZonePolicyEvaluationNewParamsPolicySource struct {
	// Inline policy-set manifest used by the `manifest` variant of PolicySource.
	// Mirrors the shape of a stored policy set version's manifest so a policy set
	// under construction (no stored id/version) can be tested without being saved
	// first. The `sha` field on each entry is ignored on input and recomputed
	// server-side.
	Manifest  ZonePolicyEvaluationNewParamsPolicySourceManifest  `json:"manifest,omitzero"`
	Policy    ZonePolicyEvaluationNewParamsPolicySourcePolicy    `json:"policy,omitzero"`
	PolicySet ZonePolicyEvaluationNewParamsPolicySourcePolicySet `json:"policy_set,omitzero"`
	Raw       ZonePolicyEvaluationNewParamsPolicySourceRaw       `json:"raw,omitzero"`
	paramObj
}

func (r ZonePolicyEvaluationNewParamsPolicySource) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicyEvaluationNewParamsPolicySource
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicyEvaluationNewParamsPolicySource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inline policy-set manifest used by the `manifest` variant of PolicySource.
// Mirrors the shape of a stored policy set version's manifest so a policy set
// under construction (no stored id/version) can be tested without being saved
// first. The `sha` field on each entry is ignored on input and recomputed
// server-side.
//
// The properties Entries, SchemaVersion are required.
type ZonePolicyEvaluationNewParamsPolicySourceManifest struct {
	Entries []PolicySetManifestEntryParam `json:"entries,omitzero" api:"required"`
	// Cedar schema version the entries are validated against.
	SchemaVersion string `json:"schema_version" api:"required"`
	paramObj
}

func (r ZonePolicyEvaluationNewParamsPolicySourceManifest) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicyEvaluationNewParamsPolicySourceManifest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicyEvaluationNewParamsPolicySourceManifest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property PolicyID is required.
type ZonePolicyEvaluationNewParamsPolicySourcePolicy struct {
	// Public ID of the policy.
	PolicyID string `json:"policy_id" api:"required"`
	// Cedar schema version override (date string, e.g. "2026-03-16"). When omitted,
	// uses the schema_version recorded on the resolved policy version or draft.
	SchemaVersion param.Opt[string] `json:"schema_version,omitzero"`
	// Public ID of a specific policy version. When omitted, the draft is used. If no
	// draft exists, returns 400.
	VersionID param.Opt[string] `json:"version_id,omitzero"`
	paramObj
}

func (r ZonePolicyEvaluationNewParamsPolicySourcePolicy) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicyEvaluationNewParamsPolicySourcePolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicyEvaluationNewParamsPolicySourcePolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property PolicySetID is required.
type ZonePolicyEvaluationNewParamsPolicySourcePolicySet struct {
	// Public ID of the policy set.
	PolicySetID string `json:"policy_set_id" api:"required"`
	// Public ID of a specific policy set version. When omitted, the draft is used. If
	// no draft exists, returns 400.
	VersionID param.Opt[string] `json:"version_id,omitzero"`
	paramObj
}

func (r ZonePolicyEvaluationNewParamsPolicySourcePolicySet) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicyEvaluationNewParamsPolicySourcePolicySet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicyEvaluationNewParamsPolicySourcePolicySet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Cedar, SchemaVersion are required.
type ZonePolicyEvaluationNewParamsPolicySourceRaw struct {
	// Raw Cedar policy text. Can contain one or more policy statements. Converted to
	// JSON before evaluation.
	Cedar string `json:"cedar" api:"required"`
	// Cedar schema version to evaluate against.
	SchemaVersion string `json:"schema_version" api:"required"`
	// When true, validate the converted policy against the schema before evaluation.
	// Validation errors are returned as 400 with details. Set to false to skip
	// validation and attempt evaluation directly.
	Validate param.Opt[bool] `json:"validate,omitzero"`
	paramObj
}

func (r ZonePolicyEvaluationNewParamsPolicySourceRaw) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicyEvaluationNewParamsPolicySourceRaw
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicyEvaluationNewParamsPolicySourceRaw) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicyEvaluationNewParamsPrincipalType string

const (
	ZonePolicyEvaluationNewParamsPrincipalTypeUser        ZonePolicyEvaluationNewParamsPrincipalType = "user"
	ZonePolicyEvaluationNewParamsPrincipalTypeApplication ZonePolicyEvaluationNewParamsPrincipalType = "application"
)

// Discriminator for the subject authority model. Determines how the subject
// relationship is expressed in the Cedar context:
//
//   - `"delegation"` — sets `context.on_behalf = true` on the actor check. The actor
//     is acting on behalf of the subject.
//   - `"impersonation"` — sets `context.impersonate = true` on the actor check. The
//     actor is impersonating the subject. When null and a subject is present,
//     defaults to `"delegation"` for backward compatibility.
type ZonePolicyEvaluationNewParamsAuthorityType string

const (
	ZonePolicyEvaluationNewParamsAuthorityTypeDelegation    ZonePolicyEvaluationNewParamsAuthorityType = "delegation"
	ZonePolicyEvaluationNewParamsAuthorityTypeImpersonation ZonePolicyEvaluationNewParamsAuthorityType = "impersonation"
)

// Additional context for Cedar authorization evaluation. Fields here are mapped
// directly into the Cedar `context` record.
type ZonePolicyEvaluationNewParamsContext struct {
	// Session identifier for the current authorization context. Mapped to
	// `context.session_id` in Cedar.
	SessionID param.Opt[string] `json:"session_id,omitzero"`
	// Token claims for the actor, subject, and/or acting principal. Mapped to
	// `context.actor_claims`, `context.subject_claims`, and
	// `context.acting_principal_claims` in Cedar.
	Claims ZonePolicyEvaluationNewParamsContextClaims `json:"claims,omitzero"`
	// OAuth scopes associated with the request. Mapped to `context.scopes` in the
	// Cedar evaluation as a `Set<String>`.
	Scopes []string `json:"scopes,omitzero"`
	paramObj
}

func (r ZonePolicyEvaluationNewParamsContext) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicyEvaluationNewParamsContext
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicyEvaluationNewParamsContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token claims for the actor, subject, and/or acting principal. Mapped to
// `context.actor_claims`, `context.subject_claims`, and
// `context.acting_principal_claims` in Cedar.
type ZonePolicyEvaluationNewParamsContextClaims struct {
	// Represents token claims passed into the Cedar context. Mapped to the Cedar
	// `Keycard::Claims` type. Known fields (`email`, `groups`, `issuer_claims`) are
	// typed; additional claims are passed through as untyped key-value pairs via
	// `additionalProperties`.
	ActingPrincipal ZonePolicyEvaluationNewParamsContextClaimsActingPrincipal `json:"acting_principal,omitzero"`
	// Represents token claims passed into the Cedar context. Mapped to the Cedar
	// `Keycard::Claims` type. Known fields (`email`, `groups`, `issuer_claims`) are
	// typed; additional claims are passed through as untyped key-value pairs via
	// `additionalProperties`.
	Actor ZonePolicyEvaluationNewParamsContextClaimsActor `json:"actor,omitzero"`
	// Represents token claims passed into the Cedar context. Mapped to the Cedar
	// `Keycard::Claims` type. Known fields (`email`, `groups`, `issuer_claims`) are
	// typed; additional claims are passed through as untyped key-value pairs via
	// `additionalProperties`.
	Subject ZonePolicyEvaluationNewParamsContextClaimsSubject `json:"subject,omitzero"`
	paramObj
}

func (r ZonePolicyEvaluationNewParamsContextClaims) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicyEvaluationNewParamsContextClaims
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZonePolicyEvaluationNewParamsContextClaims) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents token claims passed into the Cedar context. Mapped to the Cedar
// `Keycard::Claims` type. Known fields (`email`, `groups`, `issuer_claims`) are
// typed; additional claims are passed through as untyped key-value pairs via
// `additionalProperties`.
type ZonePolicyEvaluationNewParamsContextClaimsActingPrincipal struct {
	// Email claim from the authentication token.
	Email param.Opt[string] `json:"email,omitzero"`
	// Group membership claims from the authentication token.
	Groups []string `json:"groups,omitzero"`
	// Issuer-specific claims forwarded opaquely from STS (e.g. GHA OIDC assertion
	// claims). Mapped to `context.actor_claims.issuer_claims` in Cedar. All fields are
	// optional and issuer-dependent. See issuer documentation for available claim
	// fields.
	IssuerClaims map[string]any `json:"issuer_claims,omitzero"`
	ExtraFields  map[string]any `json:"-"`
	paramObj
}

func (r ZonePolicyEvaluationNewParamsContextClaimsActingPrincipal) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicyEvaluationNewParamsContextClaimsActingPrincipal
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ZonePolicyEvaluationNewParamsContextClaimsActingPrincipal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents token claims passed into the Cedar context. Mapped to the Cedar
// `Keycard::Claims` type. Known fields (`email`, `groups`, `issuer_claims`) are
// typed; additional claims are passed through as untyped key-value pairs via
// `additionalProperties`.
type ZonePolicyEvaluationNewParamsContextClaimsActor struct {
	// Email claim from the authentication token.
	Email param.Opt[string] `json:"email,omitzero"`
	// Group membership claims from the authentication token.
	Groups []string `json:"groups,omitzero"`
	// Issuer-specific claims forwarded opaquely from STS (e.g. GHA OIDC assertion
	// claims). Mapped to `context.actor_claims.issuer_claims` in Cedar. All fields are
	// optional and issuer-dependent. See issuer documentation for available claim
	// fields.
	IssuerClaims map[string]any `json:"issuer_claims,omitzero"`
	ExtraFields  map[string]any `json:"-"`
	paramObj
}

func (r ZonePolicyEvaluationNewParamsContextClaimsActor) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicyEvaluationNewParamsContextClaimsActor
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ZonePolicyEvaluationNewParamsContextClaimsActor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents token claims passed into the Cedar context. Mapped to the Cedar
// `Keycard::Claims` type. Known fields (`email`, `groups`, `issuer_claims`) are
// typed; additional claims are passed through as untyped key-value pairs via
// `additionalProperties`.
type ZonePolicyEvaluationNewParamsContextClaimsSubject struct {
	// Email claim from the authentication token.
	Email param.Opt[string] `json:"email,omitzero"`
	// Group membership claims from the authentication token.
	Groups []string `json:"groups,omitzero"`
	// Issuer-specific claims forwarded opaquely from STS (e.g. GHA OIDC assertion
	// claims). Mapped to `context.actor_claims.issuer_claims` in Cedar. All fields are
	// optional and issuer-dependent. See issuer documentation for available claim
	// fields.
	IssuerClaims map[string]any `json:"issuer_claims,omitzero"`
	ExtraFields  map[string]any `json:"-"`
	paramObj
}

func (r ZonePolicyEvaluationNewParamsContextClaimsSubject) MarshalJSON() (data []byte, err error) {
	type shadow ZonePolicyEvaluationNewParamsContextClaimsSubject
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ZonePolicyEvaluationNewParamsContextClaimsSubject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePolicyEvaluationNewParamsSubjectType string

const (
	ZonePolicyEvaluationNewParamsSubjectTypeUser        ZonePolicyEvaluationNewParamsSubjectType = "user"
	ZonePolicyEvaluationNewParamsSubjectTypeApplication ZonePolicyEvaluationNewParamsSubjectType = "application"
)

type ZonePolicyEvaluationListRequestsParams struct {
	// Cursor for forward pagination. Returned in `Pagination.after_cursor`. Mutually
	// exclusive with `before`.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination. Returned in `Pagination.before_cursor`. Mutually
	// exclusive with `after`.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Filter by shadow-evaluation presence. When `true`, returns only requests that
	// had a shadow policy set version (`shadow_psv_id IS NOT NULL`). When `false`,
	// returns only requests without one. Omit to return all.
	FilterHasShadow param.Opt[bool] `query:"filter[has_shadow],omitzero" json:"-"`
	// Maximum number of items to return per page.
	Limit            param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	XAPIVersion      param.Opt[string] `header:"X-API-Version,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	// Opt-in to additional response fields (`total_count`, `user`). Repeatable;
	// matches the `expand[]` convention used across the Keycard API.
	Expand ZonePolicyEvaluationListRequestsParamsExpandUnion `query:"expand[],omitzero" json:"-"`
	// Filter on `active_decision`. Repeatable; repeated instances OR across values.
	// See `FilterValues` in the shared spec for the full wire convention.
	//
	// Allowed values: `allow`, `deny`. Unknown values return 400 with the list of
	// allowed values. Comma-separated single values are rejected with a 400 pointing
	// at the repeated-parameter OR form.
	//
	// Note: the allowed-value enum is enforced in the handler (not as an OpenAPI
	// `items.enum`) so the server can return a targeted error for the comma-AND form
	// instead of a generic "not in allowed values" response.
	FilterActiveDecision []string `query:"filter[active_decision],omitzero" json:"-"`
	// Filter on `authority_type`. Repeatable; repeated instances OR across values. See
	// `FilterValues` in the shared spec for the full wire convention. Rows with `NULL`
	// authority_type are excluded when this filter is present.
	//
	// Allowed values: `delegation`, `impersonation`. Unknown values return 400 with
	// the list of allowed values. Comma-separated single values are rejected with a
	// 400 pointing at the repeated-parameter OR form.
	//
	// Note: the allowed-value enum is enforced in the handler.
	FilterAuthorityType []string `query:"filter[authority_type],omitzero" json:"-"`
	// Filter on `principal_id` by exact match. Repeatable; repeated instances OR
	// across values. See `FilterValues` in the shared spec for the full wire
	// convention. Comma-separated single values are rejected with a 400 pointing at
	// the repeated-parameter OR form.
	FilterPrincipalID []string `query:"filter[principal_id],omitzero" json:"-"`
	// Filter on `principal_type`. Repeatable; repeated instances OR across values. See
	// `FilterValues` in the shared spec for the full wire convention.
	//
	// Allowed values: `user`, `application`. Unknown values return 400 with the list
	// of allowed values. Comma-separated single values are rejected with a 400
	// pointing at the repeated-parameter OR form.
	//
	// Note: the allowed-value enum is enforced in the handler (not as an OpenAPI
	// `items.enum`) so the server can return a targeted error for the comma-AND form.
	FilterPrincipalType []string `query:"filter[principal_type],omitzero" json:"-"`
	// Filter on `resource_id` by exact match. Repeatable; repeated instances OR across
	// values. See `FilterValues` in the shared spec for the full wire convention.
	// Comma-separated single values are rejected with a 400 pointing at the
	// repeated-parameter OR form.
	FilterResourceID []string `query:"filter[resource_id],omitzero" json:"-"`
	// Filter on `subject_id` by exact match. Repeatable; repeated instances OR across
	// values. Rows with `NULL` subject_id are excluded when this filter is present.
	// See `FilterValues` in the shared spec for the full wire convention.
	FilterSubjectID []string `query:"filter[subject_id],omitzero" json:"-"`
	// Filter on `subject_type`. Repeatable; repeated instances OR across values. See
	// `FilterValues` in the shared spec for the full wire convention. Rows with `NULL`
	// subject_type are excluded when this filter is present.
	//
	// Allowed values: `user`, `application`. Unknown values return 400 with the list
	// of allowed values. Comma-separated single values are rejected with a 400
	// pointing at the repeated-parameter OR form.
	//
	// Note: the allowed-value enum is enforced in the handler.
	FilterSubjectType []string `query:"filter[subject_type],omitzero" json:"-"`
	// Sort direction. Default is desc (newest first).
	//
	// Any of "asc", "desc".
	Order ZonePolicyEvaluationListRequestsParamsOrder `query:"order,omitzero" json:"-"`
	// Case-insensitive substring search on `action`. Repeatable; if multiple terms are
	// supplied they are OR-ed.
	QueryAction []string `query:"query[action],omitzero" json:"-"`
	// Field to sort by.
	//
	// Any of "created_at".
	Sort ZonePolicyEvaluationListRequestsParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ZonePolicyEvaluationListRequestsParams]'s query parameters
// as `url.Values`.
func (r ZonePolicyEvaluationListRequestsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ZonePolicyEvaluationListRequestsParamsExpandUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfZonePolicyEvaluationListRequestssExpandPdpExpandField)
	OfZonePolicyEvaluationListRequestssExpandPdpExpandField param.Opt[string] `query:",omitzero,inline"`
	OfZonePolicyEvaluationListRequestssExpandArrayItemArray []string          `query:",omitzero,inline"`
	paramUnion
}

// Opt-in response field requested via `expand[]` on list endpoints. `total_count`
// adds `pagination.total_count`; `user` adds a `*_user` object beside every
// `created_by`, `updated_by` and `archived_by` field that resolves to an
// organization user.
type ZonePolicyEvaluationListRequestsParamsExpandPdpExpandField string

const (
	ZonePolicyEvaluationListRequestsParamsExpandPdpExpandFieldTotalCount ZonePolicyEvaluationListRequestsParamsExpandPdpExpandField = "total_count"
	ZonePolicyEvaluationListRequestsParamsExpandPdpExpandFieldUser       ZonePolicyEvaluationListRequestsParamsExpandPdpExpandField = "user"
)

// Sort direction. Default is desc (newest first).
type ZonePolicyEvaluationListRequestsParamsOrder string

const (
	ZonePolicyEvaluationListRequestsParamsOrderAsc  ZonePolicyEvaluationListRequestsParamsOrder = "asc"
	ZonePolicyEvaluationListRequestsParamsOrderDesc ZonePolicyEvaluationListRequestsParamsOrder = "desc"
)

// Field to sort by.
type ZonePolicyEvaluationListRequestsParamsSort string

const (
	ZonePolicyEvaluationListRequestsParamsSortCreatedAt ZonePolicyEvaluationListRequestsParamsSort = "created_at"
)
