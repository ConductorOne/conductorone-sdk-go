# TriggerAutomationDispatcher

TriggerAutomationDispatcher runs a C1 automation by id (the "Run now" path).


## Fields

| Field                                                                                                         | Type                                                                                                          | Required                                                                                                      | Description                                                                                                   |
| ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| `AutomationID`                                                                                                | `*string`                                                                                                     | :heavy_minus_sign:                                                                                            | ID of the C1 automation/workflow to run.                                                                      |
| `InputMapping`                                                                                                | map[string]`string`                                                                                           | :heavy_minus_sign:                                                                                            | Inputs passed to the automation, keyed by input name (v0: verbatim values;<br/> CEL evaluation is a later phase). |