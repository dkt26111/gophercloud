package policytypes

import "github.com/gophercloud/gophercloud"

var apiVersion = "v1"
var apiName = "policy-types"

func policyTypeGetURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL(apiVersion, apiName, id)
}
