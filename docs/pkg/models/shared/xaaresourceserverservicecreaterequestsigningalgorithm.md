# XAAResourceServerServiceCreateRequestSigningAlgorithm

JWS algorithm for grants minted for this server. UNSPECIFIED uses the
 tenant default.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.XAAResourceServerServiceCreateRequestSigningAlgorithmXaaSigningAlgorithmUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.XAAResourceServerServiceCreateRequestSigningAlgorithm("custom_value")
```


## Values

| Name                                                                                  | Value                                                                                 |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `XAAResourceServerServiceCreateRequestSigningAlgorithmXaaSigningAlgorithmUnspecified` | XAA_SIGNING_ALGORITHM_UNSPECIFIED                                                     |
| `XAAResourceServerServiceCreateRequestSigningAlgorithmXaaSigningAlgorithmEddsa`       | XAA_SIGNING_ALGORITHM_EDDSA                                                           |
| `XAAResourceServerServiceCreateRequestSigningAlgorithmXaaSigningAlgorithmRs256`       | XAA_SIGNING_ALGORITHM_RS256                                                           |
| `XAAResourceServerServiceCreateRequestSigningAlgorithmXaaSigningAlgorithmEs256`       | XAA_SIGNING_ALGORITHM_ES256                                                           |