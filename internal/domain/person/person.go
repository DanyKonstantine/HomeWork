package person

type Person struct {
	ID       int    `db:"id,omitempty"`
	Name     string `db:"name"`
	LastName string `db:"lastName"`
	Age      int    `db:"age"`
	EventId  int    `db:"eventID"`
}
