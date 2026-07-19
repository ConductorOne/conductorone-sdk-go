# Hook

Hook represents a customer-configured interception point for tool calls.

This message contains a oneof named hook_type. Only a single field of the following list may be set at a time:
  - function
  - builtinPattern



## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `BuiltinPattern`                                                         | [*shared.BuiltInPattern](../../../pkg/models/shared/builtinpattern.md)   | :heavy_minus_sign:                                                       | N/A                                                                      |
| `CreatedAt`                                                              | [*time.Time](https://pkg.go.dev/time#Time)                               | :heavy_minus_sign:                                                       | N/A                                                                      |
| `Description`                                                            | `*string`                                                                | :heavy_minus_sign:                                                       | The description field.                                                   |
| `DisplayName`                                                            | `*string`                                                                | :heavy_minus_sign:                                                       | The displayName field.                                                   |
| `Enabled`                                                                | `*bool`                                                                  | :heavy_minus_sign:                                                       | The enabled field.                                                       |
| `Event`                                                                  | [*shared.Event](../../../pkg/models/shared/event.md)                     | :heavy_minus_sign:                                                       | The event field.                                                         |
| `Filter`                                                                 | [*shared.HookFilter](../../../pkg/models/shared/hookfilter.md)           | :heavy_minus_sign:                                                       | N/A                                                                      |
| `Function`                                                               | [*shared.HookFunctionRef](../../../pkg/models/shared/hookfunctionref.md) | :heavy_minus_sign:                                                       | N/A                                                                      |
| `ID`                                                                     | `*string`                                                                | :heavy_minus_sign:                                                       | The id field.                                                            |
| `Priority`                                                               | `*int`                                                                   | :heavy_minus_sign:                                                       | The priority field.                                                      |
| `UpdatedAt`                                                              | [*time.Time](https://pkg.go.dev/time#Time)                               | :heavy_minus_sign:                                                       | N/A                                                                      |