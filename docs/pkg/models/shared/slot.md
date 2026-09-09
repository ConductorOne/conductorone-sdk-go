# Slot

Which of the object's local-policy slots this policy occupies. Part of the
 scope, and immutable with it.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SlotPolicyScopeSlotUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Slot("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `SlotPolicyScopeSlotUnspecified` | POLICY_SCOPE_SLOT_UNSPECIFIED    |
| `SlotPolicyScopeSlotEmergency`   | POLICY_SCOPE_SLOT_EMERGENCY      |