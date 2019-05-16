package drafter_test

import (
	"strings"
	"testing"

	"github.com/kppk/snowboard/adapter/drafter"
	"github.com/stretchr/testify/assert"
)

func TestDrafter_Parse(t *testing.T) {
	s := strings.NewReader("# API")
	b, err := drafter.Parse(s)
	assert.Nil(t, err)
	assert.Contains(t, string(b), "API")
}

func TestDrafter_Validate(t *testing.T) {
	s := strings.NewReader("# API")
	b, err := drafter.Validate(s)
	assert.Nil(t, err)
	assert.Empty(t, string(b))

	s = strings.NewReader("# API\n## Data Structures\n### Hello-World (object)\n+ foo: bar (string, required)")
	b, err = drafter.Validate(s)
	assert.Nil(t, err)
	assert.Contains(t, string(b), "please escape the name of the data structure using backticks")
}

func TestDrafter_Version(t *testing.T) {
	v := drafter.Version()
	assert.Equal(t, "v4.0.0-pre.4", v)
}
