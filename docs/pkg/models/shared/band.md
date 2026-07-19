# Band

Advisory size classification for the user's access graph. See
 GraphSizeBand.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.BandGraphSizeBandUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Band("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `BandGraphSizeBandUnspecified` | GRAPH_SIZE_BAND_UNSPECIFIED    |
| `BandGraphSizeBandNormal`      | GRAPH_SIZE_BAND_NORMAL         |
| `BandGraphSizeBandSummarized`  | GRAPH_SIZE_BAND_SUMMARIZED     |
| `BandGraphSizeBandTooLarge`    | GRAPH_SIZE_BAND_TOO_LARGE      |