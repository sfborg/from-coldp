package sfgarcio

import (
	"github.com/gnames/coldp/ent/coldp"
)

func (s *sfgarcio) InsertAuthors(data []coldp.Author) error {
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
	INSERT INTO author
		(
			id, source_id, alternative_id, given, family, suffix,
	  	abbreviation_botany, alternative_names, sex_id, country, birth,
	  	birth_place, death, affiliation, interest, reference_id, link,
	  	remarks, modified, modified_by
		)
	VALUES (?,?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?,?, ?,?,?)
`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, n := range data {
		_, err = stmt.Exec(
			n.ID, n.SourceID, n.AlternativeID, n.Given, n.Family, n.Suffix,
			n.AbbreviationBotany, n.AlternativeNames, n.Sex, n.Country,
			n.BirthPlace, n.Death, n.Affiliation, n.Interest, n.ReferenceID,
			n.Link, n.Remarks, n.Modified, n.ModifiedBy,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
