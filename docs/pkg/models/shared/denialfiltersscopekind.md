# DenialFiltersScopeKind

Optional budget scope to match.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DenialFiltersScopeKindSpendBlockScopeKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DenialFiltersScopeKind("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `DenialFiltersScopeKindSpendBlockScopeKindUnspecified` | SPEND_BLOCK_SCOPE_KIND_UNSPECIFIED                     |
| `DenialFiltersScopeKindSpendBlockScopeKindTenant`      | SPEND_BLOCK_SCOPE_KIND_TENANT                          |
| `DenialFiltersScopeKindSpendBlockScopeKindSubject`     | SPEND_BLOCK_SCOPE_KIND_SUBJECT                         |
| `DenialFiltersScopeKindSpendBlockScopeKindApp`         | SPEND_BLOCK_SCOPE_KIND_APP                             |
| `DenialFiltersScopeKindSpendBlockScopeKindSubjectApp`  | SPEND_BLOCK_SCOPE_KIND_SUBJECT_APP                     |