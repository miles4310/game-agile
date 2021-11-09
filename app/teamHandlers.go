package app

import (
	"encoding/json"

	"g/go/allsports/service"

	"net/http"
)

//Team Handlers
type TeamHandlers struct {
	service service.TeamService
}

//Gets all Teams from the repository
func (th *TeamHandlers) getAllTeams(w http.ResponseWriter, r *http.Request) {

	teams, _ := th.service.GetAllTeams()

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Content-Type", "applicatiom/json")
	json.NewEncoder(w).Encode(teams)

}
