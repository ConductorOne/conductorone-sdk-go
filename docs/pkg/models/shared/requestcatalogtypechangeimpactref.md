# RequestCatalogTypeChangeImpactRef

Object associated with a type-change impact.

This message contains a oneof named ref. Only a single field of the following list may be set at a time:
  - requestCatalog
  - appEntitlement
  - bundleAutomation



## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `AppEntitlement`                                                                 | [*shared.AppEntitlementRef](../../../pkg/models/shared/appentitlementref.md)     | :heavy_minus_sign:                                                               | N/A                                                                              |
| `BundleAutomation`                                                               | [*shared.BundleAutomationRef](../../../pkg/models/shared/bundleautomationref.md) | :heavy_minus_sign:                                                               | N/A                                                                              |
| `RequestCatalog`                                                                 | [*shared.RequestCatalogRef](../../../pkg/models/shared/requestcatalogref.md)     | :heavy_minus_sign:                                                               | N/A                                                                              |