# PromptInjectionScanConfig

PromptInjectionScanConfig scans tool output for prompt-injection using the
 aigov A2 judge and acts when the verdict is at or above threshold.


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `FlagOnly`                                                                        | `*bool`                                                                           | :heavy_minus_sign:                                                                | When true, a detection records the finding but does not deny (observe-only).      |
| `Threshold`                                                                       | [*shared.Threshold](../../../pkg/models/shared/threshold.md)                      | :heavy_minus_sign:                                                                | Deny (or flag) when the judge scores at or above this level. Unspecified =<br/> HIGH. |