# PersonalDeviceStatus

The device's lifecycle status. Revoked devices are retained for audit and are
 returned by Search only when the status filter requests them.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PersonalDeviceStatusPersonalDeviceStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PersonalDeviceStatus("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `PersonalDeviceStatusPersonalDeviceStatusUnspecified` | PERSONAL_DEVICE_STATUS_UNSPECIFIED                    |
| `PersonalDeviceStatusPersonalDeviceStatusActive`      | PERSONAL_DEVICE_STATUS_ACTIVE                         |
| `PersonalDeviceStatusPersonalDeviceStatusRevoked`     | PERSONAL_DEVICE_STATUS_REVOKED                        |