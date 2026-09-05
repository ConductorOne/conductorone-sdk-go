# SlackChannelTarget

SlackChannelTarget names one Slack channel. Exactly one of channel_name /
 channel_id is set; a name is resolved at send time, so an unresolvable name
 fails the dispatch rather than the rule edit.


## Fields

| Field                  | Type                   | Required               | Description            |
| ---------------------- | ---------------------- | ---------------------- | ---------------------- |
| `ChannelID`            | `*string`              | :heavy_minus_sign:     | The channelId field.   |
| `ChannelName`          | `*string`              | :heavy_minus_sign:     | The channelName field. |