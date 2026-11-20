# ScopeKind

Budget scope that denied the calls.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ScopeKindSpendScopeKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ScopeKind("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `ScopeKindSpendScopeKindUnspecified` | SPEND_SCOPE_KIND_UNSPECIFIED         |
| `ScopeKindSpendScopeKindTenant`      | SPEND_SCOPE_KIND_TENANT              |
| `ScopeKindSpendScopeKindSubject`     | SPEND_SCOPE_KIND_SUBJECT             |
| `ScopeKindSpendScopeKindApp`         | SPEND_SCOPE_KIND_APP                 |
| `ScopeKindSpendScopeKindSubjectApp`  | SPEND_SCOPE_KIND_SUBJECT_APP         |