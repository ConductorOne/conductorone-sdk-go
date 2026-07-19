# XAAScopeServiceCreateRequestClassification

Risk classification.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.XAAScopeServiceCreateRequestClassificationXaaScopeClassificationUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.XAAScopeServiceCreateRequestClassification("custom_value")
```


## Values

| Name                                                                          | Value                                                                         |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `XAAScopeServiceCreateRequestClassificationXaaScopeClassificationUnspecified` | XAA_SCOPE_CLASSIFICATION_UNSPECIFIED                                          |
| `XAAScopeServiceCreateRequestClassificationXaaScopeClassificationRead`        | XAA_SCOPE_CLASSIFICATION_READ                                                 |
| `XAAScopeServiceCreateRequestClassificationXaaScopeClassificationWrite`       | XAA_SCOPE_CLASSIFICATION_WRITE                                                |
| `XAAScopeServiceCreateRequestClassificationXaaScopeClassificationDestructive` | XAA_SCOPE_CLASSIFICATION_DESTRUCTIVE                                          |
| `XAAScopeServiceCreateRequestClassificationXaaScopeClassificationSensitive`   | XAA_SCOPE_CLASSIFICATION_SENSITIVE                                            |
| `XAAScopeServiceCreateRequestClassificationXaaScopeClassificationDangerous`   | XAA_SCOPE_CLASSIFICATION_DANGEROUS                                            |