// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SearchPoliciesRequestPolicyTypes string

const (
	SearchPoliciesRequestPolicyTypesPolicyTypeUnspecified   SearchPoliciesRequestPolicyTypes = "POLICY_TYPE_UNSPECIFIED"
	SearchPoliciesRequestPolicyTypesPolicyTypeGrant         SearchPoliciesRequestPolicyTypes = "POLICY_TYPE_GRANT"
	SearchPoliciesRequestPolicyTypesPolicyTypeRevoke        SearchPoliciesRequestPolicyTypes = "POLICY_TYPE_REVOKE"
	SearchPoliciesRequestPolicyTypesPolicyTypeCertify       SearchPoliciesRequestPolicyTypes = "POLICY_TYPE_CERTIFY"
	SearchPoliciesRequestPolicyTypesPolicyTypeAccessRequest SearchPoliciesRequestPolicyTypes = "POLICY_TYPE_ACCESS_REQUEST"
	SearchPoliciesRequestPolicyTypesPolicyTypeProvision     SearchPoliciesRequestPolicyTypes = "POLICY_TYPE_PROVISION"
)

func (e SearchPoliciesRequestPolicyTypes) ToPointer() *SearchPoliciesRequestPolicyTypes {
	return &e
}

func (e *SearchPoliciesRequestPolicyTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "POLICY_TYPE_UNSPECIFIED":
		fallthrough
	case "POLICY_TYPE_GRANT":
		fallthrough
	case "POLICY_TYPE_REVOKE":
		fallthrough
	case "POLICY_TYPE_CERTIFY":
		fallthrough
	case "POLICY_TYPE_ACCESS_REQUEST":
		fallthrough
	case "POLICY_TYPE_PROVISION":
		*e = SearchPoliciesRequestPolicyTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SearchPoliciesRequestPolicyTypes: %v", v)
	}
}

// SearchPoliciesRequest - Search Policies by a few properties.
type SearchPoliciesRequest struct {
	// Search for policies with a case insensitive match on the display name.
	DisplayName *string `json:"displayName,omitempty"`
	// The pageSize where 0 <= pageSize <= 100. Values < 10 will be set to 10. A value of 0 returns the default page size (currently 25)
	PageSize *float64 `json:"pageSize,omitempty"`
	// The pageToken field.
	PageToken *string `json:"pageToken,omitempty"`
	// The policy type to search on. This can be POLICY_TYPE_GRANT, POLICY_TYPE_REVOKE, POLICY_TYPE_CERTIFY, POLICY_TYPE_ACCESS_REQUEST, or POLICY_TYPE_PROVISION.
	PolicyTypes []SearchPoliciesRequestPolicyTypes `json:"policyTypes,omitempty"`
	// Query the policies with a fuzzy search on display name and description.
	Query *string `json:"query,omitempty"`
}

func (o *SearchPoliciesRequest) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *SearchPoliciesRequest) GetPageSize() *float64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *SearchPoliciesRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

func (o *SearchPoliciesRequest) GetPolicyTypes() []SearchPoliciesRequestPolicyTypes {
	if o == nil {
		return nil
	}
	return o.PolicyTypes
}

func (o *SearchPoliciesRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}
