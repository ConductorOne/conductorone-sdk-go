# AuthzenPolicy

AuthzenPolicy is one immutable CEL policy revision scoped to an app. A
 revision is never edited in place; activate it by setting a server's
 active_policy_id.

This message contains a oneof named engine. Only a single field of the following list may be set at a time:
  - cel



## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `AppID`                                                                    | `*string`                                                                  | :heavy_minus_sign:                                                         | App identifier.                                                            |
| `Cel`                                                                      | [*shared.AuthzenCelPolicy](../../../pkg/models/shared/authzencelpolicy.md) | :heavy_minus_sign:                                                         | N/A                                                                        |
| `CreatedAt`                                                                | [*time.Time](https://pkg.go.dev/time#Time)                                 | :heavy_minus_sign:                                                         | N/A                                                                        |
| `CreatedByUserID`                                                          | `*string`                                                                  | :heavy_minus_sign:                                                         | The ID of the user who created this revision.                              |
| `Description`                                                              | `*string`                                                                  | :heavy_minus_sign:                                                         | Free-text description of what this revision changes or is for.             |
| `DisplayName`                                                              | `*string`                                                                  | :heavy_minus_sign:                                                         | Display name for the policy revision.                                      |
| `ID`                                                                       | `*string`                                                                  | :heavy_minus_sign:                                                         | Policy identifier.                                                         |