# ListStepUpProvidersResponse

The ListStepUpProvidersResponse message.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `List`                                                                               | [][shared.StepUpProvider](../../../pkg/models/shared/stepupprovider.md)              | :heavy_minus_sign:                                                                   | The list of step-up authentication providers.                                        |
| `NextPageToken`                                                                      | `*string`                                                                            | :heavy_minus_sign:                                                                   | A token to retrieve the next page of results, or empty if there are no more results. |