# RateLimitServiceGetUsageReportRequestSortBy

The order of the entries. Defaults to SORT_BY_REQUESTS.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.RateLimitServiceGetUsageReportRequestSortBySortByUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.RateLimitServiceGetUsageReportRequestSortBy("custom_value")
```


## Values

| Name                                                           | Value                                                          |
| -------------------------------------------------------------- | -------------------------------------------------------------- |
| `RateLimitServiceGetUsageReportRequestSortBySortByUnspecified` | SORT_BY_UNSPECIFIED                                            |
| `RateLimitServiceGetUsageReportRequestSortBySortByRequests`    | SORT_BY_REQUESTS                                               |
| `RateLimitServiceGetUsageReportRequestSortBySortByThrottled`   | SORT_BY_THROTTLED                                              |