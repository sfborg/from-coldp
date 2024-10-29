package sfgarcio

import "github.com/gnames/coldp/ent/coldp"

func (s *sfgarcio) InsertTaxonConceptRelations(
	data []coldp.TaxonConceptRelation,
) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(`
	INSERT INTO taxon_concept_relation
		(
			taxon_id, related_taxon_id, source_id, type, reference_id,
			remarks, modified, modified_by
		)
	VALUES (?,?,?,?,?, ?,?,?)
`)
	if err != nil {
		return err
	}

	for _, n := range data {
		_, err = stmt.Exec(
			n.TaxonID, n.RelatedTaxonID, n.SourceID, n.Type.String(), n.ReferenceID,
			n.Remarks, n.Modified, n.ModifiedBy,
		)
		if err != nil {
			return err
		}
	}

	stmt.Close()
	return tx.Commit()
}
