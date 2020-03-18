package controllers

import (
	"api/responses"
	"net/http"
)

// ResponseVersion response
type ResponseVersion struct {
	Result string `json:"result"`
}

// Version user constroller
func Version(w http.ResponseWriter, r *http.Request) {
	version := &ResponseVersion{Result: "1.0.0"}
	responses.JSON(w, http.StatusOK, version)
}

// Login user
func Login(w http.ResponseWriter, r *http.Request) {
	// var req LoginRequestBody
	// err := json.NewDecoder(r.Body).Decode(&req)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// res := make(map[string]interface{})

	// user := models.FindUserByEmail(req.Email)
	// if user.ID != 0 {
	// 	passwordCompareResult := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	// 	if passwordCompareResult == nil {
	// 		res["user"] = user
	// 		res["token"], _ = auth.CreateToken(user.ID)
	// 		res["message"] = "Login Success."
	// 		w.WriteHeader(http.StatusOK)
	// 	} else {
	// 		res["message"] = "invalid password."
	// 		w.WriteHeader(http.StatusBadRequest)
	// 	}
	// } else {
	// 	res["message"] = "invalid email id."
	// 	w.WriteHeader(http.StatusBadRequest)
	// }
	// jData, _ := json.Marshal(res)
	// w.Write(jData)
	// return
}
