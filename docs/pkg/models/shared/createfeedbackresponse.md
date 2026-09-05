# CreateFeedbackResponse

The CreateFeedbackResponse message.


## Fields

| Field                                                                                                      | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ID`                                                                                                       | `*string`                                                                                                  | :heavy_minus_sign:                                                                                         | The stored feedback ID.                                                                                    |
| `TicketURL`                                                                                                | `*string`                                                                                                  | :heavy_minus_sign:                                                                                         | URL of the task created for this feedback. Empty while task delivery is<br/> unavailable or has not succeeded. |