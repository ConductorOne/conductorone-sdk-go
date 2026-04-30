# RequestSchemaServiceRemoveEntitlementBindingRequest

The request message for removing a single entitlement binding from a request schema.


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `AppEntitlementRef`                                                          | [*shared.AppEntitlementRef](../../../pkg/models/shared/appentitlementref.md) | :heavy_minus_sign:                                                           | The AppEntitlementRef message.                                               |
| `RequestSchemaID`                                                            | `*string`                                                                    | :heavy_minus_sign:                                                           | The unique identifier of the request schema to remove the binding from.      |