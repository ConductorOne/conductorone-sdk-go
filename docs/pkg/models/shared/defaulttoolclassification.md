# DefaultToolClassification

Classification assigned to newly discovered tools that do not declare their
 own classification (for example, read, write, or destructive).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DefaultToolClassificationToolClassificationUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DefaultToolClassification("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `DefaultToolClassificationToolClassificationUnspecified` | TOOL_CLASSIFICATION_UNSPECIFIED                          |
| `DefaultToolClassificationToolClassificationRead`        | TOOL_CLASSIFICATION_READ                                 |
| `DefaultToolClassificationToolClassificationWrite`       | TOOL_CLASSIFICATION_WRITE                                |
| `DefaultToolClassificationToolClassificationDestructive` | TOOL_CLASSIFICATION_DESTRUCTIVE                          |
| `DefaultToolClassificationToolClassificationSensitive`   | TOOL_CLASSIFICATION_SENSITIVE                            |
| `DefaultToolClassificationToolClassificationDangerous`   | TOOL_CLASSIFICATION_DANGEROUS                            |