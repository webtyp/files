# files
<img src="docs/img/badges.svg">

The webtyp contract for reading and writing **whole files by path**. A library that needs a
file (a PDF generator reading a font, a model loader reading its weights, a key-value store
persisting itself) asks for the narrowest interface it needs, and the application decides
where files actually live.

| I want to… | Use |
|---|---|
| read a file | `files.Reader` → `ReadFile(path) ([]byte, error)` |
| write (replace) a file | `files.Writer` → `WriteFile(path, data) error` |
| add to the end of a file | `files.Appender` → `AppendFile(path, data) error` |
| delete a file | `files.Remover` → `RemoveFile(path) error` (`files.ErrNotExist` when missing) |
| both read and write | `files.ReadWriter` |
| know that a file is missing | `err == files.ErrNotExist` (never wrapped) |
| test without touching disk | `mem.New()` from `webtyp.com/files/mem` |
| prove my implementation is correct | `conformance.Run(t, conformance.Factory{...})` from `webtyp.com/files/conformance` |

Implementations live in their own repositories: `webtyp/opfs` (browser, inside a Web Worker)
and a disk implementation for servers. This repository holds only the contract, the
conformance suite and the in-memory reference, the same pattern as `webtyp/storage`.

## Documentation

- [Agent guide](AGENTS.md): rules for anyone changing this library.
- [Last executed plan](docs/LAST_PLAN_EXECUTED.md): the design gate behind the contract.
