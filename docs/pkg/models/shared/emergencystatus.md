# EmergencyStatus

Search tasks that are or are not emergency access.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.EmergencyStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.EmergencyStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `EmergencyStatusUnspecified`  | UNSPECIFIED                   |
| `EmergencyStatusAll`          | ALL                           |
| `EmergencyStatusNonEmergency` | NON_EMERGENCY                 |
| `EmergencyStatusEmergency`    | EMERGENCY                     |