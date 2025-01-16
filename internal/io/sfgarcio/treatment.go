package sfgarcio

import (
	"log/slog"

	"github.com/gnames/coldp/ent/coldp"
)

func (s *sfgarcio) InsertTreatments(data []coldp.Treatment) error {
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
	INSERT INTO treatment
		(
		taxon_id, source_id, document, format, modified, modified_by
		)
	VALUES (?,?,?,?,?,?)
`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, n := range data {
		_, err = stmt.Exec(
			n.TaxonID, n.SourceID, n.Document, n.Format, n.Modified, n.ModifiedBy,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
