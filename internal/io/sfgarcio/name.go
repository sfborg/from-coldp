package sfgarcio

import "github.com/gnames/coldp/ent/coldp"

func (s *sfgarcio) InsertNames(data []coldp.Name) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(`
	INSERT INTO name
		(
			id, alternative_id, basionym_id, scientific_name, authorship,
			rank, uninomial, genus, infrageneric_epithet, specific_epithet,
			infraspecific_epithet, code, referenceID, publishedInYear, link
		)
	VALUES (?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?)
`)
	if err != nil {
		return err
	}

	for _, n := range data {
		_, err = stmt.Exec(
			n.ID, n.AlternativeID, n.BasionymID, n.ScientificName, n.Authorship,
			n.Rank, n.Uninomial, n.Genus, n.InfragenericEpithet, n.SpecificEpithet,
			n.InfraspecificEpithet, n.Code, n.ReferenceID, n.PublishedInYear, n.Link,
		)
		if err != nil {
			return err
		}
	}

	stmt.Close()
	return tx.Commit()
}
