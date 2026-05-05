# ExportServiceListEventsResponse

ExportServiceListEventsResponse is the response containing audit events for an export.


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `List`                                                                                 | []map[string]`any`                                                                     | :heavy_minus_sign:                                                                     | List contains an array of JSON OCSF events.                                            |
| `NextPageToken`                                                                        | `*string`                                                                              | :heavy_minus_sign:                                                                     | The token to retrieve the next page of results, or empty if there are no more results. |