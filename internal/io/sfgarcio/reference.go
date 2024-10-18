package sfgarcio

import "github.com/gnames/coldp/ent/coldp"

func (s *sfgarcio) InsertReferences(data []coldp.Reference) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(`
	INSERT INTO reference
		(
   id, citation, link, doi, remarks
		)
	VALUES (?,?,?,?,?)
`)
	if err != nil {
		return err
	}

	for _, d := range data {
		_, err = stmt.Exec(
			d.ID, d.Citation, d.Link, d.DOI, d.Remarks,
		)
		if err != nil {
			return err
		}
	}

	stmt.Close()
	return tx.Commit()
}
