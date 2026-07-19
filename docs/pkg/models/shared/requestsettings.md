# RequestSettings

RequestSettings holds tenant-wide configuration for the access-request flow.


## Fields

| Field                                                                                                    | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `SkipJustification`                                                                                      | `*bool`                                                                                                  | :heavy_minus_sign:                                                                                       | When true, request surfaces (webapp, Slack, MS Teams) skip prompting the<br/> requester for a justification. |