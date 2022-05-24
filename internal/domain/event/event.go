package event

type Event struct {
	ID        int     `db:"ID,omitempty"`
	Title     string  `db:"Title"`
	ShortDesc string  `db:"ShortDesc"'`
	Desc      string  `db:"Desc"`
	Long      float64 `db:"Long"`
	Let       float64 `db:"Let"`
	Images    string  `db:"Images"`
	Prewive   string  `db:"Prewive"`
}
