# SearchEmailAuditEventsResponse

The SearchEmailAuditEventsResponse message.


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `List`                                                      | []map[string]`any`                                          | :heavy_minus_sign:                                          | OCSF EmailActivity events as Struct for frontend rendering. |
| `NextPageToken`                                             | `*string`                                                   | :heavy_minus_sign:                                          | Token for next page. Empty when no more pages.              |