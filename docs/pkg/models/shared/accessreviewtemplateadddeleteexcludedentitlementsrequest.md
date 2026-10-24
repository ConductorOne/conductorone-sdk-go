# AccessReviewTemplateAddDeleteExcludedEntitlementsRequest

Incrementally adds and/or removes excluded entitlements for a user-first access review
 template without replacing the entire set (contrast
 AccessReviewSetTemplateScopeAndExcludedEntitlementsRequest).


## Fields

| Field                                                                                                               | Type                                                                                                                | Required                                                                                                            | Description                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| `ToAdd`                                                                                                             | [][shared.AccessReviewExcludedEntitlementInput](../../../pkg/models/shared/accessreviewexcludedentitlementinput.md) | :heavy_minus_sign:                                                                                                  | The entitlements to add to the template's excluded set.                                                             |
| `ToDelete`                                                                                                          | [][shared.AccessReviewExcludedEntitlementInput](../../../pkg/models/shared/accessreviewexcludedentitlementinput.md) | :heavy_minus_sign:                                                                                                  | The entitlements to remove from the template's excluded set.                                                        |