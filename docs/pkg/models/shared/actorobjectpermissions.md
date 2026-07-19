# ActorObjectPermissions

Legacy: do not use for new objects. Retained only for the existing
 AppResource / AppEntitlement / access-review consumers, which will migrate to
 c1.api.authorization.v1.ActorObjectPermissions in IGA-2331. New object views
 should reference c1.api.authorization.v1.ActorObjectPermissions instead.


## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `Delete`           | `*bool`            | :heavy_minus_sign: | The delete field.  |
| `Edit`             | `*bool`            | :heavy_minus_sign: | The edit field.    |
| `Extra`            | map[string]`bool`  | :heavy_minus_sign: | The extra field.   |
| `Read`             | `*bool`            | :heavy_minus_sign: | The read field.    |