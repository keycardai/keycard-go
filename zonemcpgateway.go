// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycardapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/keycard-api-go/internal/apijson"
	"github.com/stainless-sdks/keycard-api-go/internal/requestconfig"
	"github.com/stainless-sdks/keycard-api-go/option"
	"github.com/stainless-sdks/keycard-api-go/packages/param"
	"github.com/stainless-sdks/keycard-api-go/packages/respjson"
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
func (r *ZoneMcpGatewayService) NewServer(ctx context.Context, applicationID string, params ZoneMcpGatewayNewServerParams, opts ...option.RequestOption) (res *ZoneMcpGatewayNewServerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID == "" {
		err = errors.New("missing required zoneId parameter")
		return
	}
	if applicationID == "" {
		err = errors.New("missing required applicationId parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/mcp-gateways/%s/mcp-servers", params.ZoneID, applicationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Response containing the created upstream provider, upstream resource, and
// downstream resource for an MCP server
type ZoneMcpGatewayNewServerResponse struct {
	// A Resource is a system that exposes protected information or functionality. It
	// requires authentication of the requesting actor, which may be a user or
	// application, before allowing access.
	DownstreamResource Resource `json:"downstream_resource,required"`
	// A Provider is a system that supplies access to Resources and allows actors
	// (Users or Applications) to authenticate.
	UpstreamProvider Provider `json:"upstream_provider,required"`
	// A Resource is a system that exposes protected information or functionality. It
	// requires authentication of the requesting actor, which may be a user or
	// application, before allowing access.
	UpstreamResource Resource `json:"upstream_resource,required"`
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
func (r ZoneMcpGatewayNewServerResponse) RawJSON() string { return r.JSON.raw }
func (r *ZoneMcpGatewayNewServerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneMcpGatewayNewServerParams struct {
	ZoneID string `path:"zoneId,required" json:"-"`
	// Downstream MCP server config
	Downstream ZoneMcpGatewayNewServerParamsDownstream `json:"downstream,omitzero,required"`
	// Upstream MCP server config
	Upstream ZoneMcpGatewayNewServerParamsUpstream `json:"upstream,omitzero,required"`
	// Credential provider for the upstream connection
	UpstreamProvider ZoneMcpGatewayNewServerParamsUpstreamProvider `json:"upstream_provider,omitzero,required"`
	paramObj
}

func (r ZoneMcpGatewayNewServerParams) MarshalJSON() (data []byte, err error) {
	type shadow ZoneMcpGatewayNewServerParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneMcpGatewayNewServerParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Downstream MCP server config
type ZoneMcpGatewayNewServerParamsDownstream struct {
	// URL-safe identifier, unique within the zone
	Slug param.Opt[string] `json:"slug,omitzero"`
	paramObj
}

func (r ZoneMcpGatewayNewServerParamsDownstream) MarshalJSON() (data []byte, err error) {
	type shadow ZoneMcpGatewayNewServerParamsDownstream
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneMcpGatewayNewServerParamsDownstream) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Upstream MCP server config
//
// The properties Identifier, Name are required.
type ZoneMcpGatewayNewServerParamsUpstream struct {
	// User specified identifier, unique within the zone
	Identifier string `json:"identifier,required"`
	// Human-readable name
	Name string `json:"name,required"`
	paramObj
}

func (r ZoneMcpGatewayNewServerParamsUpstream) MarshalJSON() (data []byte, err error) {
	type shadow ZoneMcpGatewayNewServerParamsUpstream
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneMcpGatewayNewServerParamsUpstream) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credential provider for the upstream connection
//
// The properties Identifier, Name are required.
type ZoneMcpGatewayNewServerParamsUpstreamProvider struct {
	// User specified identifier, unique within the zone
	Identifier string `json:"identifier,required"`
	// Human-readable name
	Name string `json:"name,required"`
	paramObj
}

func (r ZoneMcpGatewayNewServerParamsUpstreamProvider) MarshalJSON() (data []byte, err error) {
	type shadow ZoneMcpGatewayNewServerParamsUpstreamProvider
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ZoneMcpGatewayNewServerParamsUpstreamProvider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
