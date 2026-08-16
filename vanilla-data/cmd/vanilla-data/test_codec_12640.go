//go:build protocolgen_12640

package main

import genprotocol "protocolgen/generated/1.26.40/go/protocol"

func testGeneratedVersion() string { return genprotocol.GAME_VERSION }
