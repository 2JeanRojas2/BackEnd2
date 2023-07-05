package main

import ("github.com/gin-gonic/gin"
		"FinalBack/internal/dentista"
		"FinalBack/internal/paciente"
		"FinalBack/internal/turno"
		"FinalBack/cmd/server/handler"
		"FinalBack/pkg/storage"
		"log"
		_ "github.com/go-sql-driver/mysql"
		"database/sql")

func main() {

	datasource := "root:Atalamastruga22@tcp(localhost:3306)/parcialfinalback"
	DB, err := sql.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	// Creamos una nueva instancia del repositorio de dentistas
	dentistaRepo := storage.NewSQLDentistRepository(DB)
	// Creamos una nueva instancia del servicio de dentistas
	dentistaService := dentista.NewDentistService(dentistaRepo)
	// Creamos una nueva instancia del handler de dentistas
	dentistaHandler := handler.NewDentistaHandler(*dentistaService)

	// Creamos una nueva instancia del repositorio de dentistas
	pacienteRepo := storage.NewSQLPacienteRepository(DB)
	// Creamos una nueva instancia del servicio de dentistas
	pacienteService := paciente.NewPacienteService(pacienteRepo)
	// Creamos una nueva instancia del handler de dentistas
	pacienteHandler := handler.NewPacienteHandler(*pacienteService)

		// Creamos una nueva instancia del repositorio de turnos
	turnoRepo := storage.NewSQLTurnoRepository(DB)
	// Creamos una nueva instancia del servicio de turnos
	turnoService := turno.NewTurnoService(turnoRepo)
	// Creamos una nueva instancia del handler de turnos
	turnoHandler := handler.NewTurnoHandler(*turnoService)

	// Registramos las rutas en el router
	router := gin.Default()
	dentistaRouter := router.Group("/api")
	dentistaHandler.RegisterRoutes(dentistaRouter)

	pacienteRouter := router.Group("/api")
	pacienteHandler.RegisterRoutes(pacienteRouter)

	turnoRouter := router.Group("/api")
	turnoHandler.RegisterRoutes(turnoRouter)

	// Iniciamos el servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}