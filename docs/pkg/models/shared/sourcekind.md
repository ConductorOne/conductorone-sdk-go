# SourceKind

Who authored the finding (detector, user, external).

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SourceKindFindingSourceKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SourceKind("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `SourceKindFindingSourceKindUnspecified` | FINDING_SOURCE_KIND_UNSPECIFIED          |
| `SourceKindFindingSourceKindDetector`    | FINDING_SOURCE_KIND_DETECTOR             |
| `SourceKindFindingSourceKindExternal`    | FINDING_SOURCE_KIND_EXTERNAL             |