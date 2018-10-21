package db

import (
	"fmt"

	"github.com/dannav/starship/services/internal/document"
)

// Schema represents the database schema
const Schema = `
	CREATE TABLE IF NOT EXISTS store (
		store_id UUID NOT NULL PRIMARY KEY,
		team_id TEXT NOT NULL,
		location TEXT NOT NULL,
		created TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
		updated TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
		CONSTRAINT uniq_store UNIQUE (team_id, location)
	);

	CREATE TABLE IF NOT EXISTS document_type (
		document_type_id INT NOT NULL PRIMARY KEY,
		name text NOT NULL,
		created TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
		updated TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
		CONSTRAINT uniq_doc_type UNIQUE (name)
	);

	CREATE TABLE IF NOT EXISTS document (
		document_id UUID NOT NULL PRIMARY KEY,
		document_type_id INT NOT NULL,
		team_id TEXT NOT NULL,
		name TEXT NOT NULL,
		body TEXT NOT NULL,
		created TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
		updated TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
		FOREIGN KEY (document_type_id) REFERENCES document_type (document_type_id)
	);

	CREATE TABLE IF NOT EXISTS sentence (
		sentence_id UUID NOT NULL PRIMARY KEY,
		document_id UUID NOT NULL,
		store_id UUID NOT NULL,
		annoy_id SERIAL NOT NULL,
		body TEXT NOT NULL,
		embedding JSON NOT NULL,
		created TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
		updated TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
		FOREIGN KEY (document_id) REFERENCES document (document_id),
		FOREIGN KEY (store_id) REFERENCES store (store_id)
	);
`

// Seed represents data that should exist in the database and ignores values that already exist
var Seed = []string{
	fmt.Sprintf(`
		INSERT INTO document_type(document_type_id, name)
		VALUES('%v','%v') ON CONFLICT DO NOTHING

	`, document.TypeMarkdown, "Markdown"),
	fmt.Sprintf(`
		INSERT INTO document_type(document_type_id, name)
		VALUES('%v','%v') ON CONFLICT DO NOTHING
	`, document.TypeText, "Text"),
}
