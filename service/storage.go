package service

type Storage interface {

	// get nacos cluster server list
	get() []string

	// add nacos cluster server list
	add([]string)

	// delete nacos cluster server list
	delete([]string)

	// delete all data
	deleteAll()
}
