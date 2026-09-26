# PatternKind

The patternKind field.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.PatternKindHookBuiltInPatternKindUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.PatternKind("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `PatternKindHookBuiltInPatternKindUnspecified`         | HOOK_BUILT_IN_PATTERN_KIND_UNSPECIFIED                 |
| `PatternKindHookBuiltInPatternKindPiiRedaction`        | HOOK_BUILT_IN_PATTERN_KIND_PII_REDACTION               |
| `PatternKindHookBuiltInPatternKindCreditCardBlocking`  | HOOK_BUILT_IN_PATTERN_KIND_CREDIT_CARD_BLOCKING        |
| `PatternKindHookBuiltInPatternKindQueryScopeLimit`     | HOOK_BUILT_IN_PATTERN_KIND_QUERY_SCOPE_LIMIT           |
| `PatternKindHookBuiltInPatternKindWriteAuthorization`  | HOOK_BUILT_IN_PATTERN_KIND_WRITE_AUTHORIZATION         |
| `PatternKindHookBuiltInPatternKindSensitiveFileGuard`  | HOOK_BUILT_IN_PATTERN_KIND_SENSITIVE_FILE_GUARD        |
| `PatternKindHookBuiltInPatternKindToolOutputSizeGuard` | HOOK_BUILT_IN_PATTERN_KIND_TOOL_OUTPUT_SIZE_GUARD      |
| `PatternKindHookBuiltInPatternKindSecretsMasking`      | HOOK_BUILT_IN_PATTERN_KIND_SECRETS_MASKING             |
| `PatternKindHookBuiltInPatternKindLinkFilter`          | HOOK_BUILT_IN_PATTERN_KIND_LINK_FILTER                 |
| `PatternKindHookBuiltInPatternKindEncodedContentGuard` | HOOK_BUILT_IN_PATTERN_KIND_ENCODED_CONTENT_GUARD       |
| `PatternKindHookBuiltInPatternKindPromptInjectionScan` | HOOK_BUILT_IN_PATTERN_KIND_PROMPT_INJECTION_SCAN       |