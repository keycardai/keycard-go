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

// Browse available packages and their versions.
//
// ZonePackageVersionService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZonePackageVersionService] method instead.
type ZonePackageVersionService struct {
	Options []option.RequestOption
}

// NewZonePackageVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZonePackageVersionService(opts ...option.RequestOption) (r ZonePackageVersionService) {
	r = ZonePackageVersionService{}
	r.Options = opts
	return
}

// Get a specific zone package version
func (r *ZonePackageVersionService) Get(ctx context.Context, versionID string, params ZonePackageVersionGetParams, opts ...option.RequestOption) (res *PackageVersion, err error) {
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
	if params.PackageID == "" {
		err = errors.New("missing required package_id parameter")
		return nil, err
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/packages/%s/versions/%s", url.PathEscape(params.ZoneID), url.PathEscape(params.PackageID), url.PathEscape(versionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List zone package versions
func (r *ZonePackageVersionService) List(ctx context.Context, packageID string, params ZonePackageVersionListParams, opts ...option.RequestOption) (res *PackageVersionList, err error) {
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
	if packageID == "" {
		err = errors.New("missing required package_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/packages/%s/versions", url.PathEscape(params.ZoneID), url.PathEscape(packageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type PackageVersion struct {
	ID          string    `json:"id" api:"required"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	ManifestSha string    `json:"manifest_sha" api:"required"`
	Name        string    `json:"name" api:"required"`
	// Any of "platform", "customer".
	OwnerType   PackageVersionOwnerType `json:"owner_type" api:"required"`
	Version     int64                   `json:"version" api:"required"`
	ArchivedAt  time.Time               `json:"archived_at" api:"nullable" format:"date-time"`
	CreatedBy   string                  `json:"created_by"`
	Description string                  `json:"description"`
	IconURL     string                  `json:"icon_url"`
	// Input binding for a package.
	//
	// `schema` constrains install-level inputs. `bindings` is a CEL expression that
	// assembles the flat input map — static values are CEL literals, install-provided
	// values are `pkg.inputs.X` references. Evaluated at provisioning time to produce
	// the `entities.inputs` map for entity bindings.
	Inputs PackageInputBinding  `json:"inputs"`
	Links  []PackageVersionLink `json:"links"`
	// Output binding for a package.
	//
	// `schema` describes the flat outputs surfaced on an install. `bindings` is a CEL
	// expression — a map literal whose keys match `schema.properties` and whose values
	// project fields out of the resolved entity graph. Evaluated after the provisioner
	// has resolved all entities.
	Outputs PackageOutputBinding `json:"outputs"`
	// Vocabulary-defined metadata properties, keyed by property URN.
	//
	// Known properties are declared with their schemas; additional properties with
	// custom URNs are permitted via `Record<unknown>`.
	//
	// Each property carries `x-subject-types` indicating which entity types it applies
	// to. Properties with `draft/` in the URN are experimental and carry
	// `x-internal: true`.
	Properties map[string]any `json:"properties"`
	Tags       []string       `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		ManifestSha respjson.Field
		Name        respjson.Field
		OwnerType   respjson.Field
		Version     respjson.Field
		ArchivedAt  respjson.Field
		CreatedBy   respjson.Field
		Description respjson.Field
		IconURL     respjson.Field
		Inputs      respjson.Field
		Links       respjson.Field
		Outputs     respjson.Field
		Properties  respjson.Field
		Tags        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PackageVersion) RawJSON() string { return r.JSON.raw }
func (r *PackageVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PackageVersionOwnerType string

const (
	PackageVersionOwnerTypePlatform PackageVersionOwnerType = "platform"
	PackageVersionOwnerTypeCustomer PackageVersionOwnerType = "customer"
)

// A directed, typed relationship from one entity (the subject) to another (the
// target).
//
// Follows the structure of RFC 7033 JRD link objects, adapted for intra-graph
// entity references. The subject is the entity whose `links` array contains this
// link.
type PackageVersionLink struct {
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
func (r PackageVersionLink) RawJSON() string { return r.JSON.raw }
func (r *PackageVersionLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PackageVersionList struct {
	Items []PackageVersion `json:"items" api:"required"`
	// Cursor-based pagination metadata returned alongside a list of results
	Pagination PackageVersionListPagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PackageVersionList) RawJSON() string { return r.JSON.raw }
func (r *PackageVersionList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination metadata returned alongside a list of results
type PackageVersionListPagination struct {
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
func (r PackageVersionListPagination) RawJSON() string { return r.JSON.raw }
func (r *PackageVersionListPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePackageVersionGetParams struct {
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	PackageID        string            `path:"package_id" api:"required" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ZonePackageVersionListParams struct {
	ZoneID string `path:"zone_id" api:"required" json:"-"`
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

// URLQuery serializes [ZonePackageVersionListParams]'s query parameters as
// `url.Values`.
func (r ZonePackageVersionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
