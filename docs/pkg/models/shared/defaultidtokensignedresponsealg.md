# DefaultIDTokenSignedResponseAlg

The id_token signing algorithm applied to OIDC applications that do not
 choose one. When unset, the server uses EdDSA.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DefaultIDTokenSignedResponseAlgOidcSigningAlgorithmUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DefaultIDTokenSignedResponseAlg("custom_value")
```


## Values

| Name                                                             | Value                                                            |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `DefaultIDTokenSignedResponseAlgOidcSigningAlgorithmUnspecified` | OIDC_SIGNING_ALGORITHM_UNSPECIFIED                               |
| `DefaultIDTokenSignedResponseAlgOidcSigningAlgorithmEddsa`       | OIDC_SIGNING_ALGORITHM_EDDSA                                     |
| `DefaultIDTokenSignedResponseAlgOidcSigningAlgorithmEs256`       | OIDC_SIGNING_ALGORITHM_ES256                                     |
| `DefaultIDTokenSignedResponseAlgOidcSigningAlgorithmRs256`       | OIDC_SIGNING_ALGORITHM_RS256                                     |