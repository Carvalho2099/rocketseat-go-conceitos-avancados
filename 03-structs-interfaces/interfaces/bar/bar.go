package bar

import "interfaces/foo"

func TakeFoo(i foo.Interface) {
	i.Interface()
}
