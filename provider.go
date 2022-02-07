package fluff

type provider interface {
	detect() bool
	create(vm instance) error
	destroy(vm instance) error
	start(vm instance) error
	stop(vm instance) error
	snapshot(vm instance) error
}
