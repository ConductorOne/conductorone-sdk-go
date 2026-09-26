# ExemptCredentialKinds

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ExemptCredentialKindsCredentialKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ExemptCredentialKinds("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `ExemptCredentialKindsCredentialKindUnspecified`      | CREDENTIAL_KIND_UNSPECIFIED                           |
| `ExemptCredentialKindsCredentialKindJwt`              | CREDENTIAL_KIND_JWT                                   |
| `ExemptCredentialKindsCredentialKindAPIKey`           | CREDENTIAL_KIND_API_KEY                               |
| `ExemptCredentialKindsCredentialKindPrivateKey`       | CREDENTIAL_KIND_PRIVATE_KEY                           |
| `ExemptCredentialKindsCredentialKindVendorCredential` | CREDENTIAL_KIND_VENDOR_CREDENTIAL                     |