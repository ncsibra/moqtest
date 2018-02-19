package storage

type DB interface {
	Get() Row
}

type Row interface {
	String()
}
