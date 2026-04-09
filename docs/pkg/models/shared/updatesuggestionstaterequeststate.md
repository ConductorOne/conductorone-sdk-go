# UpdateSuggestionStateRequestState

The state field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.UpdateSuggestionStateRequestStateSuggestionStateUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.UpdateSuggestionStateRequestState("custom_value")
```


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `UpdateSuggestionStateRequestStateSuggestionStateUnspecified` | SUGGESTION_STATE_UNSPECIFIED                                  |
| `UpdateSuggestionStateRequestStateSuggestionStateNew`         | SUGGESTION_STATE_NEW                                          |
| `UpdateSuggestionStateRequestStateSuggestionStateDismissed`   | SUGGESTION_STATE_DISMISSED                                    |
| `UpdateSuggestionStateRequestStateSuggestionStateAccepted`    | SUGGESTION_STATE_ACCEPTED                                     |