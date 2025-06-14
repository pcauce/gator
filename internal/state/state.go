package state

import (
	"database/sql"
	"github.com/pcauce/gator/internal/config"
	"github.com/pcauce/gator/internal/database"
)

type AppState struct {
	DBQueries *database.Queries
	Config    *config.Config
}

func LoadAppState() (AppState, error) {
	state := AppState{
		Config: config.ReadConfig(),
	}
	dbConn, err := sql.Open("postgres", state.Config.DBUrl)
	if err != nil {
		return AppState{}, err
	}
	state.DBQueries = database.New(dbConn)
	return state, nil
}
