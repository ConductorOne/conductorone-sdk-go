# ClassifierTemplateManagedBy

The managedBy field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ClassifierTemplateManagedByManagedByUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ClassifierTemplateManagedBy("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `ClassifierTemplateManagedByManagedByUnspecified` | MANAGED_BY_UNSPECIFIED                            |
| `ClassifierTemplateManagedByManagedByTenant`      | MANAGED_BY_TENANT                                 |
| `ClassifierTemplateManagedByManagedBySystem`      | MANAGED_BY_SYSTEM                                 |