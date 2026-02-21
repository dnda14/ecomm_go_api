package json

import {

}

func Write(w http.ResponseWriter, status int,data any){
	w.Header().set("Content-Type","application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
	
}