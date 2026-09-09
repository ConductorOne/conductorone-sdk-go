# DenialReason

Reason spend is denied. Present only when state is DENIED.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DenialReasonDenyReasonUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DenialReason("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `DenialReasonDenyReasonUnspecified`      | DENY_REASON_UNSPECIFIED                  |
| `DenialReasonDenyReasonTenantFrozen`     | DENY_REASON_TENANT_FROZEN                |
| `DenialReasonDenyReasonSuspendedByAdmin` | DENY_REASON_SUSPENDED_BY_ADMIN           |
| `DenialReasonDenyReasonAppSuspended`     | DENY_REASON_APP_SUSPENDED                |
| `DenialReasonDenyReasonAppPausedByYou`   | DENY_REASON_APP_PAUSED_BY_YOU            |
| `DenialReasonDenyReasonNoSupply`         | DENY_REASON_NO_SUPPLY                    |