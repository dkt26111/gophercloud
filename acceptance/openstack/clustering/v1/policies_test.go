// +build acceptance clustering policies

package v1

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/clustering/v1/policies"
	th "github.com/gophercloud/gophercloud/testhelper"
)

func TestPolicyList(t *testing.T) {
	client, err := clients.NewClusteringV1Client()
	th.AssertNoErr(t, err)

	listOpts := policies.ListOpts{
		Limit: 1,
	}

	allPages, err := policies.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	allPolicies, err := policies.ExtractPolicies(allPages)
	th.AssertNoErr(t, err)

	for _, v := range allPolicies {
		tools.PrintResource(t, v)

		if v.CreatedAt != nil {
			t.Log("Created at: " + (*v.CreatedAt).String())
		}

		if v.UpdatedAt != nil {
			t.Log("Updated at: " + (*v.UpdatedAt).String())
		}
	}
}
