# GrantEnrollmentStatus

Filters by whether the user already holds the connector's tools.
 UNSPECIFIED = all, FULLY_GRANTED = held, NOT_GRANTED = available.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

value := operations.GrantEnrollmentStatusGrantEnrollmentStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := operations.GrantEnrollmentStatus("custom_value")
```


## Values

| Name                                                         | Value                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| `GrantEnrollmentStatusGrantEnrollmentStatusUnspecified`      | GRANT_ENROLLMENT_STATUS_UNSPECIFIED                          |
| `GrantEnrollmentStatusGrantEnrollmentStatusFullyGranted`     | GRANT_ENROLLMENT_STATUS_FULLY_GRANTED                        |
| `GrantEnrollmentStatusGrantEnrollmentStatusNotGranted`       | GRANT_ENROLLMENT_STATUS_NOT_GRANTED                          |
| `GrantEnrollmentStatusGrantEnrollmentStatusPartiallyGranted` | GRANT_ENROLLMENT_STATUS_PARTIALLY_GRANTED                    |