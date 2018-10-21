package store

const (
	insertStore = `
		INSERT INTO store (store_id, team_id, location)
		SELECT :id, :teamID, :location
		WHERE NOT EXISTS (
			SELECT store_id
			FROM store
			WHERE team_id = :teamID
		) RETURNING *
	`

	getStoreByTeamID = `
		SELECT * FROM store WHERE team_id = :teamID
	`
)
