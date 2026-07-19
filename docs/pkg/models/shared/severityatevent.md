# SeverityAtEvent

The severityAtEvent field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SeverityAtEventFindingSeverityUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SeverityAtEvent("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `SeverityAtEventFindingSeverityUnspecified` | FINDING_SEVERITY_UNSPECIFIED                |
| `SeverityAtEventFindingSeverityInfo`        | FINDING_SEVERITY_INFO                       |
| `SeverityAtEventFindingSeverityLow`         | FINDING_SEVERITY_LOW                        |
| `SeverityAtEventFindingSeverityMedium`      | FINDING_SEVERITY_MEDIUM                     |
| `SeverityAtEventFindingSeverityHigh`        | FINDING_SEVERITY_HIGH                       |
| `SeverityAtEventFindingSeverityCritical`    | FINDING_SEVERITY_CRITICAL                   |