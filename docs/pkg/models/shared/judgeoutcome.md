# JudgeOutcome

Whether the routing algorithm could use this judge attempt's reply. Set
 only on judge provider-attempt rows Time Bandit recorded it for.

## Example Usage

```go
import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
)

value := shared.JudgeOutcomeTbJudgeOutcomeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.JudgeOutcome("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `JudgeOutcomeTbJudgeOutcomeUnspecified` | TB_JUDGE_OUTCOME_UNSPECIFIED            |
| `JudgeOutcomeTbJudgeOutcomeUsable`      | TB_JUDGE_OUTCOME_USABLE                 |
| `JudgeOutcomeTbJudgeOutcomeUnusable`    | TB_JUDGE_OUTCOME_UNUSABLE               |
| `JudgeOutcomeTbJudgeOutcomeCallFailed`  | TB_JUDGE_OUTCOME_CALL_FAILED            |
| `JudgeOutcomeTbJudgeOutcomeUnreported`  | TB_JUDGE_OUTCOME_UNREPORTED             |