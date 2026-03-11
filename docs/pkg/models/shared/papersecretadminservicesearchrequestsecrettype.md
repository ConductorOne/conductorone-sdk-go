# PaperSecretAdminServiceSearchRequestSecretType

Filter by secret type (optional)

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PaperSecretAdminServiceSearchRequestSecretTypeSecretTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PaperSecretAdminServiceSearchRequestSecretType("custom_value")
```


## Values

| Name                                                                  | Value                                                                 |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `PaperSecretAdminServiceSearchRequestSecretTypeSecretTypeUnspecified` | SECRET_TYPE_UNSPECIFIED                                               |
| `PaperSecretAdminServiceSearchRequestSecretTypeSecretTypeText`        | SECRET_TYPE_TEXT                                                      |
| `PaperSecretAdminServiceSearchRequestSecretTypeSecretTypeFile`        | SECRET_TYPE_FILE                                                      |