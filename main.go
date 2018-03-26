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
	changeHeaderThenServe := func(h http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			h.ServeHTTP(w, r)
		}
	}

	http.Handle("/docker/", changeHeaderThenServe(http.StripPrefix("/docker", http.FileServer(http.Dir("/var/docker/volume/swagger")))))
	// http.Handle("/docker/", http.StripPrefix("/docker", http.FileServer(http.Dir("/var/docker/volume/swagger"))))
	// http.Handle("/docker/", changeHeaderThenServe(http.StripPrefix("/docker", http.FileServer(http.Dir("C:\\docker\\gitlab\\volume\\swagger")))))
	// http.Handle("/docker2/", http.StripPrefix("/docker2", http.FileServer(http.Dir("C:\\docker\\gitlab\\volume\\swagger"))))
	http.ListenAndServe(":8080", nil)
}

func swagger(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	http.ServeFile(w, req, "swagger.json")
}
