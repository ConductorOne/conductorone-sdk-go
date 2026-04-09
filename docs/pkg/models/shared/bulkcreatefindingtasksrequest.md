# BulkCreateFindingTasksRequest

The BulkCreateFindingTasksRequest message.


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `FindingSearchRequest`                                                             | [*shared.FindingSearchRequest](../../../pkg/models/shared/findingsearchrequest.md) | :heavy_minus_sign:                                                                 | The FindingSearchRequest message.                                                  |
| `PolicyID`                                                                         | `*string`                                                                          | :heavy_minus_sign:                                                                 | The policyId field.                                                                |
| `Refs`                                                                             | [][shared.FindingRef](../../../pkg/models/shared/findingref.md)                    | :heavy_minus_sign:                                                                 | The refs field.                                                                    |