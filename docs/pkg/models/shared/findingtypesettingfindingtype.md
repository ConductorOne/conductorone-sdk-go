# FindingTypeSettingFindingType

The findingType field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FindingTypeSettingFindingTypeFindingTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.FindingTypeSettingFindingType("custom_value")
```


## Values

| Name                                                                        | Value                                                                       |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `FindingTypeSettingFindingTypeFindingTypeUnspecified`                       | FINDING_TYPE_UNSPECIFIED                                                    |
| `FindingTypeSettingFindingTypeFindingTypeSimilarUsernameMatch`              | FINDING_TYPE_SIMILAR_USERNAME_MATCH                                         |
| `FindingTypeSettingFindingTypeFindingTypeServiceAccountMisclassification`   | FINDING_TYPE_SERVICE_ACCOUNT_MISCLASSIFICATION                              |
| `FindingTypeSettingFindingTypeFindingTypeNhiUnowned`                        | FINDING_TYPE_NHI_UNOWNED                                                    |
| `FindingTypeSettingFindingTypeFindingTypeServiceAccountUnowned`             | FINDING_TYPE_SERVICE_ACCOUNT_UNOWNED                                        |
| `FindingTypeSettingFindingTypeFindingTypeDecoyCredentialUsed`               | FINDING_TYPE_DECOY_CREDENTIAL_USED                                          |
| `FindingTypeSettingFindingTypeFindingTypeCustom`                            | FINDING_TYPE_CUSTOM                                                         |
| `FindingTypeSettingFindingTypeFindingTypeConnectorAnomalyDetectionDisabled` | FINDING_TYPE_CONNECTOR_ANOMALY_DETECTION_DISABLED                           |
| `FindingTypeSettingFindingTypeFindingTypeDeactivatedOwner`                  | FINDING_TYPE_DEACTIVATED_OWNER                                              |
| `FindingTypeSettingFindingTypeFindingTypeUnusedSecret`                      | FINDING_TYPE_UNUSED_SECRET                                                  |
| `FindingTypeSettingFindingTypeFindingTypeCredentialPubliclyExposed`         | FINDING_TYPE_CREDENTIAL_PUBLICLY_EXPOSED                                    |
| `FindingTypeSettingFindingTypeFindingTypeDecoyPubliclyExposed`              | FINDING_TYPE_DECOY_PUBLICLY_EXPOSED                                         |
| `FindingTypeSettingFindingTypeFindingTypeCredentialExpiring`                | FINDING_TYPE_CREDENTIAL_EXPIRING                                            |
| `FindingTypeSettingFindingTypeFindingTypeConnectorSyncFailing`              | FINDING_TYPE_CONNECTOR_SYNC_FAILING                                         |
| `FindingTypeSettingFindingTypeFindingTypeShadowMcp`                         | FINDING_TYPE_SHADOW_MCP                                                     |