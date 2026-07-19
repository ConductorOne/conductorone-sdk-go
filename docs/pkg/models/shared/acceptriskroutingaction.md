# AcceptRiskRoutingAction

AcceptRiskRoutingAction accepts the risk for a matched finding for a
 relative duration (resolved to risk_acceptance_expires_at = now + duration
 at execution time).


## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `Duration`         | `*string`          | :heavy_minus_sign: | N/A                |
| `Reason`           | `*string`          | :heavy_minus_sign: | The reason field.  |