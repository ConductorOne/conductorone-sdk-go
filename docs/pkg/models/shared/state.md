# State

The state field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.StateSurfaceLifecycleStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.State("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `StateSurfaceLifecycleStateUnspecified` | SURFACE_LIFECYCLE_STATE_UNSPECIFIED     |
| `StateSurfaceLifecycleStateActive`      | SURFACE_LIFECYCLE_STATE_ACTIVE          |
| `StateSurfaceLifecycleStateComplete`    | SURFACE_LIFECYCLE_STATE_COMPLETE        |
| `StateSurfaceLifecycleStateDeleted`     | SURFACE_LIFECYCLE_STATE_DELETED         |