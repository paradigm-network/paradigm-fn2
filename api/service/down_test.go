package service_test

import (
	"testing"

	. "github.com/paradigm-network/paradigm-fn2/api/service"
	"github.com/stretchr/testify/assert"
)

func TestDown(t *testing.T) {
	containerId := "hello-container-id-not-exist"
	image := "world-image-name=not-exit"
	_, err := DoDown(containerId, image)
	assert.Equal(t, RemoveContainerError, err)
}
