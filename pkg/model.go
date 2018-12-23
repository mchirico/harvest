package pkg

import (
	"database/sql"
	"encoding/json"
)

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func getProducts(db *sql.DB, start, count int) ([]product, error) {
	return []product{}, nil
}

func getRoot(db *sql.DB, start, count int) ([]response2, error) {

	var r []response2

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)

	r = append(r, res)

	str = `{"page": 2, "fruits": ["pear", "orange"]}`
	json.Unmarshal([]byte(str), &res)

	r = append(r, res)

	return r, nil
}
