# PaperSecretServiceGetContentRequest

The PaperSecretServiceGetContentRequest message.


## Fields

| Field                                                                                                          | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ReaderRecipient`                                                                                              | **string*                                                                                                      | :heavy_minus_sign:                                                                                             | Client's ephemeral Age recipient (age1...) for re-encryption<br/> Server re-encrypts the content to this recipient |