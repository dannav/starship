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
		context TEXT NOT NULL,
		embedding JSON NOT NULL,
		created TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
		updated TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
		FOREIGN KEY (document_id) REFERENCES document (document_id),
		FOREIGN KEY (store_id) REFERENCES store (store_id)
	);
`

// Indexes represents possible indexes to run on tables
const Indexes = `
	CREATE INDEX IF NOT EXISTS sentence_annoyid ON sentence (annoy_id);

	CREATE INDEX IF NOT EXISTS sentence_docid ON sentence (document_id);

	CREATE INDEX IF NOT EXISTS document_teamid ON document (team_id);

	CREATE OR REPLACE FUNCTION gin_fts_fct(body text)
		RETURNS tsvector
	AS
	$BODY$
		SELECT to_tsvector(body);
	$BODY$
	LANGUAGE sql
	IMMUTABLE;

	CREATE INDEX IF NOT EXISTS idx_fts_sentence ON sentence USING gin(gin_fts_fct(body));
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
	`, document.TypePDF, "PDF"),
	fmt.Sprintf(`
		INSERT INTO document_type(document_type_id, name)
		VALUES('%v','%v') ON CONFLICT DO NOTHING
	`, document.TypeUnsupported, "Unsupported"),
}
