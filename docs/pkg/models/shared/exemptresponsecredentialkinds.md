# ExemptResponseCredentialKinds

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ExemptResponseCredentialKindsCredentialKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ExemptResponseCredentialKinds("custom_value")
```


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ExemptResponseCredentialKindsCredentialKindUnspecified`      | CREDENTIAL_KIND_UNSPECIFIED                                   |
| `ExemptResponseCredentialKindsCredentialKindJwt`              | CREDENTIAL_KIND_JWT                                           |
| `ExemptResponseCredentialKindsCredentialKindAPIKey`           | CREDENTIAL_KIND_API_KEY                                       |
| `ExemptResponseCredentialKindsCredentialKindPrivateKey`       | CREDENTIAL_KIND_PRIVATE_KEY                                   |
| `ExemptResponseCredentialKindsCredentialKindVendorCredential` | CREDENTIAL_KIND_VENDOR_CREDENTIAL                             |
| `ExemptResponseCredentialKindsCredentialKindPan`              | CREDENTIAL_KIND_PAN                                           |