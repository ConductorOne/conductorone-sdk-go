# ListRunsResponse

The ListRunsResponse message.


## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `List`                                                                                    | [][shared.RoleMiningManagementRun](../../../pkg/models/shared/roleminingmanagementrun.md) | :heavy_minus_sign:                                                                        | The list of role mining analysis runs.                                                    |
| `NextPageToken`                                                                           | `*string`                                                                                 | :heavy_minus_sign:                                                                        | Token to retrieve the next page of results, empty if no more results.                     |