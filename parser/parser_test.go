package parser_test

import (
	"strings"
	"testing"

	snowboard "github.com/bukalapak/snowboard/parser"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	s := strings.NewReader("# API")

	api, err := snowboard.Parse(s)
	assert.Nil(t, err)
	assert.Equal(t, "API", api.Title)
}

func TestParseAsJSON(t *testing.T) {
	s := strings.NewReader("# API")

	b, err := snowboard.ParseAsJSON(s)
	assert.Nil(t, err)
	assert.Contains(t, string(b), `"content": "API"`)
}

func TestLoad(t *testing.T) {
	api, err := snowboard.Load("../adapter/drafter/ext/drafter/features/fixtures/blueprint.apib")
	assert.Nil(t, err)
	assert.Equal(t, "<API name>", api.Title)
	assert.Equal(t, "<resource group name>", api.ResourceGroups[0].Title)
	assert.Equal(t, "<resource name>", api.ResourceGroups[0].Resources[0].Title)
	assert.Equal(t, "<action name>", api.ResourceGroups[0].Resources[0].Transitions[0].Title)
	assert.Equal(t, "<request name>", api.ResourceGroups[0].Resources[0].Transitions[0].Transactions[0].Request.Title)
	assert.Equal(t, "<request description>", api.ResourceGroups[0].Resources[0].Transitions[0].Transactions[0].Request.Description)
	assert.Equal(t, 200, api.ResourceGroups[0].Resources[0].Transitions[0].Transactions[0].Response.StatusCode)
	assert.Equal(t, "<response description>", api.ResourceGroups[0].Resources[0].Transitions[0].Transactions[0].Response.Description)

	api, err = snowboard.Load("../fixtures/api-blueprint/examples/10. Data Structures.md")
	assert.Nil(t, err)
	assert.False(t, api.ResourceGroups[0].Resources[1].Transitions[0].Href.Parameters[0].Required)
	assert.Equal(t, "limit", api.ResourceGroups[0].Resources[1].Transitions[0].Href.Parameters[0].Key)
	assert.Equal(t, "number", api.ResourceGroups[0].Resources[1].Transitions[0].Href.Parameters[0].Kind)
	assert.Equal(t, "10", api.ResourceGroups[0].Resources[1].Transitions[0].Href.Parameters[0].Default)

	api, err = snowboard.Load("../fixtures/examples/enum.apib")
	assert.Nil(t, err)
	assert.True(t, api.Resources[0].Transitions[0].Href.Parameters[0].Required)
	assert.Equal(t, "type", api.Resources[0].Transitions[0].Href.Parameters[0].Key)
	assert.Equal(t, "enum[string]", api.Resources[0].Transitions[0].Href.Parameters[0].Kind)
	assert.Equal(t, "foo", api.Resources[0].Transitions[0].Href.Parameters[0].Value)
	assert.Equal(t, []string{"foo", "bar", "baz"}, api.Resources[0].Transitions[0].Href.Parameters[0].Members)
}

func TestLoad_partials(t *testing.T) {
	api, err := snowboard.Load("../fixtures/partials/API.apib")
	assert.Nil(t, err)
	assert.Equal(t, "API", api.Title)
	assert.Equal(t, "Messages", api.ResourceGroups[0].Title)
	assert.Equal(t, "Users", api.ResourceGroups[1].Title)
	assert.Equal(t, "Tasks", api.ResourceGroups[2].Title)
}

func TestLoadAsJSON(t *testing.T) {
	b, err := snowboard.LoadAsJSON("../adapter/drafter/ext/drafter/features/fixtures/blueprint.apib")
	assert.Nil(t, err)
	assert.Contains(t, string(b), `"content": "<API name>"`)
}
