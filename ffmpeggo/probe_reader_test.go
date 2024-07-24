package ffmpeg_go

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)
// T
func TestProbeReader(t *testing.T) {
	f, err := os.Open(TestInputFile1)
	assert.Nil(t, err)

	data, err := ProbeReader(f, nil)
	assert.Nil(t, err)
	duration, err := probeOutputDuration(data)
	assert.Nil(t, err)

}