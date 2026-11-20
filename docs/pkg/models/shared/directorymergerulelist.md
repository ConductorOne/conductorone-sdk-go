# DirectoryMergeRuleList

DirectoryMergeRuleList wraps a full rule list so a request can tell "unchanged" (absent) from a new list.


## Fields

| Field                                                                                                                               | Type                                                                                                                                | Required                                                                                                                            | Description                                                                                                                         |
| ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| `Rules`                                                                                                                             | [][shared.DirectoryMergeRuleInput](../../../pkg/models/shared/directorymergeruleinput.md)                                           | :heavy_minus_sign:                                                                                                                  | Every rule in the desired evaluation order, built-in rules included. Same rules and limits as<br/> SetDirectoryMergeRulesRequest.rules. |