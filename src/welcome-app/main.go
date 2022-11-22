package main

import (
	"fmt"
	"html/template"
	"time"
	"net/http"
	"encoding/json"

)

type Welcome struct {
	Name string
	Time string
}

type Intern struct {
	name string
	email string
	address string

}

type College struct {
	university string
	website string
	catchPhrase string

}

type Employer struct {
	Name string	`json:"name"`
	Email string	`json:"email"`	
	Phone int	`json:"phone"`
	Details Intern	`json:"details"`

}

func main(){
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}
	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))
	result := Employer {
		Name: "Leanne Graham",
		Email: "Sinere@arpil.biz",
		Phone: 17707368031,
		Details: Intern{"Chelsey Dietrich","Lucio_Hettinger@annie.ca", `Skiles Walks Apt 351 Roscoeview LA 33263`},
		
	}

	description := College {
		university: "Keebler University",
		website: "demarco.info",
		catchPhrase: "revolutionize end-to-end systems",
	}

	fmt.Println("Employer name: ", result.Name)
	fmt.Println("Employer's email: ", result.Email)
	fmt.Println("Phone number: ", result.Phone)

	fmt.Println("Intern name: ", result.Details.name)
	fmt.Println("Intern email: ", result.Details.email)
	fmt.Println("Intern address: ", result.Details.address)

	fmt.Println("College name: ", description.university)
	fmt.Println("College website: ", description.website)
	fmt.Println("College catchphrase: ", description.catchPhrase)

	http.Handle("/static/",
	http.StripPrefix("/static/",
		http.FileServer(http.Dir("static"))))

http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	if name := r.FormValue("name"); name != "" {
		welcome.Name = name
	}
	if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
})

http.HandleFunc("/Employer", func(w http.ResponseWriter, r *http.Request)  {
	json.NewEncoder(w).Encode(result)
})
fmt.Println("Listening")
fmt.Println(http.ListenAndServe(":8080", nil))

}
