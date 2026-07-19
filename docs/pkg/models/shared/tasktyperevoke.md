# TaskTypeRevoke

The TaskTypeRevoke message indicates that a task is a revoke task and all related details.

This message contains a oneof named principal. Only a single field of the following list may be set at a time:
  - resource



## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `AppEntitlementID`                                                                   | `*string`                                                                            | :heavy_minus_sign:                                                                   | The ID of the app entitlement.                                                       |
| `AppID`                                                                              | `*string`                                                                            | :heavy_minus_sign:                                                                   | The ID of the app.                                                                   |
| `AppUserID`                                                                          | `*string`                                                                            | :heavy_minus_sign:                                                                   | The ID of the app user.                                                              |
| `IdentityUserID`                                                                     | `*string`                                                                            | :heavy_minus_sign:                                                                   | The ID of the user.                                                                  |
| `Outcome`                                                                            | [*shared.TaskTypeRevokeOutcome](../../../pkg/models/shared/tasktyperevokeoutcome.md) | :heavy_minus_sign:                                                                   | The outcome of the revoke.                                                           |
| `OutcomeTime`                                                                        | [*time.Time](https://pkg.go.dev/time#Time)                                           | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `Resource`                                                                           | [*shared.AppResourceRef](../../../pkg/models/shared/appresourceref.md)               | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `Source`                                                                             | [*shared.TaskRevokeSource](../../../pkg/models/shared/taskrevokesource.md)           | :heavy_minus_sign:                                                                   | N/A                                                                                  |