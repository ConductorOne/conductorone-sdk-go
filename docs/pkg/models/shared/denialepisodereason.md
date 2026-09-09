# DenialEpisodeReason

Reason the calls were denied.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DenialEpisodeReasonDenyReasonUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DenialEpisodeReason("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `DenialEpisodeReasonDenyReasonUnspecified`      | DENY_REASON_UNSPECIFIED                         |
| `DenialEpisodeReasonDenyReasonTenantFrozen`     | DENY_REASON_TENANT_FROZEN                       |
| `DenialEpisodeReasonDenyReasonSuspendedByAdmin` | DENY_REASON_SUSPENDED_BY_ADMIN                  |
| `DenialEpisodeReasonDenyReasonAppSuspended`     | DENY_REASON_APP_SUSPENDED                       |
| `DenialEpisodeReasonDenyReasonAppPausedByYou`   | DENY_REASON_APP_PAUSED_BY_YOU                   |
| `DenialEpisodeReasonDenyReasonNoSupply`         | DENY_REASON_NO_SUPPLY                           |