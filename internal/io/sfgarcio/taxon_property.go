package sfgarcio

import "github.com/gnames/coldp/ent/coldp"

func (s *sfgarcio) InsertTaxonProperties(
	data []coldp.TaxonProperty,
) error {
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
	INSERT INTO taxon_property
		(
		taxon_id, source_id, property, value, reference_id, page,
		ordinal, remarks, modified, modified_by
		)
	VALUES (?,?,?,?,?,?, ?,?,?,?)
`)
	if err != nil {
		return err
	}
	stmt.Close()

	for _, n := range data {
		_, err = stmt.Exec(
			n.TaxonID, n.SourceID, n.Property, n.Value, n.ReferenceID, n.Page,
			n.Ordinal, n.Remarks, n.Modified, n.ModifiedBy,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
