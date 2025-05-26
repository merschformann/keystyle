# keyStyle

Map-key styling linter plugin for golangci lint.

## Usage

Instructions for setting up a custom linter can be found in the [golangci-lint documentation](https://golangci-lint.run/plugins/module-plugins/). Replace `foo` with `keyStyle` in the example (and the paths accordingly).

Configuration:

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
          style: "camelCase"
```

## Setup and build

```bash
bash scripts/setup.sh
bash scripts/build.sh
```
