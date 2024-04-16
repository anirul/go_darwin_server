// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: client_audio.proto

package darwin_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AudioMusicEnum int32

const (
	AudioMusicEnum_AUDIO_MUSIC_NONE  AudioMusicEnum = 0 // No audio state (this is an error).
	AudioMusicEnum_AUDIO_MUSIC_TITLE AudioMusicEnum = 1 // Title audio.
	AudioMusicEnum_AUDIO_MUSIC_MENU  AudioMusicEnum = 2 // Menu audio.
	AudioMusicEnum_AUDIO_MUSIC_PLAY  AudioMusicEnum = 3 // Play audio.
	AudioMusicEnum_AUDIO_MUSIC_DEATH AudioMusicEnum = 4 // Death audio (game over).
	AudioMusicEnum_AUDIO_MUSIC_WIN   AudioMusicEnum = 5 // Win audio (wining the game).
)

// Enum value maps for AudioMusicEnum.
var (
	AudioMusicEnum_name = map[int32]string{
		0: "AUDIO_MUSIC_NONE",
		1: "AUDIO_MUSIC_TITLE",
		2: "AUDIO_MUSIC_MENU",
		3: "AUDIO_MUSIC_PLAY",
		4: "AUDIO_MUSIC_DEATH",
		5: "AUDIO_MUSIC_WIN",
	}
	AudioMusicEnum_value = map[string]int32{
		"AUDIO_MUSIC_NONE":  0,
		"AUDIO_MUSIC_TITLE": 1,
		"AUDIO_MUSIC_MENU":  2,
		"AUDIO_MUSIC_PLAY":  3,
		"AUDIO_MUSIC_DEATH": 4,
		"AUDIO_MUSIC_WIN":   5,
	}
)

func (x AudioMusicEnum) Enum() *AudioMusicEnum {
	p := new(AudioMusicEnum)
	*p = x
	return p
}

func (x AudioMusicEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AudioMusicEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_client_audio_proto_enumTypes[0].Descriptor()
}

func (AudioMusicEnum) Type() protoreflect.EnumType {
	return &file_client_audio_proto_enumTypes[0]
}

func (x AudioMusicEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AudioMusicEnum.Descriptor instead.
func (AudioMusicEnum) EnumDescriptor() ([]byte, []int) {
	return file_client_audio_proto_rawDescGZIP(), []int{0}
}

type AudioSoundEnum int32

const (
	AudioSoundEnum_AUDIO_SOUND_NONE AudioSoundEnum = 0 // No audio sound (this is an error).
	AudioSoundEnum_AUDIO_SOUND_JUMP AudioSoundEnum = 1 // Jump audio sound.
	AudioSoundEnum_AUDIO_SOUND_GOOD AudioSoundEnum = 2 // Good audio sound (you took some mass).
	AudioSoundEnum_AUDIO_SOUND_BAD  AudioSoundEnum = 3 // Bad audio sound (you lost some mass).
)

// Enum value maps for AudioSoundEnum.
var (
	AudioSoundEnum_name = map[int32]string{
		0: "AUDIO_SOUND_NONE",
		1: "AUDIO_SOUND_JUMP",
		2: "AUDIO_SOUND_GOOD",
		3: "AUDIO_SOUND_BAD",
	}
	AudioSoundEnum_value = map[string]int32{
		"AUDIO_SOUND_NONE": 0,
		"AUDIO_SOUND_JUMP": 1,
		"AUDIO_SOUND_GOOD": 2,
		"AUDIO_SOUND_BAD":  3,
	}
)

func (x AudioSoundEnum) Enum() *AudioSoundEnum {
	p := new(AudioSoundEnum)
	*p = x
	return p
}

func (x AudioSoundEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AudioSoundEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_client_audio_proto_enumTypes[1].Descriptor()
}

func (AudioSoundEnum) Type() protoreflect.EnumType {
	return &file_client_audio_proto_enumTypes[1]
}

func (x AudioSoundEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AudioSoundEnum.Descriptor instead.
func (AudioSoundEnum) EnumDescriptor() ([]byte, []int) {
	return file_client_audio_proto_rawDescGZIP(), []int{1}
}

type AudioFormatEnum int32

const (
	AudioFormatEnum_AUDIO_FORMAT_NONE    AudioFormatEnum = 0 // No audio format (this is an error).
	AudioFormatEnum_AUDIO_FORMAT_DEFAULT AudioFormatEnum = 1 // Default audio format.
	AudioFormatEnum_AUDIO_FORMAT_S16_LSB AudioFormatEnum = 2 // Signed 16-bit little-endian audio format.
	AudioFormatEnum_AUDIO_FORMAT_S16_MSB AudioFormatEnum = 3 // Signed 16-bit big-endian audio format.
	AudioFormatEnum_AUDIO_FORMAT_S16_SYS AudioFormatEnum = 4 // Signed 16-bit system-endian audio format.
)

// Enum value maps for AudioFormatEnum.
var (
	AudioFormatEnum_name = map[int32]string{
		0: "AUDIO_FORMAT_NONE",
		1: "AUDIO_FORMAT_DEFAULT",
		2: "AUDIO_FORMAT_S16_LSB",
		3: "AUDIO_FORMAT_S16_MSB",
		4: "AUDIO_FORMAT_S16_SYS",
	}
	AudioFormatEnum_value = map[string]int32{
		"AUDIO_FORMAT_NONE":    0,
		"AUDIO_FORMAT_DEFAULT": 1,
		"AUDIO_FORMAT_S16_LSB": 2,
		"AUDIO_FORMAT_S16_MSB": 3,
		"AUDIO_FORMAT_S16_SYS": 4,
	}
)

func (x AudioFormatEnum) Enum() *AudioFormatEnum {
	p := new(AudioFormatEnum)
	*p = x
	return p
}

func (x AudioFormatEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AudioFormatEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_client_audio_proto_enumTypes[2].Descriptor()
}

func (AudioFormatEnum) Type() protoreflect.EnumType {
	return &file_client_audio_proto_enumTypes[2]
}

func (x AudioFormatEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AudioFormatEnum.Descriptor instead.
func (AudioFormatEnum) EnumDescriptor() ([]byte, []int) {
	return file_client_audio_proto_rawDescGZIP(), []int{2}
}

// MusicQueue
// Next: 5
type MusicQueue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AudioFile      string         `protobuf:"bytes,1,opt,name=audio_file,json=audioFile,proto3" json:"audio_file,omitempty"`
	Title          string         `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author         string         `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	AudioMusicEnum AudioMusicEnum `protobuf:"varint,4,opt,name=audio_music_enum,json=audioMusicEnum,proto3,enum=proto.AudioMusicEnum" json:"audio_music_enum,omitempty"`
}

func (x *MusicQueue) Reset() {
	*x = MusicQueue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_audio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MusicQueue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MusicQueue) ProtoMessage() {}

func (x *MusicQueue) ProtoReflect() protoreflect.Message {
	mi := &file_client_audio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MusicQueue.ProtoReflect.Descriptor instead.
func (*MusicQueue) Descriptor() ([]byte, []int) {
	return file_client_audio_proto_rawDescGZIP(), []int{0}
}

func (x *MusicQueue) GetAudioFile() string {
	if x != nil {
		return x.AudioFile
	}
	return ""
}

func (x *MusicQueue) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MusicQueue) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *MusicQueue) GetAudioMusicEnum() AudioMusicEnum {
	if x != nil {
		return x.AudioMusicEnum
	}
	return AudioMusicEnum_AUDIO_MUSIC_NONE
}

// AudioSound
// Next: 5
type SoundQueue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AudioFile      string         `protobuf:"bytes,1,opt,name=audio_file,json=audioFile,proto3" json:"audio_file,omitempty"`
	Title          string         `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author         string         `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	AudioSoundEnum AudioSoundEnum `protobuf:"varint,4,opt,name=audio_sound_enum,json=audioSoundEnum,proto3,enum=proto.AudioSoundEnum" json:"audio_sound_enum,omitempty"`
}

func (x *SoundQueue) Reset() {
	*x = SoundQueue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_audio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SoundQueue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SoundQueue) ProtoMessage() {}

func (x *SoundQueue) ProtoReflect() protoreflect.Message {
	mi := &file_client_audio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SoundQueue.ProtoReflect.Descriptor instead.
func (*SoundQueue) Descriptor() ([]byte, []int) {
	return file_client_audio_proto_rawDescGZIP(), []int{1}
}

func (x *SoundQueue) GetAudioFile() string {
	if x != nil {
		return x.AudioFile
	}
	return ""
}

func (x *SoundQueue) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SoundQueue) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *SoundQueue) GetAudioSoundEnum() AudioSoundEnum {
	if x != nil {
		return x.AudioSoundEnum
	}
	return AudioSoundEnum_AUDIO_SOUND_NONE
}

// ClientAudio
// Next: 7
type ClientAudio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SampleRate      int32           `protobuf:"varint,1,opt,name=sample_rate,json=sampleRate,proto3" json:"sample_rate,omitempty"`
	AudioFormatEnum AudioFormatEnum `protobuf:"varint,2,opt,name=audio_format_enum,json=audioFormatEnum,proto3,enum=proto.AudioFormatEnum" json:"audio_format_enum,omitempty"`
	Channels        int32           `protobuf:"varint,3,opt,name=channels,proto3" json:"channels,omitempty"`
	ChunkSize       int32           `protobuf:"varint,4,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty"`
	FadeMusic       float64         `protobuf:"fixed64,5,opt,name=fade_music,json=fadeMusic,proto3" json:"fade_music,omitempty"`
	MusicQueues     []*MusicQueue   `protobuf:"bytes,6,rep,name=music_queues,json=musicQueues,proto3" json:"music_queues,omitempty"`
	SoundQueues     []*SoundQueue   `protobuf:"bytes,7,rep,name=sound_queues,json=soundQueues,proto3" json:"sound_queues,omitempty"`
}

func (x *ClientAudio) Reset() {
	*x = ClientAudio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_audio_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientAudio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientAudio) ProtoMessage() {}

func (x *ClientAudio) ProtoReflect() protoreflect.Message {
	mi := &file_client_audio_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientAudio.ProtoReflect.Descriptor instead.
func (*ClientAudio) Descriptor() ([]byte, []int) {
	return file_client_audio_proto_rawDescGZIP(), []int{2}
}

func (x *ClientAudio) GetSampleRate() int32 {
	if x != nil {
		return x.SampleRate
	}
	return 0
}

func (x *ClientAudio) GetAudioFormatEnum() AudioFormatEnum {
	if x != nil {
		return x.AudioFormatEnum
	}
	return AudioFormatEnum_AUDIO_FORMAT_NONE
}

func (x *ClientAudio) GetChannels() int32 {
	if x != nil {
		return x.Channels
	}
	return 0
}

func (x *ClientAudio) GetChunkSize() int32 {
	if x != nil {
		return x.ChunkSize
	}
	return 0
}

func (x *ClientAudio) GetFadeMusic() float64 {
	if x != nil {
		return x.FadeMusic
	}
	return 0
}

func (x *ClientAudio) GetMusicQueues() []*MusicQueue {
	if x != nil {
		return x.MusicQueues
	}
	return nil
}

func (x *ClientAudio) GetSoundQueues() []*SoundQueue {
	if x != nil {
		return x.SoundQueues
	}
	return nil
}

var File_client_audio_proto protoreflect.FileDescriptor

var file_client_audio_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x0a,
	0x4d, 0x75, 0x73, 0x69, 0x63, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75,
	0x64, 0x69, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x75, 0x64, 0x69, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x3f, 0x0a, 0x10, 0x61, 0x75, 0x64, 0x69, 0x6f,
	0x5f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x4d,
	0x75, 0x73, 0x69, 0x63, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x4d,
	0x75, 0x73, 0x69, 0x63, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x9a, 0x01, 0x0a, 0x0a, 0x53, 0x6f, 0x75,
	0x6e, 0x64, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x6f,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x64,
	0x69, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x12, 0x3f, 0x0a, 0x10, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x73, 0x6f,
	0x75, 0x6e, 0x64, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x53, 0x6f, 0x75, 0x6e,
	0x64, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x53, 0x6f, 0x75, 0x6e,
	0x64, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xb8, 0x02, 0x0a, 0x0b, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x41, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f,
	0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x42, 0x0a, 0x11, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0f, 0x61, 0x75, 0x64, 0x69, 0x6f,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x61, 0x64, 0x65, 0x5f, 0x6d, 0x75,
	0x73, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x66, 0x61, 0x64, 0x65, 0x4d,
	0x75, 0x73, 0x69, 0x63, 0x12, 0x34, 0x0a, 0x0c, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x0b, 0x6d,
	0x75, 0x73, 0x69, 0x63, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x0c, 0x73, 0x6f,
	0x75, 0x6e, 0x64, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x52, 0x0b, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73,
	0x2a, 0x95, 0x01, 0x0a, 0x0e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x45,
	0x6e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x5f, 0x4d, 0x55, 0x53,
	0x49, 0x43, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x55, 0x44,
	0x49, 0x4f, 0x5f, 0x4d, 0x55, 0x53, 0x49, 0x43, 0x5f, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x10, 0x01,
	0x12, 0x14, 0x0a, 0x10, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x5f, 0x4d, 0x55, 0x53, 0x49, 0x43, 0x5f,
	0x4d, 0x45, 0x4e, 0x55, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x5f,
	0x4d, 0x55, 0x53, 0x49, 0x43, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11,
	0x41, 0x55, 0x44, 0x49, 0x4f, 0x5f, 0x4d, 0x55, 0x53, 0x49, 0x43, 0x5f, 0x44, 0x45, 0x41, 0x54,
	0x48, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x5f, 0x4d, 0x55, 0x53,
	0x49, 0x43, 0x5f, 0x57, 0x49, 0x4e, 0x10, 0x05, 0x2a, 0x67, 0x0a, 0x0e, 0x41, 0x75, 0x64, 0x69,
	0x6f, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x55,
	0x44, 0x49, 0x4f, 0x5f, 0x53, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x14, 0x0a, 0x10, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x5f, 0x53, 0x4f, 0x55, 0x4e, 0x44, 0x5f,
	0x4a, 0x55, 0x4d, 0x50, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x5f,
	0x53, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x47, 0x4f, 0x4f, 0x44, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f,
	0x41, 0x55, 0x44, 0x49, 0x4f, 0x5f, 0x53, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x42, 0x41, 0x44, 0x10,
	0x03, 0x2a, 0x90, 0x01, 0x0a, 0x0f, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x5f, 0x46,
	0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14,
	0x41, 0x55, 0x44, 0x49, 0x4f, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x44, 0x45, 0x46,
	0x41, 0x55, 0x4c, 0x54, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x5f,
	0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x53, 0x31, 0x36, 0x5f, 0x4c, 0x53, 0x42, 0x10, 0x02,
	0x12, 0x18, 0x0a, 0x14, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54,
	0x5f, 0x53, 0x31, 0x36, 0x5f, 0x4d, 0x53, 0x42, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x55,
	0x44, 0x49, 0x4f, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x53, 0x31, 0x36, 0x5f, 0x53,
	0x59, 0x53, 0x10, 0x04, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x69, 0x72, 0x75, 0x6c, 0x2f, 0x67, 0x6f, 0x5f, 0x64, 0x61, 0x72,
	0x77, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x64, 0x61, 0x72, 0x77, 0x69,
	0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_audio_proto_rawDescOnce sync.Once
	file_client_audio_proto_rawDescData = file_client_audio_proto_rawDesc
)

func file_client_audio_proto_rawDescGZIP() []byte {
	file_client_audio_proto_rawDescOnce.Do(func() {
		file_client_audio_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_audio_proto_rawDescData)
	})
	return file_client_audio_proto_rawDescData
}

var file_client_audio_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_client_audio_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_client_audio_proto_goTypes = []interface{}{
	(AudioMusicEnum)(0),  // 0: proto.AudioMusicEnum
	(AudioSoundEnum)(0),  // 1: proto.AudioSoundEnum
	(AudioFormatEnum)(0), // 2: proto.AudioFormatEnum
	(*MusicQueue)(nil),   // 3: proto.MusicQueue
	(*SoundQueue)(nil),   // 4: proto.SoundQueue
	(*ClientAudio)(nil),  // 5: proto.ClientAudio
}
var file_client_audio_proto_depIdxs = []int32{
	0, // 0: proto.MusicQueue.audio_music_enum:type_name -> proto.AudioMusicEnum
	1, // 1: proto.SoundQueue.audio_sound_enum:type_name -> proto.AudioSoundEnum
	2, // 2: proto.ClientAudio.audio_format_enum:type_name -> proto.AudioFormatEnum
	3, // 3: proto.ClientAudio.music_queues:type_name -> proto.MusicQueue
	4, // 4: proto.ClientAudio.sound_queues:type_name -> proto.SoundQueue
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_client_audio_proto_init() }
func file_client_audio_proto_init() {
	if File_client_audio_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_audio_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MusicQueue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_client_audio_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SoundQueue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_client_audio_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientAudio); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_client_audio_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_audio_proto_goTypes,
		DependencyIndexes: file_client_audio_proto_depIdxs,
		EnumInfos:         file_client_audio_proto_enumTypes,
		MessageInfos:      file_client_audio_proto_msgTypes,
	}.Build()
	File_client_audio_proto = out.File
	file_client_audio_proto_rawDesc = nil
	file_client_audio_proto_goTypes = nil
	file_client_audio_proto_depIdxs = nil
}
