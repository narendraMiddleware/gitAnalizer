// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"html/template"
// 	"io/ioutil"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// type User struct  {
// 	Id int 		 `json:"id"`
// 	Name string	 `json:"name"`
// 	Email string `json:"email"`
// 	Phone string `json:"phone"`
// }

// func main() {
// 	r := mux.NewRouter()

// 	r.HandleFunc("/", handler)
// 	r.HandleFunc("/json", jsonHandler)
// 	r.HandleFunc("/template", templateHandler)
// 	r.HandleFunc("/getuser/{id}",getUserDetail).Methods("GET")
// 	http.ListenAndServe(":8081", r)
// }

// // https://api.github.com/repos/twbs/bootstrap


// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello Worlddd!")
// } 

// func getUserDetail(w http.ResponseWriter, r *http.Request){

// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)

// 	a := params["id"]
	
// 	apiURL := "https://api.github.com/users/"

// 	a = a + apiURL

// 	fmt.Fprintf(w, a)


// 	url := fmt.Sprintf("%s%s", apiURL, params)

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	body, err := ioutil.ReadAll(resp.Body)

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	sb := string(body)
// 	log.Println(sb)
// 	w.Write(body)

// }

// //jsonHandler returns http respone in JSON format.
// func jsonHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	//We Read the response body on the line below.

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	//Convert the body to type string
	
// 	sb := string(body)
// 	log.Println(sb)

// 	// user := User{Id: 1, 
// 	// 			Name: "John Doe", 
// 	// 			Email: "johndoe@gmail.com", 
// 	// 			Phone: "000099999"}
	
// 	json.NewEncoder(w).Encode(sb)	
// }

// //templateHandler renders a template and returns as http response.
// func templateHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	t, err := template.ParseFiles("template.html")
// 	if err != nil {
// 		fmt.Fprintf(w, "Unable to load template")
// 	}

// 	user := User{Id: 1, 
// 		Name: "John Doe", 
// 		Email: "johndoe@gmail.com", 
// 		Phone: "000099999"}
// 	t.Execute(w, user)
// }






package main

import (
	"log"
	"os"


	"github.com/narendraMiddleware/gitAnalizer/github"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/get-basic-details/:id", gitlab.GetProjectDetails)
	app.Get("/get-language-details/:id", gitlab.GetLanguageDetails)
	app.Get("/get-member-details/:id", gitlab.GetMemberDetails)
	app.Get("/get-users-repo/:username",gitlab.GetRepoList)
}

func main() {
	err := godotenv.Load()

	app := fiber.New()

	app.Use(cors.New())
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	setupRoutes(app)
	app.Get("*", func(c *fiber.Ctx) error {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"message": "You're at wrong route",
		})
	})

	PORT := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + PORT))
}