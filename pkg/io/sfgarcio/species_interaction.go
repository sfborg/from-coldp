package sfgarcio

import (
	"log/slog"

	"github.com/gnames/coldp/ent/coldp"
)

func (s *sfgarcio) InsertSpeciesInteractions(
	data []coldp.SpeciesInteraction,
) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			slog.Error("Cannot finish transaction", "error", err)
			tx.Rollback()
		}
	}()

	stmt, err := tx.Prepare(`
	INSERT INTO species_interaction
		(
			taxon_id, related_taxon_id, source_id, related_taxon_scientific_name,
			type, reference_id, remarks, modified, modified_by
		)
	VALUES (?,?,?,?, ?,?,?,?,?)
`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, n := range data {
		_, err = stmt.Exec(
			n.TaxonID, n.RelatedTaxonID, n.SourceID, n.RelatedTaxonScientificName,
			n.Type.ID(), n.ReferenceID, n.Remarks, n.Modified, n.ModifiedBy,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
