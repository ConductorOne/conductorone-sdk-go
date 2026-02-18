# PaperSecretServiceSetTextContentResponse

The PaperSecretServiceSetTextContentResponse message.


## Fields

| Field                                                                                                                             | Type                                                                                                                              | Required                                                                                                                          | Description                                                                                                                       |
| --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- |
| `PaperSecret`                                                                                                                     | [*shared.PaperSecret](../../../pkg/models/shared/papersecret.md)                                                                  | :heavy_minus_sign:                                                                                                                | PaperSecret is the API view of a secret (combines Vault + PaperVault fields).<br/> The vault_id is the primary identifier (Vault.id). |