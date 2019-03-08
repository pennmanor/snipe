package snipe

type Asset struct {

	ID       int    `json:"id"`
	Name     string `json:"name"`
	AssetTag string `json:"asset_tag"`
	Serial   string `json:"serial"`
	Model    struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"model"`
	ModelNumber string      `json:"model_number"`
	Eol         interface{} `json:"eol"`
	StatusLabel struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		StatusType string `json:"status_type"`
		StatusMeta string `json:"status_meta"`
	} `json:"status_label"`
	Category struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
	Manufacturer struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"manufacturer"`
	Supplier    interface{} `json:"supplier"`
	Notes       interface{} `json:"notes"`
	OrderNumber interface{} `json:"order_number"`
	Company     interface{} `json:"company"`
	Location    struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"location"`
	RtdLocation interface{} `json:"rtd_location"`
	Image       interface{} `json:"image"`
	AssignedTo  struct {
		ID             int         `json:"id"`
		Username       string      `json:"username"`
		Name           string      `json:"name"`
		FirstName      string      `json:"first_name"`
		LastName       interface{} `json:"last_name"`
		EmployeeNumber interface{} `json:"employee_number"`
		Type           string      `json:"type"`
	} `json:"assigned_to"`
	WarrantyMonths  interface{} `json:"warranty_months"`
	WarrantyExpires interface{} `json:"warranty_expires"`
	CreatedAt       struct {
		Datetime  string `json:"datetime"`
		Formatted string `json:"formatted"`
	} `json:"created_at"`
	UpdatedAt struct {
		Datetime  string `json:"datetime"`
		Formatted string `json:"formatted"`
	} `json:"updated_at"`
	LastAuditDate interface{} `json:"last_audit_date"`
	NextAuditDate interface{} `json:"next_audit_date"`
	DeletedAt     interface{} `json:"deleted_at"`
	PurchaseDate  interface{} `json:"purchase_date"`
	LastCheckout  struct {
		Datetime  string `json:"datetime"`
		Formatted string `json:"formatted"`
	} `json:"last_checkout"`
	ExpectedCheckin  interface{}   `json:"expected_checkin"`
	PurchaseCost     interface{}   `json:"purchase_cost"`
	CheckinCounter   int           `json:"checkin_counter"`
	CheckoutCounter  int           `json:"checkout_counter"`
	RequestsCounter  int           `json:"requests_counter"`
	UserCanCheckout  bool          `json:"user_can_checkout"`
	AvailableActions struct {
		Checkout bool `json:"checkout"`
		Checkin  bool `json:"checkin"`
		Clone    bool `json:"clone"`
		Restore  bool `json:"restore"`
		Update   bool `json:"update"`
		Delete   bool `json:"delete"`
	} `json:"available_actions"`

}

type Assets struct {
	Total int     `json:"total"`
	Rows  []Asset `json:"rows"`
}

type Checkin struct {
	Status		string	`json:"status"`
	Messages	string	`json:"messages"`
	Payload	struct {
		Asset	string	`json:"asset"`
	}
}
