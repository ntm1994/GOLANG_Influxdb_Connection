package shops

// Product structure of my cool shop product


type AutoGenerated struct {
	Timestamp       string `json:"Timestamp"`
	QualiaID        int    `json:"QualiaID"`
	PlateNumber     string `json:"PlateNumber"`
	DeviceID        string `json:"DeviceId"`
	Version         string `json:"Version"`
	AdvertisementID string `json:"AdvertisementId"`
	EmojiStatus     string `json:"EmojiStatus"`
	InfoCount       string `json:"InfoCount"`
	BatteryLevel    int    `json:"BatteryLevel"`
	GatewayID       string `json:"GatewayId"`
	Location        struct {
		Latitude  string `json:"Latitude"`
		Longitude string `json:"Longitude"`
	} `json:"Location"`
}

/*var allProducts = []AutoGenerated{
	AutoGenerated{
		Timestamp:       "1579698559000+",
		QualiaID:        1254,
		PlateNumber:     "BDC2246",
		DeviceID:        "12345",
		Version:         "0.1",
		AdvertisementID: "12345",
		EmojiStatus:     "2244",
		InfoCount:       "+infoClickCount+",
		BatteryLevel:    12,
		GatewayID:       "12345",
	},
}

// GetAll fetches all products in my shop
func (product AutoGenerated) GetAll() []AutoGenerated {
	return allProducts
}


// Get fetches a particular product in my shop identified By its ID
func (product AutoGenerated) Get(id int) (AutoGenerated, error) {
	if id < 0 && len(allProducts) >= id {
		log.Println(id)
		return AutoGenerated{}, errors.New("No product found")
	}
	return allProducts[id], nil*/


