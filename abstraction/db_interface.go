package abstraction

type DbContract interface {
	SelectUser(userName string) (string, error)
	InsertUser(userName string) error
	Close()
}

type Application struct {
	db DbContract
}

func NewApplication(db DbContract) *Application {
	return &Application{db: db}
}
