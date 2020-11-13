package api

import (
	"fmt"
	"net/http"

	mymetrics "github.com/fsaint-michel/demo-devops/internal/metrics"

	"github.com/rs/zerolog/log"
)

// Hello endpoint
func Hello(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Hello Endpoint Hit")

	mymetrics.HelloOK.Inc()

	log.Info().Msgf("Successfully done Hello func")
	w.Write([]byte(fmt.Sprintf("Hello Word\n")))
}

// HelloNok endpoint
func HelloNok(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Hello Endpoint Hit")

	mymetrics.HelloNOK.Inc()

	log.Info().Msgf("Failed done Hello func")
	w.Write([]byte(fmt.Sprintf("Hello Word marche pas!\n")))
}
