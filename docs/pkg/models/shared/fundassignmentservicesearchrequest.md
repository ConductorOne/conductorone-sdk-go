# FundAssignmentServiceSearchRequest

The FundAssignmentServiceSearchRequest message.


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `PageSize`                                                                | `*int`                                                                    | :heavy_minus_sign:                                                        | The pageSize field.                                                       |
| `PageToken`                                                               | `*string`                                                                 | :heavy_minus_sign:                                                        | The pageToken field.                                                      |
| `UserIds`                                                                 | []`string`                                                                | :heavy_minus_sign:                                                        | Restrict to these subjects; empty returns every assignment in the tenant. |