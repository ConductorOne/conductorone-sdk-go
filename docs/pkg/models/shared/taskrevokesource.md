# TaskRevokeSource

The TaskRevokeSource message indicates the source of the revoke task is one of expired, nonUsage, request, or review.

This message contains a oneof named origin. Only a single field of the following list may be set at a time:
  - review
  - request
  - expired
  - nonUsage



## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Expired`                                                                                  | [*shared.TaskRevokeSourceExpired](../../../pkg/models/shared/taskrevokesourceexpired.md)   | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `NonUsage`                                                                                 | [*shared.TaskRevokeSourceNonUsage](../../../pkg/models/shared/taskrevokesourcenonusage.md) | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `Request`                                                                                  | [*shared.TaskRevokeSourceRequest](../../../pkg/models/shared/taskrevokesourcerequest.md)   | :heavy_minus_sign:                                                                         | N/A                                                                                        |
| `Review`                                                                                   | [*shared.TaskRevokeSourceReview](../../../pkg/models/shared/taskrevokesourcereview.md)     | :heavy_minus_sign:                                                                         | N/A                                                                                        |