package sfgarcio

import (
	"log/slog"

	"github.com/gnames/coldp/ent/coldp"
)

func (s *sfgarcio) InsertReferences(data []coldp.Reference) error {
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
	INSERT INTO reference
		(
		id, alternative_id, source_id, citation, type,
		author, author_id, editor, editor_id, title, title_short,
		container_author, container_title, container_title_short,
		issued, accessed, collection_title, collection_editor,
		volume, issue, edition, page, publisher,
		publisher_place, version, isbn, issn, doi,
		link, remarks, modified, modified_by
		)
	VALUES (
		?,?,?,?,?, ?,?,?,?,?,?, ?,?,?, ?,?,?,?, ?,?,?,?,?, ?,?,?,?,?, ?,?,?,?
	)
`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, d := range data {
		_, err = stmt.Exec(
			d.ID, d.AlternativeID, d.SourceID, d.Citation, d.Type.String(),
			d.Author, d.AuthorID, d.Editor, d.EditorID, d.Title, d.TitleShort,
			d.ContainerAuthor, d.ContainerTitle, d.ContainerTitleShort,
			d.Issued, d.Accessed, d.CollectionTitle, d.CollectionEditor,
			d.Volume, d.Issue, d.Edition, d.Page, d.Publisher,
			d.PublisherPlace, d.Version, d.ISBN, d.ISSN, d.DOI,
			d.Link, d.Remarks, d.Modified, d.ModifiedBy,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
