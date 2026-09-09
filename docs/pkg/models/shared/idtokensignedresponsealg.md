# IDTokenSignedResponseAlg

The algorithm used to sign this application's id_token.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.IDTokenSignedResponseAlgOidcSigningAlgorithmUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.IDTokenSignedResponseAlg("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `IDTokenSignedResponseAlgOidcSigningAlgorithmUnspecified` | OIDC_SIGNING_ALGORITHM_UNSPECIFIED                        |
| `IDTokenSignedResponseAlgOidcSigningAlgorithmEddsa`       | OIDC_SIGNING_ALGORITHM_EDDSA                              |
| `IDTokenSignedResponseAlgOidcSigningAlgorithmEs256`       | OIDC_SIGNING_ALGORITHM_ES256                              |
| `IDTokenSignedResponseAlgOidcSigningAlgorithmRs256`       | OIDC_SIGNING_ALGORITHM_RS256                              |