package died_test

import (
	"testing"

	"github.com/suziziziz/died"
)

func TestExample_runs_ok(t *testing.T) {
	type Dep1 map[string]string
	type Dep2 map[string]string
	type Dep3 map[string]string

	died := died.New()

	died.Inject(func() *Dep1 { return &Dep1{"what": "DIED!"} })
	died.Inject(func() *Dep2 { return &Dep2{"is": "IS"} })
	died.Inject(func() *Dep3 { return &Dep3{"this": "THIS"} })
	died.Inject(func(dep1 *Dep1, dep2 *Dep2, dep3 *Dep3) {
		t.Logf("%s %s %s", (*dep3)["this"], (*dep2)["is"], (*dep1)["what"])
	})

	// TODO: MAKE THIS TEST WORK

	died.Invoke()
}
