# Severities

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SeveritiesFindingSeverityUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Severities("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `SeveritiesFindingSeverityUnspecified` | FINDING_SEVERITY_UNSPECIFIED           |
| `SeveritiesFindingSeverityInfo`        | FINDING_SEVERITY_INFO                  |
| `SeveritiesFindingSeverityLow`         | FINDING_SEVERITY_LOW                   |
| `SeveritiesFindingSeverityMedium`      | FINDING_SEVERITY_MEDIUM                |
| `SeveritiesFindingSeverityHigh`        | FINDING_SEVERITY_HIGH                  |
| `SeveritiesFindingSeverityCritical`    | FINDING_SEVERITY_CRITICAL              |