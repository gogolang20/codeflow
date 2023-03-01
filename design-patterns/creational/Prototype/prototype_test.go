package Prototype

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcretePrototype_Clone(t *testing.T) {
	name := "出去浪"
	proto := ConcretePrototype{name: name}
	newProto := proto.Clone()
	actualName := newProto.Name()
	fmt.Println(&proto.name,&actualName)

	assert.Equal(t, name, actualName)
}
