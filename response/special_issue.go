package response

type SpecialIssue struct {
	Abbreviation string  `json:"abbreviation"`
	CcfRanking   string  `json:"ccf_ranking"`
	IssueContent string  `json:"issue_content"`
	Description  string  `json:"description"`
	FullName     string  `json:"full_name"`
	ID           int64   `json:"id"`
	ImpactFactor float64 `json:"impact_factor"`
	Issn         string  `json:"issn"`
	JournalID    int64   `json:"journal_id"`
	Publisher    string  `json:"publisher"`
}
