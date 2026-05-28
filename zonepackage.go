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

// ZonePackageService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZonePackageService] method instead.
type ZonePackageService struct {
	Options  []option.RequestOption
	Versions ZonePackageVersionService
}

// NewZonePackageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZonePackageService(opts ...option.RequestOption) (r ZonePackageService) {
	r = ZonePackageService{}
	r.Options = opts
	r.Versions = NewZonePackageVersionService(opts...)
	return
}

// Get a zone package
func (r *ZonePackageService) Get(ctx context.Context, packageID string, params ZonePackageGetParams, opts ...option.RequestOption) (res *Package, err error) {
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
	path := fmt.Sprintf("zones/%s/packages/%s", url.PathEscape(params.ZoneID), url.PathEscape(packageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List zone packages
func (r *ZonePackageService) List(ctx context.Context, zoneID string, params ZonePackageListParams, opts ...option.RequestOption) (res *PackageList, err error) {
	if !param.IsOmitted(params.XClientRequestID) {
		opts = append(opts, option.WithHeader("X-Client-Request-ID", fmt.Sprintf("%v", params.XClientRequestID.Value)))
	}
	var preClientOpts = []option.RequestOption{requestconfig.WithSecurity(requestconfig.Security{})}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	opts = append([]option.RequestOption{option.WithBaseURL("/")}, opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/packages", url.PathEscape(zoneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Get the zone package draft
func (r *ZonePackageService) GetDraft(ctx context.Context, packageID string, params ZonePackageGetDraftParams, opts ...option.RequestOption) (res *PackageDraft, err error) {
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
	path := fmt.Sprintf("zones/%s/packages/%s/draft", url.PathEscape(params.ZoneID), url.PathEscape(packageID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Computed input state for a package — derived at response time from the package
// kind's schema and the package's input binding. Not stored.
//
// `effective_schema` is the full input schema (kind + binding required constraints
// merged). `effective_bindings` resolves the CEL binding to show actual static
// values and `{"$input": "path"}` references for install-provided fields.
type InputState struct {
	EffectiveBindings map[string]any `json:"effective_bindings"`
	// A subset of JSON Schema 2020-12 used to describe package input and output
	// shapes.
	//
	// Supported keywords:
	//
	//   - Structural: `type`, `properties`, `required`, `items`, `additionalProperties`
	//   - Annotations: `title`, `description`, `default`, `readOnly`, `writeOnly`
	//   - Constraints: `pattern`, `minLength`, `maxLength`, `minimum`, `maximum`,
	//     `minItems`, `maxItems`, `enum`, `const`, `format`
	//
	// Intentionally unsupported (reject at release time rather than silently ignore):
	//
	// - Schema combinators: `allOf`, `anyOf`, `oneOf`, `not`
	// - References: `$ref`, `$dynamicRef`
	// - `patternProperties`, `propertyNames`, `unevaluatedProperties`
	// - Custom vocabularies and `$vocabulary`
	//
	// Dialect: JSON Schema 2020-12 (implied — authors do not include `$schema`).
	EffectiveSchema InputStateEffectiveSchema `json:"effective_schema"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EffectiveBindings respjson.Field
		EffectiveSchema   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InputState) RawJSON() string { return r.JSON.raw }
func (r *InputState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A subset of JSON Schema 2020-12 used to describe package input and output
// shapes.
//
// Supported keywords:
//
//   - Structural: `type`, `properties`, `required`, `items`, `additionalProperties`
//   - Annotations: `title`, `description`, `default`, `readOnly`, `writeOnly`
//   - Constraints: `pattern`, `minLength`, `maxLength`, `minimum`, `maximum`,
//     `minItems`, `maxItems`, `enum`, `const`, `format`
//
// Intentionally unsupported (reject at release time rather than silently ignore):
//
// - Schema combinators: `allOf`, `anyOf`, `oneOf`, `not`
// - References: `$ref`, `$dynamicRef`
// - `patternProperties`, `propertyNames`, `unevaluatedProperties`
// - Custom vocabularies and `$vocabulary`
//
// Dialect: JSON Schema 2020-12 (implied — authors do not include `$schema`).
type InputStateEffectiveSchema struct {
	// Schema for properties not named in `properties`.
	AdditionalProperties any `json:"additionalProperties"`
	// Constant allowed value.
	Const any `json:"const"`
	// Default value (annotation).
	Default any `json:"default"`
	// Human-readable description (annotation).
	Description string `json:"description"`
	// Enumerated allowed values.
	Enum []any `json:"enum"`
	// Format hint (e.g., "uri", "uuid", "email", "date-time").
	Format string `json:"format"`
	// Schema for array items.
	Items     any     `json:"items"`
	Maximum   float64 `json:"maximum"`
	MaxItems  int64   `json:"maxItems"`
	MaxLength int64   `json:"maxLength"`
	Minimum   float64 `json:"minimum"`
	MinItems  int64   `json:"minItems"`
	MinLength int64   `json:"minLength"`
	Pattern   string  `json:"pattern"`
	// Property schemas, keyed by property name.
	Properties any `json:"properties"`
	// Read-only hint — server-populated, ignored on write.
	ReadOnly bool `json:"readOnly"`
	// Names of required properties.
	Required []string `json:"required"`
	// Human-readable title (annotation).
	Title string `json:"title"`
	// The `type` keyword in JSON Schema 2020-12.
	//
	// Any of "object", "array", "string", "integer", "number", "boolean", "null".
	Type string `json:"type"`
	// Write-only hint (passwords, secrets) — never returned on read.
	WriteOnly bool `json:"writeOnly"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalProperties respjson.Field
		Const                respjson.Field
		Default              respjson.Field
		Description          respjson.Field
		Enum                 respjson.Field
		Format               respjson.Field
		Items                respjson.Field
		Maximum              respjson.Field
		MaxItems             respjson.Field
		MaxLength            respjson.Field
		Minimum              respjson.Field
		MinItems             respjson.Field
		MinLength            respjson.Field
		Pattern              respjson.Field
		Properties           respjson.Field
		ReadOnly             respjson.Field
		Required             respjson.Field
		Title                respjson.Field
		Type                 respjson.Field
		WriteOnly            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InputStateEffectiveSchema) RawJSON() string { return r.JSON.raw }
func (r *InputStateEffectiveSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Package struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	Kind      string    `json:"kind" api:"required"`
	Name      string    `json:"name" api:"required"`
	// Server-populated URL-friendly identifier.
	Slug           string         `json:"slug" api:"required"`
	UpdatedAt      time.Time      `json:"updated_at" api:"required" format:"date-time"`
	CurrentVersion PackageVersion `json:"current_version"`
	Description    string         `json:"description"`
	Draft          PackageDraft   `json:"draft"`
	IconURL        string         `json:"icon_url"`
	// Computed input state for a package — derived at response time from the package
	// kind's schema and the package's input binding. Not stored.
	//
	// `effective_schema` is the full input schema (kind + binding required constraints
	// merged). `effective_bindings` resolves the CEL binding to show actual static
	// values and `{"$input": "path"}` references for install-provided fields.
	InputState InputState `json:"input_state"`
	// Input binding for a package.
	//
	// `schema` constrains install-level inputs. `bindings` is a CEL expression that
	// assembles the flat input map — static values are CEL literals, install-provided
	// values are `pkg.inputs.X` references. Evaluated at provisioning time to produce
	// the `entities.inputs` map for entity bindings.
	Inputs PackageInputBinding `json:"inputs"`
	Links  []PackageLink       `json:"links"`
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
	// Provenance info for a package originating from an ancestor catalog.
	Source PackageSource `json:"source"`
	Tags   []string      `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		Kind           respjson.Field
		Name           respjson.Field
		Slug           respjson.Field
		UpdatedAt      respjson.Field
		CurrentVersion respjson.Field
		Description    respjson.Field
		Draft          respjson.Field
		IconURL        respjson.Field
		InputState     respjson.Field
		Inputs         respjson.Field
		Links          respjson.Field
		Outputs        respjson.Field
		Properties     respjson.Field
		Source         respjson.Field
		Tags           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Package) RawJSON() string { return r.JSON.raw }
func (r *Package) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A directed, typed relationship from one entity (the subject) to another (the
// target).
//
// Follows the structure of RFC 7033 JRD link objects, adapted for intra-graph
// entity references. The subject is the entity whose `links` array contains this
// link.
type PackageLink struct {
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
func (r PackageLink) RawJSON() string { return r.JSON.raw }
func (r *PackageLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PackageDraft struct {
	ID          string    `json:"id" api:"required"`
	ManifestSha string    `json:"manifest_sha" api:"required"`
	Name        string    `json:"name" api:"required"`
	UpdatedAt   time.Time `json:"updated_at" api:"required" format:"date-time"`
	Description string    `json:"description"`
	IconURL     string    `json:"icon_url"`
	// Input binding for a package.
	//
	// `schema` constrains install-level inputs. `bindings` is a CEL expression that
	// assembles the flat input map — static values are CEL literals, install-provided
	// values are `pkg.inputs.X` references. Evaluated at provisioning time to produce
	// the `entities.inputs` map for entity bindings.
	Inputs PackageInputBinding `json:"inputs"`
	Links  []PackageDraftLink  `json:"links"`
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
		ManifestSha respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
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
func (r PackageDraft) RawJSON() string { return r.JSON.raw }
func (r *PackageDraft) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A directed, typed relationship from one entity (the subject) to another (the
// target).
//
// Follows the structure of RFC 7033 JRD link objects, adapted for intra-graph
// entity references. The subject is the entity whose `links` array contains this
// link.
type PackageDraftLink struct {
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
func (r PackageDraftLink) RawJSON() string { return r.JSON.raw }
func (r *PackageDraftLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Input binding for a package.
//
// `schema` constrains install-level inputs. `bindings` is a CEL expression that
// assembles the flat input map — static values are CEL literals, install-provided
// values are `pkg.inputs.X` references. Evaluated at provisioning time to produce
// the `entities.inputs` map for entity bindings.
type PackageInputBinding struct {
	// CEL expression assembling the flat input map from static values and
	// install-provided values (referenced via `pkg.inputs.X`).
	//
	// Scope:
	//
	// - `pkg.inputs` — install-supplied values conforming to `schema`.
	Bindings string `json:"bindings"`
	// A subset of JSON Schema 2020-12 used to describe package input and output
	// shapes.
	//
	// Supported keywords:
	//
	//   - Structural: `type`, `properties`, `required`, `items`, `additionalProperties`
	//   - Annotations: `title`, `description`, `default`, `readOnly`, `writeOnly`
	//   - Constraints: `pattern`, `minLength`, `maxLength`, `minimum`, `maximum`,
	//     `minItems`, `maxItems`, `enum`, `const`, `format`
	//
	// Intentionally unsupported (reject at release time rather than silently ignore):
	//
	// - Schema combinators: `allOf`, `anyOf`, `oneOf`, `not`
	// - References: `$ref`, `$dynamicRef`
	// - `patternProperties`, `propertyNames`, `unevaluatedProperties`
	// - Custom vocabularies and `$vocabulary`
	//
	// Dialect: JSON Schema 2020-12 (implied — authors do not include `$schema`).
	Schema PackageInputBindingSchema `json:"schema"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bindings    respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PackageInputBinding) RawJSON() string { return r.JSON.raw }
func (r *PackageInputBinding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A subset of JSON Schema 2020-12 used to describe package input and output
// shapes.
//
// Supported keywords:
//
//   - Structural: `type`, `properties`, `required`, `items`, `additionalProperties`
//   - Annotations: `title`, `description`, `default`, `readOnly`, `writeOnly`
//   - Constraints: `pattern`, `minLength`, `maxLength`, `minimum`, `maximum`,
//     `minItems`, `maxItems`, `enum`, `const`, `format`
//
// Intentionally unsupported (reject at release time rather than silently ignore):
//
// - Schema combinators: `allOf`, `anyOf`, `oneOf`, `not`
// - References: `$ref`, `$dynamicRef`
// - `patternProperties`, `propertyNames`, `unevaluatedProperties`
// - Custom vocabularies and `$vocabulary`
//
// Dialect: JSON Schema 2020-12 (implied — authors do not include `$schema`).
type PackageInputBindingSchema struct {
	// Schema for properties not named in `properties`.
	AdditionalProperties any `json:"additionalProperties"`
	// Constant allowed value.
	Const any `json:"const"`
	// Default value (annotation).
	Default any `json:"default"`
	// Human-readable description (annotation).
	Description string `json:"description"`
	// Enumerated allowed values.
	Enum []any `json:"enum"`
	// Format hint (e.g., "uri", "uuid", "email", "date-time").
	Format string `json:"format"`
	// Schema for array items.
	Items     any     `json:"items"`
	Maximum   float64 `json:"maximum"`
	MaxItems  int64   `json:"maxItems"`
	MaxLength int64   `json:"maxLength"`
	Minimum   float64 `json:"minimum"`
	MinItems  int64   `json:"minItems"`
	MinLength int64   `json:"minLength"`
	Pattern   string  `json:"pattern"`
	// Property schemas, keyed by property name.
	Properties any `json:"properties"`
	// Read-only hint — server-populated, ignored on write.
	ReadOnly bool `json:"readOnly"`
	// Names of required properties.
	Required []string `json:"required"`
	// Human-readable title (annotation).
	Title string `json:"title"`
	// The `type` keyword in JSON Schema 2020-12.
	//
	// Any of "object", "array", "string", "integer", "number", "boolean", "null".
	Type string `json:"type"`
	// Write-only hint (passwords, secrets) — never returned on read.
	WriteOnly bool `json:"writeOnly"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalProperties respjson.Field
		Const                respjson.Field
		Default              respjson.Field
		Description          respjson.Field
		Enum                 respjson.Field
		Format               respjson.Field
		Items                respjson.Field
		Maximum              respjson.Field
		MaxItems             respjson.Field
		MaxLength            respjson.Field
		Minimum              respjson.Field
		MinItems             respjson.Field
		MinLength            respjson.Field
		Pattern              respjson.Field
		Properties           respjson.Field
		ReadOnly             respjson.Field
		Required             respjson.Field
		Title                respjson.Field
		Type                 respjson.Field
		WriteOnly            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PackageInputBindingSchema) RawJSON() string { return r.JSON.raw }
func (r *PackageInputBindingSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PackageList struct {
	Items []Package `json:"items" api:"required"`
	// Cursor-based pagination metadata returned alongside a list of results
	Pagination PackageListPagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PackageList) RawJSON() string { return r.JSON.raw }
func (r *PackageList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursor-based pagination metadata returned alongside a list of results
type PackageListPagination struct {
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
func (r PackageListPagination) RawJSON() string { return r.JSON.raw }
func (r *PackageListPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Output binding for a package.
//
// `schema` describes the flat outputs surfaced on an install. `bindings` is a CEL
// expression — a map literal whose keys match `schema.properties` and whose values
// project fields out of the resolved entity graph. Evaluated after the provisioner
// has resolved all entities.
type PackageOutputBinding struct {
	// CEL expression source. Must evaluate to a map whose fields match
	// `schema.properties`.
	//
	// Scope: `entities`:
	//
	//   - `entities.inputs` — the package's input values (merged with install inputs at
	//     provisioning time).
	//   - `entities.<name>` — resolved entities in the graph, each with `href: string`
	//     and `outputs: map<string, dyn>`.
	Bindings string `json:"bindings" api:"required"`
	// A subset of JSON Schema 2020-12 used to describe package input and output
	// shapes.
	//
	// Supported keywords:
	//
	//   - Structural: `type`, `properties`, `required`, `items`, `additionalProperties`
	//   - Annotations: `title`, `description`, `default`, `readOnly`, `writeOnly`
	//   - Constraints: `pattern`, `minLength`, `maxLength`, `minimum`, `maximum`,
	//     `minItems`, `maxItems`, `enum`, `const`, `format`
	//
	// Intentionally unsupported (reject at release time rather than silently ignore):
	//
	// - Schema combinators: `allOf`, `anyOf`, `oneOf`, `not`
	// - References: `$ref`, `$dynamicRef`
	// - `patternProperties`, `propertyNames`, `unevaluatedProperties`
	// - Custom vocabularies and `$vocabulary`
	//
	// Dialect: JSON Schema 2020-12 (implied — authors do not include `$schema`).
	Schema PackageOutputBindingSchema `json:"schema" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bindings    respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PackageOutputBinding) RawJSON() string { return r.JSON.raw }
func (r *PackageOutputBinding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A subset of JSON Schema 2020-12 used to describe package input and output
// shapes.
//
// Supported keywords:
//
//   - Structural: `type`, `properties`, `required`, `items`, `additionalProperties`
//   - Annotations: `title`, `description`, `default`, `readOnly`, `writeOnly`
//   - Constraints: `pattern`, `minLength`, `maxLength`, `minimum`, `maximum`,
//     `minItems`, `maxItems`, `enum`, `const`, `format`
//
// Intentionally unsupported (reject at release time rather than silently ignore):
//
// - Schema combinators: `allOf`, `anyOf`, `oneOf`, `not`
// - References: `$ref`, `$dynamicRef`
// - `patternProperties`, `propertyNames`, `unevaluatedProperties`
// - Custom vocabularies and `$vocabulary`
//
// Dialect: JSON Schema 2020-12 (implied — authors do not include `$schema`).
type PackageOutputBindingSchema struct {
	// Schema for properties not named in `properties`.
	AdditionalProperties any `json:"additionalProperties"`
	// Constant allowed value.
	Const any `json:"const"`
	// Default value (annotation).
	Default any `json:"default"`
	// Human-readable description (annotation).
	Description string `json:"description"`
	// Enumerated allowed values.
	Enum []any `json:"enum"`
	// Format hint (e.g., "uri", "uuid", "email", "date-time").
	Format string `json:"format"`
	// Schema for array items.
	Items     any     `json:"items"`
	Maximum   float64 `json:"maximum"`
	MaxItems  int64   `json:"maxItems"`
	MaxLength int64   `json:"maxLength"`
	Minimum   float64 `json:"minimum"`
	MinItems  int64   `json:"minItems"`
	MinLength int64   `json:"minLength"`
	Pattern   string  `json:"pattern"`
	// Property schemas, keyed by property name.
	Properties any `json:"properties"`
	// Read-only hint — server-populated, ignored on write.
	ReadOnly bool `json:"readOnly"`
	// Names of required properties.
	Required []string `json:"required"`
	// Human-readable title (annotation).
	Title string `json:"title"`
	// The `type` keyword in JSON Schema 2020-12.
	//
	// Any of "object", "array", "string", "integer", "number", "boolean", "null".
	Type string `json:"type"`
	// Write-only hint (passwords, secrets) — never returned on read.
	WriteOnly bool `json:"writeOnly"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalProperties respjson.Field
		Const                respjson.Field
		Default              respjson.Field
		Description          respjson.Field
		Enum                 respjson.Field
		Format               respjson.Field
		Items                respjson.Field
		Maximum              respjson.Field
		MaxItems             respjson.Field
		MaxLength            respjson.Field
		Minimum              respjson.Field
		MinItems             respjson.Field
		MinLength            respjson.Field
		Pattern              respjson.Field
		Properties           respjson.Field
		ReadOnly             respjson.Field
		Required             respjson.Field
		Title                respjson.Field
		Type                 respjson.Field
		WriteOnly            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PackageOutputBindingSchema) RawJSON() string { return r.JSON.raw }
func (r *PackageOutputBindingSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provenance info for a package originating from an ancestor catalog.
type PackageSource struct {
	// Scope type of the catalog where the package is authored.
	//
	// Any of "global", "org", "zone".
	Scope PackageSourceScope `json:"scope" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Scope       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PackageSource) RawJSON() string { return r.JSON.raw }
func (r *PackageSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scope type of the catalog where the package is authored.
type PackageSourceScope string

const (
	PackageSourceScopeGlobal PackageSourceScope = "global"
	PackageSourceScopeOrg    PackageSourceScope = "org"
	PackageSourceScopeZone   PackageSourceScope = "zone"
)

type ZonePackageGetParams struct {
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

type ZonePackageListParams struct {
	// Cursor for forward pagination. Returned in `Pagination.after_cursor`. Mutually
	// exclusive with `before`.
	After param.Opt[string] `query:"after,omitzero" json:"-"`
	// Cursor for backward pagination. Returned in `Pagination.before_cursor`. Mutually
	// exclusive with `after`.
	Before param.Opt[string] `query:"before,omitzero" json:"-"`
	// Filter packages by slug
	FiltersSlug param.Opt[string] `query:"filters[slug],omitzero" json:"-"`
	// Filter packages by kind (comma-separated)
	Kind param.Opt[string] `query:"kind,omitzero" json:"-"`
	// Maximum number of items to return per page.
	Limit            param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ZonePackageListParams]'s query parameters as `url.Values`.
func (r ZonePackageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZonePackageGetDraftParams struct {
	ZoneID           string            `path:"zone_id" api:"required" json:"-"`
	XClientRequestID param.Opt[string] `header:"X-Client-Request-ID,omitzero" format:"uuid" json:"-"`
	paramObj
}
