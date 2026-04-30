# WellKnownRegex

WellKnownRegex specifies a common well known pattern defined as a regex.
This field is part of the `well_known` oneof.
See the documentation for `validate.StringRules` for more details.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.WellKnownRegexUnknown

// Open enum: custom values can be created with a direct type cast
custom := shared.WellKnownRegex("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `WellKnownRegexUnknown`         | UNKNOWN                         |
| `WellKnownRegexHTTPHeaderName`  | HTTP_HEADER_NAME                |
| `WellKnownRegexHTTPHeaderValue` | HTTP_HEADER_VALUE               |