# ValidationCheck

ValidationCheck for client-side validation rules.

This message contains a oneof named check. Only a single field of the following list may be set at a time:
  - call
  - and
  - or



## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `And`                                                              | [*shared.AndCheck](../../../pkg/models/shared/andcheck.md)         | :heavy_minus_sign:                                                 | N/A                                                                |
| `Call`                                                             | [*shared.FunctionCall](../../../pkg/models/shared/functioncall.md) | :heavy_minus_sign:                                                 | N/A                                                                |
| `Or`                                                               | [*shared.OrCheck](../../../pkg/models/shared/orcheck.md)           | :heavy_minus_sign:                                                 | N/A                                                                |