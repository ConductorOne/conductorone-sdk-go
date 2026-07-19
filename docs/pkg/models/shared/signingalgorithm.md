# SigningAlgorithm

JWS algorithm for grants minted for this server. UNSPECIFIED uses the
 tenant default. Minting fails if no active signing key exists for the
 resolved algorithm.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SigningAlgorithmXaaSigningAlgorithmUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SigningAlgorithm("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `SigningAlgorithmXaaSigningAlgorithmUnspecified` | XAA_SIGNING_ALGORITHM_UNSPECIFIED                |
| `SigningAlgorithmXaaSigningAlgorithmEddsa`       | XAA_SIGNING_ALGORITHM_EDDSA                      |
| `SigningAlgorithmXaaSigningAlgorithmRs256`       | XAA_SIGNING_ALGORITHM_RS256                      |
| `SigningAlgorithmXaaSigningAlgorithmEs256`       | XAA_SIGNING_ALGORITHM_ES256                      |