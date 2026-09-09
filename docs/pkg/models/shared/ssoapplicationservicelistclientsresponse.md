# SSOApplicationServiceListClientsResponse

SSOApplicationServiceListClientsResponse contains a page of App-owned OAuth
 clients.


## Fields

| Field                                                                                       | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `List`                                                                                      | [][shared.SSOApplicationOIDCClient](../../../pkg/models/shared/ssoapplicationoidcclient.md) | :heavy_minus_sign:                                                                          | App-owned clients in this page.                                                             |
| `NextPageToken`                                                                             | `*string`                                                                                   | :heavy_minus_sign:                                                                          | Pagination token for the next page, or empty when complete.                                 |