# LocalUserInvitationServiceSearchRequest

The LocalUserInvitationServiceSearchRequest message.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `DirectoryAppID`                                                   | `*string`                                                          | :heavy_minus_sign:                                                 | The directoryAppId field.                                          |
| `PageSize`                                                         | `*int`                                                             | :heavy_minus_sign:                                                 | The pageSize field.                                                |
| `PageToken`                                                        | `*string`                                                          | :heavy_minus_sign:                                                 | The pageToken field.                                               |
| `StatusFilter`                                                     | [*shared.StatusFilter](../../../pkg/models/shared/statusfilter.md) | :heavy_minus_sign:                                                 | Optional filter by invitation status.                              |