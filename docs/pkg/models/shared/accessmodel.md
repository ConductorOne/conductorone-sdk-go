# AccessModel

How this app models access. Derived during uplift from the app's resource type traits.
 Sparse ACL feature.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AccessModelAppAccessModelUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AccessModel("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `AccessModelAppAccessModelUnspecified` | APP_ACCESS_MODEL_UNSPECIFIED           |
| `AccessModelAppAccessModelClassic`     | APP_ACCESS_MODEL_CLASSIC               |
| `AccessModelAppAccessModelHybrid`      | APP_ACCESS_MODEL_HYBRID                |
| `AccessModelAppAccessModelSparse`      | APP_ACCESS_MODEL_SPARSE                |