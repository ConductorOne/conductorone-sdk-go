# BlockedMutatingClassifications

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.BlockedMutatingClassificationsToolClassificationUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.BlockedMutatingClassifications("custom_value")
```


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `BlockedMutatingClassificationsToolClassificationUnspecified` | TOOL_CLASSIFICATION_UNSPECIFIED                               |
| `BlockedMutatingClassificationsToolClassificationRead`        | TOOL_CLASSIFICATION_READ                                      |
| `BlockedMutatingClassificationsToolClassificationWrite`       | TOOL_CLASSIFICATION_WRITE                                     |
| `BlockedMutatingClassificationsToolClassificationDestructive` | TOOL_CLASSIFICATION_DESTRUCTIVE                               |
| `BlockedMutatingClassificationsToolClassificationSensitive`   | TOOL_CLASSIFICATION_SENSITIVE                                 |
| `BlockedMutatingClassificationsToolClassificationDangerous`   | TOOL_CLASSIFICATION_DANGEROUS                                 |