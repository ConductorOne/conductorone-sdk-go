# PaperSecretServiceCreateInternalRequestSecretType

Secret type: TEXT or FILE.
 TEXT secrets use SetTextContent to upload encrypted content (max 64KB).
 FILE secrets use the upload_url from CreateResponse to upload encrypted content (max 1GB).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PaperSecretServiceCreateInternalRequestSecretTypeSecretTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PaperSecretServiceCreateInternalRequestSecretType("custom_value")
```


## Values

| Name                                                                     | Value                                                                    |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `PaperSecretServiceCreateInternalRequestSecretTypeSecretTypeUnspecified` | SECRET_TYPE_UNSPECIFIED                                                  |
| `PaperSecretServiceCreateInternalRequestSecretTypeSecretTypeText`        | SECRET_TYPE_TEXT                                                         |
| `PaperSecretServiceCreateInternalRequestSecretTypeSecretTypeFile`        | SECRET_TYPE_FILE                                                         |