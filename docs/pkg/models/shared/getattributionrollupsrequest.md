# GetAttributionRollupsRequest

GetAttributionRollupsRequest specifies how to group settled spend.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `Dimension`                                                                   | [*shared.Dimension](../../../pkg/models/shared/dimension.md)                  | :heavy_minus_sign:                                                            | Dimension used to group settled calls.                                        |
| `EndTime`                                                                     | [*time.Time](https://pkg.go.dev/time#Time)                                    | :heavy_minus_sign:                                                            | N/A                                                                           |
| `PageSize`                                                                    | `*int`                                                                        | :heavy_minus_sign:                                                            | Maximum number of groups to return. The default is 25 and the maximum is 100. |
| `PageToken`                                                                   | `*string`                                                                     | :heavy_minus_sign:                                                            | Pagination token from a previous response with the same window and dimension. |
| `StartTime`                                                                   | [*time.Time](https://pkg.go.dev/time#Time)                                    | :heavy_minus_sign:                                                            | N/A                                                                           |