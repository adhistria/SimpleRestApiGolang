package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"rest_api/database"
	"rest_api/helper"
	"rest_api/model"
	"rest_api/respond"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func UploadFiles(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20) // 32MB is the default used by FormFile
	formdata := r.MultipartForm    // ok, no problem so far, read the Form data
	//get the *fileheaders
	files := formdata.File["file_uploads"] // grab the filenames
	// m := r.MultipartForm
	// header := m.

	// multipart.FileHeader(files)
	for i, _ := range files { // loop through the files one by one
		// mimeType := multipart.FileHeader(*files[i])
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		// helper.GetFileContentType(file)
		// multipart.FileHeader(file)
		// content_type, err := helper.GetFileContentType(file.(*os.File))
		// if err!= nil {
		// 	respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		// 	return
		// }
		// fmt.Println(content_type)
		// out, err := os.Create("/Users/m.adhisatria/Documents/Beego/" + files[i].Filename)
		out, err := os.Create("./tmp/" + files[i].Filename)

		defer out.Close()
		if err != nil {
			fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
			fmt.Fprintf(w, err.Error())
			return
		}

		_, err = io.Copy(out, file) // file not files[i] !

		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		fmt.Fprintf(w, "Files uploaded successfully : ")
		fmt.Fprintf(w, files[i].Filename+"\n")

	}

	// for _, file_upload := range file_uploads {
	// 	fmt.Println("masuk loop")
	// 	file, err := file_upload.Open()
	// 	fmt.Println(file)
	// 	if err!= nil {
	// 		respond.RespondWithError(w, http.StatusBadRequest, "Invalid Data")
	// 	}
	// 	helper.SaveMultipleFile(file)
	// 	// f is one of the files
	// }
	respond.RespondWithJSON(w, http.StatusOK, "product")

}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	decoder := json.NewDecoder(r.Body)
	var product model.Product
	err := decoder.Decode(&product)
	product.User_Id = int(userInfo["User_Id"].(float64))
	if err != nil {
		respond.RespondWithError(w, http.StatusBadRequest, "Invalid Payload Request")
		return
	}

	// r.ParseMultipartForm(32 << 20) // 32MB is the default used by FormFile
	// file_uploads := r.MultipartForm.File["file_uploads"]
	// for _, file_upload := range file_uploads {
	// 	file, err := file_upload.Open()

	// 	// f is one of the files
	// }

	file, handle, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}
	defer file.Close()

	mimeType := handle.Header.Get("Content-Type")
	switch mimeType {
	case "image/jpeg":
		err = helper.SaveFile(file, handle)
	case "image/png":
		err = helper.SaveFile(file, handle)
	default:
		respond.RespondWithError(w, http.StatusBadRequest, "The format file is not valid.")
		return
	}
	err = product.AddProduct(database.DB)
	if err != nil {
		respond.RespondWithError(w, http.StatusBadRequest, "Fail Add New Product")
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respond.RespondWithError(w, http.StatusBadRequest, "Invalid Product Id")
	}

	p := model.Product{Id: id}
	// var errors []error
	// res, err := p.DeleteProduct(database.DB)
	err = p.DeleteProduct(database.DB)
	// fmt.Println(res)
	if err != nil {
		// for _, err  := range errors {
		// 	arr_string_err = append(arr_string_err, err.Error())
		// }
		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, p)
	// if err{
	// 	switch err{
	// 	case sql.ErrNoRows:
	// 		respond.RespondWithError(w, http.StatusBadRequest, "Product Not Found")
	// 	default:
	// 		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
	// 	}
	// 	return
	// }
	// var make(map[]string interface{})
	// respond.RespondWithJSON(w, http.StatusOK, "Delete Success")

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	fmt.Println(id)
	if err != nil {
		// arr_string_err = append(arr_string_err,)
		fmt.Println(err.Error())
		respond.RespondWithError(w, http.StatusBadRequest, "Invalid Product Id")
		return
	}
	p := model.Product{Id: id}
	json.NewDecoder(r.Body).Decode(&p)
	p.User_Id = int(userInfo["User_Id"].(float64))
	fmt.Println(p)
	err = p.UpdateProduct(database.DB)

	json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, p)
}

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	products, err := model.GetAllProduct(database.DB)
	if err != nil {
		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// arr_string_err = append(arr_string_err,)
		respond.RespondWithError(w, http.StatusBadRequest, "Product Id")
		return
	}
	p := model.Product{Id: id}
	err = p.GetProduct(database.DB)
	fmt.Println("sebenernya udah masuk get product")
	if err != nil {
		// arr_string_err = append(arr_string_err,)
		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respond.RespondWithJSON(w, http.StatusOK, p)
	// if err!= nil {
	// 	switch err{
	// 	case sql.ErrNoRows:
	// 		respond.RespondWithError(w, http.StatusBadRequest, "Product Not Found")
	// 	default:
	// 		respond.RespondWithError(w, http.StatusBadRequest, err.Error())
	// 	}
	// 	return
	// }
	// respond.RespondWithJSON(w, http.StatusOK, p)

}
