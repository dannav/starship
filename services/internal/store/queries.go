package store

const (
	insertStore = `
		INSERT INTO store (store_id, location)
		SELECT :id, :location RETURNING *
	`

	getStore = `
		SELECT * FROM store LIMIT 1
	`
)
