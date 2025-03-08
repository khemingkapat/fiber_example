package queries

import (
	"database/sql"
	"github.com/khemingkapat/fiber_example/common"
)

var khem common.Person

func QueryPeople(db *sql.DB) ([]common.Person, error) {
	rows, err := db.Query("select * from people")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var people []common.Person

	for rows.Next() {
		var p common.Person
		err := rows.Scan(&p.ID, &p.Firstname, &p.Lastname, &p.Age, &p.Role)
		if err != nil {
			return nil, err
		}
		people = append(people, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return people, nil
}
