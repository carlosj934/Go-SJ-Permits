package models

type RawPermitData struct {
	Status          string `json:"Status"`
	Gxlocation      string `json:"gx_location"`
	ParcelNumber    string `json:"ASSESSORS_PARCEL_NUMBER"`
	Applicant       string `json:"APPLICANT"`
	OwnerName       string `json:"OWNERNAME"`
	Contractor      string `json:"CONTRACTOR"`
	FolderNumber    string `json:"FOLDERNUMBER"`
	FolderDesc      string `json:"FOLDERDESC"`
	FolderName      string `json:"FOLDERNAME"`
	SubTypeDesc     string `json:"SUBTYPEDESCRIPTION"`
	WorkDesc        string `json:"WORKDESCRIPTION"`
	PermitApprovals string `json:"PERMITAPPROVALS"`
	// will need to see the best type for the dates
	IssueDate       string `json:"ISSUEDATE"`
	FinalDate       string `json:"FINALDATE"`
	DwellingUnits   string `json:"DWELLINGUNITS"`
	PermitValuation string `json:"PERMITVALUATION"`
	SquareFootage   string `json:"SQUAREFOOTAGE"`
	FolderRSN       string `json:"FOLDERRSN"`
}

type WrapResult struct {
	Success bool        `json:"success"`
	Result  WrapRecords `json:"result"`
}

type WrapRecords struct {
	Total   int64           `json:"total"`
	Records []RawPermitData `json:"records"`
}
