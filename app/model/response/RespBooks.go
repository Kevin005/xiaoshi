package response

/**
返回books
 */
type RespBook struct {
	RespModel
	Data interface{}
}

type RespBookData struct {
	BookId     uint    `json:"book_id"`
	UserId     uint    `json:"user_id"`
	BookName   string `json:"book_name"`
	BookAuthor string `json:"book_author"`
}
