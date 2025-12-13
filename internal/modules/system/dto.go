package system

type UpsertConfigRequest struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type ReportOverviewResponse struct {
	Users         int64 `json:"users"`
	Contents      int64 `json:"contents"`
	Events        int64 `json:"events"`
	Activities    int64 `json:"activities"`
	Submissions   int64 `json:"submissions"`
	Registrations int64 `json:"registrations"`
}
