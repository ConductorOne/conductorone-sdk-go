# ResourceProfileMatch

Matches a top-level resource profile key. Scalar text matches an equal scalar or list member;
 JSON numbers and booleans match their typed equivalents; JSON arrays require the stored list
 to contain every element.


## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `Key`              | `*string`          | :heavy_minus_sign: | The key field.     |
| `Value`            | `*string`          | :heavy_minus_sign: | The value field.   |