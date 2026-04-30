# SSFReceiverEventSearchServiceSearchResponse

SSFReceiverEventSearchServiceSearchResponse contains the matching events and a pagination token.


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `List`                                                                      | [][shared.SSFReceiverEvent](../../../pkg/models/shared/ssfreceiverevent.md) | :heavy_minus_sign:                                                          | The SSF events matching the search criteria.                                |
| `NextPageToken`                                                             | `*string`                                                                   | :heavy_minus_sign:                                                          | Token to retrieve the next page. Empty when there are no more results.      |