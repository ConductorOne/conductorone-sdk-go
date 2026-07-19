# PersonalDeviceServiceSearchRequestStatusFilter

Which device statuses to return. Defaults to active devices only.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PersonalDeviceServiceSearchRequestStatusFilterPersonalDeviceStatusFilterUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PersonalDeviceServiceSearchRequestStatusFilter("custom_value")
```


## Values

| Name                                                                                  | Value                                                                                 |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `PersonalDeviceServiceSearchRequestStatusFilterPersonalDeviceStatusFilterUnspecified` | PERSONAL_DEVICE_STATUS_FILTER_UNSPECIFIED                                             |
| `PersonalDeviceServiceSearchRequestStatusFilterPersonalDeviceStatusFilterActive`      | PERSONAL_DEVICE_STATUS_FILTER_ACTIVE                                                  |
| `PersonalDeviceServiceSearchRequestStatusFilterPersonalDeviceStatusFilterRevoked`     | PERSONAL_DEVICE_STATUS_FILTER_REVOKED                                                 |
| `PersonalDeviceServiceSearchRequestStatusFilterPersonalDeviceStatusFilterAll`         | PERSONAL_DEVICE_STATUS_FILTER_ALL                                                     |