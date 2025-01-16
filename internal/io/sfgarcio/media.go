package sfgarcio

import "github.com/gnames/coldp/ent/coldp"

func (s *sfgarcio) InsertMedia(data []coldp.Media) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	stmt, err := tx.Prepare(`
	INSERT INTO media
		(
			taxon_id, source_id, url, type, format, title, created,
			creator, license, link, remarks, modified, modified_by
		)
	VALUES (?,?,?,?,?,?,?, ?,?,?,?,?,?)
`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, n := range data {
		_, err = stmt.Exec(
			n.TaxonID, n.SourceID, n.URL, n.Type, n.Format, n.Title, n.Created,
			n.Creator, n.License, n.Link, n.Remarks, n.Modified, n.ModifiedBy,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
