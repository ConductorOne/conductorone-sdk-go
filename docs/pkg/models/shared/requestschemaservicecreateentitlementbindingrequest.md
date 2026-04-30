# RequestSchemaServiceCreateEntitlementBindingRequest

The request message for creating a single entitlement binding on a request schema.


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `AppEntitlementRef`                                                          | [*shared.AppEntitlementRef](../../../pkg/models/shared/appentitlementref.md) | :heavy_minus_sign:                                                           | The AppEntitlementRef message.                                               |
| `RequestSchemaID`                                                            | `*string`                                                                    | :heavy_minus_sign:                                                           | The unique identifier of the request schema to bind the entitlement to.      |