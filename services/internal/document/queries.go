package document

const (
	insertDocument = `
		INSERT INTO document (document_id, document_type_id, team_id, name, body, download_url)
		VALUES (
			:id,
			:typeID,
			:teamID,
			:name,
			:body,
			:downloadURL
		) RETURNING *
	`

	getDocument = `
		SELECT * FROM document WHERE document_id = :id
	`

	insertSentence = `
		INSERT INTO sentence (sentence_id, document_id, store_id, body, embedding, context)
		VALUES (
			:id,
			:documentID,
			:storeID,
			:body,
			:embedding,
			:context
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
			d.download_url AS download_url,
			s.annoy_id AS annoy_id,
			s.context AS sentence_text,
			0.0 AS rel,
			s.sentence_id AS sentence_id
		FROM document d
		INNER JOIN sentence s ON s.document_id = d.document_id
		WHERE s.annoy_id IN (?)
	`

	fullTextSearchSentences = `
		SELECT document_id, annoy_id, name, context as sentence_text, (rel/(rel+1)) as rel, sentence_id, download_url FROM (
			SELECT
				document.document_id as document_id,
				document.name as name,
				document.download_url as download_url,
				context,
				sentence.annoy_id as annoy_id,
				to_tsvector(sentence.body) as v,
				ts_rank(to_tsvector(sentence.body), plainto_tsquery(:text)) as rel,
				sentence.sentence_id as sentence_id
			FROM sentence
			INNER JOIN document ON sentence.document_id = document.document_id
		) search
		WHERE search.v @@ plainto_tsquery(:text)
		ORDER BY rel DESC LIMIT 20;
	`
)
