package sfgarcio

import (
	"log/slog"

	"github.com/gnames/coldp/ent/coldp"
)

func (s *sfgarcio) InsertTaxa(data []coldp.Taxon) error {
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
	INSERT INTO taxon
		(
		id, alternative_id, source_id, parent_id, ordinal, branch_length,
		name_id, name_phrase, according_to_id, according_to_page,
		according_to_page_link, scrutinizer, scrutinizer_id,
		scrutinizer_date, provisional, reference_id, extinct,
		temporal_range_start_id, temporal_range_end_id,
		environment_id, species, section, subgenus, genus, subtribe,
		tribe, subfamily, family, superfamily, suborder, "order",
		subclass, class, subphylum, phylum, kingdom,
		link, remarks, modified, modified_by
		)
	VALUES (
		?,?,?,?,?,?, ?,?,?,?, ?,?,?, ?,?,?,?, ?,?, ?,?,?,?,?,?,
		?,?,?,?,?,?, ?,?,?,?,?, ?,?,?,?
		)
`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, t := range data {
		_, err = stmt.Exec(
			t.ID, t.AlternativeID, t.SourceID, t.ParentID, t.Ordinal, t.BranchLength,
			t.NameID, t.NamePhrase, t.AccordingToID, t.AccordingToPage,
			t.AccordingToPageLink, t.Scrutinizer, t.ScrutinizerID,
			t.ScrutinizerDate, t.Provisional, t.ReferenceID, t.Extinct,
			t.TemporalRangeStart.String(), t.TemporalRangeEnd.String(),
			t.Environment, t.Species, t.Section, t.Subgenus, t.Subtribe,
			t.Tribe, t.Subfamily, t.Family, t.Superfamily, t.Suborder, t.Order,
			t.Subclass, t.Class, t.Subphylum, t.Phylum, t.Kingdom,
			t.Link, t.Remarks, t.Modified, t.ModifiedBy,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
