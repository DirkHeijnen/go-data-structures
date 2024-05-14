package advent_of_code

import containers "avent_of_code/core/datastructures/base"

type Map interface {
	Put(key interface{}, value interface{})
	Get(key interface{}) (value, found bool)
	Has(key interface{}) bool
	Delete(key interface{})
	Keys() []interface{}

	containers.Container
}
