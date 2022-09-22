package model

import (
	"github.com/burakkuru5534/src/helper"
)

type Movie struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Typ         string `json:"typ"`
}

func (m *Movie) Create() error {

	sq := "INSERT INTO movie (name, description, typ) VALUES ($1, $2, $3) RETURNING id"
	err := helper.App.DB.QueryRow(sq, m.Name, m.Description, m.Typ).Scan(&m.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *Movie) Update(id int64) error {

	sq := "UPDATE movie SET name = $1, typ = $2, description = $3 WHERE id = $4"
	_, err := helper.App.DB.Exec(sq, m.Name, m.Typ, m.Description, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *Movie) Delete(id int64) error {

	sq := "DELETE FROM movie WHERE id = $1"
	_, err := helper.App.DB.Exec(sq, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *Movie) Get(id int64) error {

	sq := "SELECT id, name, typ, description FROM movie WHERE id = $1"
	err := helper.App.DB.QueryRow(sq, id).Scan(&m.ID, &m.Name, &m.Typ, &m.Description)
	if err != nil {
		return err
	}
	return nil
}

func (m *Movie) GetAll() ([]Movie, error) {

	rows, err := helper.App.DB.Query("SELECT id,name,description FROM movie")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// A movie slice to hold data from returned rows.
	var movies []Movie

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var movie Movie
		if err := rows.Scan(&movie.ID, &movie.Name, &movie.Description); err != nil {
			return movies, err
		}
		movies = append(movies, movie)
	}
	if err = rows.Err(); err != nil {
		return movies, err
	}
	return movies, nil
}
