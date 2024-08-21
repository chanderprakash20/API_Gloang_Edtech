package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

// Author represents the course's author
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB to store courses
var courses []Course

// Middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("API - LearnCodeOnline.in")
	r := mux.NewRouter()

	// Seeding some initial courses into the fake database
	courses = append(courses, Course{
		CourseId:    "1",
		CourseName:  "Introduction to Go",
		CoursePrice: 100,
		Author: &Author{
			Fullname: "John Doe",
			Website:  "https://johndoe.com",
		},
	})

	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "Advanced Go Programming",
		CoursePrice: 200,
		Author: &Author{
			Fullname: "Jane Smith",
			Website:  "https://janesmith.com",
		},
	})

	// Routing - defining API endpoints
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// Listen to a port and start the server
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Controllers - file

// serveHome is the controller for the home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by LearnCodeOnline</h1>"))
}

// getAllCourses returns all the courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// getOneCourse returns a single course based on the ID
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// Grab id from request
	params := mux.Vars(r)

	// Loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// Return a message if no course is found
	json.NewEncoder(w).Encode("No Course found with the given ID")
	return
}

// createOneCourse adds a new course to the fake database
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// Check if the body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// Decode the request body into a course struct
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// TODO: Check if the course title is duplicate (optional)

	// Generate a unique ID and append the new course into the courses slice
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

// updateOneCourse updates an existing course based on the ID
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// First - grab the ID from the request
	params := mux.Vars(r)

	// Loop through courses, find matching id, remove the old course, and add the updated course
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) // Remove the old course
			var newCourse Course
			_ = json.NewDecoder(r.Body).Decode(&newCourse) // Decode the new course data
			newCourse.CourseId = params["id"]              // Keep the same ID
			courses = append(courses, newCourse)           // Add the updated course
			json.NewEncoder(w).Encode(newCourse)
			return
		}
	}
	// Return a message if no course is found
	json.NewEncoder(w).Encode("No Course found with the given ID")
}

// deleteOneCourse removes a course based on the ID
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	// Grab the ID from the request
	params := mux.Vars(r)

	// Loop through courses, find matching id, and remove the course
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) // Remove the course
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}
	// Return a message if no course is found
	json.NewEncoder(w).Encode("No Course found with the given ID")
}
