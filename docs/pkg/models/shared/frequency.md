# Frequency

How often digest notifications are sent.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FrequencyDigestFrequencyUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Frequency("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `FrequencyDigestFrequencyUnspecified` | DIGEST_FREQUENCY_UNSPECIFIED          |
| `FrequencyDigestFrequencyDaily`       | DIGEST_FREQUENCY_DAILY                |
| `FrequencyDigestFrequencyWeekly`      | DIGEST_FREQUENCY_WEEKLY               |