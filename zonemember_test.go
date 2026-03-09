// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycard_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/keycardlabs/keycard-go"
	"github.com/keycardlabs/keycard-go/internal/testutil"
	"github.com/keycardlabs/keycard-go/option"
)

func TestZoneMemberGet(t *testing.T) {
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
	_, err := client.Zones.Members.Get(
		context.TODO(),
		"organizationUserId",
		keycard.ZoneMemberGetParams{
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

func TestZoneMemberUpdate(t *testing.T) {
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
	_, err := client.Zones.Members.Update(
		context.TODO(),
		"organizationUserId",
		keycard.ZoneMemberUpdateParams{
			ZoneID: "zoneId",
			Role:   keycard.ZoneRoleZoneManager,
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

func TestZoneMemberListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Members.List(
		context.TODO(),
		"zoneId",
		keycard.ZoneMemberListParams{
			After:  keycard.String("after"),
			Before: keycard.String("before"),
			Expand: keycard.ZoneMemberListParamsExpandUnion{
				OfZoneMemberListsExpandString: keycard.String("total_count"),
			},
			Limit: keycard.Int(1),
			Role:  keycard.ZoneMemberListParamsRoleZoneManager,
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

func TestZoneMemberDelete(t *testing.T) {
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
	err := client.Zones.Members.Delete(
		context.TODO(),
		"organizationUserId",
		keycard.ZoneMemberDeleteParams{
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

func TestZoneMemberAdd(t *testing.T) {
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
	_, err := client.Zones.Members.Add(
		context.TODO(),
		"zoneId",
		keycard.ZoneMemberAddParams{
			OrganizationUserID: "organization_user_id",
			Role:               keycard.ZoneRoleZoneManager,
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
