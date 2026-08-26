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

func TestZoneGroupRoleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Groups.Roles.List(
		context.TODO(),
		"groupId",
		keycard.ZoneGroupRoleListParams{
			ZoneID: "zoneId",
			After:  keycard.String("x"),
			Before: keycard.String("x"),
			Expand: keycard.ZoneGroupRoleListParamsExpandUnion{
				OfZoneGroupRoleListsExpandString: keycard.String("total_count"),
			},
			FilterID: keycard.ZoneGroupRoleListParamsFilterIDUnion{
				OfString: keycard.String("string"),
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

func TestZoneGroupRoleAddWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Groups.Roles.Add(
		context.TODO(),
		"groupId",
		keycard.ZoneGroupRoleAddParams{
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

func TestZoneGroupRoleRemoveWithOptionalParams(t *testing.T) {
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
	err := client.Zones.Groups.Roles.Remove(
		context.TODO(),
		"roleId",
		keycard.ZoneGroupRoleRemoveParams{
			ZoneID:    "zoneId",
			GroupID:   "groupId",
			ScopeID:   keycard.String("x"),
			ScopeType: keycard.String("x"),
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
