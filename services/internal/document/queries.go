package document

const (
	deleteFolder = `DELETE FROM folder WHERE folder_id = :id`

	insertFolder = `
		INSERT INTO folder (folder_id, name, path)
		VALUES (
			:id,
			:name,
			:path
		) ON CONFLICT DO NOTHING
	`

	getFolderByPath = `
			SELECT * FROM folder WHERE path = :path
	`

	deleteDocumentByPath = `DELETE FROM document WHERE path = :path`

	getDocumentByPath = `
		SELECT * FROM document WHERE path = :path
	`

	insertDocument = `
		INSERT INTO document (document_id, document_type_id, name, body, download_url, folder_id, path, object_storage_url)
		VALUES (
			:id,
			:typeID,
			:name,
			:body,
			:downloadURL,
			:folderID,
			:path,
			:objectStorageURL
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

	getIndexContent = `
		SELECT
			s.sentence_id,
			s.document_id,
			s.annoy_id,
			s.embedding
		FROM sentence s
		INNER JOIN document d ON s.document_id = d.document_id
	`

	getDocumentsBySentence = `
		SELECT
			d.name AS name,
			d.document_id AS document_id,
			d.download_url AS download_url,
			d.path AS path,
			s.annoy_id AS annoy_id,
			s.context AS sentence_text,
			0.0 AS rel,
			s.sentence_id AS sentence_id
		FROM document d
		INNER JOIN sentence s ON s.document_id = d.document_id
		WHERE s.annoy_id IN (?)
	`

	fullTextSearchSentences = `
		SELECT document_id, annoy_id, name, context as sentence_text, (rel/(rel+1)) as rel, sentence_id, download_url, "path" as "path" FROM (
			SELECT
				document.document_id as document_id,
				document.name as name,
				document.download_url as download_url,
				document.path as "path",
				context,
				sentence.annoy_id as annoy_id,
				to_tsvector(sentence.body) as v,
				ts_rank(to_tsvector(sentence.body), plainto_tsquery(:search)) as rel,
				sentence.sentence_id as sentence_id
			FROM sentence
			INNER JOIN document ON sentence.document_id = document.document_id
		) search
		WHERE search.v @@ plainto_tsquery(:search)
		ORDER BY rel DESC LIMIT 20;
	`
)
