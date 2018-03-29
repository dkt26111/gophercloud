/*
Package policytypes lists all policy types and shows details for a policy type from the OpenStack
Clustering Service.

Example of get policy type details

    policyTypeName := "senlin.policy.affinity-1.0"
    policyTypeDetail, err := policyTypes.Get(clusteringClient, policyTypeName).Extract()
    if err != nil {
        panic(err)
    }
    fmt.Printf("%+v\n", policyTypeDetail)
*/
package policytypes
