# CompositeKeyFieldSet

CompositeKeyFieldSet groups a non-empty set of composite-key
 fields as declared by one or more consumer reference sites that
 target the same Terraform type. Used in `TFSchemaMapping
 .referer_shapes` (the "inverted index" of composite-key shapes
 targeting this kind) so a multi-root producer can register
 canonical lookup keys for every shape its consumers might
 compute.

 Invariant: `fields` MUST be non-empty. The bare-id (single-id)
 form is implicit — every producer registers under
 `canonicalRefKey(id, {})` unconditionally, and consumer sites
 with empty composite_key_fields are not represented here. The
 backend's inverted-index computation skips them; including an
 empty `fields` would just round-trip to the bare-id form and
 produce a duplicate registration.


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `Fields`                                                                      | [][shared.CompositeKeyField](../../../pkg/models/shared/compositekeyfield.md) | :heavy_minus_sign:                                                            | The fields field.                                                             |