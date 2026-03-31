package models

import (
	"time"
)

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
	IssueDate       string `json:"ISSUEDATE"`
	FinalDate       string `json:"FINALDATE"`
	DwellingUnits   string `json:"DWELLINGUNITS"`
	PermitValuation string `json:"PERMITVALUATION"`
	SquareFootage   string `json:"SQUAREFOOTAGE"`
	FolderRSN       string `json:"FOLDERRSN"`
}

type Permit struct {
	Status string `json:"status"`
	Zipcode string `json:"zipcode"`
	ParcelNumber string `json:"parcel_number"`
	Applicant string `json:"applicant"`
	OwnerName string `json:"owner_name"`
	Contractor string `json:"contractor"`
	FolderNumber string `json:"folder_number"`
	FolderDesc string `json:"folder_desc"`
	FolderName string `json:"folder_name"`
	SubTypeDesc string `json:"subtype_desc"`
	WorkDesc string `json:"work_desc"`
	PermitApprovals string `json:"permit_approvals"`
	IssueDate time.Time `json:"issue_date"`
	FinalDate time.Time `json:"final_date"`
	DwellingUnits int64 `json:"dwelling_units"`
	PermitValuation float64 `json:"permit_valuation"`
	SquareFootage int64 `json:"square_footage"`
	FolderRSN string `json:"folder_rsn"`
}

type WrapResult struct {
	Success bool        `json:"success"`
	Result  WrapRecords `json:"result"`
}

type WrapRecords struct {
	Total   int64           `json:"total"`
	Records []RawPermitData `json:"records"`
}
