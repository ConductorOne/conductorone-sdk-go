# TaskTypeGrant1

The TaskTypeGrant message indicates that a task is a grant task and all related details.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `AppEntitlementID`                                                   | **string*                                                            | :heavy_minus_sign:                                                   | The ID of the app entitlement.                                       |
| `AppID`                                                              | **string*                                                            | :heavy_minus_sign:                                                   | The ID of the app.                                                   |
| `AppUserID`                                                          | **string*                                                            | :heavy_minus_sign:                                                   | The ID of the app user.                                              |
| `GrantDuration`                                                      | **string*                                                            | :heavy_minus_sign:                                                   | N/A                                                                  |
| `IdentityUserID`                                                     | **string*                                                            | :heavy_minus_sign:                                                   | The ID of the user.                                                  |
| `Outcome`                                                            | [*TaskTypeGrantOutcome](../../models/shared/tasktypegrantoutcome.md) | :heavy_minus_sign:                                                   | The outcome of the grant.                                            |
| `OutcomeTime`                                                        | [*time.Time](https://pkg.go.dev/time#Time)                           | :heavy_minus_sign:                                                   | N/A                                                                  |