# Nutrition Calculator

## Commands

```bash
go mod init nutriscorev1
go mod tidy
go work use .

go mod init nutriscorev1/components
go mod tidy
go work use .

go mod init nutriscorev1/score
go mod tidy
go work use .

go mod init nutriscorev1/types
go mod tidy
go work use .

go mod init nutriscorev1/utilities
go mod tidy
go work use .
```

## Testing

```bash
go test ./utilities
go test ./utilities -run TestGetPointsFromRange
```
