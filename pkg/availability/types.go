package availability

// Address represents the actual geographic location of a Location.
type Address struct {
	AddressId   int         `json:"address_id"`
	Address     string      `json:"address"`
	Address2    string      `json:"address2"`
	City        string      `json:"city"`
	Province    string      `json:"province"`
	Country     string      `json:"country"`
	Postal      string      `json:"postal"`
	AddressType string      `json:"address_type"`
	Phone       string      `json:"phone"`
	PhoneExt    interface{} `json:"phone_ext"`
	Fax         string      `json:"fax"`
	UserId      interface{} `json:"user_id"`
	UpdatedAt   string      `json:"updated_at"`
	CreatedAt   string      `json:"created_at"`
}

// Location represents an individual defined vaccine location.
type Location struct {
	LocId          int     `json:"loc_id"`
	LocName        string  `json:"loc_name"`
	AddressId      *int    `json:"address_id"`
	Lat            *string `json:"lat"`
	Long           *string `json:"long"`
	LocWaitlistUrl *string `json:"loc_waitlist_url"`
	Address        Address `json:"address,omitempty"`
}

type LocationStatus struct {
	Id                         int  `json:"id"`
	HasUnavailableAppointments bool `json:"hasUnavailableAppointments"`
}

type LocationStatusResp struct {
	Success bool             `json:"success"`
	Data    []LocationStatus `json:"data"`
	Message interface{}      `json:"message"`
	Icon    interface{}      `json:"icon"`
}
