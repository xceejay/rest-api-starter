package handlers

import (
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/goccy/go-json"
	"github.com/xceejay/social-network.new/api/models"
	"gorm.io/gorm"
)

// +-----------------+--------------+------+-----+----------------------+--------------------------------+
// | Field           | Type         | Null | Key | Default              | Extra                          |
// +-----------------+--------------+------+-----+----------------------+--------------------------------+
// | id              | int(11)      | NO   | PRI | <null>               | auto_increment                 |
// | name            | varchar(255) | YES  |     | <null>               |                                |
// | email           | varchar(255) | YES  | UNI | <null>               |                                |
// | email_verified  | timestamp(6) | YES  |     | <null>               |                                |
// | image           | varchar(255) | YES  |     | <null>               |                                |
// | created_at      | timestamp(6) | NO   |     | current_timestamp(6) |                                |
// | updated_at      | timestamp(6) | NO   |     | current_timestamp(6) | on update current_timestamp(6) |
// | first_name      | varchar(100) | YES  |     | <null>               |                                |
// | last_name       | varchar(100) | YES  |     | <null>               |                                |
// | profile_picture | varchar(100) | YES  |     | <null>               |                                |
// | username        | varchar(100) | YES  | UNI | <null>               |                                |
// | password        | varchar(100) | YES  |     | <null>               |                                |
// | gender          | varchar(100) | YES  |     | <null>               |                                |
// | street          | varchar(100) | YES  |     | <null>               |                                |
// | city            | varchar(100) | YES  |     | <null>               |                                |
// | state           | varchar(100) | YES  |     | <null>               |                                |
// | phone           | varchar(100) | YES  |     | <null>               |                                |
// | dob             | varchar(100) | YES  |     | <null>               |                                |
// | country         | varchar(100) | YES  |     | <null>               |                                |
// +-----------------+--------------+------+-----+----------------------+--------------------------------+

func (server *Server) signIn(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	if err := user.SignIn(server.DB); err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			respondWithError(w, http.StatusNotFound, "User not found")
			// fmt.Println(err)
		} else {
			// fmt.Println(err)
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}

		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

