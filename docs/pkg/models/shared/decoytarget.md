# DecoyTarget

DecoyTarget points at the planted decoy that produced this finding.
 Populated for findings whose subject is the decoy artifact itself
 (e.g. decoy_credential_used), giving the UI and routing rules a
 uniform handle to the decoy alongside the finding_type payload.


## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `DecoyID`          | `*string`          | :heavy_minus_sign: | The decoyId field. |