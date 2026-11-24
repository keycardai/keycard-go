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

func TestZoneApplicationRoleListWithOptionalParams(t *testing.T) {
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
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	_, err := client.Zones.Applications.Roles.List(
		context.TODO(),
		"applicationId",
		keycard.ZoneApplicationRoleListParams{
			ZoneID: "zoneId",
			After:  keycard.String("x"),
			Before: keycard.String("x"),
			Expand: keycard.ZoneApplicationRoleListParamsExpandUnion{
				OfZoneApplicationRoleListsExpandString: keycard.String("total_count"),
			},
			Limit: keycard.Int(1),
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

func TestZoneApplicationRoleAssignWithOptionalParams(t *testing.T) {
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
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	_, err := client.Zones.Applications.Roles.Assign(
		context.TODO(),
		"applicationId",
		keycard.ZoneApplicationRoleAssignParams{
			ZoneID: "zoneId",
			RoleAssignmentCreate: keycard.RoleAssignmentCreateParam{
				OwnerType:      keycard.RoleAssignmentCreateOwnerTypePlatform,
				RoleID:         keycard.String("role_id"),
				RoleIdentifier: keycard.String("role_identifier"),
				ScopeID:        keycard.String("x"),
				ScopeType:      keycard.String("x"),
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

func TestZoneApplicationRoleRevokeWithOptionalParams(t *testing.T) {
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
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	err := client.Zones.Applications.Roles.Revoke(
		context.TODO(),
		"roleId",
		keycard.ZoneApplicationRoleRevokeParams{
			ZoneID:        "zoneId",
			ApplicationID: "applicationId",
			ScopeID:       keycard.String("x"),
			ScopeType:     keycard.String("x"),
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
