package forum

import (
	"log"
)

func allForumUnames() []string {
	var allUnames []string
	rows, err := db.Query("SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var uname string
		rows.Scan(&uname)
		allUnames = append(allUnames, uname)
	}
	return allUnames
}
