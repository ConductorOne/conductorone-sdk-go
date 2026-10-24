# ModelAPIKeyServiceCreateRequest

The ModelAPIKeyServiceCreateRequest message.

This message contains a oneof named credential_format. Only a single field of the following list may be set at a time:
  - anthropic



## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `Anthropic`                                                                                  | [*shared.AnthropicCredentialFormat](../../../pkg/models/shared/anthropiccredentialformat.md) | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `AppID`                                                                                      | `*string`                                                                                    | :heavy_minus_sign:                                                                           | The app this key will access.                                                                |
| `DisplayName`                                                                                | `*string`                                                                                    | :heavy_minus_sign:                                                                           | Human-readable name for this key.                                                            |
| `EdgeID`                                                                                     | `*string`                                                                                    | :heavy_minus_sign:                                                                           | The Edge inference route for API calls with this key.                                        |
| `ExpiresTime`                                                                                | [*time.Time](https://pkg.go.dev/time#Time)                                                   | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `ScopeCeiling`                                                                               | []`string`                                                                                   | :heavy_minus_sign:                                                                           | Maximum scopes this key may use (must be subset of caller's grants and catalog).             |