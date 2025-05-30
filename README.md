# keyStyle

> [!IMPORTANT]
> This project is a first attempt at writing a custom linter and used for personal purposes mainly. I am happy if someone finds it useful, but please be aware that it is not necessarily production-ready or well-tested.

Map-key styling linter plugin for [golangci-lint](https://golangci-lint.run/).

## Example

```go
// ...
someData = LogData{
  "camelCase":     "value1",
  "IncorrectKey2": "value2",
}
// ...
```

Linter output:

```txt
example.go:10:3: Key 'IncorrectKey2' style should be camelCase (keystyle)
  "IncorrectKey2": someValue,
  ^
```

## Usage

> Instructions for setting up a custom linter can be found in the [golangci-lint documentation](https://golangci-lint.run/plugins/module-plugins/). Replace `foo` with `keyStyle` in the example (and the paths accordingly).

The following instructions are based on above documentation.

Linter plugin setup (`.custom-gcl.yml`):

```yaml
version: v2.1.6
plugins:
  # Import keystyle linter for linting of map key style.
  - module: 'github.com/merschformann/keystyle'
    import: 'github.com/merschformann/keystyle'
    version: v0.0.1
```

Linter configuration (`.golangci.yml`):

```yaml
version: "2"
linters:
  default: none
  enable:
    - keystyle
  settings:
    custom:
      keystyle:
        type: "module"
        description: Keystyle preference.
        settings:
          checks:
            - style: "camelCase"
              type-name: "LogData"
            - style: "custom"
              type-name: "OtherData"
              regex: "^[a-z][a-zA-Z0-9]*$"
```
