# ToolGatesServiceCreateRequest

The ToolGatesServiceCreateRequest message.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Description`                                                          | `*string`                                                              | :heavy_minus_sign:                                                     | The description field.                                                 |
| `DisplayName`                                                          | `string`                                                               | :heavy_check_mark:                                                     | The displayName field.                                                 |
| `Enabled`                                                              | `*bool`                                                                | :heavy_minus_sign:                                                     | The enabled field.                                                     |
| `Filter`                                                               | [*shared.ToolGateFilter](../../../pkg/models/shared/toolgatefilter.md) | :heavy_minus_sign:                                                     | N/A                                                                    |
| `GrantPolicyID`                                                        | `string`                                                               | :heavy_check_mark:                                                     | The grantPolicyId field.                                               |
| `ManagedByGuardrails`                                                  | `*bool`                                                                | :heavy_minus_sign:                                                     | The managedByGuardrails field.                                         |
| `Priority`                                                             | `*int`                                                                 | :heavy_minus_sign:                                                     | The priority field.                                                    |