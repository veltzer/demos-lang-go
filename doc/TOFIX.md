# TOFIX

Findings from a code scan on 2026-10-04.

## Medium

- `slides/solutions/17-relay/relay.go:31` - `fmt.Printn` is a typo for `fmt.Println`; the solution does not compile.
- `slides/solutions/14-method/bank.go:20` - embeds `Customer` and uses `b.name`, but `Customer` is only defined in `slides/solutions/23-bank/customer.go`, while `23-bank` has the tests but no `bank.go` (its tests fail with `undefined: CreateAccount`); the two directories were split wrongly - put `bank.go` and `customer.go` together (or copy `customer.go` into `14-method`).
- `slides/solutions/23-stack/one_test.go:6` - tests call `Create()`/`Create(size)` and read `s.items`, but `stack.go` has no `Create` and its field is `Items [MaxSize]int`, so `go test` in this solution does not compile; make the tests and the solution match.
- `slides/solutions/24-example/example_test.go:3` - `ExampleMMandP` refers to an unknown identifier, so `go vet` rejects it and `go test` will not run it; rename it to `ExampleMakeMapAndPrint`.
- `slides/solutions/22-server2/server.go:18` - imports the sibling package as bare `"chanlist"` (GOPATH-era layout) and keeps `chanlist.go` (package `chanlist`) in the same directory as `package main`, so it cannot build under modules; move `chanlist.go` into a `chanlist/` subdirectory and add a `go.mod`.
- `rsconstruct.toml:54` - the generator only compiles `src/`, so nothing under `slides/solutions/` or `standalone/multiapp/` is ever built, which is how the broken solutions above went unnoticed; add the solutions (with `go vet`/`go test` per directory) and `standalone/multiapp` to the build. Several `slides/solutions` directories (`11-slice`, `13-map`, `16-files`, `others`) also hold several `package main` files with their own `main`, so they need building per file.

## Low

- `slides/README.md:5` - points to `talks.godoc.org`, which Google shut down, and `godoc.org` (line 3), which now only redirects to pkg.go.dev; link `golang.org/x/tools/present` on pkg.go.dev and drop the talks.godoc.org hosting suggestion.
- `slides/present.sh:3` - runs `present presentation.slide` relative to the current directory, so it only works when started from `slides/`; `cd "$(dirname "$0")"` first.
- `rsconstruct.toml:28` - `ruff`/`mypy` (line 32) list `src` and `config`, and `shellcheck` (line 37) lists `src` and `config`, in `src_dirs`, but `src/` holds only `.go` and `config/` only `.lua` files; list only `scripts` for ruff/mypy and drop `src`/`config` from shellcheck.
