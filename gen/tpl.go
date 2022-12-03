package gen

type DayTempl struct {
	Day  string
	Year int
}

func DayTemplate() []byte {
	return []byte(`
package {{ .Day }}

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
package {{ .Day }}_test

import (
	"testing"

	"github.com/jonavdm/aoc-{{ .Year }}/{{ .Day }}"
	_ "github.com/jonavdm/aoc-{{ .Year }}/testing"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	assert.Equal(t, [2]interface{}{0, 0}, {{ .Day }}.Run())
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		{{ .Day }}.Run()
	}
}

	`)
}
