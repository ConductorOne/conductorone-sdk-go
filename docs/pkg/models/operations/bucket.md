# Bucket

Defaults to TB_TIME_BUCKET_HOUR when unset.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/operations"
)

value := operations.BucketTbTimeBucketUnspecified

// Open enum: custom values can be created with a direct type cast
custom := operations.Bucket("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `BucketTbTimeBucketUnspecified` | TB_TIME_BUCKET_UNSPECIFIED      |
| `BucketTbTimeBucketHour`        | TB_TIME_BUCKET_HOUR             |
| `BucketTbTimeBucketDay`         | TB_TIME_BUCKET_DAY              |