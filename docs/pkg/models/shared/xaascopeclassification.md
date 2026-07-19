# XAAScopeClassification

Risk classification.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.XAAScopeClassificationXaaScopeClassificationUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.XAAScopeClassification("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `XAAScopeClassificationXaaScopeClassificationUnspecified` | XAA_SCOPE_CLASSIFICATION_UNSPECIFIED                      |
| `XAAScopeClassificationXaaScopeClassificationRead`        | XAA_SCOPE_CLASSIFICATION_READ                             |
| `XAAScopeClassificationXaaScopeClassificationWrite`       | XAA_SCOPE_CLASSIFICATION_WRITE                            |
| `XAAScopeClassificationXaaScopeClassificationDestructive` | XAA_SCOPE_CLASSIFICATION_DESTRUCTIVE                      |
| `XAAScopeClassificationXaaScopeClassificationSensitive`   | XAA_SCOPE_CLASSIFICATION_SENSITIVE                        |
| `XAAScopeClassificationXaaScopeClassificationDangerous`   | XAA_SCOPE_CLASSIFICATION_DANGEROUS                        |