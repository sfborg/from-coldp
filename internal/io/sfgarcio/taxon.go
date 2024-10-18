package sfgarcio

import "github.com/gnames/coldp/ent/coldp"

func (s *sfgarcio) InsertTaxa(data []coldp.Taxon) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(`
	INSERT INTO taxon
		(
   id, name_id, parent_id, according_to_id, scrutinizer,
   scrutinizer_id, scrutinizer_date, reference_id, link
		)
	VALUES (?,?,?,?,?, ?,?,?,?)
`)
	if err != nil {
		return err
	}

	for _, t := range data {
		_, err = stmt.Exec(
			t.ID, t.NameID, t.ParentID, t.AccordingToID, t.Scrutinizer,
			t.ScrutinizerID, t.ScrutinizerDate, t.ReferenceID, t.Link,
		)
		if err != nil {
			return err
		}
	}

	stmt.Close()
	return tx.Commit()
}
