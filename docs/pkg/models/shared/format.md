# Format

Wire format the provider expects. Defaults to
 FORMAT_JSON_OBJECT.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.FormatFormatJSONObject

// Open enum: custom values can be created with a direct type cast
custom := shared.Format("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `FormatFormatJSONObject`          | FORMAT_JSON_OBJECT                |
| `FormatFormatColonSeparated`      | FORMAT_COLON_SEPARATED            |
| `FormatFormatUnderscoreSeparated` | FORMAT_UNDERSCORE_SEPARATED       |