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
	"github.com/keycardai/keycard-go/internal/requestconfig"
	"github.com/keycardai/keycard-go/option"
	"github.com/keycardai/keycard-go/packages/param"
	"github.com/keycardai/keycard-go/packages/respjson"
)

// ZoneMcpGatewayService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneMcpGatewayService] method instead.
type ZoneMcpGatewayService struct {
	Options []option.RequestOption
}

// NewZoneMcpGatewayService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneMcpGatewayService(opts ...option.RequestOption) (r ZoneMcpGatewayService) {
	r = ZoneMcpGatewayService{}
	r.Options = opts
	return
}

// Creates all resources required to access an MCP server through an MCP gateway
func (r *ZoneMcpGatewayService) NewMcpServer(ctx context.Context, applicationID string, params ZoneMcpGatewayNewMcpServerParams, opts ...option.RequestOption) (res *ZoneMcpGatewayNewMcpServerResponse, err error) {
	var preClientOpts = []option.RequestOption{requestconfig.WithBearerAuthSecurity()}
	opts = slices.Concat(preClientOpts, r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return nil, err
	}
	if applicationID == "" {
		err = errors.New("missing required applicationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/mcp-gateways/%s/mcp-servers", url.PathEscape(params.ZoneID), url.PathEscape(applicationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Response containing the created upstream provider, upstream resource, and
// downstream resource for an MCP server
type ZoneMcpGatewayNewMcpServerResponse struct {
	// A Resource is a system that exposes protected information or functionality. It
	// requires authentication of the requesting actor, which may be a user or
	// application, before allowing access.
	DownstreamResource Resource `json:"downstream_resource" api:"required"`
	// A Provider is a system that supplies access to Resources and allows actors
	// (Users or Applications) to authenticate.
	UpstreamProvider Provider `json:"upstream_provider" api:"required"`
	// A Resource is a system that exposes protected information or functionality. It
	// requires authentication of the requesting actor, which may be a user or
	// application, before allowing access.
	UpstreamResource Resource `json:"upstream_resource" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DownstreamResource respjson.Field
		UpstreamProvider   respjson.Field
		UpstreamResource   respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ZoneMcpGatewayNewMcpServerResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneMcpGatewayNewMcpServerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneMcpGatewayNewMcpServerParams struct {
	ZoneID string `path:"zoneId" api:"required" json:"-"`
	// Downstream MCP server config
	Downstream ZoneMcpGatewayNewMcpServerParamsDownstream `json:"downstream,omitzero" api:"required"`
	// Upstream MCP server config
	Upstream ZoneMcpGatewayNewMcpServerParamsUpstream `json:"upstream,omitzero" api:"required"`
	// Credential provider for the upstream connection
	UpstreamProvider ZoneMcpGatewayNewMcpServerParamsUpstreamProvider `json:"upstream_provider,omitzero" api:"required"`
	paramObj
}

func (r ZoneMcpGatewayNewMcpServerParams) MarshalJSON() (data []byte, err error) {
	type shadow ZoneMcpGatewayNewMcpServerParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneMcpGatewayNewMcpServerParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Downstream MCP server config
type ZoneMcpGatewayNewMcpServerParamsDownstream struct {
	// URL-safe identifier, unique within the zone
	Slug param.Opt[string] `json:"slug,omitzero"`
	paramObj
}

func (r ZoneMcpGatewayNewMcpServerParamsDownstream) MarshalJSON() (data []byte, err error) {
	type shadow ZoneMcpGatewayNewMcpServerParamsDownstream
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneMcpGatewayNewMcpServerParamsDownstream) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Upstream MCP server config
//
// The properties Identifier, Name are required.
type ZoneMcpGatewayNewMcpServerParamsUpstream struct {
	// User specified identifier, unique within the zone
	Identifier string `json:"identifier" api:"required"`
	// Human-readable name
	Name string `json:"name" api:"required"`
	paramObj
}

func (r ZoneMcpGatewayNewMcpServerParamsUpstream) MarshalJSON() (data []byte, err error) {
	type shadow ZoneMcpGatewayNewMcpServerParamsUpstream
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneMcpGatewayNewMcpServerParamsUpstream) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credential provider for the upstream connection
//
// The properties Identifier, Name are required.
type ZoneMcpGatewayNewMcpServerParamsUpstreamProvider struct {
	// User specified identifier, unique within the zone
	Identifier string `json:"identifier" api:"required"`
	// Human-readable name
	Name string `json:"name" api:"required"`
	paramObj
}

func (r ZoneMcpGatewayNewMcpServerParamsUpstreamProvider) MarshalJSON() (data []byte, err error) {
	type shadow ZoneMcpGatewayNewMcpServerParamsUpstreamProvider
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneMcpGatewayNewMcpServerParamsUpstreamProvider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
