# AppUserMergeRulesServiceListSourcesResponse

The AppUserMergeRulesServiceListSourcesResponse message contains the records that make up an account.


## Fields

| Field                                                                                                         | Type                                                                                                          | Required                                                                                                      | Description                                                                                                   |
| ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| `Sources`                                                                                                     | [][shared.AppUserSource](../../../pkg/models/shared/appusersource.md)                                         | :heavy_minus_sign:                                                                                            | The account's records: records from the application's connectors first, then records from other applications. |