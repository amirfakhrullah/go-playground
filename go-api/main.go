package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course
type Course struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Name    string `json:"name"`
	Website string `json:"website"`
}

// fake DB
var courses []Course

// middleware/helper
func (c *Course) IsEmpty() bool {
	return c.Title == ""
}

// controllers

// routes handler
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to this API!</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// grab id for request
	params := mux.Vars(r)
	courseId := params["id"]

	if len(courseId) == 0 {
		json.NewEncoder(w).Encode("ID is required")
		return
	}

	for _, course := range courses {
		if course.Id == courseId {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Course contents are required")
		return
	}

	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if (err != nil || course.IsEmpty()) {
		json.NewEncoder(w).Encode("Data provided are improper")
		return
	}

	// generate ID
	rand.Seed(time.Now().UnixNano())
	course.Id = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	courseId := params["id"]
	if len(courseId) == 0 {
		json.NewEncoder(w).Encode("ID is required")
		return
	}

	var updatedCourse Course
	err := json.NewDecoder(r.Body).Decode(&updatedCourse)
	if (err != nil || updatedCourse.IsEmpty()) {
		json.NewEncoder(w).Encode("Data provided are improper")
		return
	}

	for index, course := range courses {
		if course.Id == courseId {
			courses = append(courses[:index], courses[index+1:]...)
			updatedCourse.Id = courseId
			break
		}
	}
	courses = append(courses, updatedCourse)
	json.NewEncoder(w).Encode(updatedCourse)
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	courseId := params["id"]
	if len(courseId) == 0 {
		json.NewEncoder(w).Encode("ID is required")
		return
	}

	for index, course := range courses {
		if course.Id == courseId {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Successfully Deleted")
		}
	}
}

func main() {
	r := mux.NewRouter()

	// seeding dummy data
	courses = append(courses, Course{
		Id:    "123",
		Title: "TypeScript",
		Price: 2.99,
		Author: &Author{
			Name:    "Amir",
			Website: "amrf.me",
		},
	})

	// routes ======================
	// GET
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getCourseById).Methods("GET")
	// POST
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	// PUT
	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	// DELETE
	r.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
