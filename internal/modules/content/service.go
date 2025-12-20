package content

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{r}
}
