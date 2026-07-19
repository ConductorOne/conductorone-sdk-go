# FindingTypes

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FindingTypesFindingTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.FindingTypes("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `FindingTypesFindingTypeUnspecified`                       | FINDING_TYPE_UNSPECIFIED                                   |
| `FindingTypesFindingTypeSimilarUsernameMatch`              | FINDING_TYPE_SIMILAR_USERNAME_MATCH                        |
| `FindingTypesFindingTypeServiceAccountMisclassification`   | FINDING_TYPE_SERVICE_ACCOUNT_MISCLASSIFICATION             |
| `FindingTypesFindingTypeNhiUnowned`                        | FINDING_TYPE_NHI_UNOWNED                                   |
| `FindingTypesFindingTypeServiceAccountUnowned`             | FINDING_TYPE_SERVICE_ACCOUNT_UNOWNED                       |
| `FindingTypesFindingTypeDecoyCredentialUsed`               | FINDING_TYPE_DECOY_CREDENTIAL_USED                         |
| `FindingTypesFindingTypeCustom`                            | FINDING_TYPE_CUSTOM                                        |
| `FindingTypesFindingTypeConnectorAnomalyDetectionDisabled` | FINDING_TYPE_CONNECTOR_ANOMALY_DETECTION_DISABLED          |