package respond

import (
	"encoding/json"
	// "fmt"
	"net/http"
	// "rest_api/model"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	error_message := map[string]interface{}{
		"Message": message,
		"Status":  0,
	}
	RespondWithJSON(w, code, error_message)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	if code != 400{
		payload = map[string]interface{}{
			"Message" : "Success",
			"Status"  : 1,
			"Data"    : payload,
		}
	}

		
	
	

	// cannot range over payload (type interface {})

	// response, _ := json.Marshal(payload)
	// if rec, ok := payload.(map[string]interface{}); ok {
	//     for key, val := range rec {
	// 		fmt.Println(key, val)
	//     }
	// } else {
	//     fmt.Printf("record not a map[string]interface{}: %v\n", payload)
	// }
	// fmt.Println(response)

	// for _, item := range payload.(map[string]interface{}){

	// }
	// fmt.Println(payload)
	// success_message := make(map[string]interface{})
	// v := reflect.ValueOf(payload)
	// fmt.Println("in respond")
	// if v.Kind() == reflect.Map {
	// 	fmt.Println("reflect map")
	// 	for _, key := range v.MapKeys() {
	// 		strct := v.MapIndex(key)
	// 		fmt.Println(key.Interface(), strct.Interface())
	// 		success_message[key.Interface().(string)] = strct.Interface()
	// 	}
	// }
	// fmt.Println(success_message)
	// for k, value := range payload {
	//     if _, ok := payload[k]; ok {
	//         success_message[k] = value
	//     }
	// }
	
	response, _ := json.Marshal(payload)
	// json.NewDecoder(response)
	// json.Unmarshal([]byte(response), &user)
	// fmt.Println("inituh user")
	// fmt.Println(user)
	// out["message"] = "Success"
	// out["status"] = 1
	// response, _ = json.Marshal(out)

	// fmt.Println(out)

	// // if rec, ok := payload.(map[string]interface{}); ok {
	//     for key, val := range response {
	// 		fmt.Println(key, val)
	//     }
	// // } else {
	// //     fmt.Printf("record not a map[string]interface{}: %v\n", payload)
	// // }
	// fmt.Println(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
