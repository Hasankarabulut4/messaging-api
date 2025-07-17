package employee

type Message struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Message      string `json:"message"`
}

type DtoMessage struct {
	Username     string `json:"username"`
	Message      string `json:"message"`
}

type Error struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
}


type Data struct {
	Messages   []Message   `json:"messages"`
	NextId     int         `json:"nextId"`
}