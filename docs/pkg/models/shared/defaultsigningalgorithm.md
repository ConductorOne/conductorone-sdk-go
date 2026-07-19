# DefaultSigningAlgorithm

Tenant-default signing algorithm. UNSPECIFIED resolves to ES256.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DefaultSigningAlgorithmXaaSigningAlgorithmUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DefaultSigningAlgorithm("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `DefaultSigningAlgorithmXaaSigningAlgorithmUnspecified` | XAA_SIGNING_ALGORITHM_UNSPECIFIED                       |
| `DefaultSigningAlgorithmXaaSigningAlgorithmEddsa`       | XAA_SIGNING_ALGORITHM_EDDSA                             |
| `DefaultSigningAlgorithmXaaSigningAlgorithmRs256`       | XAA_SIGNING_ALGORITHM_RS256                             |
| `DefaultSigningAlgorithmXaaSigningAlgorithmEs256`       | XAA_SIGNING_ALGORITHM_ES256                             |