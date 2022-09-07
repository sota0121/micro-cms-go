// ====================================
// Model for API
// ====================================

package mcmslib

type BasicResponse struct {
	Message string `json:"message"`
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Tag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Tags []Tag

type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  int    `json:"user_id"`
	Tags    Tags   `json:"tags"`
}
