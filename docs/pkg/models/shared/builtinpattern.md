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
  - secretsMasking
  - linkFilter
  - encodedContentGuard
  - promptInjectionScan
  - blockOutput
  - blockToolCall
  - preToolBlock



## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `BlockOutput`                                                                                | [*shared.BlockOutputConfig](../../../pkg/models/shared/blockoutputconfig.md)                 | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `BlockToolCall`                                                                              | [*shared.BlockToolCallConfig](../../../pkg/models/shared/blocktoolcallconfig.md)             | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `CreditCardBlocking`                                                                         | [*shared.CreditCardBlockingConfig](../../../pkg/models/shared/creditcardblockingconfig.md)   | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `EncodedContentGuard`                                                                        | [*shared.EncodedContentGuardConfig](../../../pkg/models/shared/encodedcontentguardconfig.md) | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `LinkFilter`                                                                                 | [*shared.LinkFilterConfig](../../../pkg/models/shared/linkfilterconfig.md)                   | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `PiiRedaction`                                                                               | [*shared.PIIRedactionConfig](../../../pkg/models/shared/piiredactionconfig.md)               | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `PreToolBlock`                                                                               | [*shared.PreToolBlockConfig](../../../pkg/models/shared/pretoolblockconfig.md)               | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `PromptInjectionScan`                                                                        | [*shared.PromptInjectionScanConfig](../../../pkg/models/shared/promptinjectionscanconfig.md) | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `QueryScopeLimit`                                                                            | [*shared.QueryScopeLimitConfig](../../../pkg/models/shared/queryscopelimitconfig.md)         | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `SecretsMasking`                                                                             | [*shared.SecretsMaskingConfig](../../../pkg/models/shared/secretsmaskingconfig.md)           | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `SensitiveFileGuard`                                                                         | [*shared.SensitiveFileGuardConfig](../../../pkg/models/shared/sensitivefileguardconfig.md)   | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `ToolOutputSizeGuard`                                                                        | [*shared.ToolOutputSizeGuardConfig](../../../pkg/models/shared/tooloutputsizeguardconfig.md) | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `WriteAuthorization`                                                                         | [*shared.WriteAuthorizationConfig](../../../pkg/models/shared/writeauthorizationconfig.md)   | :heavy_minus_sign:                                                                           | N/A                                                                                          |