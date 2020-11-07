package domain

// URL domain model
type URL struct {
	ID        int    `json:"id" db:"id"`
	Origin    string `json:"origin" db:"origin"`
	Short     string `json:"short" db:"short"`
	IP        string `json:"ip" db:"ip"`
	UserAgent string `json:"agent" db:"agent"`
	Hits      int    `json:"hits" db:"hits"`
}

// CreateURL - DTO
type CreateURL struct {
	Origin    string `json:"origin" validate:"required,uri"`
	Short     string
	IP        string
	UserAgent string
	Hits      int
}

func (c *CreateURL) MapToModel() *URL {
	return &URL{
		Origin: c.Origin,
		Short:  c.Short,
	}
}

// URLRepository - repository
type URLRepository interface {
	FindByShort(string) (*URL, error)
	Create(*CreateURL) (*URL, error)
	Hit(int) error
}

// URLService - usecase
type URLService interface {
	FindOrigin(string) (string, error)
	Create(*CreateURL) (*URL, error)
}
