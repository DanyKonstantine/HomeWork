package event

type Event struct {
	ID        int64   `db:"id"`
	Title     string  `db:"title"`
	ShortDesc string  `db:"shortDesc"'`
	Desc      string  `db:"desc"`
	Long      float64 `db:"long"`
	Let       float64 `db:"let"`
	Images    string  `db:"images"`
	Preview   string  `db:"preview"`
}
