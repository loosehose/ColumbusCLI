package domain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"
)

func TLD(domainSuffix string) {
	url := fmt.Sprintf("https://columbus.elmasy.com/api/tld/%s", domainSuffix)

	resp, err := http.Get(url)
	if err != nil {
		log.Error().Err(err).Msgf("Error making GET request to %s", url)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error().Msgf("Unexpected status code %d from %s", resp.StatusCode, url)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("Error reading response body")
		return
	}

	var tldEntries []string
	err = json.Unmarshal(body, &tldEntries)
	if err != nil {
		log.Error().Err(err).Msg("Error unmarshalling JSON response")
		return
	}

	// Sanitize and display the TLD list
	for _, entry := range tldEntries {
		if entry != "" {
			fmt.Println(entry)
		}
	}
}
