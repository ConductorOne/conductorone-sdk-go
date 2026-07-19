# NhiType

The NHI classification (K3 spine) for this resource. Populated for
 non-human-identity resources; UNSPECIFIED for everything else. Mirrors
 agent_trait: read-only and translated from the model enum at the API boundary.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.NhiTypeNhiTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.NhiType("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `NhiTypeNhiTypeUnspecified`     | NHI_TYPE_UNSPECIFIED            |
| `NhiTypeNhiTypeAppRegistration` | NHI_TYPE_APP_REGISTRATION       |
| `NhiTypeNhiTypeAssumableRole`   | NHI_TYPE_ASSUMABLE_ROLE         |
| `NhiTypeNhiTypeManagedIdentity` | NHI_TYPE_MANAGED_IDENTITY       |