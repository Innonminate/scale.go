package source_test

import (
	"github.com/Innonminate/scale.go/source"
	"testing"
)

func TestLoadTypeRegistry(t *testing.T) {
	source.LoadTypeRegistry([]byte(source.BaseType))
}
