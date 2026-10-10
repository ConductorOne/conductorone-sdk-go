# ManagedBy

The managedBy field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ManagedByManagedByUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ManagedBy("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `ManagedByManagedByUnspecified` | MANAGED_BY_UNSPECIFIED          |
| `ManagedByManagedByTenant`      | MANAGED_BY_TENANT               |
| `ManagedByManagedBySystem`      | MANAGED_BY_SYSTEM               |