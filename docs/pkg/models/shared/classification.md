# Classification

Tool risk classification for policy decisions.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ClassificationToolClassificationUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Classification("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `ClassificationToolClassificationUnspecified` | TOOL_CLASSIFICATION_UNSPECIFIED               |
| `ClassificationToolClassificationRead`        | TOOL_CLASSIFICATION_READ                      |
| `ClassificationToolClassificationWrite`       | TOOL_CLASSIFICATION_WRITE                     |
| `ClassificationToolClassificationDestructive` | TOOL_CLASSIFICATION_DESTRUCTIVE               |
| `ClassificationToolClassificationSensitive`   | TOOL_CLASSIFICATION_SENSITIVE                 |
| `ClassificationToolClassificationDangerous`   | TOOL_CLASSIFICATION_DANGEROUS                 |