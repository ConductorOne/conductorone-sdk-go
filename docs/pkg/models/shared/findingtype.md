# FindingType

The findingType field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FindingTypeFindingTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.FindingType("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `FindingTypeFindingTypeUnspecified`                       | FINDING_TYPE_UNSPECIFIED                                  |
| `FindingTypeFindingTypeSimilarUsernameMatch`              | FINDING_TYPE_SIMILAR_USERNAME_MATCH                       |
| `FindingTypeFindingTypeServiceAccountMisclassification`   | FINDING_TYPE_SERVICE_ACCOUNT_MISCLASSIFICATION            |
| `FindingTypeFindingTypeNhiUnowned`                        | FINDING_TYPE_NHI_UNOWNED                                  |
| `FindingTypeFindingTypeServiceAccountUnowned`             | FINDING_TYPE_SERVICE_ACCOUNT_UNOWNED                      |
| `FindingTypeFindingTypeDecoyCredentialUsed`               | FINDING_TYPE_DECOY_CREDENTIAL_USED                        |
| `FindingTypeFindingTypeCustom`                            | FINDING_TYPE_CUSTOM                                       |
| `FindingTypeFindingTypeConnectorAnomalyDetectionDisabled` | FINDING_TYPE_CONNECTOR_ANOMALY_DETECTION_DISABLED         |
| `FindingTypeFindingTypeDeactivatedOwner`                  | FINDING_TYPE_DEACTIVATED_OWNER                            |
| `FindingTypeFindingTypeUnusedSecret`                      | FINDING_TYPE_UNUSED_SECRET                                |
| `FindingTypeFindingTypeCredentialPubliclyExposed`         | FINDING_TYPE_CREDENTIAL_PUBLICLY_EXPOSED                  |
| `FindingTypeFindingTypeDecoyPubliclyExposed`              | FINDING_TYPE_DECOY_PUBLICLY_EXPOSED                       |
| `FindingTypeFindingTypeCredentialExpiring`                | FINDING_TYPE_CREDENTIAL_EXPIRING                          |
| `FindingTypeFindingTypeConnectorSyncFailing`              | FINDING_TYPE_CONNECTOR_SYNC_FAILING                       |
| `FindingTypeFindingTypeShadowMcp`                         | FINDING_TYPE_SHADOW_MCP                                   |