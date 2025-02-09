package database


type Db interface {
	useDb(name string, create bool)
	createTable(name string)
	createColumns(name string, table string)
	executeRead(query string)
	executeWrite(query string)
}