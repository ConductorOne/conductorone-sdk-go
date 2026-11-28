# ConflictMonitorCreateRequestSeverity

Configured severity for new findings. Defaults to medium when unset.
 The value is stored, but monitor responses and new findings use medium until
 configured severity is enabled for your tenant.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ConflictMonitorCreateRequestSeverityFindingSeverityUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ConflictMonitorCreateRequestSeverity("custom_value")
```


## Values

| Name                                                             | Value                                                            |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ConflictMonitorCreateRequestSeverityFindingSeverityUnspecified` | FINDING_SEVERITY_UNSPECIFIED                                     |
| `ConflictMonitorCreateRequestSeverityFindingSeverityInfo`        | FINDING_SEVERITY_INFO                                            |
| `ConflictMonitorCreateRequestSeverityFindingSeverityLow`         | FINDING_SEVERITY_LOW                                             |
| `ConflictMonitorCreateRequestSeverityFindingSeverityMedium`      | FINDING_SEVERITY_MEDIUM                                          |
| `ConflictMonitorCreateRequestSeverityFindingSeverityHigh`        | FINDING_SEVERITY_HIGH                                            |
| `ConflictMonitorCreateRequestSeverityFindingSeverityCritical`    | FINDING_SEVERITY_CRITICAL                                        |