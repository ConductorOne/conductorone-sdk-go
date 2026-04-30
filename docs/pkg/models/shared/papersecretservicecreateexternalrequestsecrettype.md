# PaperSecretServiceCreateExternalRequestSecretType

Secret type: TEXT or FILE.
 TEXT secrets use SetTextContent to upload encrypted content (max 64KB).
 FILE secrets use the upload_url from CreateResponse to upload encrypted content (max 1GB).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PaperSecretServiceCreateExternalRequestSecretTypeSecretTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PaperSecretServiceCreateExternalRequestSecretType("custom_value")
```


## Values

| Name                                                                     | Value                                                                    |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `PaperSecretServiceCreateExternalRequestSecretTypeSecretTypeUnspecified` | SECRET_TYPE_UNSPECIFIED                                                  |
| `PaperSecretServiceCreateExternalRequestSecretTypeSecretTypeText`        | SECRET_TYPE_TEXT                                                         |
| `PaperSecretServiceCreateExternalRequestSecretTypeSecretTypeFile`        | SECRET_TYPE_FILE                                                         |