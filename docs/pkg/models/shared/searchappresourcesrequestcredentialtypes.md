# SearchAppResourcesRequestCredentialTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SearchAppResourcesRequestCredentialTypesCredentialTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SearchAppResourcesRequestCredentialTypes("custom_value")
```


## Values

| Name                                                                  | Value                                                                 |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `SearchAppResourcesRequestCredentialTypesCredentialTypeUnspecified`   | CREDENTIAL_TYPE_UNSPECIFIED                                           |
| `SearchAppResourcesRequestCredentialTypesCredentialTypeStaticSecret`  | CREDENTIAL_TYPE_STATIC_SECRET                                         |
| `SearchAppResourcesRequestCredentialTypesCredentialTypeAsymmetricKey` | CREDENTIAL_TYPE_ASYMMETRIC_KEY                                        |
| `SearchAppResourcesRequestCredentialTypesCredentialTypeCertificate`   | CREDENTIAL_TYPE_CERTIFICATE                                           |