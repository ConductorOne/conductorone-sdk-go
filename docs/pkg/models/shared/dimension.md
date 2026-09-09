# Dimension

Dimension used to group settled calls.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DimensionAttributionDimensionUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Dimension("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `DimensionAttributionDimensionUnspecified`       | ATTRIBUTION_DIMENSION_UNSPECIFIED                |
| `DimensionAttributionDimensionUser`              | ATTRIBUTION_DIMENSION_USER                       |
| `DimensionAttributionDimensionDepartment`        | ATTRIBUTION_DIMENSION_DEPARTMENT                 |
| `DimensionAttributionDimensionAgent`             | ATTRIBUTION_DIMENSION_AGENT                      |
| `DimensionAttributionDimensionSession`           | ATTRIBUTION_DIMENSION_SESSION                    |
| `DimensionAttributionDimensionModel`             | ATTRIBUTION_DIMENSION_MODEL                      |
| `DimensionAttributionDimensionCallSource`        | ATTRIBUTION_DIMENSION_CALL_SOURCE                |
| `DimensionAttributionDimensionSubjectFundSource` | ATTRIBUTION_DIMENSION_SUBJECT_FUND_SOURCE        |
| `DimensionAttributionDimensionApp`               | ATTRIBUTION_DIMENSION_APP                        |