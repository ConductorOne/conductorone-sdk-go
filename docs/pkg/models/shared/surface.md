# Surface

The surface field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SurfaceClassifierTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Surface("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `SurfaceClassifierTypeUnspecified` | CLASSIFIER_TYPE_UNSPECIFIED        |
| `SurfaceClassifierTypeAgent`       | CLASSIFIER_TYPE_AGENT              |
| `SurfaceClassifierTypeGateway`     | CLASSIFIER_TYPE_GATEWAY            |