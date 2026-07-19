# XAAResourceServerServiceListResponse

XAAResourceServerServiceListResponse returns a page of resource servers.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `NextPageToken`                                                               | `*string`                                                                     | :heavy_minus_sign:                                                            | Token for the next page, or empty if there are no more results.               |
| `ResourceServers`                                                             | [][shared.XAAResourceServer](../../../pkg/models/shared/xaaresourceserver.md) | :heavy_minus_sign:                                                            | The page of resource servers.                                                 |