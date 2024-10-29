package sfgarcio

import "github.com/gnames/coldp/ent/coldp"

func (s *sfgarcio) InsertSynonyms(data []coldp.Synonym) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(`
	INSERT INTO synonym
	(
		id, taxon_id, source_id, name_id, name_phrase,
		according_to_id, status_id, reference_id,
		link, remarks, modified, modified_by
	)
	VALUES (?,?,?,?,?, ?,?,?, ?,?,?,?)
`)
	if err != nil {
		return err
	}

	for _, d := range data {
		_, err = stmt.Exec(
			d.ID, d.TaxonID, d.SourceID, d.NameID, d.NamePhrase,
			d.AccordingToID, d.Status.String(), d.ReferenceID,
			d.Link, d.Remarks, d.Modified, d.ModifiedBy,
		)
		if err != nil {
			return err
		}
	}

	stmt.Close()
	return tx.Commit()
}
