# DeliveryMethod

Delivery configuration.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.DeliveryMethodSsfDeliveryMethodUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.DeliveryMethod("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `DeliveryMethodSsfDeliveryMethodUnspecified` | SSF_DELIVERY_METHOD_UNSPECIFIED              |
| `DeliveryMethodSsfDeliveryMethodPush`        | SSF_DELIVERY_METHOD_PUSH                     |
| `DeliveryMethodSsfDeliveryMethodPoll`        | SSF_DELIVERY_METHOD_POLL                     |