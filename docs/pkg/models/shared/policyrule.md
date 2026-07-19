# PolicyRule

PolicyRule is one rung of the ordered sign-in cascade. Rules are evaluated
 top to bottom; the first enforced rule whose condition matches supplies the
 outcome.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Description`                                                          | `*string`                                                              | :heavy_minus_sign:                                                     | A human-readable description shown in the admin UI.                    |
| `ID`                                                                   | `*string`                                                              | :heavy_minus_sign:                                                     | A stable identifier for the rule, surfaced in audit.                   |
| `MatchCel`                                                             | `*string`                                                              | :heavy_minus_sign:                                                     | A boolean condition expression evaluated against the sign-in context.  |
| `Mode`                                                                 | [*shared.PolicyRuleMode](../../../pkg/models/shared/policyrulemode.md) | :heavy_minus_sign:                                                     | Whether the rule is live, evaluated-only, or skipped.                  |
| `Outcome`                                                              | [*shared.PolicyOutcome](../../../pkg/models/shared/policyoutcome.md)   | :heavy_minus_sign:                                                     | N/A                                                                    |