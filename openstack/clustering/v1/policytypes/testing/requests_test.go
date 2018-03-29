package testing

import (
	"testing"

	"github.com/gophercloud/gophercloud/openstack/clustering/v1/policytypes"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

func TestGetPolicyType(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandlePolicyTypeGet(t)

	actual, err := policytypes.Get(fake.ServiceClient(), "1234").Extract()
	th.AssertNoErr(t, err)

	th.AssertDeepEquals(t, ExpectedPolicyTypeDetail, actual)
}
