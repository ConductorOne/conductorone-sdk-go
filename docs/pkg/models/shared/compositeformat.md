# CompositeFormat

Wire format the provider expects. Defaults to
 FORMAT_JSON_OBJECT.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CompositeFormatFormatJSONObject

// Open enum: custom values can be created with a direct type cast
custom := shared.CompositeFormat("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `CompositeFormatFormatJSONObject`          | FORMAT_JSON_OBJECT                         |
| `CompositeFormatFormatColonSeparated`      | FORMAT_COLON_SEPARATED                     |
| `CompositeFormatFormatUnderscoreSeparated` | FORMAT_UNDERSCORE_SEPARATED                |