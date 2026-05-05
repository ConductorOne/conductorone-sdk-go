# CanonicalType

C1 canonical outcome (what C1 understood and did).
 The normalized event type after mapping from the wire event type.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CanonicalTypeSsfCanonicalEventTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.CanonicalType("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `CanonicalTypeSsfCanonicalEventTypeUnspecified`             | SSF_CANONICAL_EVENT_TYPE_UNSPECIFIED                        |
| `CanonicalTypeSsfCanonicalEventTypeUnrecognized`            | SSF_CANONICAL_EVENT_TYPE_UNRECOGNIZED                       |
| `CanonicalTypeSsfCanonicalEventTypeSessionRevoked`          | SSF_CANONICAL_EVENT_TYPE_SESSION_REVOKED                    |
| `CanonicalTypeSsfCanonicalEventTypeCredentialChanged`       | SSF_CANONICAL_EVENT_TYPE_CREDENTIAL_CHANGED                 |
| `CanonicalTypeSsfCanonicalEventTypeTokenClaimsChanged`      | SSF_CANONICAL_EVENT_TYPE_TOKEN_CLAIMS_CHANGED               |
| `CanonicalTypeSsfCanonicalEventTypeAssuranceLevelChanged`   | SSF_CANONICAL_EVENT_TYPE_ASSURANCE_LEVEL_CHANGED            |
| `CanonicalTypeSsfCanonicalEventTypeDeviceComplianceChanged` | SSF_CANONICAL_EVENT_TYPE_DEVICE_COMPLIANCE_CHANGED          |
| `CanonicalTypeSsfCanonicalEventTypeRiskLevelChanged`        | SSF_CANONICAL_EVENT_TYPE_RISK_LEVEL_CHANGED                 |
| `CanonicalTypeSsfCanonicalEventTypeSessionEstablished`      | SSF_CANONICAL_EVENT_TYPE_SESSION_ESTABLISHED                |
| `CanonicalTypeSsfCanonicalEventTypeSessionPresented`        | SSF_CANONICAL_EVENT_TYPE_SESSION_PRESENTED                  |
| `CanonicalTypeSsfCanonicalEventTypeAccountDisabled`         | SSF_CANONICAL_EVENT_TYPE_ACCOUNT_DISABLED                   |
| `CanonicalTypeSsfCanonicalEventTypeAccountEnabled`          | SSF_CANONICAL_EVENT_TYPE_ACCOUNT_ENABLED                    |
| `CanonicalTypeSsfCanonicalEventTypeAccountPurged`           | SSF_CANONICAL_EVENT_TYPE_ACCOUNT_PURGED                     |
| `CanonicalTypeSsfCanonicalEventTypeCredentialCompromise`    | SSF_CANONICAL_EVENT_TYPE_CREDENTIAL_COMPROMISE              |
| `CanonicalTypeSsfCanonicalEventTypeRecoveryActivated`       | SSF_CANONICAL_EVENT_TYPE_RECOVERY_ACTIVATED                 |
| `CanonicalTypeSsfCanonicalEventTypeIdentifierChanged`       | SSF_CANONICAL_EVENT_TYPE_IDENTIFIER_CHANGED                 |
| `CanonicalTypeSsfCanonicalEventTypeVerification`            | SSF_CANONICAL_EVENT_TYPE_VERIFICATION                       |
| `CanonicalTypeSsfCanonicalEventTypeStreamUpdated`           | SSF_CANONICAL_EVENT_TYPE_STREAM_UPDATED                     |