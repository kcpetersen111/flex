package persist

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect(user, password, address string) (*sql.DB, error) {
	//eventually we will need to figure out how to not disable ssl
	connstr := fmt.Sprintf("postgres://%v:%v@%v/flexDB?sslmode=disable", user, password, address)
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		return nil, err
	}
	Setup(db)
	return db, nil
}
