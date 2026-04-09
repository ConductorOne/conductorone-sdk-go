# CertifyOutcomes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CertifyOutcomesCertifyOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.CertifyOutcomes("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `CertifyOutcomesCertifyOutcomeUnspecified`  | CERTIFY_OUTCOME_UNSPECIFIED                 |
| `CertifyOutcomesCertifyOutcomeCertified`    | CERTIFY_OUTCOME_CERTIFIED                   |
| `CertifyOutcomesCertifyOutcomeDecertified`  | CERTIFY_OUTCOME_DECERTIFIED                 |
| `CertifyOutcomesCertifyOutcomeError`        | CERTIFY_OUTCOME_ERROR                       |
| `CertifyOutcomesCertifyOutcomeCancelled`    | CERTIFY_OUTCOME_CANCELLED                   |
| `CertifyOutcomesCertifyOutcomeWaitTimedOut` | CERTIFY_OUTCOME_WAIT_TIMED_OUT              |