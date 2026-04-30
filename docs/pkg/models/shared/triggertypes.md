# TriggerTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TriggerTypesTriggerTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TriggerTypes("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `TriggerTypesTriggerTypeUnspecified`       | TRIGGER_TYPE_UNSPECIFIED                   |
| `TriggerTypesTriggerTypeUserProfileChange` | TRIGGER_TYPE_USER_PROFILE_CHANGE           |
| `TriggerTypesTriggerTypeAppUserCreate`     | TRIGGER_TYPE_APP_USER_CREATE               |
| `TriggerTypesTriggerTypeAppUserUpdate`     | TRIGGER_TYPE_APP_USER_UPDATE               |
| `TriggerTypesTriggerTypeUnusedAccess`      | TRIGGER_TYPE_UNUSED_ACCESS                 |
| `TriggerTypesTriggerTypeUserCreated`       | TRIGGER_TYPE_USER_CREATED                  |
| `TriggerTypesTriggerTypeGrantFound`        | TRIGGER_TYPE_GRANT_FOUND                   |
| `TriggerTypesTriggerTypeGrantDeleted`      | TRIGGER_TYPE_GRANT_DELETED                 |
| `TriggerTypesTriggerTypeWebhook`           | TRIGGER_TYPE_WEBHOOK                       |
| `TriggerTypesTriggerTypeSchedule`          | TRIGGER_TYPE_SCHEDULE                      |
| `TriggerTypesTriggerTypeForm`              | TRIGGER_TYPE_FORM                          |
| `TriggerTypesTriggerTypeScheduleAppUser`   | TRIGGER_TYPE_SCHEDULE_APP_USER             |
| `TriggerTypesTriggerTypeAccessConflict`    | TRIGGER_TYPE_ACCESS_CONFLICT               |
| `TriggerTypesTriggerTypeScheduleNoUser`    | TRIGGER_TYPE_SCHEDULE_NO_USER              |