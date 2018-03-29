package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/openstack/clustering/v1/policytypes"
	th "github.com/gophercloud/gophercloud/testhelper"
	fake "github.com/gophercloud/gophercloud/testhelper/client"
)

const PolicyTypeDetailBody = `
{
    "policy_type": {
        "name": "senlin.policy.affinity-1.0",
        "schema": {
            "availability_zone": {
                "description": "Name of the availability zone to place the nodes.",
                "required": false,
                "type": "String",
                "updatable": false
            },
            "enable_drs_extension": {
                "default": false,
                "description": "Enable vSphere DRS extension.",
                "required": false,
                "type": "Boolean",
                "updatable": false
            },
            "servergroup": {
                "description": "Properties of the VM server group",
                "required": false,
                "schema": {
                    "name": {
                        "description": "The name of the server group",
                        "required": false,
                        "type": "String",
                        "updatable": false
                    },
                    "policies": {
                        "constraints": [
                            {
                                "constraint": [
                                    "affinity",
                                    "anti-affinity"
                                ],
                                "type": "AllowedValues"
                            }
                        ],
                        "default": "anti-affinity",
                        "description": "The server group policies.",
                        "required": false,
                        "type": "String",
                        "updatable": false
                    }
                },
                "type": "Map",
                "updatable": false
            }
        },
        "support_status": {
            "1.0": [
                {
                    "status": "SUPPORTED",
                    "since": "2016.10"
                }
            ]
        }
    }
}
`

var (
	ExpectedPolicyTypeDetail = &policytypes.PolicyTypeDetail{
		Name: "senlin.policy.affinity-1.0",
		Schema: policytypes.SchemaType{
			AvailabilityZone: map[string]interface{}{
				"description": "Name of the availability zone to place the nodes.",
				"required":    false,
				"type":        "String",
				"updatable":   false,
			},
			EnableDrsExtension: map[string]interface{}{
				"default":     false,
				"description": "Enable vSphere DRS extension.",
				"required":    false,
				"type":        "Boolean",
				"updatable":   false,
			},
			Servergroup: map[string]interface{}{
				"description": "Properties of the VM server group",
				"required":    false,
				"schema": map[string]interface{}{
					"name": map[string]interface{}{
						"description": "The name of the server group",
						"required":    false,
						"type":        "String",
						"updatable":   false,
					},
					"policies": map[string]interface{}{
						"constraints": []interface{}{
							map[string]interface{}{
								"constraint": []interface{}{
									"affinity",
									"anti-affinity",
								},
								"type": "AllowedValues",
							},
						},
						"default":     "anti-affinity",
						"description": "The server group policies.",
						"required":    false,
						"type":        "String",
						"updatable":   false,
					},
				},
				"type":      "Map",
				"updatable": false,
			},
		},
		SupportStatus: map[string][]policytypes.SupportStatusType{
			"1.0": []policytypes.SupportStatusType{
				{
					Status: "SUPPORTED",
					Since:  "2016.10",
				},
			},
		},
	}
)

func HandlePolicyTypeGet(t *testing.T) {
	th.Mux.HandleFunc("/v1/policy-types/1234", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, PolicyTypeDetailBody)
	})
}
