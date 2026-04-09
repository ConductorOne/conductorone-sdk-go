# StatusFilter

Optional filter by invitation status.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.StatusFilterLocalInvitationStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.StatusFilter("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `StatusFilterLocalInvitationStatusUnspecified` | LOCAL_INVITATION_STATUS_UNSPECIFIED            |
| `StatusFilterLocalInvitationStatusPending`     | LOCAL_INVITATION_STATUS_PENDING                |
| `StatusFilterLocalInvitationStatusAccepted`    | LOCAL_INVITATION_STATUS_ACCEPTED               |
| `StatusFilterLocalInvitationStatusRevoked`     | LOCAL_INVITATION_STATUS_REVOKED                |
| `StatusFilterLocalInvitationStatusExpired`     | LOCAL_INVITATION_STATUS_EXPIRED                |