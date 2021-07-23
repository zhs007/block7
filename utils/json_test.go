package block7utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetJsonString(t *testing.T) {
	s, err := GetJsonString([]byte(`{"abc":"123"}`), "abc")
	assert.NoError(t, err)
	assert.Equal(t, s, "123")

	s, err = GetJsonString([]byte(`{"abc":123}`), "abc")
	assert.NoError(t, err)
	assert.Equal(t, s, "123")

	s, err = GetJsonString([]byte(`{"abc":123.456}`), "abc")
	assert.NoError(t, err)
	assert.Equal(t, s, "123.456")

	t.Logf("Test_GetJsonString OK")
}

func Test_GetJsonInt(t *testing.T) {
	i64, err := GetJsonInt([]byte(`{"abc":"123"}`), "abc")
	assert.NoError(t, err)
	assert.Equal(t, i64, int64(123))

	i64, err = GetJsonInt([]byte(`{"abc":123}`), "abc")
	assert.NoError(t, err)
	assert.Equal(t, i64, int64(123))

	i64, err = GetJsonInt([]byte(`{"abc":123.456}`), "abc")
	assert.NoError(t, err)
	assert.Equal(t, i64, int64(123))

	i64, err = GetJsonInt([]byte(`{"abc":"123.456"}`), "abc")
	assert.NoError(t, err)
	assert.Equal(t, i64, int64(123))

	t.Logf("Test_GetJsonInt OK")
}
