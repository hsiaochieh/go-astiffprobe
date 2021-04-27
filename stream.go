package astiffprobe

import (
	"context"
	"fmt"
)

// Channel layouts
const (
	ChannelLayout51     = "5.1"
	ChannelLayout51Side = "5.1(side)"
	ChannelLayoutMono   = "mono"
	ChannelLayoutStereo = "stereo"
)

// Codec names
const (
	CodecNameAC3         = "ac3"
	CodecNameDVBSub      = "dvbsub"
	CodecNameDVBTeletext = "dvb_teletext"
	CodecNameH264        = "h264"
)

// Codec types
const (
	CodecTypeAudio    = "audio"
	CodecTypeSubtitle = "subtitle"
	CodecTypeVideo    = "video"
)

// Pixel formats
const (
	PixelFormatYuv420p = "yuv420p"
)

// Profiles
const (
	ProfileBaseline = "baseline"
	ProfileHight    = "high"
	ProfileMain     = "main"
)

// Stream represents a stream
type Stream struct {
	AvgFramerate       Rational          `json:"avg_frame_rate"`
	Bitrate            int64             `json:"bit_rate,string"`
	BitsPerRawSample   int64             `json:"bits_per_raw_sample,string"`
	ChannelLayout      string            `json:"channel_layout"`
	Channels           int               `json:"channels"`
	ChromaLocation     string            `json:"chroma_location"`
	CodecLongName      string            `json:"codec_long_name"`
	CodecName          string            `json:"codec_name"`
	CodecTag           string            `json:"codec_tag"`
	CodecTagString     string            `json:"codec_tag_string"`
	CodecTimeBase      string            `json:"codec_time_base"`
	CodecType          string            `json:"codec_type"`
	CodedHeight        int               `json:"coded_height"`
	CodedWidth         int               `json:"coded_width"`
	ColorPrimaries     string            `json:"color_primaries"`
	ColorRange         string            `json:"color_range"`
	ColorSpace         string            `json:"color_space"`
	ColorTransfer      string            `json:"color_transfer"`
	DisplayAspectRatio Rational          `json:"display_aspect_ratio"`
	Disposition        Disposition       `json:"disposition"`
	Duration           Duration          `json:"duration"`
	DurationTs         int64             `json:"duration_ts"`
	HasBFrames         int               `json:"has_b_frames"`
	Height             int               `json:"height"`
	ID                 Hexadecimal       `json:"id"`
	Index              int               `json:"index"`
	IsAVC              bool              `json:"is_avc,string"`
	Level              int               `json:"level"`
	NalLengthSize      int               `json:"nal_length_size,string"`
	NbFrames           int64             `json:"nb_frames,string"`
	PixFmt             string            `json:"pix_fmt"`
	Profile            string            `json:"profile"`
	Refs               int64             `json:"refs"`
	RFrameRate         Rational          `json:"r_frame_rate"`
	SampleAspectRatio  Rational          `json:"sample_aspect_ratio"`
	SampleFmt          string            `json:"sample_fmt"`
	SampleRate         int               `json:"sample_rate,string"`
	StartPts           int64             `json:"start_pts"`
	StartTime          Duration          `json:"start_time"`
	Tags               map[string]string `json:"tags"`
	TimeBase           Rational          `json:"time_base"`
	Width              int               `json:"width"`
}

// Disposition represents a stream disposition
type Disposition struct {
	AttachedPic     Bool `json:"attached_pic"`
	CleanEffects    Bool `json:"clean_effects"`
	Comment         Bool `json:"comment"`
	Default         Bool `json:"default"`
	Dub             Bool `json:"dub"`
	Forced          Bool `json:"forced"`
	HearingImpaired Bool `json:"hearing_impaired"`
	Karaoke         Bool `json:"karaoke"`
	Lyrics          Bool `json:"lyrics"`
	Original        Bool `json:"original"`
	TimedThumbnails Bool `json:"timed_thumbnails"`
	VisualImpaired  Bool `json:"visual_impaired"`
}

// Streams returns the streams of a video
func (f *FFProbe) Streams(ctx context.Context, src string) (ss []Stream, err error) {
	// Execute
	var o Output
	if o, err = f.exec(ctx, f.binaryPath, "-loglevel", "error", "-show_streams", "-print_format", "json", src); err != nil {
		err = fmt.Errorf("astiffprobe: executing failed: %w", err)
		return
	}
	return o.Streams, nil
}
