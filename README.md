<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-refresh-toc -->
**Table of Contents**

- [adventofcode](#adventofcode)
    - [Development](#development)
        - [Setup dev environment](#setup-dev-environment)
        - [Prepare to work with the project](#prepare-to-work-with-the-project)
        - [Run solutions from sources](#run-solutions-from-sources)
        - [Other commands](#other-commands)
    - [License](#license)

<!-- markdown-toc end -->

# adventofcode
Solutions to [Advent of Code](https://adventofcode.com/) puzzles.

## Development
### Setup dev environment
1. Install [`gofumpt`](https://github.com/mvdan/gofumpt#installation) for improved code formatting.
2. Install [`golangci-lint`](https://golangci-lint.run/usage/install/) for code linting.
3. Install [`pre-commit`](https://pre-commit.com/#install) to run linters before commit.

### Prepare to work with the project
1. Install `pre-commit` hooks by running:
   ```bash
   pre-commit install
   ```

### Run solutions from sources
1. Get your puzzle input and put it into folder of particular solution, e.g. `2023/1/input.txt`.

2. Run the solution:
   ```bash
   go run 2023/1/main.go
   ```

### Other commands
To get full list of available commands run:
```bash
make help
```

## License
Copyright (c) 2023 Alexander Kurbatov

Licensed under the [GPLv3](LICENSE).
