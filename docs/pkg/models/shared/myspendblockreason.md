# MySpendBlockReason

Reason the calls were denied.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MySpendBlockReasonDenyReasonUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MySpendBlockReason("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `MySpendBlockReasonDenyReasonUnspecified`      | DENY_REASON_UNSPECIFIED                        |
| `MySpendBlockReasonDenyReasonTenantFrozen`     | DENY_REASON_TENANT_FROZEN                      |
| `MySpendBlockReasonDenyReasonSuspendedByAdmin` | DENY_REASON_SUSPENDED_BY_ADMIN                 |
| `MySpendBlockReasonDenyReasonAppSuspended`     | DENY_REASON_APP_SUSPENDED                      |
| `MySpendBlockReasonDenyReasonAppPausedByYou`   | DENY_REASON_APP_PAUSED_BY_YOU                  |
| `MySpendBlockReasonDenyReasonNoSupply`         | DENY_REASON_NO_SUPPLY                          |