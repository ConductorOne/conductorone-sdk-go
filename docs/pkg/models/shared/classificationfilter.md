# ClassificationFilter

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ClassificationFilterToolClassificationUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ClassificationFilter("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `ClassificationFilterToolClassificationUnspecified` | TOOL_CLASSIFICATION_UNSPECIFIED                     |
| `ClassificationFilterToolClassificationRead`        | TOOL_CLASSIFICATION_READ                            |
| `ClassificationFilterToolClassificationWrite`       | TOOL_CLASSIFICATION_WRITE                           |
| `ClassificationFilterToolClassificationDestructive` | TOOL_CLASSIFICATION_DESTRUCTIVE                     |
| `ClassificationFilterToolClassificationSensitive`   | TOOL_CLASSIFICATION_SENSITIVE                       |
| `ClassificationFilterToolClassificationDangerous`   | TOOL_CLASSIFICATION_DANGEROUS                       |