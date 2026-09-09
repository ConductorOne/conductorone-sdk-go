# ScopeObjectType

When scope_view is POLICY_SCOPE_VIEW_SCOPED, narrow local policies to a
 coarse object type (app-local vs entitlement-local).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ScopeObjectTypePolicyScopeObjectTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ScopeObjectType("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `ScopeObjectTypePolicyScopeObjectTypeUnspecified` | POLICY_SCOPE_OBJECT_TYPE_UNSPECIFIED              |
| `ScopeObjectTypePolicyScopeObjectTypeApp`         | POLICY_SCOPE_OBJECT_TYPE_APP                      |
| `ScopeObjectTypePolicyScopeObjectTypeEntitlement` | POLICY_SCOPE_OBJECT_TYPE_ENTITLEMENT              |