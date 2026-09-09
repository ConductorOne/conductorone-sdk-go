# ScopeView

Which policies to return based on scope. Defaults to global-only, so
 app/entitlement-scoped policies never appear unless explicitly requested.
 Ignored when refs are provided (explicit ID lookups always resolve).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ScopeViewPolicyScopeViewUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ScopeView("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `ScopeViewPolicyScopeViewUnspecified`     | POLICY_SCOPE_VIEW_UNSPECIFIED             |
| `ScopeViewPolicyScopeViewGlobal`          | POLICY_SCOPE_VIEW_GLOBAL                  |
| `ScopeViewPolicyScopeViewScoped`          | POLICY_SCOPE_VIEW_SCOPED                  |
| `ScopeViewPolicyScopeViewAll`             | POLICY_SCOPE_VIEW_ALL                     |
| `ScopeViewPolicyScopeViewGlobalAndObject` | POLICY_SCOPE_VIEW_GLOBAL_AND_OBJECT       |