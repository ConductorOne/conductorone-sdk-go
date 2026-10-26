# DenialEpisodeReason

Reason the calls were denied.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DenialEpisodeReasonSpendDenyReasonUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DenialEpisodeReason("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `DenialEpisodeReasonSpendDenyReasonUnspecified`      | SPEND_DENY_REASON_UNSPECIFIED                        |
| `DenialEpisodeReasonSpendDenyReasonTenantFrozen`     | SPEND_DENY_REASON_TENANT_FROZEN                      |
| `DenialEpisodeReasonSpendDenyReasonSuspendedByAdmin` | SPEND_DENY_REASON_SUSPENDED_BY_ADMIN                 |
| `DenialEpisodeReasonSpendDenyReasonAppSuspended`     | SPEND_DENY_REASON_APP_SUSPENDED                      |
| `DenialEpisodeReasonSpendDenyReasonAppPausedByYou`   | SPEND_DENY_REASON_APP_PAUSED_BY_YOU                  |
| `DenialEpisodeReasonSpendDenyReasonNoSupply`         | SPEND_DENY_REASON_NO_SUPPLY                          |