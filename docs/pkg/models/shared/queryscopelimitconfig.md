# QueryScopeLimitConfig

QueryScopeLimitConfig caps numeric fields (e.g. limit, page_size) in tool
 input so callers cannot request unbounded data.


## Fields

| Field               | Type                | Required            | Description         |
| ------------------- | ------------------- | ------------------- | ------------------- |
| `Fields`            | []`string`          | :heavy_minus_sign:  | The fields field.   |
| `MaxLimit`          | `*int`              | :heavy_minus_sign:  | The maxLimit field. |