// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keycard

import (
	"github.com/keycardai/keycard-go/option"
)

// OrganizationUserService contains methods and other services that help with
// interacting with the keycard-api API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationUserService] method instead.
type OrganizationUserService struct {
	Options []option.RequestOption
}

// NewOrganizationUserService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationUserService(opts ...option.RequestOption) (r OrganizationUserService) {
	r = OrganizationUserService{}
	r.Options = opts
	return
}

// User's role in the organization
type OrganizationRole string

const (
	OrganizationRoleOrgAdmin  OrganizationRole = "org_admin"
	OrganizationRoleOrgMember OrganizationRole = "org_member"
	OrganizationRoleOrgViewer OrganizationRole = "org_viewer"
)
