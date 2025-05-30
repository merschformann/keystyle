# keyStyle

Map-key styling linter plugin for golangci lint.

## Usage

Instructions for setting up a custom linter can be found in the [golangci-lint documentation](https://golangci-lint.run/plugins/module-plugins/). Replace `foo` with `keyStyle` in the example (and the paths accordingly).

Configuration (`.golangci.yml`):

```yaml
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

## Setup and build

```bash
bash scripts/setup.sh
bash scripts/build.sh
```
