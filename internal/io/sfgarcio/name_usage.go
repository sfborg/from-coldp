package sfgarcio

import (
	"log/slog"

	"github.com/gnames/coldp/ent/coldp"
)

func (s *sfgarcio) InsertNameUsages(data []coldp.NameUsage) error {
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

	tStmt, err := tx.Prepare(`
	INSERT INTO taxon
		(
		id, alternative_id, source_id, parent_id, ordinal, branch_length,
		name_id, name_phrase, according_to_id, according_to_page,
		according_to_page_link, scrutinizer, scrutinizer_id,
		scrutinizer_date, reference_id, extinct,
		temporal_range_start_id, temporal_range_end_id,
		environment_id, species, section, subgenus, genus, subtribe,
		tribe, subfamily, family, superfamily, suborder, "order",
		subclass, class, subphylum, phylum, kingdom,
		link, remarks, modified, modified_by
		)
	VALUES (
		?,?,?,?,?,?, ?,?,?,?, ?,?,?, ?,?,?, ?,?, ?,?,?,?,?,?,
		?,?,?,?,?,?, ?,?,?,?,?, ?,?,?,?
		)
`)
	if err != nil {
		return err
	}
	defer tStmt.Close()

	nStmt, err := tx.Prepare(`
	INSERT INTO name
		(
		id, alternative_id, source_id, scientific_name, authorship,
		rank_id, uninomial, genus, infrageneric_epithet,
		specific_epithet,infraspecific_epithet, cultivar_epithet,
		notho_id, original_spelling, combination_authorship,
		combination_authorship_id, combination_ex_authorship,
		combination_ex_authorship_id, combination_authorship_year,
		basionym_authorship, basionym_authorship_id,
		basionym_ex_authorship, basionym_ex_authorship_id,
		basionym_authorship_year, code_id, status_id, reference_id,
		published_in_year, published_in_page, published_in_page_link,
		gender_id, gender_agreement, etymology,
		link, remarks, modified, modified_by,
		gn_full_scientific_name
		)
	VALUES (?,?,?,?,?, ?,?,?,?, ?,?,?, ?,?,?, ?,?, ?,?, ?,?, ?,?, ?,?,?,?, ?,?,?,
		?,?,?, ?,?,?,?, ?) 
`)
	if err != nil {
		return err
	}
	defer nStmt.Close()

	sStmt, err := tx.Prepare(`
	INSERT INTO synonym
	  (
		id, taxon_id, source_id, name_id, name_phrase, according_to_id,
		status_id, reference_id, link, remarks, modified, modified_by
	  )
	VALUES (?,?,?,?,?,?, ?,?,?,?,?,?)`)
	if err != nil {
		return err
	}
	defer sStmt.Close()

	basStmt, err := tx.Prepare(`
	INSERT INTO name_relation
		(name_id, related_name_id, type_id)
	VALUES (?, ?, ?)
`)
	if err != nil {
		return err
	}
	defer basStmt.Close()

	for _, d := range data {
		switch d.TaxonomicStatus {
		case coldp.SynonymTS, coldp.AmbiguousSynonymTS, coldp.MisappliedTS:
			_, err = sStmt.Exec(
				d.ID, d.ParentID, d.SourceID, d.ID, d.NamePhrase, d.AccordingToID,
				d.TaxonomicStatus, d.ReferenceID, d.Link, d.NameRemarks,
				d.Modified, d.ModifiedBy,
			)
			if err != nil {
				return err
			}
		default:
			_, err = tStmt.Exec(
				d.ID, d.AlternativeID, d.SourceID, d.ParentID, d.Ordinal, d.BranchLength,
				d.ID, d.NamePhrase, d.AccordingToID, d.AccordingToPage,
				d.AccordingToPageLink, d.Scrutinizer, d.ScrutinizerID,
				d.ScrutinizerDate, d.ReferenceID, d.Extinct,
				d.TemporalRangeStart, d.TemporalRangeEnd,
				d.Environment, d.Species, d.Section, d.Subgenus, d.Genus, d.Subtribe,
				d.Tribe, d.Subfamily, d.Family, d.Superfamily, d.Suborder, d.Order,
				d.Subclass, d.Class, d.Subphylum, d.Phylum, d.Kingdom,
				d.Link, d.Remarks, d.Modified, d.ModifiedBy,
			)
			if err != nil {
				return err
			}
		}

		gsn := d.ScientificName
		if d.Authorship != "" {
			gsn = gsn + " " + d.Authorship
		}

		_, err = nStmt.Exec(
			d.ID, d.NameAlternativeID, d.SourceID, d.ScientificName, d.Authorship,
			d.Rank, d.Uninomial, d.GenericName, d.InfragenericEpithet,
			d.SpecificEpithet, d.InfraspecificEpithet, d.CultivarEpithet,
			d.Notho, d.OriginalSpelling, d.CombinationAuthorship,
			d.CombinationAuthorshipID, d.CombinationExAuthorship,
			d.CombinationExAuthorshipID, d.CombinationAuthorshipYear,
			d.BasionymAuthorship, d.BasionymAuthorshipID,
			d.BasionymExAuthorship, d.BasionymExAuthorshipID,
			d.BasionymAuthorshipYear, d.Code, d.NameStatus, d.NameReferenceID,
			d.PublishedInYear, d.PublishedInPage, d.PublishedInPageLink,
			d.Gender, d.GenderAgreement, d.Etymology,
			d.Link, d.NameRemarks, d.Modified, d.ModifiedBy,
			gsn,
		)

		if d.BasionymID == "" {
			continue
		}
		basStmt.Exec(
			d.ID, d.BasionymID, coldp.Basionym.String(),
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
