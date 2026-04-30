# ListAutomationsResponse

The ListAutomationsResponse message.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `List`                                                                        | [][shared.Automation](../../../pkg/models/shared/automation.md)               | :heavy_minus_sign:                                                            | The page of automations.                                                      |
| `NextPageToken`                                                               | `*string`                                                                     | :heavy_minus_sign:                                                            | Token to retrieve the next page of results, empty when no more results exist. |