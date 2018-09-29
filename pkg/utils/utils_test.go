package utils_test

import (
	"os"
	"testing"

	. "github.com/paradigm-network/paradigm-fn2/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestEnsureFile(t *testing.T) {
	fullPath := "/tmp/a/b/c.json"
	err := EnsureFile(fullPath)
	assert.Nil(t, err)

	_, err = os.Stat(fullPath)
	assert.Nil(t, err)

	os.Remove(fullPath)
}

func TestPairsToParams(t *testing.T) {
	pairs := []string{
		"a=1",
		"b=2",
	}
	params := PairsToParams(pairs)
	assert.Equal(t, map[string]string{"a": "1", "b": "2"}, params)
}

func TestOutputJSON(t *testing.T) {
	obj := map[string]string{
		"name": "minghe",
	}
	err := OutputJSON(obj)
	assert.Nil(t, err)
}
