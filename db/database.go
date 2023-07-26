package db

// DB is the interface that wraps the basic methods of a database.
type DB interface {
	Close()
}
