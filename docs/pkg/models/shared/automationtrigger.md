# AutomationTrigger

Automation Triggers

This message contains a oneof named kind. Only a single field of the following list may be set at a time:
  - userProfileChange
  - appUserCreated
  - appUserUpdated
  - usageBasedRevocation
  - userCreated
  - grantFound
  - grantDeleted
  - webhook
  - schedule
  - scheduleAppUser
  - accessConflict
  - scheduleNoUser



## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `AccessConflict`                                                                                 | [*shared.AccessConflictTrigger](../../../pkg/models/shared/accessconflicttrigger.md)             | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `AppUserCreated`                                                                                 | [*shared.AppUserCreatedTrigger](../../../pkg/models/shared/appusercreatedtrigger.md)             | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `AppUserUpdated`                                                                                 | [*shared.AppUserUpdatedTrigger](../../../pkg/models/shared/appuserupdatedtrigger.md)             | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `GrantDeleted`                                                                                   | [*shared.GrantDeletedTrigger](../../../pkg/models/shared/grantdeletedtrigger.md)                 | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `GrantFound`                                                                                     | [*shared.GrantFoundTrigger](../../../pkg/models/shared/grantfoundtrigger.md)                     | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `Schedule`                                                                                       | [*shared.ScheduleTrigger](../../../pkg/models/shared/scheduletrigger.md)                         | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `ScheduleAppUser`                                                                                | [*shared.ScheduleTriggerAppUser](../../../pkg/models/shared/scheduletriggerappuser.md)           | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `ScheduleNoUser`                                                                                 | [*shared.ScheduleTriggerNoUser](../../../pkg/models/shared/scheduletriggernouser.md)             | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `UsageBasedRevocation`                                                                           | [*shared.UsageBasedRevocationTrigger](../../../pkg/models/shared/usagebasedrevocationtrigger.md) | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `UserCreated`                                                                                    | [*shared.UserCreatedTrigger](../../../pkg/models/shared/usercreatedtrigger.md)                   | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `UserProfileChange`                                                                              | [*shared.UserProfileChangeTrigger](../../../pkg/models/shared/userprofilechangetrigger.md)       | :heavy_minus_sign:                                                                               | N/A                                                                                              |
| `Webhook`                                                                                        | [*shared.WebhookAutomationTrigger](../../../pkg/models/shared/webhookautomationtrigger.md)       | :heavy_minus_sign:                                                                               | N/A                                                                                              |