# FundsMoney

Money is a non-negative monetary value used by spending controls.


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `CurrencyCode`                                                           | `*string`                                                                | :heavy_minus_sign:                                                       | ISO 4217 currency code.                                                  |
| `Nanos`                                                                  | `*int`                                                                   | :heavy_minus_sign:                                                       | Fractional nano-units, from zero through 999999999.                      |
| `Units`                                                                  | `*int64`                                                                 | :heavy_minus_sign:                                                       | Whole currency units. The bound keeps nano-unit conversion within int64. |