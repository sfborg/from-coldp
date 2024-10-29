package sfgarcio

import "github.com/gnames/coldp/ent/coldp"

func (s *sfgarcio) InsertTypeMaterials(data []coldp.TypeMaterial) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(`
	INSERT INTO type_material
		(
		id, source_id, name_id, citation, status_id, institution_code,
		catalog_number, reference_id, locality, country, latitude,
		longitude, altitude, host, sex_id, date, collector,
		associated_sequences, link, remarks, modified, modified_by
		)
	VALUES (?,?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?,?, ?,?,?,?,?)
`)
	if err != nil {
		return err
	}

	for _, n := range data {
		_, err = stmt.Exec(
			n.ID, n.SourceID, n.NameID, n.Citation, n.Status, n.InstitutionCode,
			n.CatalogNumber, n.ReferenceID, n.Locality, n.Country, n.Latitude,
			n.Longitude, n.Altitude, n.Host, n.Sex, n.Date, n.Collector,
			n.AssociatedSequences, n.Link, n.Remarks, n.Modified, n.ModifiedBy,
		)
		if err != nil {
			return err
		}
	}

	stmt.Close()
	return tx.Commit()
}
