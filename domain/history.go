package domain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type DNSRecord struct {
	Type  int    `json:"type"`
	Value string `json:"value"`
	Time  int64  `json:"time"`
}

type HistoryEntry struct {
	Domain  string      `json:"Domain"`
	Records []DNSRecord `json:"Records"`
}

func History(domain string) {
	url := fmt.Sprintf("https://columbus.elmasy.com/api/history/%s?days=-1", domain)

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

	var historyEntries []HistoryEntry
	err = json.Unmarshal(body, &historyEntries)
	if err != nil {
		log.Error().Err(err).Msg("Error unmarshalling JSON response")
		return
	}

	// Display sanitized data
	for _, entry := range historyEntries {
		if len(entry.Records) > 0 {
			fmt.Println("\nDomain:", entry.Domain)
			for _, record := range entry.Records {
				fmt.Printf("  Type: %d, Value: %s, Time: %s\n", record.Type, record.Value, time.Unix(record.Time, 0))
			}
		}
	}
}
