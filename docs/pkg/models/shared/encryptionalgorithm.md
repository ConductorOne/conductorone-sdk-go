# EncryptionAlgorithm

The algorithm used when encrypt_assertions is set.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EncryptionAlgorithmSamlEncryptionAlgorithmUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.EncryptionAlgorithm("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `EncryptionAlgorithmSamlEncryptionAlgorithmUnspecified` | SAML_ENCRYPTION_ALGORITHM_UNSPECIFIED                   |
| `EncryptionAlgorithmSamlEncryptionAlgorithmAes256Gcm`   | SAML_ENCRYPTION_ALGORITHM_AES256_GCM                    |
| `EncryptionAlgorithmSamlEncryptionAlgorithmAes128Gcm`   | SAML_ENCRYPTION_ALGORITHM_AES128_GCM                    |
| `EncryptionAlgorithmSamlEncryptionAlgorithmAes256Cbc`   | SAML_ENCRYPTION_ALGORITHM_AES256_CBC                    |