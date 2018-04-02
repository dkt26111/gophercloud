package policies

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// ListOptsBuilder Builder.
type ListOptsBuilder interface {
	ToPolicyListQuery() (string, error)
}

// ListOpts params
type ListOpts struct {
	// Limit instructs List to refrain from sending excessively large lists of profiles
	Limit int `q:"limit"`

	// Marker and Limit control paging. Marker instructs List where to start listing from.
	Marker string `q:"marker"`

	// Marker and Limit control paging. Marker instructs List where to start listing from.
	Sort string `q:"sort"`

	// GlobalProject indicates whether to include resources for all projects or resources for the current project
	GlobalProject bool `q:"global_project"`

	// Name to filter the response by the specified name property of the object
	Name string `q:"name"`

	// Filter the response by the specified type property of the object
	Type string `q:"type"`
}

// ToPolicyListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToPolicyListQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	return q.String(), err
}

// ListDetail instructs OpenStack to provide a list of policies.
func ListDetail(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := policyListURL(client)
	if opts != nil {
		query, err := opts.ToPolicyListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return PolicyPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

