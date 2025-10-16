package di

import "go.uber.org/dig"

type DI struct {
	Dig *dig.Container
}

func New() *DI {
	return &DI{
		Dig: dig.New(),
	}
}

func (d *DI) InjectDependencies(database string) {
	switch database {
	case "postgres":
		d.InjectPostgres()
		d.InjectPostgresRepos()
	default:
		panic("Database não suportado")
	}

	// d.injetaGrpc()
	// d.injetaServices()
	d.InjectControllers()
}
