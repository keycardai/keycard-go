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

func TestZoneNewWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	_, err := client.Zones.New(context.TODO(), keycard.ZoneNewParams{
		Name:                         "x",
		DefaultMcpGatewayApplication: keycard.Bool(true),
		Description:                  keycard.String("description"),
		EncryptionKey: keycard.EncryptionKeyAwsKmsConfigParam{
			Arn:  "x",
			Type: keycard.EncryptionKeyAwsKmsConfigTypeAws,
		},
		LoginFlow: keycard.ZoneNewParamsLoginFlowDefault,
		Protocols: keycard.ZoneNewParamsProtocols{
			Oauth2: keycard.ZoneNewParamsProtocolsOauth2{
				DcrEnabled:   keycard.Bool(true),
				PkceRequired: keycard.Bool(true),
			},
		},
		RequiresInvitation: keycard.Bool(true),
	})
	if err != nil {
		var apierr *keycard.Error
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
	client := keycard.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	_, err := client.Zones.Get(
		context.TODO(),
		"zoneId",
		keycard.ZoneGetParams{
			Expand: keycard.ZoneGetParamsExpandUnion{
				OfZoneGetsExpandString: keycard.String("permissions"),
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

func TestZoneUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	_, err := client.Zones.Update(
		context.TODO(),
		"zoneId",
		keycard.ZoneUpdateParams{
			DefaultMcpGatewayApplicationID: keycard.String("default_mcp_gateway_application_id"),
			Description:                    keycard.String("description"),
			EncryptionKey: keycard.ZoneUpdateParamsEncryptionKey{
				Arn:  "x",
				Type: "aws",
			},
			LoginFlow: keycard.ZoneUpdateParamsLoginFlowDefault,
			Name:      keycard.String("x"),
			Protocols: keycard.ZoneUpdateParamsProtocols{
				Oauth2: keycard.ZoneUpdateParamsProtocolsOauth2{
					DcrEnabled:   keycard.Bool(true),
					PkceRequired: keycard.Bool(true),
				},
			},
			RequiresInvitation:     keycard.Bool(true),
			UserIdentityProviderID: keycard.String("user_identity_provider_id"),
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

func TestZoneListWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	_, err := client.Zones.List(context.TODO(), keycard.ZoneListParams{
		After:  keycard.String("x"),
		Before: keycard.String("x"),
		Cursor: keycard.String("cursor"),
		Expand: keycard.ZoneListParamsExpandUnion{
			OfZoneListsExpandString: keycard.String("total_count"),
		},
		Limit: keycard.Int(1),
		Slug:  keycard.String("slug"),
	})
	if err != nil {
		var apierr *keycard.Error
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
	client := keycard.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	err := client.Zones.Delete(context.TODO(), "zoneId")
	if err != nil {
		var apierr *keycard.Error
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
	client := keycard.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	err := client.Zones.DeleteMcpServer(
		context.TODO(),
		"downstreamId",
		keycard.ZoneDeleteMcpServerParams{
			ZoneID: "zoneId",
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

func TestZoneListSessionResourceAccessWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	_, err := client.Zones.ListSessionResourceAccess(
		context.TODO(),
		"zoneId",
		keycard.ZoneListSessionResourceAccessParams{
			After:  keycard.String("x"),
			Before: keycard.String("x"),
			Expand: keycard.ZoneListSessionResourceAccessParamsExpandUnion{
				OfZoneListSessionResourceAccesssExpandString: keycard.String("total_count"),
			},
			Limit:          keycard.Int(1),
			ResourceID:     keycard.String("resource_id"),
			RollupChildren: keycard.Bool(true),
			SessionID:      keycard.String("session_id"),
			UserID:         keycard.String("user_id"),
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
