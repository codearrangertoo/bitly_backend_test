package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	openapi "github.com/bitly/openapi"
	log "github.com/sirupsen/logrus"
)

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {

	token := strings.Fields(r.Header.Get("Authorization"))

	if len(token) != 2 {
		w.WriteHeader(http.StatusUnauthorized)
		log.Error("Incorrect number of fields in Authorization header")
		return
	}

	if token[0] != "Bearer" {
		w.WriteHeader(http.StatusUnauthorized)
		log.Error("No Bearer in Authorization header")
		return
	}

	client := openapi.NewAPIClient(openapi.NewConfiguration())
	ctx := context.WithValue(r.Context(), openapi.ContextAccessToken, token[1])

	group_guid, err := getDefaultGroupGuid(ctx, client)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Error(err)
		return
	}

	metrics, err := buildMetrics(ctx, client, group_guid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)

}

func buildMetrics(ctx context.Context, client *openapi.APIClient, group_guid string) (map[string]float64, error) {

	metrics := map[string]float64{}

	links, err := getBitLinksByGroup(ctx, client, group_guid)
	if err != nil {
		return metrics, err
	}

	for _, link := range links {
		clickMetrics, err := getMetricsForBitlinkByCountries(ctx, client, *link.Id)
		if err != nil {
			return metrics, err
		}
		for _, metric := range clickMetrics {
			metrics[*metric.Value] += float64(*metric.Clicks)
		}

	}

	for key, value := range metrics {
		metrics[key] = value / 30
	}

	return metrics, nil

}

func getDefaultGroupGuid(ctx context.Context, client *openapi.APIClient) (string, error) {
	user, _, err := client.UserApi.GetUser(ctx).Execute()
	if err != nil {
		return "", err
	}
	return *user.DefaultGroupGuid, nil
}

func getBitLinksByGroup(ctx context.Context, client *openapi.APIClient, group_guid string) ([]openapi.BitlinkBody, error) {

	var page int32 = 1
	var links []openapi.BitlinkBody

	// iterate through each page
	for {

		bitlinks, _, err := client.BitlinksApi.GetBitlinksByGroup(ctx, group_guid).Size(100).Page(page).Execute()
		if err != nil {
			return links, err
		}

		links = append(links, *bitlinks.Links...)

		if bitlinks.Pagination.GetNext() == "" {
			break
		}
		page = bitlinks.Pagination.GetPage() + 1

	}
	return links, nil

}

func getMetricsForBitlinkByCountries(ctx context.Context, client *openapi.APIClient, bitlink string) ([]openapi.ClickMetric, error) {

	var empty []openapi.ClickMetric

	resp, _, err := client.BitlinksApi.GetMetricsForBitlinkByCountries(ctx, bitlink).Unit("day").Units(int32(30)).Execute()
	if err != nil {
		return empty, err
	}

	return *resp.Metrics, nil
}
