package sfgarcio

import (
	"log/slog"

	"github.com/gnames/coldp/ent/coldp"
)

func (s *sfgarcio) InsertVernaculars(data []coldp.Vernacular) error {
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
	INSERT INTO vernacular
		(
			taxon_id, source_id, name, transliteration, language, preferred,
	   	country, area, sex_id, reference_id, remarks, modified,
	  	modified_by
		)
	VALUES (?,?,?,?,?,?, ?,?,?,?,?,?, ?)
`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, d := range data {
		_, err = stmt.Exec(
			d.TaxonID, d.SourceID, d.Name, d.Transliteration, d.Language, d.Preferred,
			d.Country, d.Area, d.Sex.ID(), d.ReferenceID, d.Remarks, d.Modified,
			d.ModifiedBy,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
