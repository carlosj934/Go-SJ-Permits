package enricher

import(
	"time"
	"strings"
	"strconv"

	"permit-proxy/internal/models"
)

func Enrich(rawPermit []models.RawPermitData) ([]models.Permit, error) {
	e := make([]models.Permit, 0)

	for _, p := range rawPermit {
		issueDate, err := parseDate(p.IssueDate)
		if err != nil {
			return nil, err
		}
		
		finalDate, err := parseDate(p.FinalDate)
		if err != nil {
			return nil, err
		}

		zipCode := extractZip(p.Gxlocation)

		dwellingUnits, err := parseInt(p.DwellingUnits)
		if err != nil {
			return nil, err
		}

		permitValuation, err := parseFloat(p.PermitValuation)
		if err != nil {
			return nil, err
		}

		squareFootage, err := parseInt(p.SquareFootage)
		if err != nil {
			return nil, err
		}

		permit := models.Permit{
				Status: p.Status,
				Zipcode: zipCode,
				ParcelNumber: p.ParcelNumber,
				Applicant: p.Applicant,
				OwnerName: p.OwnerName,
				Contractor: p.Contractor,
				FolderNumber: p.FolderNumber,
				FolderDesc: p.FolderDesc,
				FolderName: p.FolderName,
				SubTypeDesc: p.SubTypeDesc,
				WorkDesc: p.WorkDesc,
				PermitApprovals: p.PermitApprovals,
				IssueDate: issueDate,
				FinalDate: finalDate,
				DwellingUnits: dwellingUnits,
				PermitValuation: permitValuation, 
				SquareFootage: squareFootage,
				FolderRSN: p.FolderRSN,
		}

		e = append(e, permit)
	}

	return e, nil
}

func parseDate(s string) (time.Time, error) {
	if s == "" {
		return time.Time{}, nil
	}

	t, err := time.Parse("1/2/2006 3:04:05 PM", s)
	if err != nil {
		return time.Time{}, err
	}
	
	return t, nil
}

func extractZip(gxLocation string) string {
	z := strings.Fields(gxLocation)

	if len(z) == 0 {
		return ""
	}
	return z[len(z)-1]
}

func parseFloat(s string) (float64, error) {
	if s == "" {
		return 0, nil
	}

	v, err := strconv.ParseFloat(s, 64)

	if err != nil {
		return 0, err
	}

	return v, nil
}

func parseInt(s string) (int64, error) {
	if s == "" {
		return 0, nil
	}

	i, err := strconv.ParseInt(s, 10, 64)
	
	if err != nil {
		return 0, err
	}

	return i, nil
}
