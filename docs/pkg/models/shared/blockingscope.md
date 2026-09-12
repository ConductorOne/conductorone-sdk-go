# BlockingScope

The blockingScope field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.BlockingScopeSpendScopeKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.BlockingScope("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `BlockingScopeSpendScopeKindUnspecified` | SPEND_SCOPE_KIND_UNSPECIFIED             |
| `BlockingScopeSpendScopeKindTenant`      | SPEND_SCOPE_KIND_TENANT                  |
| `BlockingScopeSpendScopeKindSubject`     | SPEND_SCOPE_KIND_SUBJECT                 |
| `BlockingScopeSpendScopeKindApp`         | SPEND_SCOPE_KIND_APP                     |
| `BlockingScopeSpendScopeKindSubjectApp`  | SPEND_SCOPE_KIND_SUBJECT_APP             |