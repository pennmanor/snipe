package snipe

import "encoding/json"

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
	ExpectedCheckin  interface{} `json:"expected_checkin"`
	PurchaseCost     interface{} `json:"purchase_cost"`
	CheckinCounter   int         `json:"checkin_counter"`
	CheckoutCounter  int         `json:"checkout_counter"`
	RequestsCounter  int         `json:"requests_counter"`
	UserCanCheckout  bool        `json:"user_can_checkout"`
	CustomFields     interface{} `json:"custom_fields"`
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

type Users struct {
	Total int    `json:"total"`
	Rows  []User `json:"rows"`
}

type User struct {
	ID          int         `json:"id"`
	Avatar      string      `json:"avatar"`
	Name        string      `json:"name"`
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	Username    string      `json:"username"`
	EmployeeNum string      `json:"employee_num"`
	Manager     interface{} `json:"manager"`
	Jobtitle    interface{} `json:"jobtitle"`
	Phone       interface{} `json:"phone"`
	Address     interface{} `json:"address"`
	City        interface{} `json:"city"`
	State       interface{} `json:"state"`
	Country     interface{} `json:"country"`
	Zip         interface{} `json:"zip"`
	Email       string      `json:"email"`
	Department  interface{} `json:"department"`
	Location    struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"location"`
	Notes       string `json:"notes"`
	Permissions struct {
		Superuser             json.Number `json:"superuser,Number"`
		Admin                 json.Number `json:"admin,Number"`
		ReportsView           json.Number `json:"reports.view,Number"`
		AssetsView            json.Number `json:"assets.view,Number"`
		AssetsCreate          json.Number `json:"assets.create,Number"`
		AssetsEdit            json.Number `json:"assets.edit,Number"`
		AssetsDelete          json.Number `json:"assets.delete,Number"`
		AssetsCheckin         json.Number `json:"assets.checkin,Number"`
		AssetsCheckout        json.Number `json:"assets.checkout,Number"`
		AssetsAudit           json.Number `json:"assets.audit,Number"`
		AssetsViewRequestable json.Number `json:"assets.view.requestable,Number"`
		AccessoriesView       json.Number `json:"accessories.view,Number"`
		AccessoriesCreate     json.Number `json:"accessories.create,Number"`
		AccessoriesEdit       json.Number `json:"accessories.edit,Number"`
		AccessoriesDelete     json.Number `json:"accessories.delete,Number"`
		AccessoriesCheckout   json.Number `json:"accessories.checkout,Number"`
		AccessoriesCheckin    json.Number `json:"accessories.checkin,Number"`
		ConsumablesView       json.Number `json:"consumables.view,Number"`
		ConsumablesCreate     json.Number `json:"consumables.create,Number"`
		ConsumablesEdit       json.Number `json:"consumables.edit,Number"`
		ConsumablesDelete     json.Number `json:"consumables.delete,Number"`
		ConsumablesCheckout   json.Number `json:"consumables.checkout,Number"`
		LicensesView          json.Number `json:"licenses.view,Number"`
		LicensesCreate        json.Number `json:"licenses.create,Number"`
		LicensesEdit          json.Number `json:"licenses.edit,Number"`
		LicensesDelete        json.Number `json:"licenses.delete,Number"`
		LicensesCheckout      json.Number `json:"licenses.checkout,Number"`
		LicensesKeys          json.Number `json:"licenses.keys,Number"`
		ComponentsView        json.Number `json:"components.view,Number"`
		ComponentsCreate      json.Number `json:"components.create,Number"`
		ComponentsEdit        json.Number `json:"components.edit,Number"`
		ComponentsDelete      json.Number `json:"components.delete,Number"`
		ComponentsCheckout    json.Number `json:"components.checkout,Number"`
		ComponentsCheckin     json.Number `json:"components.checkin,Number"`
		UsersView             json.Number `json:"users.view,Number"`
		UsersCreate           json.Number `json:"users.create,Number"`
		UsersEdit             json.Number `json:"users.edit,Number"`
		UsersDelete           json.Number `json:"users.delete,Number"`
		ModelsView            json.Number `json:"models.view,Number"`
		ModelsCreate          json.Number `json:"models.create,Number"`
		ModelsEdit            json.Number `json:"models.edit,Number"`
		ModelsDelete          json.Number `json:"models.delete,Number"`
		CategoriesView        json.Number `json:"categories.view,Number"`
		CategoriesCreate      json.Number `json:"categories.create,Number"`
		CategoriesEdit        json.Number `json:"categories.edit,Number"`
		CategoriesDelete      json.Number `json:"categories.delete,Number"`
		DepartmentsView       json.Number `json:"departments.view,Number"`
		DepartmentsCreate     json.Number `json:"departments.create,Number"`
		DepartmentsEdit       json.Number `json:"departments.edit,Number"`
		DepartmentsDelete     json.Number `json:"departments.delete,Number"`
		StatuslabelsView      json.Number `json:"statuslabels.view,Number"`
		StatuslabelsCreate    json.Number `json:"statuslabels.create,Number"`
		StatuslabelsEdit      json.Number `json:"statuslabels.edit,Number"`
		StatuslabelsDelete    json.Number `json:"statuslabels.delete,Number"`
		CustomfieldsView      json.Number `json:"customfields.view,Number"`
		CustomfieldsCreate    json.Number `json:"customfields.create,Number"`
		CustomfieldsEdit      json.Number `json:"customfields.edit,Number"`
		CustomfieldsDelete    json.Number `json:"customfields.delete,Number"`
		SuppliersView         json.Number `json:"suppliers.view,Number"`
		SuppliersCreate       json.Number `json:"suppliers.create,Number"`
		SuppliersEdit         json.Number `json:"suppliers.edit,Number"`
		SuppliersDelete       json.Number `json:"suppliers.delete,Number"`
		ManufacturersView     json.Number `json:"manufacturers.view,Number"`
		ManufacturersCreate   json.Number `json:"manufacturers.create,Number"`
		ManufacturersEdit     json.Number `json:"manufacturers.edit,Number"`
		ManufacturersDelete   json.Number `json:"manufacturers.delete,Number"`
		DepreciationsView     json.Number `json:"depreciations.view,Number"`
		DepreciationsCreate   json.Number `json:"depreciations.create,Number"`
		DepreciationsEdit     json.Number `json:"depreciations.edit,Number"`
		DepreciationsDelete   json.Number `json:"depreciations.delete,Number"`
		LocationsView         json.Number `json:"locations.view,Number"`
		LocationsCreate       json.Number `json:"locations.create,Number"`
		LocationsEdit         json.Number `json:"locations.edit,Number"`
		LocationsDelete       json.Number `json:"locations.delete,Number"`
		CompaniesView         json.Number `json:"companies.view,Number"`
		CompaniesCreate       json.Number `json:"companies.create,Number"`
		CompaniesEdit         json.Number `json:"companies.edit,Number"`
		CompaniesDelete       json.Number `json:"companies.delete,Number"`
		SelfTwoFactor         json.Number `json:"self.two_factor,Number"`
		SelfAPI               json.Number `json:"self.api,Number"`
		SelfEditLocation      json.Number `json:"self.edit_location,Number"`
	} `json:"permissions"`
	Activated          bool `json:"activated"`
	TwoFactorActivated bool `json:"two_factor_activated"`
	AssetsCount        int  `json:"assets_count"`
	LicensesCount      int  `json:"licenses_count"`
	AccessoriesCount   int  `json:"accessories_count"`
	ConsumablesCount   int  `json:"consumables_count"`
	Company            struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"company"`
	CreatedAt struct {
		Datetime  string `json:"datetime"`
		Formatted string `json:"formatted"`
	} `json:"created_at"`
	UpdatedAt struct {
		Datetime  string `json:"datetime"`
		Formatted string `json:"formatted"`
	} `json:"updated_at"`
	LastLogin        interface{} `json:"last_login"`
	AvailableActions struct {
		Update  bool `json:"update"`
		Delete  bool `json:"delete"`
		Clone   bool `json:"clone"`
		Restore bool `json:"restore"`
	} `json:"available_actions"`
	Groups struct {
		Total int     `json:"total"`
		Rows  []Group `json:"rows"`
	} `json:"groups"`
}

type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Checkout struct {
	Status   string `json:"status"`
	Messages string `json:"messages"`
	Payload  struct {
		Asset string `json:"asset"`
	}
}

type Checkin struct {
	Status   string `json:"status"`
	Messages string `json:"messages"`
	Payload  struct {
		Asset string `json:"asset"`
	}
}

type Activities struct {
	Total int        `json:"total"`
	Rows  []Activity `json:"rows"`
}

type Activity struct {
	ID   int         `json:"id"`
	Icon string      `json:"icon"`
	File interface{} `json:"file"`
	Item struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"item"`
	Location struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"location"`
	CreatedAt struct {
		Datetime  string `json:"datetime"`
		Formatted string `json:"formatted"`
	} `json:"created_at"`
	UpdatedAt struct {
		Datetime  string `json:"datetime"`
		Formatted string `json:"formatted"`
	} `json:"updated_at"`
	NextAuditDate struct {
		Date      string `json:"date"`
		Formatted string `json:"formatted"`
	} `json:"next_audit_date"`
	DaysToNextAudit int    `json:"days_to_next_audit"`
	ActionType      string `json:"action_type"`
	Admin           struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"admin"`
	Target struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"target"`
	Note          interface{} `json:"note"`
	SignatureFile interface{} `json:"signature_file"`
	LogMeta       interface{} `json:"log_meta"`
}

type Locations struct {
	Total int        `json:"total"`
	Rows  []Location `json:"rows"`
}

type Location struct {
	ID                  int         `json:"id"`
	Name                string      `json:"name"`
	Image               interface{} `json:"image"`
	Address             interface{} `json:"address"`
	Address2            interface{} `json:"address2"`
	City                interface{} `json:"city"`
	State               interface{} `json:"state"`
	Country             interface{} `json:"country"`
	Zip                 interface{} `json:"zip"`
	AssignedAssetsCount int         `json:"assigned_assets_count"`
	AssetsCount         int         `json:"assets_count"`
	UsersCount          int         `json:"users_count"`
	Currency            interface{} `json:"currency"`
	CreatedAt           struct {
		Datetime  string `json:"datetime"`
		Formatted string `json:"formatted"`
	} `json:"created_at"`
	UpdatedAt struct {
		Datetime  string `json:"datetime"`
		Formatted string `json:"formatted"`
	} `json:"updated_at"`
	Parent           interface{}   `json:"parent"`
	Manager          interface{}   `json:"manager"`
	Children         []interface{} `json:"children"`
	AvailableActions struct {
		Update bool `json:"update"`
		Delete bool `json:"delete"`
	} `json:"available_actions"`
}

type Update struct {
	Status   string `json:"status"`
	Messages string `json:"messages"`
	Payload  struct {
		ID                        int         `json:"id"`
		Name                      string      `json:"name"`
		AssetTag                  string      `json:"asset_tag"`
		ModelID                   int         `json:"model_id"`
		Serial                    string      `json:"serial"`
		PurchaseDate              interface{} `json:"purchase_date"`
		PurchaseCost              interface{} `json:"purchase_cost"`
		OrderNumber               string      `json:"order_number"`
		AssignedTo                interface{} `json:"assigned_to"`
		Notes                     string      `json:"notes"`
		Image                     interface{} `json:"image"`
		UserID                    interface{} `json:"user_id"`
		CreatedAt                 string      `json:"created_at"`
		UpdatedAt                 string      `json:"updated_at"`
		Physical                  int         `json:"physical"`
		DeletedAt                 interface{} `json:"deleted_at"`
		StatusID                  int         `json:"status_id"`
		Archived                  int         `json:"archived"`
		WarrantyMonths            interface{} `json:"warranty_months"`
		Depreciate                interface{} `json:"depreciate"`
		SupplierID                interface{} `json:"supplier_id"`
		Requestable               int         `json:"requestable"`
		RtdLocationID             interface{} `json:"rtd_location_id"`
		SnipeitEthernetMac1       string      `json:"_snipeit_ethernet_mac_1"`
		Accepted                  interface{} `json:"accepted"`
		LastCheckout              interface{} `json:"last_checkout"`
		ExpectedCheckin           interface{} `json:"expected_checkin"`
		CompanyID                 int         `json:"company_id"`
		AssignedType              interface{} `json:"assigned_type"`
		LastAuditDate             string      `json:"last_audit_date"`
		NextAuditDate             string      `json:"next_audit_date"`
		LocationID                interface{} `json:"location_id"`
		CheckinCounter            int         `json:"checkin_counter"`
		CheckoutCounter           int         `json:"checkout_counter"`
		RequestsCounter           int         `json:"requests_counter"`
		SnipeitWirelessMac2       string      `json:"_snipeit_wireless_mac_2"`
		SnipeitPuppetRole3        string      `json:"_snipeit_puppet_role_3"`
		SnipeitPuppetEnvironment4 string      `json:"_snipeit_puppet_environment_4"`
		SnipeitDrcOu5             interface{} `json:"_snipeit_drc_ou_5"`
		SnipeitLoanerNumber6      interface{} `json:"_snipeit_loaner_number_6"`
	} `json:"payload"`
}

type UserPatch struct {
	Status   string `json:"status"`
	Messages string `json:"messages"`
	User     User   `json:"payload"`
}
