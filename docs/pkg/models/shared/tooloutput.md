# ToolOutput


## Supported Types

### 

```go
toolOutput := shared.CreateToolOutputStr(string{/* values here */})
```

### 

```go
toolOutput := shared.CreateToolOutputNumber(float64{/* values here */})
```

### GatedToolCallTarget3

```go
toolOutput := shared.CreateToolOutputGatedToolCallTarget3(shared.GatedToolCallTarget3{/* values here */})
```

### 

```go
toolOutput := shared.CreateToolOutputArrayOfAny([]any{/* values here */})
```

### 

```go
toolOutput := shared.CreateToolOutputBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch toolOutput.Type {
	case shared.ToolOutputTypeStr:
		// toolOutput.Str is populated
	case shared.ToolOutputTypeNumber:
		// toolOutput.Number is populated
	case shared.ToolOutputTypeGatedToolCallTarget3:
		// toolOutput.GatedToolCallTarget3 is populated
	case shared.ToolOutputTypeArrayOfAny:
		// toolOutput.ArrayOfAny is populated
	case shared.ToolOutputTypeBoolean:
		// toolOutput.Boolean is populated
}
```
