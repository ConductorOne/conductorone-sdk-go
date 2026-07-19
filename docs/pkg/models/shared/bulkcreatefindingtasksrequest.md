# BulkCreateFindingTasksRequest

The BulkCreateFindingTasksRequest message.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `PolicyID`                                                                           | `*string`                                                                            | :heavy_minus_sign:                                                                   | Optional policy ID to use for the created tasks. Defaults to the app's grant policy. |
| `Refs`                                                                               | [][shared.FindingRef](../../../pkg/models/shared/findingref.md)                      | :heavy_minus_sign:                                                                   | Individual finding references to create tasks for (by-ID mode).                      |
| `SearchRequest`                                                                      | [*shared.FindingSearchRequest](../../../pkg/models/shared/findingsearchrequest.md)   | :heavy_minus_sign:                                                                   | N/A                                                                                  |