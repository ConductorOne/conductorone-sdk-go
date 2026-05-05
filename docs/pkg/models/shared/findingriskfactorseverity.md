# FindingRiskFactorSeverity

The severity field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FindingRiskFactorSeverityFindingSeverityUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.FindingRiskFactorSeverity("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `FindingRiskFactorSeverityFindingSeverityUnspecified` | FINDING_SEVERITY_UNSPECIFIED                          |
| `FindingRiskFactorSeverityFindingSeverityInfo`        | FINDING_SEVERITY_INFO                                 |
| `FindingRiskFactorSeverityFindingSeverityLow`         | FINDING_SEVERITY_LOW                                  |
| `FindingRiskFactorSeverityFindingSeverityMedium`      | FINDING_SEVERITY_MEDIUM                               |
| `FindingRiskFactorSeverityFindingSeverityHigh`        | FINDING_SEVERITY_HIGH                                 |
| `FindingRiskFactorSeverityFindingSeverityCritical`    | FINDING_SEVERITY_CRITICAL                             |