package spice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisplay(t *testing.T) {
	expected := "-spice unix=on,addr=D51A7383-9D70-4B18-B34F-B7AB5DBD9971.spice,disable-ticketing=on,image-compression=off,playback-compression=off,streaming-video=off,gl=off"

	result := Display(
		NewDisplayProperty("unix", true),
		WithIPAddress("D51A7383-9D70-4B18-B34F-B7AB5DBD9971.spice"),
		IsTicketingDisabled(true),
		WithImageCompressionType(ImageCompressionOff),
		IsAudioStreamCompression(false),
		WithVideoStreamDetection(VideoStreamDetectionOff),
		IsOpenGL(false),
	).ArgsString()

	assert.Equal(t, result, expected)
}
