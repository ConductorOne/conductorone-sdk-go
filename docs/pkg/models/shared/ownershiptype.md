# OwnershipType

The type of ownership.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.OwnershipTypeUserOwnershipTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.OwnershipType("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `OwnershipTypeUserOwnershipTypeUnspecified` | USER_OWNERSHIP_TYPE_UNSPECIFIED             |
| `OwnershipTypeUserOwnershipTypeApp`         | USER_OWNERSHIP_TYPE_APP                     |
| `OwnershipTypeUserOwnershipTypeResource`    | USER_OWNERSHIP_TYPE_RESOURCE                |
| `OwnershipTypeUserOwnershipTypeEntitlement` | USER_OWNERSHIP_TYPE_ENTITLEMENT             |