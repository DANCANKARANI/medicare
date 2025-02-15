package routes

import (
	"github.com/dancankarani/medicare/api/routes/admin/doctor"
	"github.com/dancankarani/medicare/api/routes/admin/pharmacist"
	"github.com/dancankarani/medicare/api/routes/inventory"
	"github.com/dancankarani/medicare/api/routes/patient"
	"github.com/dancankarani/medicare/api/routes/role"
	"github.com/dancankarani/medicare/api/routes/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func RegisterEndpoints() {
	app := fiber.New()
    app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins, change this to specific origins in production
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization", 
	}))
	doctor.SetDoctorsRoutes(app)
    user.SetUserRoutes(app)
    role.SetRoleRoutes(app)
    patient.SetPatientRoutes(app)
    pharmacist.SetPharmacistRoutes(app)
	inventory.SetInventoryRoutes(app)
    //
    app.Listen(":8000")
}