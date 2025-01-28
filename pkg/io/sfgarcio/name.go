package sfgarcio

import (
	"log/slog"

	"github.com/gnames/coldp/ent/coldp"
)

func (s *sfgarcio) InsertNames(data []coldp.Name) error {
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
		gn_scientific_name_string
		)
	VALUES (?,?,?,?,?, ?,?,?,?, ?,?,?, ?,?,?, ?,?, ?,?, ?,?, ?,?, ?,?,?,?, ?,?,?,
		?,?,?, ?,?,?,?, ?) 
`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	basStmt, err := tx.Prepare(`
	INSERT INTO name_relation
		(name_id, related_name_id, type_id)
	VALUES (?, ?, ?)
`)
	if err != nil {
		return err
	}
	defer basStmt.Close()

	for _, n := range data {
		_, err = stmt.Exec(
			n.ID, n.AlternativeID, n.SourceID, n.ScientificName, n.Authorship,
			n.Rank.ID(), n.Uninomial, n.Genus, n.InfragenericEpithet,
			n.SpecificEpithet, n.InfraspecificEpithet, n.CultivarEpithet,
			n.Notho.ID(), n.OriginalSpelling, n.CombinationAuthorship,
			n.CombinationAuthorshipID, n.CombinationExAuthorship,
			n.CombinationExAuthorshipID, n.CombinationAuthorshipYear,
			n.BasionymAuthorship, n.BasionymAuthorshipID,
			n.BasionymExAuthorship, n.BasionymExAuthorshipID,
			n.BasionymAuthorshipYear, n.Code.ID(), n.Status.ID(), n.ReferenceID,
			n.PublishedInYear, n.PublishedInPage, n.PublishedInPageLink,
			n.Gender.ID(), n.GenderAgreement, n.Etymology,
			n.Link, n.Remarks, n.Modified, n.ModifiedBy,
			n.ScientificNameString,
		)
		if err != nil {
			return err
		}

		if n.BasionymID == "" {
			continue
		}
		basStmt.Exec(
			n.ID, n.BasionymID, coldp.Basionym.ID(),
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
