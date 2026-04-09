# CreateAppRequestIdentityMatching

Define the app user identity matching strategy for this app.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.CreateAppRequestIdentityMatchingAppUserIdentityMatchingUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.CreateAppRequestIdentityMatching("custom_value")
```


## Values

| Name                                                                 | Value                                                                |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `CreateAppRequestIdentityMatchingAppUserIdentityMatchingUnspecified` | APP_USER_IDENTITY_MATCHING_UNSPECIFIED                               |
| `CreateAppRequestIdentityMatchingAppUserIdentityMatchingStrict`      | APP_USER_IDENTITY_MATCHING_STRICT                                    |
| `CreateAppRequestIdentityMatchingAppUserIdentityMatchingDisplayName` | APP_USER_IDENTITY_MATCHING_DISPLAY_NAME                              |
| `CreateAppRequestIdentityMatchingAppUserIdentityMatchingCustom`      | APP_USER_IDENTITY_MATCHING_CUSTOM                                    |