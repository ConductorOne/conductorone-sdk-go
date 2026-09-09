# MySpendBlockScopeKind

Budget scope that denied the calls.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.MySpendBlockScopeKindSpendBlockScopeKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.MySpendBlockScopeKind("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `MySpendBlockScopeKindSpendBlockScopeKindUnspecified` | SPEND_BLOCK_SCOPE_KIND_UNSPECIFIED                    |
| `MySpendBlockScopeKindSpendBlockScopeKindTenant`      | SPEND_BLOCK_SCOPE_KIND_TENANT                         |
| `MySpendBlockScopeKindSpendBlockScopeKindSubject`     | SPEND_BLOCK_SCOPE_KIND_SUBJECT                        |
| `MySpendBlockScopeKindSpendBlockScopeKindApp`         | SPEND_BLOCK_SCOPE_KIND_APP                            |
| `MySpendBlockScopeKindSpendBlockScopeKindSubjectApp`  | SPEND_BLOCK_SCOPE_KIND_SUBJECT_APP                    |