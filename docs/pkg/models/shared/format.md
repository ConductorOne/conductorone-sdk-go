# Format

The format field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FormatExportFormatUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Format("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `FormatExportFormatUnspecified`  | EXPORT_FORMAT_UNSPECIFIED        |
| `FormatExportFormatOcsfJSONZstd` | EXPORT_FORMAT_OCSF_JSON_ZSTD     |
| `FormatExportFormatOcsfJSONGzip` | EXPORT_FORMAT_OCSF_JSON_GZIP     |