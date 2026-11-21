# ResponseInspectionProfile

Default inspection profile for hosted egress response bodies. Unspecified
 means strict. Requests always use strict inspection. Web balanced requires
 a runtime that supports it; older runtimes retain strict inspection.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ResponseInspectionProfileEdgeResponseInspectionProfileUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ResponseInspectionProfile("custom_value")
```


## Values

| Name                                                                | Value                                                               |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ResponseInspectionProfileEdgeResponseInspectionProfileUnspecified` | EDGE_RESPONSE_INSPECTION_PROFILE_UNSPECIFIED                        |
| `ResponseInspectionProfileEdgeResponseInspectionProfileStrict`      | EDGE_RESPONSE_INSPECTION_PROFILE_STRICT                             |
| `ResponseInspectionProfileEdgeResponseInspectionProfileWebBalanced` | EDGE_RESPONSE_INSPECTION_PROFILE_WEB_BALANCED                       |