# RequestSchemaServiceFindBindingForAppEntitlementResponse

The response message containing the binding for the specified app entitlement.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `AppEntitlementRef`                                                            | [*shared.AppEntitlementRef](../../../pkg/models/shared/appentitlementref.md)   | :heavy_minus_sign:                                                             | The AppEntitlementRef message.                                                 |
| `RequestSchemaID`                                                              | `*string`                                                                      | :heavy_minus_sign:                                                             | The unique identifier of the request schema bound to this entitlement, if any. |