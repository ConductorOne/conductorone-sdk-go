# ScopeSlot

When scope_view narrows to one object, only return that object's local
 policies in this slot. Ignored when no object is identified by
 scope_app_id, which lists every local policy regardless of slot.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ScopeSlotPolicyScopeSlotUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ScopeSlot("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `ScopeSlotPolicyScopeSlotUnspecified` | POLICY_SCOPE_SLOT_UNSPECIFIED         |
| `ScopeSlotPolicyScopeSlotEmergency`   | POLICY_SCOPE_SLOT_EMERGENCY           |