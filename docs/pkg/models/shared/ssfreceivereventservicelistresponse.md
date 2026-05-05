# SSFReceiverEventServiceListResponse

SSFReceiverEventServiceListResponse contains a page of received SSF events.


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `List`                                                                      | [][shared.SSFReceiverEvent](../../../pkg/models/shared/ssfreceiverevent.md) | :heavy_minus_sign:                                                          | The SSF events in the current page.                                         |
| `NextPageToken`                                                             | `*string`                                                                   | :heavy_minus_sign:                                                          | Token to retrieve the next page. Empty when there are no more results.      |