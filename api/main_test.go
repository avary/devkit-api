package api

import (
	"os"
	"testing"

	"github.com/darwishdev/devkit-api/config"
	"github.com/darwishdev/devkit-api/db"
	"github.com/darwishdev/devkit-api/proto_gen/devkit/v1/devkitv1connect"
	"github.com/rs/zerolog/log"
)

var testConfig config.Config

func newTestApi(store db.Store) devkitv1connect.DevkitServiceHandler {
	api, err := NewApi(testConfig, store)
	if err != nil {
		log.Fatal().Err(err).Msg("canot start the api")
	}
	return api
}
func TestMain(m *testing.M) {
	_, err := config.LoadState("../config")
	if err != nil {
		panic(err)
	}
	testConfig, err = config.LoadConfig("../config", "test")
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}