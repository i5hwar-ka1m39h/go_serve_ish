package responsemodels



type ShowRes struct{
	Title string `json:"title"`
	Description string `json:"desc"`
	Rating int16 `json:"rating"`
	Genre []string `json:"genre"`
}