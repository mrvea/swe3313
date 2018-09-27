package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	_ "strconv"

	"github.com/class/pizza/env"
	_ "github.com/class/pizza/model"
	_ "github.com/eyesore/httprouter"
)

func UserRedirectHome(e *env.Env, w http.ResponseWriter, r *http.Request) error {
	http.Redirect(w, r, "/home", http.StatusPermanentRedirect)
	return nil
}

func UserHome(e *env.Env, w http.ResponseWriter, r *http.Request) error {
	t, err := template.ParseFiles("templates/dash.html.tpl")
	if err != nil {
		return err
	}
	// u := GetLoggedInUser(r)
	t.Execute(w, e)
	return nil
}

func GetCurrentUser(e *env.Env, w http.ResponseWriter, r *http.Request) error {
	u := GetLoggedInUser(r)
	err := json.NewEncoder(w).Encode(u)
	return err
}

// //UserAdd handles POST data to create a user in the database
// func UserAdd(e *env.Env, w http.ResponseWriter, r *http.Request) error {
// 	users := model.UserTable()
// 	ctx := r.Context()
// 	jsonData := ctx.Value(ParsedJSON)
// 	if jsonData != nil {
// 		loggedInUser := GetLoggedInUser(r)
// 		user, err := users.InsertUser(jsonData.(map[string]interface{}), loggedInUser, e.AB)
// 		if err != nil {
// 			return err
// 		}
// 		err = json.NewEncoder(w).Encode(user)
// 		return err
// 	}
// 	return nil
// }

// //UserEdit handles PUT data to edit an existing user in the database
// func UserEdit(e *env.Env, w http.ResponseWriter, r *http.Request) error {
// 	users := model.UserTable()
// 	ctx := r.Context()
// 	jsonData := ctx.Value(ParsedJSON)
// 	if jsonData != nil {
// 		loggedInUser := GetLoggedInUser(r)
// 		user, err := users.EditUser(jsonData.(map[string]interface{}), loggedInUser, e.AB)
// 		if err != nil {
// 			return err
// 		}
// 		err = json.NewEncoder(w).Encode(user)
// 		return err
// 	}
// 	return nil
// }

// //UserGet handles a GET request in order to retreive a user
// func UserGet(e *env.Env, w http.ResponseWriter, r *http.Request) error {
// 	params := httprouter.ParamsFromContext(r.Context())
// 	id, _ := strconv.Atoi(params.ByName("id"))
// 	users := model.UserTable()
// 	user, err := users.GetByID(id)
// 	if err == nil {
// 		err = json.NewEncoder(w).Encode(user)
// 	}

// 	return err
// }

// //UserDelete handles a DELETE request in order to remove a user from the database
// func UserDelete(e *env.Env, w http.ResponseWriter, r *http.Request) error {
// 	params := httprouter.ParamsFromContext(r.Context())
// 	id, _ := strconv.Atoi(params.ByName("id"))
// 	users := model.UserTable()
// 	err := users.DeleteUser(id)
// 	if err == nil {
// 		err = json.NewEncoder(w).Encode("true") //TODO update to default response
// 	}
// 	log.Debug(err)
// 	return err
// }

// //UserIndex handles a GET request in order to retreive a list of users in the database
// func UserIndex(e *env.Env, w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }

// func GetUsers(e *env.Env, w http.ResponseWriter, r *http.Request) error {
// 	log.Debug("getting Users")
// 	users, err := model.UserTable().GetAll()
// 	if err != nil {
// 		return err
// 	}
// 	err = json.NewEncoder(w).Encode(users)
// 	return err
// }
