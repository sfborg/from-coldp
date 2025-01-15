package sfgarcio

import (
	"github.com/gnames/coldp/ent/coldp"
)

func (s *sfgarcio) InsertVernaculars(data []coldp.Vernacular) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
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

	for _, d := range data {
		var pref int
		if d.Preferred {
			pref = 1
		}
		_, err = stmt.Exec(
			d.TaxonID, d.SourceID, d.Name, d.Transliteration, d.Language, pref,
			d.Country, d.Area, d.Sex, d.ReferenceID, d.Remarks, d.Modified,
			d.ModifiedBy,
		)
		if err != nil {
			return err
		}
	}

	stmt.Close()
	return tx.Commit()
}
