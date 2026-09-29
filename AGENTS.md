# Agent Guide — `webtyp/files`

Constraints for agents working on this library. **Read this before any change.**

## What this library is

The contract for whole-file reading and writing (`Reader`, `Writer`, `Appender`, `ReadWriter`,
`ErrNotExist`), its conformance suite (`conformance/`) and an in-memory reference (`mem/`). It
must never contain a real backend (disk, OPFS, HTTP). Each lives in its own repository and runs
`conformance.Run` against itself.

## Rules of the contract

- A missing file is `ErrNotExist`, returned **unwrapped** so callers compare with `==`.
- Implementations copy: data passed to `WriteFile` and returned by `ReadFile` is never shared
  with the caller.
- Interfaces stay one method each. A new capability (listing, deleting) is a **new** interface,
  not a method added to an existing one, so existing implementations keep compiling.

## The builds that define "done"

```bash
go vet ./...
gotest
gotest -tinygo
```

## Never import these

| Never | Use instead | Why |
|---|---|---|
| `os`, `io`, `io/fs` | nothing: this is the contract, not a backend | backends live in their own repositories |
| `fmt`, `errors`, `strings` | nothing (the one error is a local type) | zero dependencies keeps every consumer's binary small |
| `map[K]V` | a slice | TinyGo's map runtime is a size tax |
