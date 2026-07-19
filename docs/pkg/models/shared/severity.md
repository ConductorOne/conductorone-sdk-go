# Severity

Severity of the finding. Must be a defined, non-unspecified value.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SeverityFindingSeverityUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Severity("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `SeverityFindingSeverityUnspecified` | FINDING_SEVERITY_UNSPECIFIED         |
| `SeverityFindingSeverityInfo`        | FINDING_SEVERITY_INFO                |
| `SeverityFindingSeverityLow`         | FINDING_SEVERITY_LOW                 |
| `SeverityFindingSeverityMedium`      | FINDING_SEVERITY_MEDIUM              |
| `SeverityFindingSeverityHigh`        | FINDING_SEVERITY_HIGH                |
| `SeverityFindingSeverityCritical`    | FINDING_SEVERITY_CRITICAL            |