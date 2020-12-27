// Package bins is a client for the Bedford Borough Council bin collections API.
package bins

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Client is the *http.Client used to request the API.
var Client *http.Client
var bedford *time.Location

const baseURL = "https://bbaz-as-prod-bartecapi.azurewebsites.net/api/bincollections/residential/getbyuprn/"

func init() {
	Client = &http.Client{
		Timeout: time.Second * 10, // Maximum of 10 secs
	}

	var err error
	bedford, err = time.LoadLocation("Europe/London")
	if err != nil {
		panic(err)
	}
}

// binResponse is a type to hold a single bins API response
type binResponse struct {
	BinCollectionDays []string
	BinCollections    []binCollection
}

type binCollection []struct {
	BinType           string
	JobScheduledStart customTime
}

// Collection represents all collections on a single day, which can have multiple bin colours.
type Collection struct {
	BinTypes []BinType
	Start    time.Time
}

type customTime struct {
	time.Time
}

func (t *customTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		t.Time = time.Time{}
		return
	}
	t.Time, err = time.ParseInLocation("2006-01-02T15:04:05", s, bedford)
	return
}

// BinType is the colour of bin collected.
type BinType string

const (
	// Black bin
	Black BinType = "Black bin"
	// Orange bin
	Orange BinType = "Orange bin"
	// Green bin
	Green BinType = "Green bin"
)

// GetByUPRN fetches bin collection information for a specified UPRN.
func GetByUPRN(ctx context.Context, uprn uint64) (bins []Collection, err error) {
	url := baseURL + strconv.FormatUint(uprn, 10)

	return getURL(ctx, url)
}

func getURL(ctx context.Context, url string) (bins []Collection, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return bins, err
	}

	req = req.WithContext(ctx)

	req.Header.Set("User-Agent", "BinsBot/0.1")

	res, getErr := Client.Do(req)
	if getErr != nil {
		return bins, getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return bins, readErr
	}

	binres := binResponse{}
	jsonErr := json.Unmarshal(body, &binres)
	if jsonErr != nil {
		return bins, jsonErr
	}

	for _, col := range binres.BinCollections {
		bins = append(bins, Collection{
			Start:    col.Start(),
			BinTypes: col.BinTypes(),
		})
	}

	return bins, nil
}

// Start returns the time that this BinCollection is due to start
func (c binCollection) Start() time.Time {
	return c[0].JobScheduledStart.Time
}

// BinTypes returns the types of bins due to be collected.
func (c binCollection) BinTypes() []BinType {
	result := []BinType{}

	for _, job := range c {
		for _, colour := range []BinType{Black, Orange, Green} {
			if job.BinType == string(colour) {
				result = append(result, colour)
			}
		}
	}

	return result
}
