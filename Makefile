SHELL := /bin/sh

GO_CACHE ?= /tmp/go-build-cache
GO := GOCACHE=$(GO_CACHE) go
PROTOCOLGEN := $(GO) run ./cmd/protocolgen

MANIFEST := generated/1.26.40/manifest.json
NAMING := generated/1.26.40/naming.json
DOMAINS := generated/1.26.40/domains.json
DOCS := generated/1.26.40/docs.json
SOURCE_LOCK := generated/1.26.40/source-lock.json
MOJANG_DIR ?=
ENDSTONE_DIR ?=
GOPHERTUNNEL_DIR ?=
ORACLE_REPORT ?= /tmp/protocolgen-gophertunnel-report.json
GOPHER_ARGS = $(if $(GOPHERTUNNEL_DIR),-gophertunnel $(GOPHERTUNNEL_DIR))

.PHONY: regen differential verify

differential:
	$(GO) -C differential test ./...

regen:
	@test -n "$(MOJANG_DIR)" || (echo "MOJANG_DIR is required" >&2; exit 2)
	@test -n "$(ENDSTONE_DIR)" || (echo "ENDSTONE_DIR is required" >&2; exit 2)
	$(PROTOCOLGEN) reconcile \
		-lock $(SOURCE_LOCK) \
		-directions generated/1.26.40/directions.json \
		-nbt-encodings generated/1.26.40/nbt-encodings.json \
		-mojang $(MOJANG_DIR) \
		-mojang-corrections generated/1.26.40/corrections/mojang \
		-endstone $(ENDSTONE_DIR) \
		-endstone-corrections generated/1.26.40/corrections/endstone \
		-adjudications generated/1.26.40/adjudications.json \
		-out $(MANIFEST)
	$(PROTOCOLGEN) validate -manifest $(MANIFEST)
	$(PROTOCOLGEN) emit-go \
		-manifest $(MANIFEST) \
		-naming $(NAMING) \
		-domains $(DOMAINS) \
		-docs $(DOCS) \
		-out generated/1.26.40/go \
		-protocol-import protocolgen/generated/1.26.40/go/protocol
	$(PROTOCOLGEN) emit-rust \
		-manifest $(MANIFEST) \
		-naming $(NAMING) \
		-domains $(DOMAINS) \
		-docs $(DOCS) \
		-out generated/1.26.40/rust
	$(PROTOCOLGEN) verify-gophertunnel \
		-manifest $(MANIFEST) \
		-report $(ORACLE_REPORT) $(GOPHER_ARGS)
	$(PROTOCOLGEN) parity \
		-manifest testdata/parity/v2-small.json \
		-axolotl testdata/parity/axolotl-v1-small.json
	$(GO) build ./...
	$(GO) test ./...
	$(GO) vet ./...

verify: regen differential
	@test -z "$$(git status --porcelain)" || (echo "regeneration produced drift:" >&2; git status --short >&2; exit 1)
