# ResolvedLimit

ResolvedLimit is funds.v1.SpendLimit minus `blocked` — a blocked scope fails
 closed during resolution and never creates an account — with Money flattened
 to nano.

This message contains a oneof named kind. Only a single field of the following list may be set at a time:
  - unlimited
  - amount



## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `Amount`                                                                     | [*shared.ResolvedAmount](../../../pkg/models/shared/resolvedamount.md)       | :heavy_minus_sign:                                                           | N/A                                                                          |
| `Unlimited`                                                                  | [*shared.ResolvedUnlimited](../../../pkg/models/shared/resolvedunlimited.md) | :heavy_minus_sign:                                                           | N/A                                                                          |