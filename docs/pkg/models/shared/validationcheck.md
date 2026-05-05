# ValidationCheck

ValidationCheck for client-side validation rules.

This message contains a oneof named check. Only a single field of the following list may be set at a time:
  - call
  - and
  - or



## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `AndCheck`                                                         | [*shared.AndCheck](../../../pkg/models/shared/andcheck.md)         | :heavy_minus_sign:                                                 | AndCheck requires all checks to pass.                              |
| `FunctionCall`                                                     | [*shared.FunctionCall](../../../pkg/models/shared/functioncall.md) | :heavy_minus_sign:                                                 | FunctionCall represents a client-side function invocation.         |
| `OrCheck`                                                          | [*shared.OrCheck](../../../pkg/models/shared/orcheck.md)           | :heavy_minus_sign:                                                 | OrCheck requires at least one check to pass.                       |