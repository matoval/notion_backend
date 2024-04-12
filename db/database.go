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
		cover JSON,
		createdby JSON,
		createdtime TEXT,
		icon JSON,
		intrash BOOL,
		lasteditedby JSON,
		lasteditedtime TEXT,
		object TEXT,
		parent JSON,
		properties JSON,
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

	stmt, err := db.Prepare("INSERT INTO notion(id, archived, cover, createdby, createdtime, icon, intrash, lasteditedby, lasteditedtime, object, parent, properties, publicurl, url) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Printf("Database prepare failed with error: %v", err)
	}

	for _, data := range notionData {
		fmt.Appendf("%v \n", string(data))
		_, err = stmt.Exec(data.Id, data.Archived, data.Cover, data.CreatedBy, data.CreatedTime, data.Icon, data.InTrash, data.LastEditedBy, data.LastEditedTime, data.Object, data.Parent, data.Properties, data.PublicUrl, data.Url)
		if err != nil {
			fmt.Printf("Database failed to add notion data with error: %v", err)
		}
	}

}
