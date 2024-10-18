package sfgarcio

import (
	"log/slog"
	"strings"

	"github.com/gnames/coldp/ent/coldp"
)

func (s *sfgarcio) InsertMeta(m *coldp.Meta) error {
	var id int
	q := `INSERT INTO metadata
			(
   doi, title, alias, description, issued, version, keywords,
   geographic_scope, taxonomic_scope, temporal_scope, confidence,
   completeness, license, url, logo, label, citation, private
		  )
		VALUES
			(?, ?, ?, ?, ?, ?, ?,  ?, ?, ?, ?,  ?, ?, ?, ?, ?, ?, ?)`

	keywords := strings.Join(m.Keywords, ",")
	private := "no"
	if m.Private {
		private = "yes"
	}

	_, err := s.db.Exec(q,
		m.DOI, m.Title, m.Alias, m.Description, m.Issued, m.Version, keywords,
		m.GeographicScope, m.TaxonomicScope, m.TemporalScope, m.Confidence,
		m.Completeness, m.License, m.URL, m.Logo, m.Label, m.Citation, private,
	)
	if err != nil {
		slog.Error("Error inserting metadata", "error", err)
		return err
	}

	err = s.db.QueryRow("SELECT last_insert_rowid()").Scan(&id)
	if err != nil {
		slog.Error("Error getting ID for inserted metadata", "error", err)
		return err
	}

	if m.Contact != nil {
		err = s.addActor(m.Contact, id, "contact")
		if err != nil {
			return err
		}
	}
	if m.Publisher != nil {
		err = s.addActor(m.Publisher, id, "publisher")
		if err != nil {
			return err
		}
	}

	for _, v := range m.Editors {
		err = s.addActor(&v, id, "editor")
		if err != nil {
			return err
		}
	}

	// for _, v := range m.Creators {
	// 	err = s.addActor(&v, id, "creator")
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	for _, v := range m.Contributors {
		err = s.addActor(&v, id, "contributor")
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *sfgarcio) addActor(cnt *coldp.Actor, metaID int, table string) error {
	q := `INSERT INTO ` + table + `
      (metadata_id, orcid, given, family, rorid, name, email, url, note)
			VALUES
      (?,?,?,?,?,?,?,?,?)`

	_, err := s.db.Exec(q,
		metaID, cnt.Orcid, cnt.Given, cnt.Family, cnt.RorID, cnt.Organization,
		cnt.Email, cnt.URL, cnt.Note,
	)
	if err != nil {
		slog.Error("Error inserting metadata contact", "error", err)
		return err
	}
	return nil
}
