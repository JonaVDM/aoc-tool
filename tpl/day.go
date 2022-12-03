package tpl

type DayTempl struct {
	Day  int
	Year int
}

func DayTemplate() []byte {
	return []byte(`
package day{{ .Day }}

import (
	"fmt"

	"github.com/jonavdm/aoc-{{ .Year }}/utils"
)

func Run() [2]interface{} {
	data := utils.ReadFile("{{ .Day }}")
	fmt.Println(len(data))

	return [2]interface{}{
		0,
		0,
	}
}

	`)
}

func DayTestTemplate() []byte {
	return []byte(`
package day{{ .Day }}_test

import (
	"testing"

	"github.com/jonavdm/aoc-{{ .Year }}/template"
	_ "github.com/jonavdm/aoc-{{ .Year }}/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{0, 0}, day{{ .Day }}.Run())
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day{{ .Day }}.Run()
	}
}

	`)
}
