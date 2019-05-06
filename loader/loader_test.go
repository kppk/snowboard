package loader_test

import (
	"testing"

	"github.com/bukalapak/snowboard/loader"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	b, err := loader.Load("../fixtures/extensions/html-comment.apib")
	assert.Nil(t, err)
	assert.Contains(t, string(b), `{class="ui table"}`)
	assert.Contains(t, string(b), `{id="awesome-table" class="ui small table"}`)
}

func TestLoad_seed(t *testing.T) {
	b, err := loader.Load("../fixtures/seeds/API.apib")
	assert.Nil(t, err)
	assert.Contains(t, string(b), `200`)
	assert.Contains(t, string(b), `seeds usage`)
	assert.Contains(t, string(b), `user-related`)
}

func TestRead_helperFuncs(t *testing.T) {
	b, err := loader.Load("../fixtures/extensions/helper-funcs.apib")
	assert.Nil(t, err)
	assert.Contains(t, string(b), `"type": "object",`)
	assert.Contains(t, string(b), `            {`) // indented by 12 spaces
}
