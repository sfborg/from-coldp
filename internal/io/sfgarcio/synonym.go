package sfgarcio

import "github.com/gnames/coldp/ent/coldp"

func (s *sfgarcio) InsertSynonyms(data []coldp.Synonym) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(`
	INSERT INTO synonym
		(id, taxon_id, name_id, according_to_id, reference_id, link)
	VALUES (?,?,?,?,?,?)
`)
	if err != nil {
		return err
	}

	for _, d := range data {
		_, err = stmt.Exec(
			d.ID, d.TaxonID, d.NameID, d.AccordingToID, d.ReferenceID, d.Link,
		)
		if err != nil {
			return err
		}
	}

	stmt.Close()
	return tx.Commit()
}
