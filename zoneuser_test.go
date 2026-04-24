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

func TestZoneUserGet(t *testing.T) {
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
	_, err := client.Zones.Users.Get(
		context.TODO(),
		"id",
		keycard.ZoneUserGetParams{
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

func TestZoneUserListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Users.List(
		context.TODO(),
		"zoneId",
		keycard.ZoneUserListParams{
			After:  keycard.String("x"),
			Before: keycard.String("x"),
			Expand: keycard.ZoneUserListParamsExpandUnion{
				OfZoneUserListsExpandString: keycard.String("total_count"),
			},
			FilterEmail: keycard.ZoneUserListParamsFilterEmailUnion{
				OfString: keycard.String("dev@stainless.com"),
			},
			Limit: keycard.Int(1),
			Query: keycard.ZoneUserListParamsQueryUnion{
				OfString: keycard.String("x"),
			},
			QueryEmail: keycard.ZoneUserListParamsQueryEmailUnion{
				OfString: keycard.String("x"),
			},
			QueryIdentifier: keycard.ZoneUserListParamsQueryIdentifierUnion{
				OfString: keycard.String("x"),
			},
			QuerySubject: keycard.ZoneUserListParamsQuerySubjectUnion{
				OfString: keycard.String("x"),
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
