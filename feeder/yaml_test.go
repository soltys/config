package feeder_test

import (
	"github.com/soltys/config/feeder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Yaml_Feed_Not_Existing_File_It_Should_Return_Error(t *testing.T) {
	j := feeder.Yaml{Path: "/404.yml"}

	_, err := j.Feed()

	assert.Error(t, err)
}

func Test_Yaml_Feed_Sample1(t *testing.T) {
	j := feeder.Yaml{Path: "test/yaml/config.yml"}

	m, err := j.Feed()

	assert.NoError(t, err)
	assert.Equal(t, "MyAppUsingConfig", m["name"])
	assert.Equal(t, 3.14, m["version"])

	userMap, _ := m["user"].(map[string]interface{})
	addressMap, _ := userMap["address"].(map[string]interface{})
	assert.Equal(t, "Iran", addressMap["country"])
}
