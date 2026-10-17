# DenialReason

Reason spend is denied. Present only when state is DENIED.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DenialReasonSpendDenyReasonUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DenialReason("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `DenialReasonSpendDenyReasonUnspecified`      | SPEND_DENY_REASON_UNSPECIFIED                 |
| `DenialReasonSpendDenyReasonTenantFrozen`     | SPEND_DENY_REASON_TENANT_FROZEN               |
| `DenialReasonSpendDenyReasonSuspendedByAdmin` | SPEND_DENY_REASON_SUSPENDED_BY_ADMIN          |
| `DenialReasonSpendDenyReasonAppSuspended`     | SPEND_DENY_REASON_APP_SUSPENDED               |
| `DenialReasonSpendDenyReasonAppPausedByYou`   | SPEND_DENY_REASON_APP_PAUSED_BY_YOU           |
| `DenialReasonSpendDenyReasonNoSupply`         | SPEND_DENY_REASON_NO_SUPPLY                   |