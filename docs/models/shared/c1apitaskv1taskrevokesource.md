# C1APITaskV1TaskRevokeSource

The TaskRevokeSource message.

This message contains a oneof named origin. Only a single field of the following list may be set at a time:
  - review
  - request
  - expired
  - nonUsage



## Fields

| Field                                                                                              | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `Expired`                                                                                          | [*C1APITaskV1TaskRevokeSourceExpired](../../models/shared/c1apitaskv1taskrevokesourceexpired.md)   | :heavy_minus_sign:                                                                                 | The TaskRevokeSourceExpired message.                                                               |
| `NonUsage`                                                                                         | [*C1APITaskV1TaskRevokeSourceNonUsage](../../models/shared/c1apitaskv1taskrevokesourcenonusage.md) | :heavy_minus_sign:                                                                                 | The TaskRevokeSourceNonUsage message.                                                              |
| `Request`                                                                                          | [*C1APITaskV1TaskRevokeSourceRequest](../../models/shared/c1apitaskv1taskrevokesourcerequest.md)   | :heavy_minus_sign:                                                                                 | The TaskRevokeSourceRequest message.                                                               |
| `Review`                                                                                           | [*C1APITaskV1TaskRevokeSourceReview](../../models/shared/c1apitaskv1taskrevokesourcereview.md)     | :heavy_minus_sign:                                                                                 | The TaskRevokeSourceReview message.                                                                |