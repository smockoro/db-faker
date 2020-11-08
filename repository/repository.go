package repository

// FakerRepository : Hidden Database system.
type FakerRepository interface {
	Ping() error
}
