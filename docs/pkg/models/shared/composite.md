# Composite

Composite import IDs combine values from multiple component fields
 per the declared `format`.


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `Fields`                                                                | [][shared.CompositeField](../../../pkg/models/shared/compositefield.md) | :heavy_minus_sign:                                                      | Component fields, in the order they participate in the import<br/> ID.  |
| `Format`                                                                | [*shared.Format](../../../pkg/models/shared/format.md)                  | :heavy_minus_sign:                                                      | Wire format the provider expects. Defaults to<br/> FORMAT_JSON_OBJECT.  |