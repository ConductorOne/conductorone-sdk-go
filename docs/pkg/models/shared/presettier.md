# PresetTier

Non-UNSPECIFIED only on the three system presets (Shadow/Balanced/Strict);
 identifies and strictness-orders them for the default-classifier picker.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PresetTierPresetTierUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PresetTier("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `PresetTierPresetTierUnspecified` | PRESET_TIER_UNSPECIFIED           |
| `PresetTierPresetTierShadow`      | PRESET_TIER_SHADOW                |
| `PresetTierPresetTierBalanced`    | PRESET_TIER_BALANCED              |
| `PresetTierPresetTierStrict`      | PRESET_TIER_STRICT                |