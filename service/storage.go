package service

type Storage interface {

	// get nacos cluster server list
	get() ([]string, error)

	// add nacos cluster server list
	add([]string) error

	// delete nacos cluster server list
	delete([]string) error

	// delete all data
	deleteAll() error
}
