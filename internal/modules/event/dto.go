package event

type CreateEventRequest struct {
	Title    string `json:"title" binding:"required"`
	Time     string `json:"time" binding:"required"`
	Location string `json:"location"`
	Image    string `json:"image"`
}

type UpdateEventRequest struct {
	Title string `json:"title"`
	Image string `json:"image"`
}

type ManualCheckinRequest struct {
	UserID uint `json:"user_id" binding:"required"`
}
