package sfgarcio

import "github.com/gnames/coldp/ent/coldp"

func (s *sfgarcio) InsertNameUsages(data []coldp.NameUsage) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(`
	INSERT INTO name
		(
		)
	VALUES (?,?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?,?, ?,?,?)
`)
	if err != nil {
		return err
	}

	for _, n := range data {
		_ = n
		_, err = stmt.Exec()
		if err != nil {
			return err
		}
	}

	stmt.Close()
	return tx.Commit()
}
