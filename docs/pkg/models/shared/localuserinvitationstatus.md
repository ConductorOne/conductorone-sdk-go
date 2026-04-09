# LocalUserInvitationStatus

Current lifecycle status. Read-only.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.LocalUserInvitationStatusLocalInvitationStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.LocalUserInvitationStatus("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `LocalUserInvitationStatusLocalInvitationStatusUnspecified` | LOCAL_INVITATION_STATUS_UNSPECIFIED                         |
| `LocalUserInvitationStatusLocalInvitationStatusPending`     | LOCAL_INVITATION_STATUS_PENDING                             |
| `LocalUserInvitationStatusLocalInvitationStatusAccepted`    | LOCAL_INVITATION_STATUS_ACCEPTED                            |
| `LocalUserInvitationStatusLocalInvitationStatusRevoked`     | LOCAL_INVITATION_STATUS_REVOKED                             |
| `LocalUserInvitationStatusLocalInvitationStatusExpired`     | LOCAL_INVITATION_STATUS_EXPIRED                             |