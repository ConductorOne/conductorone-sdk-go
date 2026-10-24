# TBAuditAttestationKeyStatus

The status field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.TBAuditAttestationKeyStatusTbAuditAttestationKeyStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.TBAuditAttestationKeyStatus("custom_value")
```


## Values

| Name                                                                | Value                                                               |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `TBAuditAttestationKeyStatusTbAuditAttestationKeyStatusUnspecified` | TB_AUDIT_ATTESTATION_KEY_STATUS_UNSPECIFIED                         |
| `TBAuditAttestationKeyStatusTbAuditAttestationKeyStatusCurrent`     | TB_AUDIT_ATTESTATION_KEY_STATUS_CURRENT                             |
| `TBAuditAttestationKeyStatusTbAuditAttestationKeyStatusPrevious`    | TB_AUDIT_ATTESTATION_KEY_STATUS_PREVIOUS                            |