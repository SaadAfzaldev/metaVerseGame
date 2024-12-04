package userHandler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/SaadAfzaldev/metaVerseGame/prisma/db"
)



func BulkMetaDataHandler (w http.ResponseWriter, r * http.Request) {

	userIDs := strings.Split(r.URL.Query().Get("ids"),",")

	client := db.NewClient()

	if err := client.Connect(); err != nil {
		http.Error(w,"Error connecting database", http.StatusBadRequest)
		return
	}

	defer client.Disconnect()

	users,err := client.User.FindMany(
		db.User.ID.In(userIDs),
	).Exec(r.Context())
	 
	fmt.Println(users)
		
	if err != nil {
		http.Error(w,"Failed to fetch users",http.StatusBadRequest)
		return
	}

	// This route is not completed yet
	



}
