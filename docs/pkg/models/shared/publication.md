# Publication

Which C1-published route of a followed tier this is. Set by the server; send back the value you read.
 UNSPECIFIED for a tenant-authored route.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PublicationEdgeInferenceRoutePublicationUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Publication("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `PublicationEdgeInferenceRoutePublicationUnspecified` | EDGE_INFERENCE_ROUTE_PUBLICATION_UNSPECIFIED          |
| `PublicationEdgeInferenceRoutePublicationTier`        | EDGE_INFERENCE_ROUTE_PUBLICATION_TIER                 |
| `PublicationEdgeInferenceRoutePublicationModelAlias`  | EDGE_INFERENCE_ROUTE_PUBLICATION_MODEL_ALIAS          |