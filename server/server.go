package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"msgbroker/serverutils"
	"msgbroker/utils"
	"net/http"
)

var server *serverutils.Server = serverutils.NewServer()

func caller(qs map[string]*utils.Queue, d utils.Data) {
	fmt.Println(d)
	for _, v := range qs {
		v.Channel <- d
	}

}
func consumequeue(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	if server.Queues[params["queue"]] == nil {
		fmt.Fprintf(w, "No Queue found")
		return
	}
	msg := <-server.Queues[params["queue"]].Channel
	json.NewEncoder(w).Encode(msg)
}

func publish(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Publishing ...")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	exchangetype := params["exchange"]

	if server.Exchanges[exchangetype] == nil {
		fmt.Fprintf(w, "No Exchange found")
		return
	}
	d := utils.Data{}
	e := json.NewDecoder(r.Body).Decode(&d)
	fmt.Println(e)
	go server.Exchanges[exchangetype].Call(server.Exchanges[exchangetype].Queues, d)
	fmt.Println("Published")
	json.NewEncoder(w).Encode(d)
}

func addqueue(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	fmt.Println("Added queue", params["queue"])
	server.AddQueue(params["queue"])
}

func bindqueue(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	fmt.Println("Bound", params["queue"], "to", params["exchange"])
	server.BindQueue(params["exchange"], params["queue"])
}
func main() {

	r := mux.NewRouter()
	server.Exchanges["direct"] = utils.NewExchange(caller)
	r.HandleFunc("/consume/{queue}", consumequeue)
	r.HandleFunc("/publish/{exchange}", publish).Methods("POST")
	r.HandleFunc("/addqueue/{queue}", addqueue)
	r.HandleFunc("/bind/{queue}/{exchange}", bindqueue)
	http.ListenAndServe(":8000", r)
}
