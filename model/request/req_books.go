package request

/**
请求的books
 */
type ReqBooks struct {
	Private   bool   `json:"is_private"`
	BookName  string `json:"book_name"`
	Author    string `json:"author"`
	Image     string `json:"image"`
	PageTotal int    `json:"page_total"`
}
