# DenialMoneySnapshot

DenialMoneySnapshot captures account amounts at the first refusal.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ConsumedNano`                                                       | `*int64`                                                             | :heavy_minus_sign:                                                   | Spend already settled when the first refusal occurred.               |
| `CurrencyCode`                                                       | `*string`                                                            | :heavy_minus_sign:                                                   | Currency used for monetary values.                                   |
| `LimitNano`                                                          | `*int64`                                                             | :heavy_minus_sign:                                                   | Effective limit that caused the refusal.                             |
| `RequestedNano`                                                      | `*int64`                                                             | :heavy_minus_sign:                                                   | Spend requested by the refused call.                                 |
| `ReservedNano`                                                       | `*int64`                                                             | :heavy_minus_sign:                                                   | Spend reserved by in-progress calls when the first refusal occurred. |