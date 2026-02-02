// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycardapi_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/keycard-api-go"
	"github.com/stainless-sdks/keycard-api-go/internal/testutil"
	"github.com/stainless-sdks/keycard-api-go/option"
)

func TestZoneMcpGatewayNewServerWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := keycardapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	_, err := client.Zones.McpGateways.NewServer(
		context.TODO(),
		"applicationId",
		keycardapi.ZoneMcpGatewayNewServerParams{
			ZoneID: "zoneId",
			Downstream: keycardapi.ZoneMcpGatewayNewServerParamsDownstream{
				Slug: keycardapi.String("slug"),
			},
			Upstream: keycardapi.ZoneMcpGatewayNewServerParamsUpstream{
				Identifier: "x",
				Name:       "x",
			},
			UpstreamProvider: keycardapi.ZoneMcpGatewayNewServerParamsUpstreamProvider{
				Identifier: "x",
				Name:       "x",
			},
		},
	)
	if err != nil {
		var apierr *keycardapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
