package sfgarcio

import "github.com/gnames/coldp/ent/coldp"

func (s *sfgarcio) InsertNameRelations(data []coldp.NameRelation) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(`
	INSERT INTO name_relation
		(
		name_id, related_name_id, source_id, type_id, page, reference_id,
		remarks, modified, modified_by,
		)
	VALUES (?,?,?,?,?,?, ?,?,?)
`)
	if err != nil {
		return err
	}

	for _, n := range data {
		_, err = stmt.Exec(
			n.NameID, n.RelatedNameID, n.SourceID, n.Type, n.Page, n.ReferenceID,
			n.Remarks, n.Modified, n.ModifiedBy,
		)
		if err != nil {
			return err
		}
	}

	stmt.Close()
	return tx.Commit()
}
