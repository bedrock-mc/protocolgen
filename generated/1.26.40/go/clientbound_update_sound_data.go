// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol2168

type ClientboundUpdateSoundData struct {
	ServerSoundHandle ServerSoundHandle
	Stop              *SoundDataEvent
	SetVolume         *SoundDataEvent
	SetPitch          *SoundDataEvent
	Fade              *SoundDataEvent
	SeekTo            *SoundDataEvent
	Pause             *SoundDataEvent
	Resume            *SoundDataEvent
}
