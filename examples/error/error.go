package main

import (
	"fmt"
	"log"

	"github.com/suziziziz/died"
)

func main() {
	died := died.New()

	// As you can see, it is possible to pass `app` before the required dependencies.
	died.Inject(app)
	died.Inject(dep4)
	// died.Inject(dep2)
	died.Inject(dep1)
	died.Inject(dep3)

	// After providing all dependencies, run `Invoke`.
	died.Invoke()

	// You can require a dependency from the context after invoke
	dep := died.Require(func(dep *Dep1) {}).(*Dep1)
	log.Println((*dep)["what"])
}

type Dep1 map[string]string
type Dep2 map[string]string
type Dep3 map[string]string
type Dep4 string

func dep1() *Dep1 {
	return &Dep1{"what": "DIED!"}
}

func dep2() *Dep2 {
	return &Dep2{"is": "IS"}
}

func dep3() *Dep3 {
	return &Dep3{"this": "THIS"}
}

func dep4(dep1 *Dep1, dep2 *Dep2, dep3 *Dep3) *Dep4 {
	t := Dep4(fmt.Sprintf("%s %s %s", (*dep3)["this"], (*dep2)["is"], (*dep1)["what"]))

	return &t
}

func app(dep3 *Dep3, dep4 *Dep4) {
	log.Println(*dep4)
}
