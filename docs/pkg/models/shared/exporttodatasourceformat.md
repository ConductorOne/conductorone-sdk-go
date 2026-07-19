# ExportToDatasourceFormat

The format field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ExportToDatasourceFormatExportFormatUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ExportToDatasourceFormat("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `ExportToDatasourceFormatExportFormatUnspecified`  | EXPORT_FORMAT_UNSPECIFIED                          |
| `ExportToDatasourceFormatExportFormatOcsfJSONZstd` | EXPORT_FORMAT_OCSF_JSON_ZSTD                       |
| `ExportToDatasourceFormatExportFormatOcsfJSONGzip` | EXPORT_FORMAT_OCSF_JSON_GZIP                       |