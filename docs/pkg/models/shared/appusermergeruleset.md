# AppUserMergeRuleSet

The ordered rules that decide which existing account an incoming account record
 belongs to. Rules are tried in order and the first enabled rule that matches wins.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `AppID`                                                                          | `*string`                                                                        | :heavy_minus_sign:                                                               | The ID of the application these rules belong to.                                 |
| `CreatedAt`                                                                      | [*time.Time](https://pkg.go.dev/time#Time)                                       | :heavy_minus_sign:                                                               | N/A                                                                              |
| `RuleSetVersion`                                                                 | `*int64`                                                                         | :heavy_minus_sign:                                                               | The version of the rule set. It increases by one every time the rules are saved. |
| `Rules`                                                                          | [][shared.AppUserMergeRule](../../../pkg/models/shared/appusermergerule.md)      | :heavy_minus_sign:                                                               | The rules, in the order they are tried.                                          |
| `UpdatedAt`                                                                      | [*time.Time](https://pkg.go.dev/time#Time)                                       | :heavy_minus_sign:                                                               | N/A                                                                              |