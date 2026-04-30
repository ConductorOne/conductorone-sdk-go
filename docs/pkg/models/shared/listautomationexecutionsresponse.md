# ListAutomationExecutionsResponse

The ListAutomationExecutionsResponse message.


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `AutomationExecutions`                                                            | [][shared.AutomationExecution](../../../pkg/models/shared/automationexecution.md) | :heavy_minus_sign:                                                                | The page of automation executions.                                                |
| `NextPageToken`                                                                   | `*string`                                                                         | :heavy_minus_sign:                                                                | Token to retrieve the next page of results, empty when no more results exist.     |