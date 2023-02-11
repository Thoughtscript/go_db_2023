package domain

import (
	"log"
	"goserver/database"
)

// Uses int type for id since not a production app.
// Capitalize these to make them visible to JSON encoder when serializing.
// Enforce json lowercase capitalization.
// https://stackoverflow.com/questions/8270816/converting-go-struct-to-json
type Example struct {
	ID		int		`json:"id"`
	Name	any		`json:"name"`
	Val		any		`json:"val"`
}

func GetExamples() map[int]Example {
	sql := database.GetInstance()
	rows, err := sql.Query("SELECT * FROM examples")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	var results = make(map[int]Example)
	for rows.Next() {
		var id int
		var name, val any

		err = rows.Scan(&id, &name, &val)
		if err != nil {
			log.Fatal(err)
		}

		results[id] = Example{
			ID:          id,
			Name:    	 name,
			Val:       	 val,
		}

		// log.Println(results[id])
	}

	return results
}
