package conf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigLoader_computeFileName(t *testing.T) {
	assert.Equal(t, "/usr/bin/config.json", computeFileName("/usr/bin/someExe", "config.json"))
	assert.Equal(t, "config.json", computeFileName("./someExe", "config.json"))
	assert.Equal(t, "../somedir/config.json", computeFileName("../somedir/someExe", "config.json"))
}

func TestConfigLoader_LoadInto(t *testing.T) {
	type Config struct {
		IntParam    int
		StringParam string
	}
	var config Config
	loader := NewJSONConfig("")
	cliPrintConfig = true
	loader.LoadInto(config)
}
