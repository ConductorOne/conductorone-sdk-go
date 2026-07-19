# BuiltInPattern

BuiltInPattern references a ConductorOne-maintained DLP pattern.
 The specific pattern and its configuration are encoded as a oneof.

This message contains a oneof named config. Only a single field of the following list may be set at a time:
  - piiRedaction
  - creditCardBlocking
  - queryScopeLimit
  - writeAuthorization
  - sensitiveFileGuard
  - toolOutputSizeGuard



## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `CreditCardBlocking`                                                                         | [*shared.CreditCardBlockingConfig](../../../pkg/models/shared/creditcardblockingconfig.md)   | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `PiiRedaction`                                                                               | [*shared.PIIRedactionConfig](../../../pkg/models/shared/piiredactionconfig.md)               | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `QueryScopeLimit`                                                                            | [*shared.QueryScopeLimitConfig](../../../pkg/models/shared/queryscopelimitconfig.md)         | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `SensitiveFileGuard`                                                                         | [*shared.SensitiveFileGuardConfig](../../../pkg/models/shared/sensitivefileguardconfig.md)   | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `ToolOutputSizeGuard`                                                                        | [*shared.ToolOutputSizeGuardConfig](../../../pkg/models/shared/tooloutputsizeguardconfig.md) | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `WriteAuthorization`                                                                         | [*shared.WriteAuthorizationConfig](../../../pkg/models/shared/writeauthorizationconfig.md)   | :heavy_minus_sign:                                                                           | N/A                                                                                          |