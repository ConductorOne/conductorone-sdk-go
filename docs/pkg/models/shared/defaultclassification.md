# DefaultClassification

Default tool classification from MCP config (system-managed during discovery).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DefaultClassificationToolClassificationUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DefaultClassification("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `DefaultClassificationToolClassificationUnspecified` | TOOL_CLASSIFICATION_UNSPECIFIED                      |
| `DefaultClassificationToolClassificationRead`        | TOOL_CLASSIFICATION_READ                             |
| `DefaultClassificationToolClassificationWrite`       | TOOL_CLASSIFICATION_WRITE                            |
| `DefaultClassificationToolClassificationDestructive` | TOOL_CLASSIFICATION_DESTRUCTIVE                      |
| `DefaultClassificationToolClassificationSensitive`   | TOOL_CLASSIFICATION_SENSITIVE                        |
| `DefaultClassificationToolClassificationDangerous`   | TOOL_CLASSIFICATION_DANGEROUS                        |