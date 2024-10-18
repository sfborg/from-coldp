package sfgarcio

import "github.com/gnames/coldp/ent/coldp"

func (s *sfgarcio) InsertName(n coldp.Name) error {
	q := `
	INSERT INTO name
	(
	id, alternative_id, basionym_id, scientific_name, authorship,
	rank, uninomial, genus, infrageneric_epithet, specific_epithet,
	infraspecific_epithet, code, referenceID, publishedInYear, link
	)
	VALUES (?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?)
`
	_, err := s.db.Exec(q,
		n.ID, n.AlternativeID, n.BasionymID, n.ScientificName, n.Authorship,
		n.Rank, n.Uninomial, n.Genus, n.InfragenericEpithet, n.SpecificEpithet,
		n.InfraspecificEpithet, n.Code, n.ReferenceID, n.PublishedInYear, n.Link,
	)
	if err != nil {
		return err
	}

	return nil
}
