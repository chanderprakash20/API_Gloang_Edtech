# API_Gloang_Edtech
# API_Golang_Edtech

This project is a simple RESTful API built using Golang and the Gorilla Mux router. It provides functionality for managing a collection of courses, including creating, retrieving, updating, and deleting courses.

## Features

 Create a Course**: Add a new course to the collection.
 Retrieve All Courses**: Get a list of all available courses.
 Retrieve a Specific Course**: Get details of a single course by its ID.
 Update a Course**: Modify the details of an existing course.
 Delete a Course**: Remove a course from the collection.

## Endpoints

### 1. Get All Courses
  Endpoint : `/courses`
  Method : `GET`
  Description : Retrieves all courses in the collection.
  Response Example :
    ```json
    [
        {
            "courseid": "1",
            "coursename": "Introduction to Go",
            "price": 100,
            "author": {
                "fullname": "John Doe",
                "website": "https://johndoe.com"
            }
        },
        {
            "courseid": "2",
            "coursename": "Advanced Go Programming",
            "price": 200,
            "author": {
                "fullname": "Jane Smith",
                "website": "https://janesmith.com"
            }
        }
    ]
    ```

### 2. Get a Course by ID
 Endpoint : `/course/{id}`
 Method : `GET`
 Description : Retrieves a single course by its ID.
 Response Example :
    ```json
    {
        "courseid": "1",
        "coursename": "Introduction to Go",
        "price": 100,
        "author": {
            "fullname": "John Doe",
            "website": "https://johndoe.com"
        }
    }
    ```

### 3. Create a New Course
  Endpoint : `/course`
  Method : `POST`
  Description : Adds a new course to the collection.
  Request Body Example :
    ```json
    {
        "coursename": "Python for Beginners",
        "price": 150,
        "author": {
            "fullname": "Alice Johnson",
            "website": "https://alicejohnson.com"
        }
    }
    ```

### 4. Update a Course
 Endpoint : `/course/{id}`
 Method : `PUT`
 Description : Updates an existing course.
 Request Body Example :
    ```json
    {
        "coursename": "Introduction to Go (Updated)",
        "price": 120,
        "author": {
            "fullname": "John Doe",
            "website": "https://johndoe.com"
        }
    }
    ```

### 5. Delete a Course
   Endpoint : `/course/{id}`
   Method : `DELETE`
   Description : Deletes a course by its ID.

## Setup Instructions

1.  Clone the Repository :
    bash
   git clone https://github.com/chanderprakash20/API_Golang_Edtech.git
   cd API_Golang_Edtech
