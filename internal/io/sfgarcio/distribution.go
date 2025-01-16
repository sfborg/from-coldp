package sfgarcio

import "github.com/gnames/coldp/ent/coldp"

func (s *sfgarcio) InsertDistributions(data []coldp.Distribution) error {
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
	INSERT INTO distribution
		(
			taxon_id, source_id, area, area_id, gazetteer_id,
			status_id, reference_id, remarks, modified, modified_by
		)
	VALUES (?,?,?,?,?, ?,?,?,?,?)
`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, n := range data {
		_, err = stmt.Exec(
			n.TaxonID, n.SourceID, n.Area, n.AreaID, n.Gazetteer.String(),
			n.Status.String(), n.ReferenceID, n.Remarks, n.Modified, n.ModifiedBy,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
