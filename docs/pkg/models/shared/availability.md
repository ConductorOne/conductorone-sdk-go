# Availability

Data availability for the requested reporting window.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AvailabilityReportingAvailabilityUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Availability("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `AvailabilityReportingAvailabilityUnspecified`      | REPORTING_AVAILABILITY_UNSPECIFIED                  |
| `AvailabilityReportingAvailabilityComplete`         | REPORTING_AVAILABILITY_COMPLETE                     |
| `AvailabilityReportingAvailabilityPartialRetention` | REPORTING_AVAILABILITY_PARTIAL_RETENTION            |
| `AvailabilityReportingAvailabilityOutsideRetention` | REPORTING_AVAILABILITY_OUTSIDE_RETENTION            |
| `AvailabilityReportingAvailabilityNoActivity`       | REPORTING_AVAILABILITY_NO_ACTIVITY                  |