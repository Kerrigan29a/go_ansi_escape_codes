package ansi_escape_codes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCSISequenceString(t *testing.T) {
	assert.Equal(t, "34;1;41;1m", FgBlue.Compose(Bright).Compose(BgRed).Compose(Bright).ParamsString())
}

func TestCSISequenceCompose(t *testing.T) {
	assert.Equal(t, CSISequence{"30", ";", "1", ";", "4", "m"}, FgBlack.Compose(Bright).Compose(Underline))
	assert.Equal(t, CSISequence{"31", ";", "1", ";", "4", "m"}, FgRed.Compose(Bright).Compose(Underline))
	assert.Equal(t, CSISequence{"32", ";", "1", ";", "4", "m"}, FgGreen.Compose(Bright).Compose(Underline))
	assert.Equal(t, CSISequence{"33", ";", "1", ";", "4", "m"}, FgYellow.Compose(Bright).Compose(Underline))
	assert.Equal(t, CSISequence{"34", ";", "1", ";", "4", "m"}, FgBlue.Compose(Bright).Compose(Underline))
	assert.Equal(t, CSISequence{"35", ";", "1", ";", "4", "m"}, FgMagenta.Compose(Bright).Compose(Underline))
	assert.Equal(t, CSISequence{"36", ";", "1", ";", "4", "m"}, FgCyan.Compose(Bright).Compose(Underline))
	assert.Equal(t, CSISequence{"37", ";", "1", ";", "4", "m"}, FgWhite.Compose(Bright).Compose(Underline))
}
