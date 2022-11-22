package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}


/*
import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"encoding/json"
	
)

 Create a struct that holds information to be displayed in our HTML file
type Welcome struct {
	Name string
	Time string
}

type JsonResponse struct{
	Value1 string `json:"key1"`
	Value2 string `json:"key2"`
	JsonNested JsonNested `json:"JsoneNested"`
}

type JsonNested struct {
	NestedValue1 string `json:"nestedkey1"`
	NestedValue2 string `json:"nestedkey2"`
}

 Go application entrypoint
 func main() {
	fetch("https://jsonplaceholder.ir/users");{
	
		,then(response => response.json())
		.then(json => console.log(json)),
	}
	Instantiate a Welcome struct object and pass in some random information.
	We shall get the name of the user as a query parameter from the URL
	 welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	We tell Go exactly where we can find our html file. We ask Go to parse the html file (Notice
	 the relative path). We wrap it in a call to template.Must() which handles any errors and halts if there are fatal errors

	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	nested := JsonNested{
		NestedValue1: "first nested value",
		NestedValue2: "second nested value",
	}

	JsonResp := JsonResponse{
		Value1: "some Data",
		Value2: "other Data",
		JsonNested: nested,
	}

	

	Our HTML comes with CSS that go needs to provide when we run the app. Here we tell go to create
	 a handle that looks in the static directory, go then uses the "/static/" as a url that our
	html can refer to when looking for our css and other files.

	http.Handle("/static/", //final url can be anything
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")))) Go looks in the relative "static" directory first using http.FileServer(), then matches it to a
	url of our choice as shown in http.Handle("/static/"). This url is what we need when referencing our css files
	once the server begins. Our html code would therefore be <link rel="stylesheet"  href="/static/stylesheet/...">
	It is important to note the url in http.Handle can be whatever we like, so long as we are consistent.

	This method takes in the URL path "/" and a function that takes in a response writer, and a http request.
	 **** THIS IS THE MAIN PATH / 
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		Takes the name from the URL query e.g ?name=Martin, will set welcome.Name = Martin.
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}
		If errors show an internal server error message
		I also pass the welcome struct to the welcome-template.html file.
		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/jsonResponse", func(w http.ResponseWriter, r *http.Request)  {
		json.NewEncoder(w).Encode(JsonResp)
	})

	Start the web server, set the port to listen to 8080. Without a path it assumes localhost
	Print any errors from starting the webserver using fmt
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
} 
*/