# ApplicableType

The applicableType field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ApplicableTypeClassifierTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ApplicableType("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `ApplicableTypeClassifierTypeUnspecified` | CLASSIFIER_TYPE_UNSPECIFIED               |
| `ApplicableTypeClassifierTypeAgent`       | CLASSIFIER_TYPE_AGENT                     |
| `ApplicableTypeClassifierTypeGateway`     | CLASSIFIER_TYPE_GATEWAY                   |