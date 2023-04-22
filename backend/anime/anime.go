package anime

type Anime struct {
	ID          string `json:"id" dynamodbav:"id"`
	Name        string `json:"name"`
	Episodes    int    `json:"episodes"`
	Genre       string `json:"genre"`
	Autor       string `json:"author"`
	Description string `json:"description"`
	Image       string `json:"image"`
	IsActive    bool   `json:"is_active"`
	RelaseDate  string `json:"release_date"`
	CreateAt    string `json:"created_at"`
	UpdateAt    string `json:"updated_at"`
}
