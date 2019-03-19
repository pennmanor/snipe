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

type Checkout struct {
	Status		string	`json:"status"`
	Messages	string	`json:"messages"`
	Payload	struct {
		Asset	string	`json:"asset"`
	}
}

type Checkin struct {
	Status		string	`json:"status"`
	Messages	string	`json:"messages"`
	Payload	struct {
		Asset	string	`json:"asset"`
	}
}

type Users struct {
	Total int     `json:"total"`
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
		Superuser             string `json:"superuser"`
		Admin                 string `json:"admin"`
		ReportsView           string `json:"reports.view"`
		AssetsView            string `json:"assets.view"`
		AssetsCreate          string `json:"assets.create"`
		AssetsEdit            string `json:"assets.edit"`
		AssetsDelete          string `json:"assets.delete"`
		AssetsCheckin         string `json:"assets.checkin"`
		AssetsCheckout        string `json:"assets.checkout"`
		AssetsAudit           string `json:"assets.audit"`
		AssetsViewRequestable string `json:"assets.view.requestable"`
		AccessoriesView       string `json:"accessories.view"`
		AccessoriesCreate     string `json:"accessories.create"`
		AccessoriesEdit       string `json:"accessories.edit"`
		AccessoriesDelete     string `json:"accessories.delete"`
		AccessoriesCheckout   string `json:"accessories.checkout"`
		AccessoriesCheckin    string `json:"accessories.checkin"`
		ConsumablesView       string `json:"consumables.view"`
		ConsumablesCreate     string `json:"consumables.create"`
		ConsumablesEdit       string `json:"consumables.edit"`
		ConsumablesDelete     string `json:"consumables.delete"`
		ConsumablesCheckout   string `json:"consumables.checkout"`
		LicensesView          string `json:"licenses.view"`
		LicensesCreate        string `json:"licenses.create"`
		LicensesEdit          string `json:"licenses.edit"`
		LicensesDelete        string `json:"licenses.delete"`
		LicensesCheckout      string `json:"licenses.checkout"`
		LicensesKeys          string `json:"licenses.keys"`
		ComponentsView        string `json:"components.view"`
		ComponentsCreate      string `json:"components.create"`
		ComponentsEdit        string `json:"components.edit"`
		ComponentsDelete      string `json:"components.delete"`
		ComponentsCheckout    string `json:"components.checkout"`
		ComponentsCheckin     string `json:"components.checkin"`
		UsersView             string `json:"users.view"`
		UsersCreate           string `json:"users.create"`
		UsersEdit             string `json:"users.edit"`
		UsersDelete           string `json:"users.delete"`
		ModelsView            string `json:"models.view"`
		ModelsCreate          string `json:"models.create"`
		ModelsEdit            string `json:"models.edit"`
		ModelsDelete          string `json:"models.delete"`
		CategoriesView        string `json:"categories.view"`
		CategoriesCreate      string `json:"categories.create"`
		CategoriesEdit        string `json:"categories.edit"`
		CategoriesDelete      string `json:"categories.delete"`
		DepartmentsView       string `json:"departments.view"`
		DepartmentsCreate     string `json:"departments.create"`
		DepartmentsEdit       string `json:"departments.edit"`
		DepartmentsDelete     string `json:"departments.delete"`
		StatuslabelsView      string `json:"statuslabels.view"`
		StatuslabelsCreate    string `json:"statuslabels.create"`
		StatuslabelsEdit      string `json:"statuslabels.edit"`
		StatuslabelsDelete    string `json:"statuslabels.delete"`
		CustomfieldsView      string `json:"customfields.view"`
		CustomfieldsCreate    string `json:"customfields.create"`
		CustomfieldsEdit      string `json:"customfields.edit"`
		CustomfieldsDelete    string `json:"customfields.delete"`
		SuppliersView         string `json:"suppliers.view"`
		SuppliersCreate       string `json:"suppliers.create"`
		SuppliersEdit         string `json:"suppliers.edit"`
		SuppliersDelete       string `json:"suppliers.delete"`
		ManufacturersView     string `json:"manufacturers.view"`
		ManufacturersCreate   string `json:"manufacturers.create"`
		ManufacturersEdit     string `json:"manufacturers.edit"`
		ManufacturersDelete   string `json:"manufacturers.delete"`
		DepreciationsView     string `json:"depreciations.view"`
		DepreciationsCreate   string `json:"depreciations.create"`
		DepreciationsEdit     string `json:"depreciations.edit"`
		DepreciationsDelete   string `json:"depreciations.delete"`
		LocationsView         string `json:"locations.view"`
		LocationsCreate       string `json:"locations.create"`
		LocationsEdit         string `json:"locations.edit"`
		LocationsDelete       string `json:"locations.delete"`
		CompaniesView         string `json:"companies.view"`
		CompaniesCreate       string `json:"companies.create"`
		CompaniesEdit         string `json:"companies.edit"`
		CompaniesDelete       string `json:"companies.delete"`
		SelfTwoFactor         string `json:"self.two_factor"`
		SelfAPI               string `json:"self.api"`
		SelfEditLocation      string `json:"self.edit_location"`
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
	Groups interface{} `json:"groups"`
}

type Activity struct {
	Total int `json:"total"`
	Rows  []struct {
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
	} `json:"rows"`
}

type AssetList struct {
	Total int `json:"total"`
	Rows  []struct {
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
		Company     struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"company"`
		Location struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"location"`
		RtdLocation struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"rtd_location"`
		Image           interface{} `json:"image"`
		AssignedTo      interface{} `json:"assigned_to"`
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
		ExpectedCheckin interface{} `json:"expected_checkin"`
		PurchaseCost    interface{} `json:"purchase_cost"`
		CheckinCounter  int         `json:"checkin_counter"`
		CheckoutCounter int         `json:"checkout_counter"`
		RequestsCounter int         `json:"requests_counter"`
		UserCanCheckout bool        `json:"user_can_checkout"`
		CustomFields    struct {
			EthernetMAC struct {
				Field       string `json:"field"`
				Value       string `json:"value"`
				FieldFormat string `json:"field_format"`
			} `json:"Ethernet MAC"`
			WirelessMAC struct {
				Field       string `json:"field"`
				Value       string `json:"value"`
				FieldFormat string `json:"field_format"`
			} `json:"Wireless MAC"`
			PuppetRole struct {
				Field       string      `json:"field"`
				Value       interface{} `json:"value"`
				FieldFormat string      `json:"field_format"`
			} `json:"puppet_role"`
			PuppetEnvionment struct {
				Field       string      `json:"field"`
				Value       interface{} `json:"value"`
				FieldFormat string      `json:"field_format"`
			} `json:"puppet_envionment"`
		} `json:"custom_fields"`
		AvailableActions struct {
			Checkout bool `json:"checkout"`
			Checkin  bool `json:"checkin"`
			Clone    bool `json:"clone"`
			Restore  bool `json:"restore"`
			Update   bool `json:"update"`
			Delete   bool `json:"delete"`
		} `json:"available_actions"`
	} `json:"rows"`
}
