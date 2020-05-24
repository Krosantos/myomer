package manager

// Army -- A template for loading preset units
type Army struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	Cohort    string `json:"cohort"`
	Auxiliary string `json:"auxiliary"`
}
