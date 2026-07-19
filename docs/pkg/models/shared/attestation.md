# Attestation

How strictly the authenticator's origin must be attested.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AttestationAttestationRequirementUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Attestation("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `AttestationAttestationRequirementUnspecified` | ATTESTATION_REQUIREMENT_UNSPECIFIED            |
| `AttestationAttestationRequirementNone`        | ATTESTATION_REQUIREMENT_NONE                   |
| `AttestationAttestationRequirementIndirect`    | ATTESTATION_REQUIREMENT_INDIRECT               |
| `AttestationAttestationRequirementDirect`      | ATTESTATION_REQUIREMENT_DIRECT                 |
| `AttestationAttestationRequirementEnterprise`  | ATTESTATION_REQUIREMENT_ENTERPRISE             |