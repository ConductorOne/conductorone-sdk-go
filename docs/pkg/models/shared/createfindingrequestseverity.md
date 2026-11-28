# CreateFindingRequestSeverity

Severity of the finding. Must be a defined, non-unspecified value.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CreateFindingRequestSeverityFindingSeverityUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.CreateFindingRequestSeverity("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `CreateFindingRequestSeverityFindingSeverityUnspecified` | FINDING_SEVERITY_UNSPECIFIED                             |
| `CreateFindingRequestSeverityFindingSeverityInfo`        | FINDING_SEVERITY_INFO                                    |
| `CreateFindingRequestSeverityFindingSeverityLow`         | FINDING_SEVERITY_LOW                                     |
| `CreateFindingRequestSeverityFindingSeverityMedium`      | FINDING_SEVERITY_MEDIUM                                  |
| `CreateFindingRequestSeverityFindingSeverityHigh`        | FINDING_SEVERITY_HIGH                                    |
| `CreateFindingRequestSeverityFindingSeverityCritical`    | FINDING_SEVERITY_CRITICAL                                |