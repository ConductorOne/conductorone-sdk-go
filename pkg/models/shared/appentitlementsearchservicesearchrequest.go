// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AppEntitlementSearchServiceSearchRequest -  Search app entitlements by a variety of filters.
type AppEntitlementSearchServiceSearchRequest struct {
	//  The app entitlement expand mask allows the user to get additonal information when getting responses containing app entitlement views.
	//
	AppEntitlementExpandMask *AppEntitlementExpandMask `json:"expandMask,omitempty"`
	//  Search for app entitlements that are being reviewed as part of this access review campaign.
	//
	AccessReviewID *string `json:"accessReviewId,omitempty"`
	//  Search for app entitlements that have this alias (exact match).
	//
	Alias *string `json:"alias,omitempty"`
	//  Search for app entitlements contained in any of these apps.
	//
	AppIds []string `json:"appIds,omitempty"`
	//  Search for app entitlements that are granted to any of these app user ids.
	//
	AppUserIds []string `json:"appUserIds,omitempty"`
	//  Search for app entitlements that are part of these compliace frameworks.
	//
	ComplianceFrameworkIds []string `json:"complianceFrameworkIds,omitempty"`
	//  Exclude app entitlements from the results that are in these app IDs.
	//
	ExcludeAppIds []string `json:"excludeAppIds,omitempty"`
	//  Exclude app entitlements from the results that these app users have granted.
	//
	ExcludeAppUserIds []string `json:"excludeAppUserIds,omitempty"`
	//  Restrict results to only those who have expiring app entitlement user bindings.
	//
	OnlyGetExpiring *bool `json:"onlyGetExpiring,omitempty"`
	//  The pageSize where 0 <= pageSize <= 100. Values < 10 will be set to 10. A value of 0 returns the default page size (currently 25)
	//
	PageSize *float64 `json:"pageSize,omitempty"`
	//  The pageToken field.
	//
	PageToken *string `json:"pageToken,omitempty"`
	//  Query the app entitlements with a fuzzy search on display name and description.
	//
	Query *string `json:"query,omitempty"`
	//  Search for app entitlements that are for items on these resource types.
	//
	ResourceTypeIds []string `json:"resourceTypeIds,omitempty"`
	//  Search for app entitlements with these risk levels.
	//
	RiskLevelIds []string `json:"riskLevelIds,omitempty"`
}

func (o *AppEntitlementSearchServiceSearchRequest) GetAppEntitlementExpandMask() *AppEntitlementExpandMask {
	if o == nil {
		return nil
	}
	return o.AppEntitlementExpandMask
}

func (o *AppEntitlementSearchServiceSearchRequest) GetAccessReviewID() *string {
	if o == nil {
		return nil
	}
	return o.AccessReviewID
}

func (o *AppEntitlementSearchServiceSearchRequest) GetAlias() *string {
	if o == nil {
		return nil
	}
	return o.Alias
}

func (o *AppEntitlementSearchServiceSearchRequest) GetAppIds() []string {
	if o == nil {
		return nil
	}
	return o.AppIds
}

func (o *AppEntitlementSearchServiceSearchRequest) GetAppUserIds() []string {
	if o == nil {
		return nil
	}
	return o.AppUserIds
}

func (o *AppEntitlementSearchServiceSearchRequest) GetComplianceFrameworkIds() []string {
	if o == nil {
		return nil
	}
	return o.ComplianceFrameworkIds
}

func (o *AppEntitlementSearchServiceSearchRequest) GetExcludeAppIds() []string {
	if o == nil {
		return nil
	}
	return o.ExcludeAppIds
}

func (o *AppEntitlementSearchServiceSearchRequest) GetExcludeAppUserIds() []string {
	if o == nil {
		return nil
	}
	return o.ExcludeAppUserIds
}

func (o *AppEntitlementSearchServiceSearchRequest) GetOnlyGetExpiring() *bool {
	if o == nil {
		return nil
	}
	return o.OnlyGetExpiring
}

func (o *AppEntitlementSearchServiceSearchRequest) GetPageSize() *float64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *AppEntitlementSearchServiceSearchRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

func (o *AppEntitlementSearchServiceSearchRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *AppEntitlementSearchServiceSearchRequest) GetResourceTypeIds() []string {
	if o == nil {
		return nil
	}
	return o.ResourceTypeIds
}

func (o *AppEntitlementSearchServiceSearchRequest) GetRiskLevelIds() []string {
	if o == nil {
		return nil
	}
	return o.RiskLevelIds
}
