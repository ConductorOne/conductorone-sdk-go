# ResourceType

The resourceType field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ResourceTypeRole

// Open enum: custom values can be created with a direct type cast
custom := shared.ResourceType("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `ResourceTypeRole`        | ROLE                      |
| `ResourceTypeGroup`       | GROUP                     |
| `ResourceTypeLicense`     | LICENSE                   |
| `ResourceTypeProject`     | PROJECT                   |
| `ResourceTypeCatalog`     | CATALOG                   |
| `ResourceTypeCustom`      | CUSTOM                    |
| `ResourceTypeVault`       | VAULT                     |
| `ResourceTypeProfileType` | PROFILE_TYPE              |