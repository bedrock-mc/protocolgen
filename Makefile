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

VANILLA_VERSION ?= 1.26.44
VANILLA_MANIFEST := ../generated/$(VANILLA_VERSION)/manifest.json
VANILLA_SOURCE := ../generated/$(VANILLA_VERSION)/vanilla-source.json
VANILLA_OUT := ../generated/$(VANILLA_VERSION)/vanilla-data
VANILLA_GO_ARGS := $(if $(filter 1.26.40,$(VANILLA_VERSION)),-modfile go-1.26.40.mod -tags protocolgen_12640,)
BDS_ADDRESS ?= 127.0.0.1:19132
BDS_BINARY ?=

HOTFIX_MANIFEST := generated/1.26.44/manifest.json
HOTFIX_SPEC := generated/1.26.44/hotfix.json
HOTFIX_NAMING := generated/1.26.44/naming.json
HOTFIX_DOMAINS := generated/1.26.44/domains.json
HOTFIX_DOCS := generated/1.26.44/docs.json

TARGET_12650 := generated/1.26.50
CLAIMS_12650_MOJANG ?= /tmp/protocolgen-1.26.50-mojang-claims.json
CLAIMS_12650_ENDSTONE ?= /tmp/protocolgen-1.26.50-endstone-claims.json

.PHONY: regen regen-1.26.50 hotfix vanilla-data differential verify verify-1.26.50

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

regen-1.26.50:
	@test -n "$(MOJANG_DIR)" || (echo "MOJANG_DIR is required" >&2; exit 2)
	@test -n "$(ENDSTONE_DIR)" || (echo "ENDSTONE_DIR is required" >&2; exit 2)
	$(PROTOCOLGEN) ingest \
		-lock $(TARGET_12650)/source-lock.json -kind mojang -id mojang \
		-root $(MOJANG_DIR) -corrections $(TARGET_12650)/corrections/mojang \
		-out $(CLAIMS_12650_MOJANG)
	$(PROTOCOLGEN) ingest \
		-lock $(TARGET_12650)/source-lock.json -kind endstone -id endstone \
		-root $(ENDSTONE_DIR) -corrections $(TARGET_12650)/corrections/endstone \
		-out $(CLAIMS_12650_ENDSTONE)
	$(PROTOCOLGEN) reconcile-claims \
		-lock $(TARGET_12650)/source-lock.json \
		-claims $(CLAIMS_12650_MOJANG) \
		-claims $(CLAIMS_12650_ENDSTONE) \
		-claims $(TARGET_12650)/lens-claims.json \
		-adjudications $(TARGET_12650)/adjudications.json \
		-directions $(TARGET_12650)/directions.json \
		-nbt-encodings $(TARGET_12650)/nbt-encodings.json \
		-out $(TARGET_12650)/manifest.json
	$(PROTOCOLGEN) validate -manifest $(TARGET_12650)/manifest.json
	$(PROTOCOLGEN) emit-go \
		-manifest $(TARGET_12650)/manifest.json \
		-naming $(TARGET_12650)/naming.json \
		-domains $(TARGET_12650)/domains.json \
		-docs $(TARGET_12650)/docs.json \
		-out $(TARGET_12650)/go \
		-protocol-import protocolgen/generated/1.26.50/go/protocol
	$(PROTOCOLGEN) emit-rust \
		-manifest $(TARGET_12650)/manifest.json \
		-naming $(TARGET_12650)/naming.json \
		-domains $(TARGET_12650)/domains.json \
		-docs $(TARGET_12650)/docs.json \
		-out $(TARGET_12650)/rust
	cargo fmt --manifest-path $(TARGET_12650)/rust/Cargo.toml
	$(GO) test ./generated/1.26.50/go/... -count=1
	cargo test --manifest-path $(TARGET_12650)/rust/Cargo.toml

hotfix:
	$(PROTOCOLGEN) hotfix \
		-base $(MANIFEST) \
		-spec $(HOTFIX_SPEC) \
		-out $(HOTFIX_MANIFEST)
	$(PROTOCOLGEN) validate -manifest $(HOTFIX_MANIFEST)
	$(PROTOCOLGEN) emit-go \
		-manifest $(HOTFIX_MANIFEST) \
		-naming $(HOTFIX_NAMING) \
		-domains $(HOTFIX_DOMAINS) \
		-docs $(HOTFIX_DOCS) \
		-out generated/1.26.44/go \
		-protocol-import protocolgen/generated/1.26.44/go/protocol
	$(PROTOCOLGEN) emit-rust \
		-manifest $(HOTFIX_MANIFEST) \
		-naming $(HOTFIX_NAMING) \
		-domains $(HOTFIX_DOMAINS) \
		-docs $(HOTFIX_DOCS) \
		-out generated/1.26.44/rust

vanilla-data:
	@test -n "$(BDS_BINARY)" || (echo "BDS_BINARY is required" >&2; exit 2)
	$(GO) -C vanilla-data run $(VANILLA_GO_ARGS) ./cmd/vanilla-data \
		-manifest $(VANILLA_MANIFEST) \
		-source $(VANILLA_SOURCE) \
		-out $(VANILLA_OUT) \
		-bds-binary $(BDS_BINARY) \
		-address $(BDS_ADDRESS)

verify: regen hotfix differential
	@test -z "$$(git status --porcelain)" || (echo "regeneration produced drift:" >&2; git status --short >&2; exit 1)

verify-1.26.50: regen-1.26.50
	@git diff --exit-code -- $(TARGET_12650) || (echo "1.26.50 regeneration produced drift" >&2; exit 1)
