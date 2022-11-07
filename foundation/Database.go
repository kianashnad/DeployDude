package foundation

import (
	"DeployDude/ent"
	"context"
)

type Project struct {
	ID      int
	Title   string
	GitURL  string
	HashID  string
	DirPath string
}

func GetDBClient() *ent.Client {
	client, err := ent.Open("sqlite3", "file:"+GetENV("DEDU_DB_FILENAME")+".sqlite3?mode=rwc&_fk=1")
	if err != nil {
		panic(err)
	}

	return client
}

func GenerateDBSchema() {
	client := GetDBClient()
	err := client.Schema.Create(context.Background())
	if err != nil {
		panic(err)
	}
	defer client.Close()
}
