# PrimaryTriggerType

The primaryTriggerType field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PrimaryTriggerTypeTriggerTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PrimaryTriggerType("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `PrimaryTriggerTypeTriggerTypeUnspecified`       | TRIGGER_TYPE_UNSPECIFIED                         |
| `PrimaryTriggerTypeTriggerTypeUserProfileChange` | TRIGGER_TYPE_USER_PROFILE_CHANGE                 |
| `PrimaryTriggerTypeTriggerTypeAppUserCreate`     | TRIGGER_TYPE_APP_USER_CREATE                     |
| `PrimaryTriggerTypeTriggerTypeAppUserUpdate`     | TRIGGER_TYPE_APP_USER_UPDATE                     |
| `PrimaryTriggerTypeTriggerTypeUnusedAccess`      | TRIGGER_TYPE_UNUSED_ACCESS                       |
| `PrimaryTriggerTypeTriggerTypeUserCreated`       | TRIGGER_TYPE_USER_CREATED                        |
| `PrimaryTriggerTypeTriggerTypeGrantFound`        | TRIGGER_TYPE_GRANT_FOUND                         |
| `PrimaryTriggerTypeTriggerTypeGrantDeleted`      | TRIGGER_TYPE_GRANT_DELETED                       |
| `PrimaryTriggerTypeTriggerTypeWebhook`           | TRIGGER_TYPE_WEBHOOK                             |
| `PrimaryTriggerTypeTriggerTypeSchedule`          | TRIGGER_TYPE_SCHEDULE                            |
| `PrimaryTriggerTypeTriggerTypeForm`              | TRIGGER_TYPE_FORM                                |
| `PrimaryTriggerTypeTriggerTypeScheduleAppUser`   | TRIGGER_TYPE_SCHEDULE_APP_USER                   |
| `PrimaryTriggerTypeTriggerTypeAccessConflict`    | TRIGGER_TYPE_ACCESS_CONFLICT                     |