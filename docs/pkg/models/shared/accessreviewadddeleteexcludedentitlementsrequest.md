# AccessReviewAddDeleteExcludedEntitlementsRequest

Incrementally adds and/or removes excluded entitlements for a user-first campaign without
 replacing the entire set (contrast AccessReviewSetCampaignScopeAndExcludedEntitlementsRequest).


## Fields

| Field                                                                                                               | Type                                                                                                                | Required                                                                                                            | Description                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| `ToAdd`                                                                                                             | [][shared.AccessReviewExcludedEntitlementInput](../../../pkg/models/shared/accessreviewexcludedentitlementinput.md) | :heavy_minus_sign:                                                                                                  | The entitlements to add to the campaign's excluded set.                                                             |
| `ToDelete`                                                                                                          | [][shared.AccessReviewExcludedEntitlementInput](../../../pkg/models/shared/accessreviewexcludedentitlementinput.md) | :heavy_minus_sign:                                                                                                  | The entitlements to remove from the campaign's excluded set.                                                        |