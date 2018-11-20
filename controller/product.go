package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"omahihrom/database"
	"omahihrom/helper"
	"omahihrom/model"
	"omahihrom/respond"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func AddProduct(w http.ResponseWriter, req *http.Request) {

	userInfo := req.Context().Value("userInfo").(jwt.MapClaims)

	req.ParseMultipartForm(32 << 20)

	formdata := req.MultipartForm
	files := formdata.File["file_uploads"]
	price := req.FormValue("price")
	name := req.Form.Get("name")
	userId := int(userInfo["UserId"].(float64))
	convertedPrice, err := strconv.Atoi(price)

	helper.CheckError(w, err)

	product := model.Product{Name: name, UserId: userId, Price: convertedPrice}
	err = product.AddProduct(database.DB)

	for i, _ := range files {

		file, err := files[i].Open()

		defer file.Close()

		helper.CheckError(w, err)

		mimeType := files[i].Header.Get("Content-Type")
		err = helper.CheckMimeType(mimeType)
		helper.CheckError(w, err)

		extension := filepath.Ext(files[i].Filename)
		unixTimeStamp := int32(time.Now().Unix())
		imgUrl := "/static/" + helper.ConvertToString(unixTimeStamp) + extension
		image := model.Image{Url: imgUrl, ProductId: product.Id}

		err = image.AddImage(database.DB)
		helper.CheckError(w, err)
		// helper.SaveImage(imgUrl, file )

		imgPath := "." + imgUrl
		out, err := os.Create(imgPath)

		defer out.Close()
		helper.CheckError(w, err)

		_, err = io.Copy(out, file)

		helper.CheckError(w, err)
	}

	err = product.GetProduct(database.DB)
	helper.CheckError(w, err)
	respond.RespondWithJSON(w, http.StatusOK, product)

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	helper.CheckError(w, err)

	p := model.Product{Id: id}
	err = p.DeleteProduct(database.DB)

	helper.CheckError(w, err)

	respond.RespondWithJSON(w, http.StatusOK, p)

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	helper.CheckError(w, err)
	p := model.Product{Id: id}
	json.NewDecoder(r.Body).Decode(&p)
	p.UserId = int(userInfo["UserId"].(float64))

	err = p.UpdateProduct(database.DB)
	helper.CheckError(w, err)

	respond.RespondWithJSON(w, http.StatusOK, p)
}

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	products, err := model.GetAllProduct(database.DB)

	helper.CheckError(w, err)

	respond.RespondWithJSON(w, http.StatusOK, products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	helper.CheckError(w, err)

	p := model.Product{Id: id}
	err = p.GetProduct(database.DB)

	helper.CheckError(w, err)

	respond.RespondWithJSON(w, http.StatusOK, p)
}
