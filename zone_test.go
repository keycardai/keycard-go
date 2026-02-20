// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycardapi_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/keycardai/keycard-go"
	"github.com/keycardai/keycard-go/internal/testutil"
	"github.com/keycardai/keycard-go/option"
)

func TestZoneNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Zones.New(context.TODO(), keycardapi.ZoneNewParams{
		Name:                         "x",
		DefaultMcpGatewayApplication: keycardapi.Bool(true),
		Description:                  keycardapi.String("description"),
		DirectoryOpenSignupsEnabled:  keycardapi.Bool(true),
		EncryptionKey: keycardapi.EncryptionKeyAwsKmsConfigParam{
			Arn:  "x",
			Type: keycardapi.EncryptionKeyAwsKmsConfigTypeAws,
		},
		LoginFlow: keycardapi.ZoneNewParamsLoginFlowDefault,
		Protocols: keycardapi.ZoneNewParamsProtocols{
			Oauth2: keycardapi.ZoneNewParamsProtocolsOauth2{
				DcrEnabled:   keycardapi.Bool(true),
				PkceRequired: keycardapi.Bool(true),
			},
		},
	})
	if err != nil {
		var apierr *keycardapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneGetWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Zones.Get(
		context.TODO(),
		"zoneId",
		keycardapi.ZoneGetParams{
			Expand: keycardapi.ZoneGetParamsExpandUnion{
				OfZoneGetsExpandString: keycardapi.String("permissions"),
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

func TestZoneUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Zones.Update(
		context.TODO(),
		"zoneId",
		keycardapi.ZoneUpdateParams{
			DefaultMcpGatewayApplicationID: keycardapi.String("default_mcp_gateway_application_id"),
			Description:                    keycardapi.String("description"),
			DirectoryOpenSignupsEnabled:    keycardapi.Bool(true),
			EncryptionKey: keycardapi.ZoneUpdateParamsEncryptionKey{
				Arn:  "x",
				Type: "aws",
			},
			LoginFlow: keycardapi.ZoneUpdateParamsLoginFlowDefault,
			Name:      keycardapi.String("x"),
			Protocols: keycardapi.ZoneUpdateParamsProtocols{
				Oauth2: keycardapi.ZoneUpdateParamsProtocolsOauth2{
					DcrEnabled:   keycardapi.Bool(true),
					PkceRequired: keycardapi.Bool(true),
				},
			},
			UserIdentityProviderID: keycardapi.String("user_identity_provider_id"),
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

func TestZoneListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Zones.List(context.TODO(), keycardapi.ZoneListParams{
		After:  keycardapi.String("x"),
		Before: keycardapi.String("x"),
		Cursor: keycardapi.String("cursor"),
		Expand: keycardapi.ZoneListParamsExpandUnion{
			OfZoneListsExpandString: keycardapi.String("total_count"),
		},
		Limit: keycardapi.Int(1),
		Slug:  keycardapi.String("slug"),
	})
	if err != nil {
		var apierr *keycardapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	err := client.Zones.Delete(context.TODO(), "zoneId")
	if err != nil {
		var apierr *keycardapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneDeleteMcpServer(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	err := client.Zones.DeleteMcpServer(
		context.TODO(),
		"downstreamId",
		keycardapi.ZoneDeleteMcpServerParams{
			ZoneID: "zoneId",
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

func TestZoneListSessionResourceAccessWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Zones.ListSessionResourceAccess(
		context.TODO(),
		"zoneId",
		keycardapi.ZoneListSessionResourceAccessParams{
			After:  keycardapi.String("x"),
			Before: keycardapi.String("x"),
			Expand: keycardapi.ZoneListSessionResourceAccessParamsExpandUnion{
				OfZoneListSessionResourceAccesssExpandString: keycardapi.String("total_count"),
			},
			Limit:      keycardapi.Int(1),
			ResourceID: keycardapi.String("resource_id"),
			SessionID:  keycardapi.String("session_id"),
			UserID:     keycardapi.String("user_id"),
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
