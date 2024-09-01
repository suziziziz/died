package main

import (
	"log"

	"github.com/suziziziz/died"
)

func main() {
	died := died.New()

	died.Inject(dep1)
	died.Inject(dep2)
	died.Inject(dep3)
	died.Inject(app)

	died.Invoke()
}

type Dep1 map[string]string
type Dep2 map[string]string
type Dep3 map[string]string

func dep1() *Dep1 {
	return &Dep1{"what": "DIED!"}
}

func dep2() *Dep2 {
	return &Dep2{"is": "IS"}
}

func dep3() *Dep3 {
	return &Dep3{"this": "THIS"}
}

func app(dep1 *Dep1, dep2 *Dep2, dep3 *Dep3) {
	log.Printf("%s %s %s", (*dep3)["this"], (*dep2)["is"], (*dep1)["what"])
}
