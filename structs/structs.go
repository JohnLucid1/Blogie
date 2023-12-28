package structs

type Author struct {
	AuthorID     int    `json:"author_id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Is_admin     bool   `json:"is_admin"`
	Picture_path string `json:"picture_path"`
}

type Image struct {
	Image_id   int    `json:"image_id"`
	Post_id    int    `json:"post_id"` // References post id
	Image_path string `json:"image_path"`
}

type Category struct {
	Category_id   int    `json:"category_id"`
	Category_name string `json:"category_name"`
}

type Post struct {
	Post_id      int      `json:"post_id"`
	Title        string   `json:"title"`
	Subtitle     string   `json:"subtitle"`
	Categories   []string `json:"categories"`
	Author_id    int      `json:"author_id"`
	Content      string   `json:"content"`
	Likes        int      `json:"likes"`
	Dislikes     int      `json:"dislikes"`
	Date_created string   `json:"date_created"`
}

type Comment struct {
	Comment_id      int    `json:"comment_id"`
	Author_id       int    `json:"author_id"`
	Foreign_post_id int    `json:"foreign_post_id"`
	Content         string `json:"content"`
	Likes           int    `json:"lieks"`
	Dislikes        int    `json:"dislikes"`
}
