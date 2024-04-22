package db

import (
	"database/sql"
	"fmt"
	"notion_backend/models"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDatabase() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Printf("Database failed to open %v", err)
	}
	defer db.Close()

	sql := `CREATE TABLE IF NOT EXISTS notion (
		id INTEGER PRIMARY KEY,
		archived BOOL,
		coverurl TEXT,
		coverexpirytime TEXT,
		createdbyid TEXT,
		createdbyobject TEXT
		createdtime TEXT,
		icontype TEXT,
		iconurl TEXT,
		iconexpirytime TEXT,
		iconemoji TEXT,
		intrash BOOL,
		lasteditedbyid TEXT,
		lasteditedbyobject TEXT,
		lasteditedtime TEXT,
		object TEXT,
		parentdatabaseid TEXT,
		parenttype TEXT,
		properties BLOB,
		publicurl TEXT,
		url TEXT
	);`

	_, err = db.Exec(sql)
	if err != nil {
		fmt.Printf("Database failed to create Notion table, Error: %v", err)
	}
}

func AddNotionData(notionData []models.NotionData) {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Printf("Database failed to open %v", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO notion(id, archived, coverurl, coverexpirytime, createdbyid, createdbyobject, createdtime, icontype, iconurl, iconexpirytime, iconemoji, intrash, lasteditedbyid, lastediedbyobject, lasteditedtime, object, parentdatabaseid, parenttype, properties, publicurl, url) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Printf("Database prepare failed with error: %v", err)
	}

	for _, data := range notionData {
		_, err = stmt.Exec(data.Id, data.Archived, data.Cover.Url, data.Cover.ExpiryTime, data.CreatedBy.Id, data.CreatedBy.Object, data.CreatedTime, data.Icon.Type, data.Icon.Url, data.Icon.ExpiryTime, data.Icon.Emoji, data.InTrash, data.LastEditedBy.Id, data.LastEditedBy.Object, data.LastEditedTime, data.Object, data.Parent.DatabaseId, data.Parent.Type, data.Properties, data.PublicUrl, data.Url)
		if err != nil {
			fmt.Printf("Database failed to add notion data with error: %v", err)
		}
	}

}
