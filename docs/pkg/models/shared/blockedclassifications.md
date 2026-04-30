# BlockedClassifications

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.BlockedClassificationsToolClassificationUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.BlockedClassifications("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `BlockedClassificationsToolClassificationUnspecified` | TOOL_CLASSIFICATION_UNSPECIFIED                       |
| `BlockedClassificationsToolClassificationRead`        | TOOL_CLASSIFICATION_READ                              |
| `BlockedClassificationsToolClassificationWrite`       | TOOL_CLASSIFICATION_WRITE                             |
| `BlockedClassificationsToolClassificationDestructive` | TOOL_CLASSIFICATION_DESTRUCTIVE                       |
| `BlockedClassificationsToolClassificationSensitive`   | TOOL_CLASSIFICATION_SENSITIVE                         |
| `BlockedClassificationsToolClassificationDangerous`   | TOOL_CLASSIFICATION_DANGEROUS                         |