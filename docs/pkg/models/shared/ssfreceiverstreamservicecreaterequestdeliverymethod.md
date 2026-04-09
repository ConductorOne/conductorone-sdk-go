# SSFReceiverStreamServiceCreateRequestDeliveryMethod

The deliveryMethod field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SSFReceiverStreamServiceCreateRequestDeliveryMethodSsfDeliveryMethodUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SSFReceiverStreamServiceCreateRequestDeliveryMethod("custom_value")
```


## Values

| Name                                                                              | Value                                                                             |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `SSFReceiverStreamServiceCreateRequestDeliveryMethodSsfDeliveryMethodUnspecified` | SSF_DELIVERY_METHOD_UNSPECIFIED                                                   |
| `SSFReceiverStreamServiceCreateRequestDeliveryMethodSsfDeliveryMethodPush`        | SSF_DELIVERY_METHOD_PUSH                                                          |
| `SSFReceiverStreamServiceCreateRequestDeliveryMethodSsfDeliveryMethodPoll`        | SSF_DELIVERY_METHOD_POLL                                                          |