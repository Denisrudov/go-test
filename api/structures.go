package api


const (
	RANDOM_NAME string = "http://uinames.com/api/"
	CHUCKNORRIS_JOKE string = "http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=[nerdy]"
)

type Name struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Gender  string `json:"gender"`
	Region  string `json:"region"`
}

type ChuckNorris struct {
	Type string `json:"type"`
	Value struct {
		ID         int `json:"id"`
		Joke       string `json:"joke"`
		Categories []string `json:"categories"`
	} `json:"value"`
}

type Result struct {
	Name        Name `json:"randomName"`
	ChuckNorris ChuckNorris `json:"chuckPhrase"`
}
