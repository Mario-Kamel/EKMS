package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Mario-Kamel/EKMS/pkg/controllers"
	"github.com/Mario-Kamel/EKMS/pkg/repositories"
	"github.com/Mario-Kamel/EKMS/pkg/service"
)

func main() {
	gotenv.Load("./.env")
	client := getClient()
	defer client.Disconnect(context.Background())

	// personRepo := repositories.NewPersonRepo(client)
	// serviceRepo := repositories.NewServiceRepo(client)

	//Create a Person
	// p, err := personRepo.CreatePerson(context.Background(), models.Person{
	// 	Name:     "Mario Kamel",
	// 	Birthday: time.Date(1999, 10, 11, 0, 0, 0, 0, time.UTC),
	// 	Phone:    "01206032004",
	// 	Address:  "Cairo, Egypt",
	// 	Fr:       "Timo",
	// 	Degree:   "BSc",
	// })

	// fmt.Println("Created this\n", p, err)

	//Get all Persons
	// persons, err := personRepo.GetAllPersons(context.Background())
	// fmt.Println("All persons\n", persons, err)

	//Get a Person by ID
	// person, err := personRepo.GetPersonById(context.Background(), "6546075376a3e3d86900bdc7")
	// fmt.Println("Person\n", person, err)

	//Update a Person
	// 	person, err := personRepo.UpdatePerson(context.Background(), "6546075376a3e3d86900bdc7", models.Person{
	// 		Name:     "Mario Medhat",
	// 		Birthday: time.Date(1999, 10, 11, 0, 0, 0, 0, time.UTC),
	// 		Phone:    "01206032004",
	// 		Address:  "Alex, Arm",
	// 		Fr:       "Timo",
	// 		Degree:   "BSc of Engineering",
	// 	})
	// 	fmt.Println("Updated this\n", person, err)

	//Delete a Person
	// err := personRepo.DeletePerson(context.Background(), "6546075376a3e3d86900bdc7")
	// fmt.Println("Deleted this\n", err)

	// Create a Service
	// s, err := serviceRepo.CreateService(context.Background(), models.Service{
	// 	Date:         time.Date(2020, 10, 11, 0, 0, 0, 0, time.UTC),
	// 	Subject:      "Test",
	// 	Speaker:      "Mario Kamel",
	// 	BibleChapter: "John 3:16",
	// })
	// fmt.Println("Created this\n", s, err)

	//Get all Services
	// services, err := serviceRepo.GetAllServices(context.Background())
	// fmt.Println("All services\n", services, err)

	// Get a Service by ID
	// service, err := serviceRepo.GetServiceById(context.Background(), "65461a951196e8ebdc5135cf")
	// fmt.Println("Service\n", service, err)

	// Update a Service
	// service, err := serviceRepo.UpdateService(context.Background(), "65461a951196e8ebdc5135cf", models.Service{
	// 	Date:         time.Date(2023, 11, 3, 0, 0, 0, 0, time.UTC),
	// 	Subject:      "Updated",
	// 	Speaker:      "Mona Mounir",
	// 	BibleChapter: "John 3:16",
	// })
	// fmt.Println("Updated this\n", service, err)

	// Delete a Service
	// err := serviceRepo.DeleteService(context.Background(), "65461a951196e8ebdc5135cf")
	// fmt.Println("Deleted this\n", err)

	// Add an Attendant
	// sid, _ := primitive.ObjectIDFromHex("65462077c2f9683d8fc9c555")
	// pid, err := primitive.ObjectIDFromHex("65460fe85230151703f840b1")
	// if err != nil {
	// 	fmt.Println("Error while converting ID to ObjectID", err)
	// }
	// service, err := serviceRepo.AddAttendanceRecord(context.Background(), sid, models.AttendanceRecord{
	// 	PersonID: pid,
	// 	Time:     time.Now(),
	// 	Status:   "Present",
	// })

	// fmt.Println("Added this\n", service, err)

	personRepo := repositories.NewPersonRepo(client)
	personService := service.NewPersonService(personRepo)
	personController := controllers.NewPersonController(personService)

	serviceRepo := repositories.NewServiceRepo(client)
	serviceService := service.NewServiceService(serviceRepo)
	serviceController := controllers.NewServiceController(serviceService)

	assignmentRepo := repositories.NewAssignmentRepo(client)
	assignmentService := service.NewAssignmentService(assignmentRepo)
	assignmentController := controllers.NewAssignmentController(assignmentService)

	r := mux.NewRouter()
	r.HandleFunc("/persons", personController.GetAllPersons).Methods("GET")
	r.HandleFunc("/persons/{id}", personController.GetPersonById).Methods("GET")
	r.HandleFunc("/persons", personController.CreatePerson).Methods("POST")
	r.HandleFunc("/persons/{id}", personController.UpdatePerson).Methods("PUT")
	r.HandleFunc("/persons/{id}", personController.DeletePerson).Methods("DELETE")

	r.HandleFunc("/services", serviceController.GetAllServices).Methods("GET")
	r.HandleFunc("/services/{id}", serviceController.GetServiceById).Methods("GET")
	r.HandleFunc("/services", serviceController.CreateService).Methods("POST")
	r.HandleFunc("/services/{id}", serviceController.UpdateService).Methods("PUT")
	r.HandleFunc("/services/{id}", serviceController.DeleteService).Methods("DELETE")

	r.HandleFunc("/services/{id}/attendance", serviceController.AddAttendanceRecord).Methods("POST")
	r.HandleFunc("/services/{id}/attendance", serviceController.EditAttendanceRecord).Methods("PUT")
	r.HandleFunc("/services/{id}/attendance", serviceController.DeleteAttendanceRecord).Methods("DELETE")

	r.HandleFunc("/assignments", assignmentController.GetAllAssignments).Methods("GET")
	r.HandleFunc("/assignments/{id}", assignmentController.GetAssignmentById).Methods("GET")
	r.HandleFunc("/assignments", assignmentController.CreateAssignment).Methods("POST")
	r.HandleFunc("/assignments/{id}", assignmentController.UpdateAssignment).Methods("PUT")
	r.HandleFunc("/assignments/{id}", assignmentController.DeleteAssignment).Methods("DELETE")

	server := http.Server{
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         ":8080",
		Handler:      r,
	}

	log.Fatal(server.ListenAndServe())
}

func getClient() *mongo.Client {
	fmt.Println("Connecting to MongoDB...", os.Getenv("MONGO_URI"))
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}
