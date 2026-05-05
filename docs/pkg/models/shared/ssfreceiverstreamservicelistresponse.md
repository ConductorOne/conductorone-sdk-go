# SSFReceiverStreamServiceListResponse

SSFReceiverStreamServiceListResponse contains a page of SSF receiver streams.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `List`                                                                        | [][shared.SSFReceiverStream](../../../pkg/models/shared/ssfreceiverstream.md) | :heavy_minus_sign:                                                            | The SSF receiver streams in the current page.                                 |
| `NextPageToken`                                                               | `*string`                                                                     | :heavy_minus_sign:                                                            | Token to retrieve the next page. Empty when there are no more results.        |