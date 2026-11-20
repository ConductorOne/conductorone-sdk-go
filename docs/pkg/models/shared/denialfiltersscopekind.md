# DenialFiltersScopeKind

Optional budget scope to match.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DenialFiltersScopeKindSpendScopeKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DenialFiltersScopeKind("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `DenialFiltersScopeKindSpendScopeKindUnspecified` | SPEND_SCOPE_KIND_UNSPECIFIED                      |
| `DenialFiltersScopeKindSpendScopeKindTenant`      | SPEND_SCOPE_KIND_TENANT                           |
| `DenialFiltersScopeKindSpendScopeKindSubject`     | SPEND_SCOPE_KIND_SUBJECT                          |
| `DenialFiltersScopeKindSpendScopeKindApp`         | SPEND_SCOPE_KIND_APP                              |
| `DenialFiltersScopeKindSpendScopeKindSubjectApp`  | SPEND_SCOPE_KIND_SUBJECT_APP                      |