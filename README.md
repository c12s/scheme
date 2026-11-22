# scheme — legacy shared proto repository

The original c12s "one place for all `.proto` files" repository: per-service
protobuf definitions (`apollo`, `blackhole`, `celestial`, `core`, `flusher`,
`gravity`, `magnetar`, `meridian`, `stellar`) with the generated Go code
(`*.pb.go`, old `protoc-gen-go` style) committed next to them.

## Role in c12s

**Historical / legacy.** The active services no longer import this module:
each repo carries its own protos and generated code (`magnetar/pkg/api`,
`gravity/pkg/api`, `kuiper/pkg/api`, `star/pkg/api`, `apollo/proto1`,
`oort/pkg/api`), and consumers wire them via `go.mod` `replace` directives to
the sibling directories in this workspace.

Remaining users are the legacy repos only:

| Repo | Imports |
|---|---|
| `celestial`, `blackhole`, `stellar`, `stellar-go` | `github.com/c12s/scheme/...` |

The protos here also **diverge from today's services**: e.g.
`scheme/magnetar/magnetar.proto` still declares `MagnetarService`
(`Reserve`/`List`/`Free`/`Query`) while the current Magnetar gRPC API is
`GetFromNodePool`/`ClaimOwnership`/`SetDesiredOsImage`/… and lives in
`magnetar/pkg/api/proto/`; `scheme/gravity/gravity.proto` is the old
`GravityService.PutTask`, not the current `AgentQueue`.

## Contents

```
<service>/<service>.proto     definition
<service>/<service>.pb.go     generated code (committed)
gen.sh                        regeneration script
```

## Regeneration (as shipped — stale)

```bash
# from the repo root, with $GOPATH set and the ancient toolchain:
protoc -I. --go_out=plugins=grpc:$GOPATH/src/. magnetar/magnetar.proto
# ... see gen.sh for all services
```
