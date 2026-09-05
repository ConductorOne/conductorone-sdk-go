# SpendExtension

SpendExtension replaces the row's total with a temporary one until
 expires_at. It never changes the period, and it never expresses a refusal —
 a temporary refusal is a SpendSuspension.


## Fields

| Field                                                                                                               | Type                                                                                                                | Required                                                                                                            | Description                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| `ExpiresAt`                                                                                                         | [*time.Time](https://pkg.go.dev/time#Time)                                                                          | :heavy_minus_sign:                                                                                                  | N/A                                                                                                                 |
| `Limit`                                                                                                             | [*shared.SpendLimit](../../../pkg/models/shared/spendlimit.md)                                                      | :heavy_minus_sign:                                                                                                  | N/A                                                                                                                 |
| `Reason`                                                                                                            | `*string`                                                                                                           | :heavy_minus_sign:                                                                                                  | Subject-visible: "why do I have this bump". Mutation rationale rides the<br/> history change_reason annotation instead. |