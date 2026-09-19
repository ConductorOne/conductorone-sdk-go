# QueryParamBucket

Defaults to TB_TIME_BUCKET_HOUR when unset.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

value := operations.QueryParamBucketTbTimeBucketUnspecified

// Open enum: custom values can be created with a direct type cast
custom := operations.QueryParamBucket("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `QueryParamBucketTbTimeBucketUnspecified` | TB_TIME_BUCKET_UNSPECIFIED                |
| `QueryParamBucketTbTimeBucketHour`        | TB_TIME_BUCKET_HOUR                       |
| `QueryParamBucketTbTimeBucketDay`         | TB_TIME_BUCKET_DAY                        |