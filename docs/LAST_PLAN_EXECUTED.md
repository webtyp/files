# Plan (executed locally) — `webtyp.com/files`: one contract for whole-file I/O

## Context

The same idea existed three times with three vocabularies: `kvdb.Store`
(`GetFile`/`SetFile`/`AddToFile`), `pdf/fpdf`'s `WriteFileFunc`/`ReadFileFunc` options, and
the model loaders about to be written (`qwen`, `weights`, over `webtyp/opfs`). This repository
is the one contract they all use.

## Design gate (api-design)

1. **Prior art.** Go's `os.ReadFile`/`os.WriteFile` (the verbs), `io/fs.ReadFileFS` (a
   one-method read interface), `afero.Fs` and `go-billy` (filesystem abstractions swappable per
   environment), Node's `fs.readFile`/`writeFile`. We keep Go's verbs and `fs.ReadFileFS`'s
   one-method shape, and we do not import `io`/`io/fs`, whose `fs.FS` brings `Open`/`File`/`Stat`
   semantics that neither OPFS nor TinyGo consumers need.
2. **Novice-name test.** `files.Reader`, `files.Writer`, `files.Appender`,
   `files.ErrNotExist`. The package name is plural, like `strings`, `bytes` and `errors`, so it
   does not collide with the `file` variable every Go program has.
3. **Complexity ledger.** Concepts +4 / −3 (kvdb.Store's methods, fpdf's two function options).
   Ways to do the same thing +0 / −2 after kvdb and pdf migrate.
4. **Where it belongs.** A contract owned by no consumer and no backend, in its own
   repository, with conformance and a mem reference (the `webtyp/storage` pattern).
5. **What it deletes.** Downstream: `kvdb.Store`'s bespoke method names and fpdf's
   `WriteFileFunc`/`ReadFileFunc`.
