# FindingSeverity

The severity field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FindingSeverityFindingSeverityUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.FindingSeverity("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `FindingSeverityFindingSeverityUnspecified` | FINDING_SEVERITY_UNSPECIFIED                |
| `FindingSeverityFindingSeverityInfo`        | FINDING_SEVERITY_INFO                       |
| `FindingSeverityFindingSeverityLow`         | FINDING_SEVERITY_LOW                        |
| `FindingSeverityFindingSeverityMedium`      | FINDING_SEVERITY_MEDIUM                     |
| `FindingSeverityFindingSeverityHigh`        | FINDING_SEVERITY_HIGH                       |
| `FindingSeverityFindingSeverityCritical`    | FINDING_SEVERITY_CRITICAL                   |