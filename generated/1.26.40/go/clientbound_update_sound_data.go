// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ClientboundUpdateSoundData struct {
	ServerSoundHandle ServerSoundHandle
	Stop              Optional[SoundDataEvent]
	SetVolume         Optional[SoundDataEvent]
	SetPitch          Optional[SoundDataEvent]
	Fade              Optional[SoundDataEvent]
	SeekTo            Optional[SoundDataEvent]
	Pause             Optional[SoundDataEvent]
	Resume            Optional[SoundDataEvent]
}
