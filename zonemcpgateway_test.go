// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycard_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/keycardai/keycard-go"
	"github.com/keycardai/keycard-go/internal/testutil"
	"github.com/keycardai/keycard-go/option"
)

func TestZoneMcpGatewayNewMcpServerWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := keycard.NewClient(
		option.WithBaseURL(baseURL),
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	_, err := client.Zones.McpGateways.NewMcpServer(
		context.TODO(),
		"applicationId",
		keycard.ZoneMcpGatewayNewMcpServerParams{
			ZoneID: "zoneId",
			Downstream: keycard.ZoneMcpGatewayNewMcpServerParamsDownstream{
				Slug: keycard.String("slug"),
			},
			Upstream: keycard.ZoneMcpGatewayNewMcpServerParamsUpstream{
				Identifier: "x",
				Name:       "x",
			},
			UpstreamProvider: keycard.ZoneMcpGatewayNewMcpServerParamsUpstreamProvider{
				Identifier: "x",
				Name:       "x",
			},
		},
	)
	if err != nil {
		var apierr *keycard.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
