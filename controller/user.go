package controller

import (
	// "database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rest_api/database"
	"rest_api/model"
	"rest_api/respond"
	"strconv"

	"bytes"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	// "context"
)

func GetUserLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("masukz xcoy")
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	fmt.Println(userInfo)

	// var user model.User
	message := fmt.Sprintf("hello %s", userInfo["Username"])
	w.Write([]byte(message))
}

func ReadUser(w http.ResponseWriter, r *http.Request) {
	// read params
	vars := mux.Vars(r)
	// convert string to integer
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respond.RespondWithError(w, http.StatusBadRequest, "Invalid User Id")
		return
	}
	u := model.User{Id: id}
	err = u.GetUser(database.DB)
	if err != nil {
		// arr_string_err = append(arr_string_err,)
		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, u)
	return
}

type MyResponseWriter struct {
	http.ResponseWriter
	buf *bytes.Buffer
}

func (mrw *MyResponseWriter) Write(p []byte) (int, error) {
	return mrw.buf.Write(p)
}

func Login(w http.ResponseWriter, req *http.Request) {
	// err := req.ParseForm()
	//     username := req.Form.Get("username")
	//     password := req.Form.Get("password")
	//     fmt.Println(username, password)
	// fmt.Printf("username %s\n", req.FormValue("username"))
	// fmt.Printf("password %s\n", req.FormValue("password"))

	// if err != nil {
	// 	fmt.Println("data bukan form data")
	// }
	// for key, values := range req.PostForm {
	// 	fmt.Println("key", key)
	// 	fmt.Println("value", values)
	// }
	var user model.User
	// user.Username = req.FormValue("username")
	// user.Password = req.FormValue("password")
	decoder := json.NewDecoder(req.Body)
	_ = decoder.Decode(&user)
	fmt.Println(user)
	err := user.Login(database.DB)
	if err != nil {
		// arr_string_err = append(arr_string_err,)
		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	// claims = jwt.Claims()
	claims := model.Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "Omah Ihrom Janti Claims",
			ExpiresAt: time.Now().Add(time.Duration(1) * time.Hour).Unix(),
		},
		Username: user.Username,
		UserId:   user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString([]byte("omahihromjantisupersecretboy"))

	json.NewEncoder(w).Encode(model.JwtToken{Token: tokenString})

	// v := req.Form
	// fmt.Println(req.Body["username"])

	// b, err := ioutil.ReadAll(req.Body)
	// fmt.Println(b)
	// arr_string_err = arr_string_err[:0]
	// var user2 model.User
	// decoder2 := json.NewDecoder(req.Body)
	// err = decoder2.Decode(&user2)
	// // err = json.NewDecoder(req.Body).Decode(&user2)
	// fmt.Println(user2)
	// if err != nil {
	// 	fmt.Println("masuk error")
	// 	arr_string_err = append(arr_string_err, err.Error())
	// 	respond.RespondWithError(w, http.StatusBadRequest, arr_string_err)
	// 	return
	// }

	// error := user.Login(database.DB)
	// if error != nil {
	// 	arr_string_err = append(arr_string_err, err.Error())
	// }

	// claims := model.Claims{
	// 	StandardClaims: jwt.StandardClaims{
	// 		Issuer:    "Omah Ihrom Janti Claims",
	// 		ExpiresAt: time.Now().Add(time.Duration(1) * time.Hour).Unix(),
	// 	},
	// 	Username: user.Username,
	// 	// Email:    userInfo["email"].(string),
	// 	// Group:    userInfo["group"].(string),
	// 	// "username": user.Username,
	// 	// "password": user.Password,
	// }

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// tokenString, error := token.SignedString([]byte("omahihromjantisupersecretboy"))

	// json.NewEncoder(w).Encode(model.JwtToken{Token: tokenString})
}

func ReadUsers(w http.ResponseWriter, r *http.Request) {
	users, err := model.GetUsers(database.DB)
	if err != nil {
		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// arr_string_err = append(arr_string_err,)
		respond.RespondWithError(w, http.StatusBadRequest, "Invalid User Id")
		return
	}
	u := model.User{Id: id}
	err = u.DeleteUser(database.DB)
	if err != nil {
		// arr_string_err = append(arr_string_err,)
		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, "User deleted")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respond.RespondWithError(w, http.StatusBadRequest, "Invalid User Id")
		return
	}

	var u model.User
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&u)
	if err != nil {
		respond.RespondWithError(w, http.StatusBadRequest, "Invalid Payload")
		return
	}
	defer r.Body.Close()
	u.Id = id
	err = u.UpdateUser(database.DB)
	if err != nil {
		// arr_string_err = append(arr_string_err,)
		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, u)

	// Kalo pake form
	// r.ParseForm()
	// for key, value := range r.Form {
	// 	fmt.Printf("%s = %s\n", key, value)
	// }
	// fmt.Printf("Name => %s\n", r.FormValue("Name"))
	// fmt.Printf("Age => %s\n", r.FormValue("Age"))
	// // age := r.FormValue("Age")
	// var u User
	// u.Id = id
	// u.Name = r.FormValue("Name")
	// u.Age, err = strconv.Atoi(r.FormValue("Age"))
	// if err != nil {
	// 	RespondWithError(w, http.StatusBadRequest, "Invalid Age Person")
	// 	return
	// }
	// err = u.UpdateUser(database.DB)
	// if err != nil {
	// 	switch err {
	// 	case sql.ErrNoRows:
	// 		RespondWithError(w, http.StatusNotFound, "User Not Found")
	// 	default:
	// 		RespondWithError(w, http.StatusBadRequest, err.Error())
	// 		return
	// 	}
	// }
	// RespondWithJSON(w, http.StatusOK, u)
	//

}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var u model.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	if err != nil {
		respond.RespondWithError(w, http.StatusBadRequest, "Invalid Payload")
		return
	}
	defer r.Body.Close()
	err = u.AddUser(database.DB)
	if err != nil {
		// arr_string_err = append(arr_string_err,)
		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, u)
}

func GetUserProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("User Product")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// arr_string_err = append(arr_string_err,)
		respond.RespondWithError(w, http.StatusBadRequest, "Invalid User Id")
		return
	}
	u := model.User{Id: id}
	fmt.Println("get users")
	err = u.GetUserProduct(database.DB)
	fmt.Println("get products")
	if err != nil {
		// arr_string_err = append(arr_string_err,)
		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, u)

}
