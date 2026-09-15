# Mode

Configures the intended evaluation behavior: DISABLED and UNSPECIFIED
 deny requests, OBSERVE records policy decisions without enforcing them,
 and ENFORCE applies policy decisions. Evaluation endpoints are not served
 by this control-plane API.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ModeAuthzenModeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Mode("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `ModeAuthzenModeUnspecified` | AUTHZEN_MODE_UNSPECIFIED     |
| `ModeAuthzenModeDisabled`    | AUTHZEN_MODE_DISABLED        |
| `ModeAuthzenModeObserve`     | AUTHZEN_MODE_OBSERVE         |
| `ModeAuthzenModeEnforce`     | AUTHZEN_MODE_ENFORCE         |