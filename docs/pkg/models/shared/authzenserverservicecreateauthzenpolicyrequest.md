# AuthzenServerServiceCreateAuthzenPolicyRequest

AuthzenServerServiceCreateAuthzenPolicyRequest creates a new CEL policy
 revision for an app.

This message contains a oneof named engine. Only a single field of the following list may be set at a time:
  - cel



## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Cel`                                                                      | [*shared.AuthzenCelPolicy](../../../pkg/models/shared/authzencelpolicy.md) | :heavy_minus_sign:                                                         | N/A                                                                        |
| `Description`                                                              | `*string`                                                                  | :heavy_minus_sign:                                                         | Free-text description of what this revision changes or is for.             |
| `DisplayName`                                                              | `*string`                                                                  | :heavy_minus_sign:                                                         | Display name for the policy revision.                                      |