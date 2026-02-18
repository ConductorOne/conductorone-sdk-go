# A2UIComponent

A2UIComponent - typed discriminated union for UI components.
 Uses oneof for proper TypeScript discrimination.

This message contains a oneof named component. Only a single field of the following list may be set at a time:
  - text
  - textField
  - checkBox
  - choicePicker
  - dateTimeInput
  - slider
  - button
  - row
  - column
  - card
  - divider
  - c1StatusIndicator
  - c1CodeBlock
  - c1ResourcePicker



## Fields

| Field                                                                                          | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ButtonComponent`                                                                              | [*shared.ButtonComponent](../../../pkg/models/shared/buttoncomponent.md)                       | :heavy_minus_sign:                                                                             | ButtonComponent triggers actions.                                                              |
| `C1CodeBlockComponent`                                                                         | [*shared.C1CodeBlockComponent](../../../pkg/models/shared/c1codeblockcomponent.md)             | :heavy_minus_sign:                                                                             | C1CodeBlockComponent displays code with syntax highlighting.                                   |
| `C1ResourcePickerComponent`                                                                    | [*shared.C1ResourcePickerComponent](../../../pkg/models/shared/c1resourcepickercomponent.md)   | :heavy_minus_sign:                                                                             | C1ResourcePickerComponent allows selecting C1 resources.                                       |
| `C1StatusIndicatorComponent`                                                                   | [*shared.C1StatusIndicatorComponent](../../../pkg/models/shared/c1statusindicatorcomponent.md) | :heavy_minus_sign:                                                                             | C1StatusIndicatorComponent shows agent progress status.                                        |
| `CardComponent`                                                                                | [*shared.CardComponent](../../../pkg/models/shared/cardcomponent.md)                           | :heavy_minus_sign:                                                                             | CardComponent is a container with styling.                                                     |
| `CheckBoxComponent`                                                                            | [*shared.CheckBoxComponent](../../../pkg/models/shared/checkboxcomponent.md)                   | :heavy_minus_sign:                                                                             | CheckBoxComponent is a boolean checkbox.                                                       |
| `ChoicePickerComponent`                                                                        | [*shared.ChoicePickerComponent](../../../pkg/models/shared/choicepickercomponent.md)           | :heavy_minus_sign:                                                                             | ChoicePickerComponent allows selection from predefined choices.                                |
| `ColumnComponent`                                                                              | [*shared.ColumnComponent](../../../pkg/models/shared/columncomponent.md)                       | :heavy_minus_sign:                                                                             | ColumnComponent arranges children vertically.                                                  |
| `DateTimeInputComponent`                                                                       | [*shared.DateTimeInputComponent](../../../pkg/models/shared/datetimeinputcomponent.md)         | :heavy_minus_sign:                                                                             | DateTimeInputComponent for date/time selection.                                                |
| `DividerComponent`                                                                             | [*shared.DividerComponent](../../../pkg/models/shared/dividercomponent.md)                     | :heavy_minus_sign:                                                                             | DividerComponent is a visual separator.                                                        |
| `RowComponent`                                                                                 | [*shared.RowComponent](../../../pkg/models/shared/rowcomponent.md)                             | :heavy_minus_sign:                                                                             | RowComponent arranges children horizontally.                                                   |
| `SliderComponent`                                                                              | [*shared.SliderComponent](../../../pkg/models/shared/slidercomponent.md)                       | :heavy_minus_sign:                                                                             | SliderComponent for numeric range selection.                                                   |
| `TextComponent`                                                                                | [*shared.TextComponent](../../../pkg/models/shared/textcomponent.md)                           | :heavy_minus_sign:                                                                             | TextComponent displays text content.                                                           |
| `TextFieldComponent`                                                                           | [*shared.TextFieldComponent](../../../pkg/models/shared/textfieldcomponent.md)                 | :heavy_minus_sign:                                                                             | TextFieldComponent is a text input field.                                                      |
| `ID`                                                                                           | **string*                                                                                      | :heavy_minus_sign:                                                                             | The id field.                                                                                  |
| `Weight`                                                                                       | **int*                                                                                         | :heavy_minus_sign:                                                                             | The weight field.                                                                              |