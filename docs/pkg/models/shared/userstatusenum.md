# UserStatusEnum

The userStatusEnum field.
This field is part of the `user_status` oneof.
See the documentation for `c1.api.automations.v1.UpdateUser` for more details.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.UserStatusEnumUnknown

// Open enum: custom values can be created with a direct type cast
custom := shared.UserStatusEnum("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `UserStatusEnumUnknown`  | UNKNOWN                  |
| `UserStatusEnumEnabled`  | ENABLED                  |
| `UserStatusEnumDisabled` | DISABLED                 |
| `UserStatusEnumDeleted`  | DELETED                  |