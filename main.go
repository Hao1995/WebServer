// @APIVersion 1.0.0
// @APITitle Teamwork Desk
// @APIDescription Bend Teamwork Desk to your will using these read and write endpoints
// @Contact support@teamwork.com
// @TermsOfServiceUrl https://www.teamwork.com/termsofservice
// @License BSD
// @LicenseUrl http://opensource.org/licenses/BSD-2-Clause

package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/swagger.json", swagger)
	http.HandleFunc("/index.json", index)
	http.ListenAndServe(":8080", nil)
}

func swagger(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	http.ServeFile(w, req, "swagger.json")
}

func index(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "index.json")
}
