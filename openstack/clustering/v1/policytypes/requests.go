package policytypes

import (
	"github.com/gophercloud/gophercloud"
)

// Get makes a request against the API to get details for a policy type
func Get(client *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(policyTypeGetURL(client, id), &r.Body,
		&gophercloud.RequestOpts{OkCodes: []int{200}})
	return
}
