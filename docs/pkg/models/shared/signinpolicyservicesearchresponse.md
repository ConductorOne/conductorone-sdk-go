# SignInPolicyServiceSearchResponse

The SignInPolicyServiceSearchResponse message.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `List`                                                                 | [][shared.SignInPolicy](../../../pkg/models/shared/signinpolicy.md)    | :heavy_minus_sign:                                                     | The page of matching policies.                                         |
| `NextPageToken`                                                        | `*string`                                                              | :heavy_minus_sign:                                                     | A token to fetch the next page, or empty if there are no more results. |