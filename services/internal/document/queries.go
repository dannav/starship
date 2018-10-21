package document

const (
	insertDocument = `
		INSERT INTO document (document_id, document_type_id, team_id, name, body)
		VALUES (
			:id,
			:typeID,
			:teamID,
			:name,
			:body
		) RETURNING *
	`

	getDocument = `
		SELECT * FROM document WHERE document_id = :id
	`

	insertSentence = `
		INSERT INTO sentence (sentence_id, document_id, store_id, body, embedding)
		VALUES (
			:id,
			:documentID,
			:storeID,
			:body,
			:embedding
		) RETURNING *
	`

	getDocumentSentences = `
		SELECT * FROM sentence WHERE document_id = ?
	`

	getIndexContentForTeam = `
		SELECT
			s.sentence_id,
			s.document_id,
			s.annoy_id,
			s.embedding
		FROM sentence s
		INNER JOIN document d ON s.document_id = d.document_id AND d.team_id = :teamID
	`

	getDocumentsBySentence = `
		SELECT
			d.name AS name,
			d.document_id AS document_id,
			s.annoy_id AS annoy_id,
			s.body AS sentence_text
		FROM document d
		INNER JOIN sentence s ON s.document_id = d.document_id
		WHERE s.annoy_id IN (?)
	`
)
