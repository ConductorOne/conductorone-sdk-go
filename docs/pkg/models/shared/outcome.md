# Outcome

The outcome field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.OutcomeGoLinkResolveOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.Outcome("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `OutcomeGoLinkResolveOutcomeUnspecified`         | GO_LINK_RESOLVE_OUTCOME_UNSPECIFIED              |
| `OutcomeGoLinkResolveOutcomeRedirect`            | GO_LINK_RESOLVE_OUTCOME_REDIRECT                 |
| `OutcomeGoLinkResolveOutcomePrerequisiteMissing` | GO_LINK_RESOLVE_OUTCOME_PREREQUISITE_MISSING     |
| `OutcomeGoLinkResolveOutcomeChooser`             | GO_LINK_RESOLVE_OUTCOME_CHOOSER                  |
| `OutcomeGoLinkResolveOutcomeAction`              | GO_LINK_RESOLVE_OUTCOME_ACTION                   |
| `OutcomeGoLinkResolveOutcomeNotFound`            | GO_LINK_RESOLVE_OUTCOME_NOT_FOUND                |
| `OutcomeGoLinkResolveOutcomeDisabled`            | GO_LINK_RESOLVE_OUTCOME_DISABLED                 |