# GoLinkServiceResolveResponseOutcome

The outcome field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.GoLinkServiceResolveResponseOutcomeGoLinkResolveOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.GoLinkServiceResolveResponseOutcome("custom_value")
```


## Values

| Name                                                                         | Value                                                                        |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `GoLinkServiceResolveResponseOutcomeGoLinkResolveOutcomeUnspecified`         | GO_LINK_RESOLVE_OUTCOME_UNSPECIFIED                                          |
| `GoLinkServiceResolveResponseOutcomeGoLinkResolveOutcomeRedirect`            | GO_LINK_RESOLVE_OUTCOME_REDIRECT                                             |
| `GoLinkServiceResolveResponseOutcomeGoLinkResolveOutcomePrerequisiteMissing` | GO_LINK_RESOLVE_OUTCOME_PREREQUISITE_MISSING                                 |
| `GoLinkServiceResolveResponseOutcomeGoLinkResolveOutcomeChooser`             | GO_LINK_RESOLVE_OUTCOME_CHOOSER                                              |
| `GoLinkServiceResolveResponseOutcomeGoLinkResolveOutcomeAction`              | GO_LINK_RESOLVE_OUTCOME_ACTION                                               |
| `GoLinkServiceResolveResponseOutcomeGoLinkResolveOutcomeNotFound`            | GO_LINK_RESOLVE_OUTCOME_NOT_FOUND                                            |
| `GoLinkServiceResolveResponseOutcomeGoLinkResolveOutcomeDisabled`            | GO_LINK_RESOLVE_OUTCOME_DISABLED                                             |