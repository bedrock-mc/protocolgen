package main

import "testing"

func TestPacketIDsByNumber(t *testing.T) {
	const src = `package packet

const (
	IDLogin = iota + 1
	IDPlayStatus
	_
	IDDisconnect
)`
	ids, err := packetIDsByNumber([]byte(src))
	if err != nil {
		t.Fatalf("packetIDsByNumber: %v", err)
	}
	if ids[1] != "IDLogin" || ids[2] != "IDPlayStatus" || ids[4] != "IDDisconnect" {
		t.Errorf("ids = %v, want 1=IDLogin 2=IDPlayStatus 4=IDDisconnect", ids)
	}
	if _, ok := ids[3]; ok {
		t.Errorf("blank identifier should not be recorded at 3")
	}
}
