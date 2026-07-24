// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycard

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"slices"

	"github.com/keycardai/keycard-go/internal/apiform"
	"github.com/keycardai/keycard-go/internal/requestconfig"
	"github.com/keycardai/keycard-go/option"
	"github.com/keycardai/keycard-go/packages/param"
)

// Per-user Policy Bundle resource. Allows clients (typically the Keycard CLI) to
// GET, PUT, and DELETE the effective Policy Set for the calling user on a zone.
// The bundle is encoded with a content-negotiated codec (currently only
// `application/vnd.keycard.policy-bundle.v1+tar+gzip`).
//
// ## Archive layout
//
// The bundle is a gzip-compressed tar archive with this logical layout:
//
// | Entry                        | Required on PUT | Notes                                                                                                                                                                                  |
// | ---------------------------- | --------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
// | `manifest.json`              | **Yes**         | See `PolicyBundleManifest`. The only source of the authoritative `schema.version`.                                                                                                     |
// | `schema.cedarschema`         | No              | Convenience snapshot of the Cedar schema. **Ignored on PUT** — the server validates policies against its own attested schema for `manifest.schema.version`. **Always present on GET.** |
// | `policies/<public_id>.cedar` | —               | One Cedar policy per file; the filename stem is the policy's public ID.                                                                                                                |
//
// Decode rules: duplicate entries and unrecognized/nested entries are rejected
// (`bundle_invalid`). On PUT the manifest's `sha` fields and `policies[]` list are
// advisory — the server recomputes every digest from the archived bytes and
// derives the policy set from the `policies/` files. On GET every digest is
// authoritative.
//
// PolicyBundleService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPolicyBundleService] method instead.
type PolicyBundleService struct {
	Options []option.RequestOption
}

// NewPolicyBundleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPolicyBundleService(opts ...option.RequestOption) (r PolicyBundleService) {
	r = PolicyBundleService{}
	r.Options = opts
	return
}

// Returns the effective Policy Bundle for the user identified by the zone-issued
// resource-scoped token. When no user-scope binding exists, one will be generated
// from the default set.
//
// The response body is a binary archive in the codec selected via the `Accept`
// header. The only codec supported today is
// `application/vnd.keycard.policy-bundle.v1+tar+gzip`. Clients SHOULD send an
// explicit `Accept` header; absent one, the server defaults to the tar+gzip codec.
//
// Supports conditional fetch via `If-None-Match`: when the supplied ETag matches
// the current bundle, the server responds `304 Not Modified` with no body.
func (r *PolicyBundleService) Get(ctx context.Context, query PolicyBundleGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(query.IfNoneMatch) {
		opts = append(opts, option.WithHeader("If-None-Match", fmt.Sprintf("%v", query.IfNoneMatch.Value)))
	}
	if !param.IsOmitted(query.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", query.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithBearerAuthSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.keycard.policy-bundle.v1+tar+gzip")}, opts...)
	path := "policy/bundle"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Accepts an edited Policy Bundle archive and applies it as the active user-scope
// PolicySetVersion for the calling user.
//
// The user's policy set is seeded from the system-default policies on first
// access, forked into customer-owned policies; a user bundle therefore contains
// only customer-owned policies. Applying an edit creates a new version of the
// affected policy, and a `new_policy` entry adds a further customer-owned policy.
// Platform-owned catalog policies are never edited in place by this operation.
//
// The request body codec is determined from `Content-Type`. The only codec
// supported today is `application/vnd.keycard.policy-bundle.v1+tar+gzip`.
//
// Supports optimistic concurrency via `If-Match`: when supplied, the server
// applies the bundle only if the supplied ETag matches the current bundle ETag;
// otherwise responds `412 Precondition Failed`.
//
// On success the server returns the materialized bundle (in the same codec) and
// its new `ETag`.
func (r *PolicyBundleService) Update(ctx context.Context, params PolicyBundleUpdateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.IfMatch) {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", params.IfMatch.Value)))
	}
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithBearerAuthSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.keycard.policy-bundle.v1+tar+gzip")}, opts...)
	path := "policy/bundle"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Archives the PolicySet for the calling user (if any), causing subsequent
// `GET /policy/bundle` requests to fall back to the default user policies.
// Idempotent: returns `204 No Content` even when no user-scope binding exists.
func (r *PolicyBundleService) Reset(ctx context.Context, body PolicyBundleResetParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", body.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithBearerAuthSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "policy/bundle"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type PolicyBundleGetParams struct {
	IfNoneMatch      param.Opt[string] `header:"If-None-Match,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

type PolicyBundleUpdateParams struct {
	// tar+gzip Policy Bundle archive. `manifest.json` is **required** (see
	// `PolicyBundleManifest`); `schema.cedarschema` is **optional and ignored** — the
	// server validates against its attested schema for `manifest.schema.version`. The
	// manifest's `policies[]` list is authoritative for the resulting set: each entry
	// must have a matching `policies/<public_id>.cedar` (or, for a `new_policy` entry,
	// `policies/<new_policy>.cedar`) member, and a member with no manifest entry is
	// dropped. Only the `sha` fields are advisory and recomputed server-side.
	// Duplicate or unrecognized entries are rejected with `bundle_invalid`. See the
	// **PolicyBundle** tag for the layout.
	Body             io.Reader
	IfMatch          param.Opt[string] `header:"If-Match,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r PolicyBundleUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r.Body, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type PolicyBundleResetParams struct {
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}
