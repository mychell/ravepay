package rave

// Bank is a type of rave bank resources
type Bank struct {
	Code            string `json:"bankcode"`
	Name            string `json:"bankname"`
	Internetbanking bool   `json:"internetbanking"`
}

// FIXME: Done to enable testing
var listBanksURL = buildURL(ListBanksURL)

// ListBanks returns list of banks from the rave api
// https://flutterwavedevelopers.readme.io/v1.0/reference#list-of-banks
func ListBanks() ([]Bank, error) {
	banks := []Bank{}

	err := sendRequestAndParseResponse("GET", listBanksURL, nil, &banks)
	return banks, err
}
