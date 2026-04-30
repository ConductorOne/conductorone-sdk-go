# ManuallyManagedUsersResponse

The ManuallyManagedUsersResponse message.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `BulkActionID`                                                         | `*string`                                                              | :heavy_minus_sign:                                                     | The ID of the bulk action created to process the membership additions. |
| `FailedUsersErrorMap`                                                  | map[string]`string`                                                    | :heavy_minus_sign:                                                     | A map of user IDs to error messages for users that could not be added. |