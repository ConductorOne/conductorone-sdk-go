# RequestCatalogManagementServiceAddAccessEntitlementsRequest

The RequestCatalogManagementServiceAddAccessEntitlementsRequest message is used to add access entitlements to a request
 catalog to determine which users can view the request catalog.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `AccessEntitlements`                                                          | [][shared.AppEntitlementRef](../../../pkg/models/shared/appentitlementref.md) | :heavy_minus_sign:                                                            | List of entitlements to add to the request catalog as access entitlements.    |