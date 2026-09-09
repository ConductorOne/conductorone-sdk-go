# ScopeKind

Budget scope that denied the calls.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ScopeKindSpendBlockScopeKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ScopeKind("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `ScopeKindSpendBlockScopeKindUnspecified` | SPEND_BLOCK_SCOPE_KIND_UNSPECIFIED        |
| `ScopeKindSpendBlockScopeKindTenant`      | SPEND_BLOCK_SCOPE_KIND_TENANT             |
| `ScopeKindSpendBlockScopeKindSubject`     | SPEND_BLOCK_SCOPE_KIND_SUBJECT            |
| `ScopeKindSpendBlockScopeKindApp`         | SPEND_BLOCK_SCOPE_KIND_APP                |
| `ScopeKindSpendBlockScopeKindSubjectApp`  | SPEND_BLOCK_SCOPE_KIND_SUBJECT_APP        |