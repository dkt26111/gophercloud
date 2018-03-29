// +build acceptance clustering policytypes

package v1

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/clustering/v1/policytypes"
	th "github.com/gophercloud/gophercloud/testhelper"
)

func TestPolicyTypeGet(t *testing.T) {
	client, err := clients.NewClusteringV1Client()
	th.AssertNoErr(t, err)

	policyType, err := policytypes.Get(client, "senlin.policy.batch-1.0").Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, policyType)
}
