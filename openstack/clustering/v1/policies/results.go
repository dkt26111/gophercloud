package policies

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// Policy represents a clustering policy in the Openstack cloud
type Policy struct {
	CreatedAt   *time.Time             `json:"-"`
	Data        map[string]interface{} `json:"data"`
	DomainUUID  string                 `json:"domain"`
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	ProjectUUID string                 `json:"project"`
	Spec        map[string]interface{} `json:"spec"`
	Type        string                 `json:"type"`
	UpdatedAt   *time.Time             `json:"-"`
	UserUUID    string                 `json:"user"`
}

// ExtractPolicies interprets a page of results as a slice of Policy.
func ExtractPolicies(r pagination.Page) ([]Policy, error) {
	var s struct {
		Policies []Policy `json:"policies"`
	}
	err := (r.(PolicyPage)).ExtractInto(&s)
	return s.Policies, err
}

// PolicyPage contains a list page of all policies from a List call.
type PolicyPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines if a PolicyPage contains any results.
func (page PolicyPage) IsEmpty() (bool, error) {
	policies, err := ExtractPolicies(page)
	return len(policies) == 0, err
}

func (r *Policy) UnmarshalJSON(b []byte) error {
	type tmp Policy
	var s struct {
		tmp
		CreatedAt *gophercloud.JSONRFC3339MilliNoZ `json:"created_at,omitempty"`
		UpdatedAt *gophercloud.JSONRFC3339MilliNoZ `json:"updated_at,omitempty"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		fmt.Println("Error Unmarshal")
		fmt.Printf("Detail Unmarshal Error: %v", err)
		return err
	}
	*r = Policy(s.tmp)

	if s.CreatedAt != nil {
		r.CreatedAt = new(time.Time)
		*r.CreatedAt = time.Time(*s.CreatedAt)
	} else {
		r.CreatedAt = nil
	}

	if s.UpdatedAt != nil {
		r.UpdatedAt = new(time.Time)
		*r.UpdatedAt = time.Time(*s.UpdatedAt)
	} else {
		r.UpdatedAt = nil
	}

	return nil
}
