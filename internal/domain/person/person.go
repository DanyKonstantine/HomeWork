package person

type Person struct {
	ID       int    `db:"ID,omitempty"`
	Name     string `db:"Name"`
	Sername  string `db:"Sername"`
	Age      int    `db:"Age"`
	Event_id int    `db:"Event_ID"`
}
