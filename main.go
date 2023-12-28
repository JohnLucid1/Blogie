package main

import (
	"blog/db"
	"blog/handlers"
	"log"

	"net/http"

	"github.com/labstack/echo/v4"
)

func TestConnection(c echo.Context) error {
	return c.String(http.StatusOK, "HAIIII")
}

// TODO: And write [IMAGE: parsing thingy]
// TODO: write image in posts how to do that
// TODO: maybe add post views
// TODO: Write functions for posts and so on

func main() {
	e := echo.New()
	DB, err := db.InitDatabase()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(DB.Ping())

	e.GET("/test", handlers.TestConnection(DB))
	e.POST("/like", handlers.Like(DB))
	e.POST("/dislike", handlers.Dislike(DB))
	e.POST("/login", handlers.Login(DB))
	e.POST("/create_author", handlers.CreateAuthor(DB))
	e.POST("/create_post/", handlers.CreatePost(DB))
	e.POST("/create_edit/", handlers.CreateEdit(DB))
	e.POST("/create_comment/", handlers.CreateComment(DB))
	e.GET("/all_authors/:user_id", handlers.GetAllAuthors(DB))
	e.Start("0.0.0.0:8080")
}

//| post_id | user_id | title          | content                                          | created_at          |
//|---------|---------|----------------|--------------------------------------------------|---------------------|
//| 1       | 1       | First Post     | Some text [IMAGE:1] some text [IMAGE:2] ...       | 2023-01-03 15:45:00 |
//| 2       | 2       | Second Post    | Some text [IMAGE:3] some text [IMAGE:4] ...       | 2023-01-04 09:15:00 |

//| image_id | post_id | image_url                |
//|----------|---------|--------------------------|
//| 1        | 1       | /images/post1_image1.jpg |
//| 2        | 1       | /images/post1_image2.jpg |
//| 3        | 2       | /images/post2_image1.jpg |
//| 4        | 2       | /images/post2_image2.jpg |
