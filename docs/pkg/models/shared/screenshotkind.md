# ScreenshotKind

The kind field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.ScreenshotKindKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.ScreenshotKind("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `ScreenshotKindKindUnspecified` | KIND_UNSPECIFIED                |
| `ScreenshotKindKindScreen`      | KIND_SCREEN                     |
| `ScreenshotKindKindRegion`      | KIND_REGION                     |
| `ScreenshotKindKindUpload`      | KIND_UPLOAD                     |