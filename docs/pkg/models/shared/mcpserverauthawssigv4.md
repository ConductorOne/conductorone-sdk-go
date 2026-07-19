# MCPServerAuthAWSSigV4

MCPServerAuthAWSSigV4 provides AWS Signature Version 4 authentication.
 Used for hosted MCP servers backed by AWS service impls (the
 amazonaws_com_* catalog entries). Outbound requests from the gateway are
 signed per-request with the supplied access key + secret, against the
 service+region scope sourced from the impl bundle's connect.auth.

 Only the SHARED token-sharing model is supported — every caller in the
 tenant signs with the same admin-configured credentials. Per-user AWS
 (STS / Web Identity / IAM Identity Center) is a separate future surface.


## Fields

| Field                                                                                                                                               | Type                                                                                                                                                | Required                                                                                                                                            | Description                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| `AccessKeyID`                                                                                                                                       | `*string`                                                                                                                                           | :heavy_minus_sign:                                                                                                                                  | AWS access key ID (the IAM user / role's public identifier, e.g.<br/> "AKIAIOSFODNN7EXAMPLE"). Persisted in plaintext form; the secret half<br/> is sealed. |
| `SecretAccessKey`                                                                                                                                   | `*string`                                                                                                                                           | :heavy_minus_sign:                                                                                                                                  | AWS secret access key. Sealed by the backend on write; never returned on<br/> read.                                                                 |
| `SessionToken`                                                                                                                                      | `*string`                                                                                                                                           | :heavy_minus_sign:                                                                                                                                  | Optional AWS session token. Set only when the credential is a<br/> short-lived STS temporary credential. Static IAM-user keys leave this<br/> empty. |