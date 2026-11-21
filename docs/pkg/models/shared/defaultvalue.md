# DefaultValue


## Supported Types

### 

```go
defaultValue := shared.CreateDefaultValueStr(string{/* values here */})
```

### 

```go
defaultValue := shared.CreateDefaultValueNumber(float64{/* values here */})
```

### Three

```go
defaultValue := shared.CreateDefaultValueThree(shared.Three{/* values here */})
```

### 

```go
defaultValue := shared.CreateDefaultValueArrayOfAny([]any{/* values here */})
```

### 

```go
defaultValue := shared.CreateDefaultValueBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch defaultValue.Type {
	case shared.DefaultValueTypeStr:
		// defaultValue.Str is populated
	case shared.DefaultValueTypeNumber:
		// defaultValue.Number is populated
	case shared.DefaultValueTypeThree:
		// defaultValue.Three is populated
	case shared.DefaultValueTypeArrayOfAny:
		// defaultValue.ArrayOfAny is populated
	case shared.DefaultValueTypeBoolean:
		// defaultValue.Boolean is populated
}
```
