package feeder_test

import (
	"github.com/soltys/config/feeder"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_JsonDirectory_Feed_Not_Existing_Dir_It_Should_Return_Error(t *testing.T) {
	j := feeder.JsonDirectory{Path: "/404"}

	_, err := j.Feed()

	assert.Error(t, err)
}

func Test_JsonDirectory_Feed_Invalid_JSON_Dir_It_Should_Return_Error(t *testing.T) {
	j := feeder.JsonDirectory{Path: "test/invalid-json"}

	_, err := j.Feed()

	assert.Error(t, err)
}

func Test_JsonDirectory_Feed_Sample1(t *testing.T) {
	j := feeder.JsonDirectory{Path: "test/json"}

	_ = os.Mkdir("test/json/empty", os.ModePerm)

	m, err := j.Feed()

	assert.NoError(t, err)
	assert.Equal(t, m["app"].(map[string]interface{})["name"], "MyAppUsingConfig")
	assert.Equal(t, m["app"].(map[string]interface{})["version"], 3.14)
}
