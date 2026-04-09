# SuggestionState

The suggestionState field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.SuggestionStateSuggestionStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.SuggestionState("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `SuggestionStateSuggestionStateUnspecified` | SUGGESTION_STATE_UNSPECIFIED                |
| `SuggestionStateSuggestionStateNew`         | SUGGESTION_STATE_NEW                        |
| `SuggestionStateSuggestionStateDismissed`   | SUGGESTION_STATE_DISMISSED                  |
| `SuggestionStateSuggestionStateAccepted`    | SUGGESTION_STATE_ACCEPTED                   |