# ConflictMonitorUpdateRequestSeverity

Configured severity for new findings. Unset preserves the stored value.
 Explicit values replace the stored configuration, even when responses use medium.
 Monitor responses and new findings use medium until configured severity is
 enabled for your tenant.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ConflictMonitorUpdateRequestSeverityFindingSeverityUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ConflictMonitorUpdateRequestSeverity("custom_value")
```


## Values

| Name                                                             | Value                                                            |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ConflictMonitorUpdateRequestSeverityFindingSeverityUnspecified` | FINDING_SEVERITY_UNSPECIFIED                                     |
| `ConflictMonitorUpdateRequestSeverityFindingSeverityInfo`        | FINDING_SEVERITY_INFO                                            |
| `ConflictMonitorUpdateRequestSeverityFindingSeverityLow`         | FINDING_SEVERITY_LOW                                             |
| `ConflictMonitorUpdateRequestSeverityFindingSeverityMedium`      | FINDING_SEVERITY_MEDIUM                                          |
| `ConflictMonitorUpdateRequestSeverityFindingSeverityHigh`        | FINDING_SEVERITY_HIGH                                            |
| `ConflictMonitorUpdateRequestSeverityFindingSeverityCritical`    | FINDING_SEVERITY_CRITICAL                                        |