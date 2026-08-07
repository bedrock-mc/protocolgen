// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ServerSoundHandle struct {
	ServerSoundHandle uint64
}

// Marshal reads or writes ServerSoundHandle using its canonical wire layout.
func (x *ServerSoundHandle) Marshal(io IO) {
	io.Uint64(&x.ServerSoundHandle)
}
