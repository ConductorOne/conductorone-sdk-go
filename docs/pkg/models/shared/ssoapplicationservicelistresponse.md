# SSOApplicationServiceListResponse

SSOApplicationServiceListResponse returns a page of SSO applications.


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `List`                                                                     | [][shared.SSOApplication](../../../pkg/models/shared/ssoapplication.md)    | :heavy_minus_sign:                                                         | The page of SSO applications.                                              |
| `NextPageToken`                                                            | `*string`                                                                  | :heavy_minus_sign:                                                         | Pagination token for the next page, or empty if there are no more results. |