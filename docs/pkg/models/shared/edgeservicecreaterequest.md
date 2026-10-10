# EdgeServiceCreateRequest

EdgeServiceCreateRequest creates an Edge for an app. The new Edge starts
 disabled and denies every capability until enabled.


## Fields

| Field                                                                                                               | Type                                                                                                                | Required                                                                                                            | Description                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| `DisplayName`                                                                                                       | `*string`                                                                                                           | :heavy_minus_sign:                                                                                                  | Display name for the Edge. Defaults to "Edge" when empty.                                                           |
| `EnabledKinds`                                                                                                      | [][shared.EdgeServiceCreateRequestEnabledKinds](../../../pkg/models/shared/edgeservicecreaterequestenabledkinds.md) | :heavy_minus_sign:                                                                                                  | The capabilities this Edge will serve. Each mints its own entitlement<br/> stack; there is no default.              |