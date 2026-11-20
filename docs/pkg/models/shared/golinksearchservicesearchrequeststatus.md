# GoLinkSearchServiceSearchRequestStatus

Filter by status. If not set, returns all non-deleted GoLinks.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.GoLinkSearchServiceSearchRequestStatusGoLinkStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.GoLinkSearchServiceSearchRequestStatus("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `GoLinkSearchServiceSearchRequestStatusGoLinkStatusUnspecified` | GO_LINK_STATUS_UNSPECIFIED                                      |
| `GoLinkSearchServiceSearchRequestStatusGoLinkStatusEnabled`     | GO_LINK_STATUS_ENABLED                                          |
| `GoLinkSearchServiceSearchRequestStatusGoLinkStatusDisabled`    | GO_LINK_STATUS_DISABLED                                         |