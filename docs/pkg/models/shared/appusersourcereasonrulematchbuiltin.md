# AppUserSourceReasonRuleMatchBuiltin

The built-in rule that matched.
This field is part of the `rule` oneof.
See the documentation for `c1.api.app.v1.AppUserSourceReasonRuleMatch` for more details.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AppUserSourceReasonRuleMatchBuiltinAppUserMergeBuiltinRuleUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AppUserSourceReasonRuleMatchBuiltin("custom_value")
```


## Values

| Name                                                                     | Value                                                                    |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `AppUserSourceReasonRuleMatchBuiltinAppUserMergeBuiltinRuleUnspecified`  | APP_USER_MERGE_BUILTIN_RULE_UNSPECIFIED                                  |
| `AppUserSourceReasonRuleMatchBuiltinAppUserMergeBuiltinRulePrimaryEmail` | APP_USER_MERGE_BUILTIN_RULE_PRIMARY_EMAIL                                |
| `AppUserSourceReasonRuleMatchBuiltinAppUserMergeBuiltinRuleAnyEmail`     | APP_USER_MERGE_BUILTIN_RULE_ANY_EMAIL                                    |
| `AppUserSourceReasonRuleMatchBuiltinAppUserMergeBuiltinRuleUsername`     | APP_USER_MERGE_BUILTIN_RULE_USERNAME                                     |