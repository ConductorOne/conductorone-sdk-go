# FindingSettingsEntryFindingType

The finding type to configure. Must be a detector-backed type.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FindingSettingsEntryFindingTypeFindingTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.FindingSettingsEntryFindingType("custom_value")
```


## Values

| Name                                                                          | Value                                                                         |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `FindingSettingsEntryFindingTypeFindingTypeUnspecified`                       | FINDING_TYPE_UNSPECIFIED                                                      |
| `FindingSettingsEntryFindingTypeFindingTypeSimilarUsernameMatch`              | FINDING_TYPE_SIMILAR_USERNAME_MATCH                                           |
| `FindingSettingsEntryFindingTypeFindingTypeServiceAccountMisclassification`   | FINDING_TYPE_SERVICE_ACCOUNT_MISCLASSIFICATION                                |
| `FindingSettingsEntryFindingTypeFindingTypeNhiUnowned`                        | FINDING_TYPE_NHI_UNOWNED                                                      |
| `FindingSettingsEntryFindingTypeFindingTypeServiceAccountUnowned`             | FINDING_TYPE_SERVICE_ACCOUNT_UNOWNED                                          |
| `FindingSettingsEntryFindingTypeFindingTypeDecoyCredentialUsed`               | FINDING_TYPE_DECOY_CREDENTIAL_USED                                            |
| `FindingSettingsEntryFindingTypeFindingTypeCustom`                            | FINDING_TYPE_CUSTOM                                                           |
| `FindingSettingsEntryFindingTypeFindingTypeConnectorAnomalyDetectionDisabled` | FINDING_TYPE_CONNECTOR_ANOMALY_DETECTION_DISABLED                             |
| `FindingSettingsEntryFindingTypeFindingTypeDeactivatedOwner`                  | FINDING_TYPE_DEACTIVATED_OWNER                                                |
| `FindingSettingsEntryFindingTypeFindingTypeUnusedSecret`                      | FINDING_TYPE_UNUSED_SECRET                                                    |
| `FindingSettingsEntryFindingTypeFindingTypeCredentialPubliclyExposed`         | FINDING_TYPE_CREDENTIAL_PUBLICLY_EXPOSED                                      |
| `FindingSettingsEntryFindingTypeFindingTypeDecoyPubliclyExposed`              | FINDING_TYPE_DECOY_PUBLICLY_EXPOSED                                           |
| `FindingSettingsEntryFindingTypeFindingTypeCredentialExpiring`                | FINDING_TYPE_CREDENTIAL_EXPIRING                                              |
| `FindingSettingsEntryFindingTypeFindingTypeConnectorSyncFailing`              | FINDING_TYPE_CONNECTOR_SYNC_FAILING                                           |
| `FindingSettingsEntryFindingTypeFindingTypeShadowMcp`                         | FINDING_TYPE_SHADOW_MCP                                                       |