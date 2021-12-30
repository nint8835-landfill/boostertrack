package availability

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/publicsuffix"
	"net/http"
	"net/http/cookiejar"
)

const beginBookingUrl = "https://portal.healthmyself.net/nleasternhealth/guest/booking/form/abec68ea-5a99-421b-b137-6e83cf7a3231"
const locationStatusUrl = "https://portal.healthmyself.net/nleasternhealth/guest/booking/type/5892/locations"
const startPageUrl = "https://portal.healthmyself.net/nleasternhealth/forms/P2E"

func GetAvailabilities() ([]LocationStatus, error) {
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, fmt.Errorf("error creating cookie jar: %w", err)
	}

	client := http.Client{Jar: jar}

	beginBookingReq, err := http.NewRequest(http.MethodGet, beginBookingUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("error constructing begin booking request: %w", err)
	}
	beginBookingReq.Header.Set("Referer", startPageUrl)
	beginBookingResp, err := client.Do(beginBookingReq)
	if err != nil {
		return nil, fmt.Errorf("error making begin booking request: %w", err)
	}
	defer beginBookingResp.Body.Close()

	fmt.Println(beginBookingResp.StatusCode)

	locationStatusResp, err := client.Get(locationStatusUrl)
	if err != nil {
		return nil, fmt.Errorf("error making location status request: %w", err)
	}
	defer locationStatusResp.Body.Close()

	var respData LocationStatusResp
	err = json.NewDecoder(locationStatusResp.Body).Decode(&respData)
	if err != nil {
		return nil, fmt.Errorf("error decoding location status json: %w", err)
	}

	return respData.Data, nil
}
