package actions

import (
	"github.com/gobuffalo/buffalo"
	uuid "github.com/satori/go.uuid"
)

type Blog struct {
	AuthorName string `json:"name"`
	AuthorTitle string `json:"title"`
	ID uuid.UUID `json:"id"`
}

var db = make(map[uuid.UUID]Blog)

// var blogs = []Blog {
// 	Blog {Author_name: "Henry James", Author_title: "Daisy Miller", ID: 1},
// 	Blog {Author_name: "Thomas McCarthy", Author_title: "Bo Country for Old Men", ID: 2},
// }
// BlogList default implementation.
// func BlogList(c buffalo.Context) error {
// 	return c.Render(http.StatusOK, r.JSON(blogs))
// }

type BlogResource struct{}

func (br BlogResource) List(c buffalo.Context) error {
	return c.Render(200, r.JSON(db))
}

// Create Blog.
func (br BlogResource) Create(c buffalo.Context) error {
	// new Blog
	blog := &Blog{
		AuthorName:"",
		AuthorTitle:"",
		ID: uuid.NewV4(),
	}
	// add in database
	db[blog.ID] = *blog

	return c.Render(201, r.JSON(blog))
}

// To display a specific blog
func (br BlogResource) Show(c buffalo.Context) error {
	// get id and format to uuid
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		// if id isnt uuid
		return c.Render(500, r.String("id is not uuid v4"))
	}

	// get blog in database
	blog, ok := db[id]
	if ok {
		// if exist return blog
		return c.Render(200, r.JSON(blog))
	}

	// if not exist return not found
	return c.Render(404, r.String("blog not found"))
}

// BlogGetByID default implementation.
// func BlogGetByID(c buffalo.Context) error {
// 	params := mux.Vars(c.Request())
// 	id, _ := strconv.Atoi(params["id"])
// 	for _, blog := range blogs {
//         if blog.ID == id {
//             return c.Render(http.StatusOK, r.JSON(blog))
//         }
//     }
// 	return c.Render(http.StatusNotFound, nil)
// }
