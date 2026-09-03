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

func TestZoneGroupNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Groups.New(
		context.TODO(),
		"zoneId",
		keycard.ZoneGroupNewParams{
			GroupCreate: keycard.GroupCreateParam{
				Name:       "x",
				Identifier: keycard.String("x"),
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

func TestZoneGroupGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Groups.Get(
		context.TODO(),
		"groupId",
		keycard.ZoneGroupGetParams{
			ZoneID: "zoneId",
			Expand: keycard.ZoneGroupGetParamsExpandUnion{
				OfZoneGroupGetsExpandString: keycard.String("member_count"),
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

func TestZoneGroupUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Groups.Update(
		context.TODO(),
		"groupId",
		keycard.ZoneGroupUpdateParams{
			ZoneID: "zoneId",
			GroupUpdate: keycard.GroupUpdateParam{
				Identifier: keycard.String("x"),
				Name:       keycard.String("x"),
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

func TestZoneGroupListWithOptionalParams(t *testing.T) {
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
	_, err := client.Zones.Groups.List(
		context.TODO(),
		"zoneId",
		keycard.ZoneGroupListParams{
			After:  keycard.String("x"),
			Before: keycard.String("x"),
			Expand: keycard.ZoneGroupListParamsExpandUnion{
				OfZoneGroupListsExpandString: keycard.String("total_count"),
			},
			FilterID: keycard.ZoneGroupListParamsFilterIDUnion{
				OfString: keycard.String("string"),
			},
			FilterIdentifier: keycard.ZoneGroupListParamsFilterIdentifierUnion{
				OfString: keycard.String("string"),
			},
			Limit: keycard.Int(1),
			Query: keycard.ZoneGroupListParamsQueryUnion{
				OfString: keycard.String("x"),
			},
			Sort: keycard.String("-identifier,\t\r\r \tidentifier,\n\f\t\f\ncreated_at"),
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

func TestZoneGroupDelete(t *testing.T) {
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
	err := client.Zones.Groups.Delete(
		context.TODO(),
		"groupId",
		keycard.ZoneGroupDeleteParams{
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
