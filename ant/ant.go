package ant

import "github.com/panjf2000/ants/v2"

func demo() {
	pool, _ := ants.NewPool(100)
	pool.Submit()
}
