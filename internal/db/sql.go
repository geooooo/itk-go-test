package db

const (
	createTableSQL = `
		DROP TABLE IF EXISTS wallets.wallet;
		
		CREATE TABLE wallets.wallet (
			uuid VARCHAR(36) PRIMARY KEY NOT NULL,
			sum INTEGER NOT NULL
		)
	`

	fillDataSQL = `
		INSERT INTO wallets.wallet
			(uuid, sum) 
		VALUES 
			('336351b8-6c80-4498-b080-57e34f315aab', 0),
			('96efb4a3-3632-45b6-9e40-47d78b45ef98', 1000),
			('875cd5ee-d0f0-4605-8c87-ca5974420c91', 5000)
	`

	getWalletSumSQL = `
		SELECT sum 
		FROM wallets.wallet 
		WHERE uuid=$1
	`

	upateWalletSumSQL = `
		UPDATE wallets.wallet 
		SET sum=$1
		WHERE uuid=$2;
	`
)
