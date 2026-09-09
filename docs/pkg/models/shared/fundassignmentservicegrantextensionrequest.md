# FundAssignmentServiceGrantExtensionRequest

The FundAssignmentServiceGrantExtensionRequest message.


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ExpiresAt`                                                    | [*time.Time](https://pkg.go.dev/time#Time)                     | :heavy_minus_sign:                                             | N/A                                                            |
| `Limit`                                                        | [*shared.SpendLimit](../../../pkg/models/shared/spendlimit.md) | :heavy_minus_sign:                                             | N/A                                                            |
| `Reason`                                                       | `*string`                                                      | :heavy_minus_sign:                                             | Subject-visible: "why do I have this bump".                    |