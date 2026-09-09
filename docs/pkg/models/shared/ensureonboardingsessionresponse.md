# EnsureOnboardingSessionResponse

Returns the active onboarding conversation and whether this call created it.


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ConversationID`                                               | `*string`                                                      | :heavy_minus_sign:                                             | The active onboarding conversation ID.                         |
| `Created`                                                      | `*bool`                                                        | :heavy_minus_sign:                                             | True only when this call created and started the conversation. |