package sfgarcio

import "github.com/gnames/coldp/ent/coldp"

func (s *sfgarcio) InsertSpeciesEstimates(data []coldp.SpeciesEstimate) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(`
	INSERT INTO species_estimate
		(
		taxon_id, source_id, estimate, type, reference_id,
		remarks, modified, modified_by
		)
	VALUES (?,?,?,?,?, ?,?,?)
`)
	if err != nil {
		return err
	}

	for _, n := range data {
		_, err = stmt.Exec(
			n.TaxonID, n.SourceID, n.Estimate, n.Type.String(), n.ReferenceID,
			n.Remarks, n.Modified, n.ModifiedBy,
		)
		if err != nil {
			return err
		}
	}

	stmt.Close()
	return tx.Commit()
}
