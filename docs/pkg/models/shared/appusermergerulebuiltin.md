# AppUserMergeRuleBuiltin

A built-in rule. Each built-in rule may appear at most once.
This field is part of the `kind` oneof.
See the documentation for `c1.api.app.v1.AppUserMergeRule` for more details.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.AppUserMergeRuleBuiltinAppUserMergeBuiltinRuleUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AppUserMergeRuleBuiltin("custom_value")
```


## Values

| Name                                                         | Value                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| `AppUserMergeRuleBuiltinAppUserMergeBuiltinRuleUnspecified`  | APP_USER_MERGE_BUILTIN_RULE_UNSPECIFIED                      |
| `AppUserMergeRuleBuiltinAppUserMergeBuiltinRulePrimaryEmail` | APP_USER_MERGE_BUILTIN_RULE_PRIMARY_EMAIL                    |
| `AppUserMergeRuleBuiltinAppUserMergeBuiltinRuleAnyEmail`     | APP_USER_MERGE_BUILTIN_RULE_ANY_EMAIL                        |
| `AppUserMergeRuleBuiltinAppUserMergeBuiltinRuleUsername`     | APP_USER_MERGE_BUILTIN_RULE_USERNAME                         |