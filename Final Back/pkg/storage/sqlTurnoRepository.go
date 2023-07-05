package storage

import (
	"FinalBack/internal/domain"
	"FinalBack/pkg/web"
	"database/sql"
	"fmt"
)

type SQLTurnoRepository struct {
	db *sql.DB
}

func NewSQLTurnoRepository(db *sql.DB) *SQLTurnoRepository {
	return &SQLTurnoRepository{db}
}

func (r *SQLTurnoRepository) Create(turno *domain.TurnoPost) error {
	query := "INSERT INTO turnos (dentistas_id, pacientes_id, fechayhora, Descripcion) VALUES (?, ?, ?, ?)"
	_, err := r.db.Exec(query, turno.DentistaId, turno.PacienteId, turno.FechaYHora, turno.Descripcion)
	if err != nil {
		return fmt.Errorf("could not create turno: %v", err)
	}
	return nil
}

func (r *SQLTurnoRepository) GetById(id int) (*domain.Turno, error) {
	query := `SELECT t.id, t.fechayhora, t.descripcion, p.dni as paciente_dni, p.nombre as paciente_nombre, d.matricula as dentista_matricula, d.nombre as dentista_nombre
								FROM turnos t 
								INNER JOIN pacientes p ON t.paciente_id = p.id 
								INNER JOIN dentistas d ON t.dentista_id = d.id 
								WHERE p.id = ?`
	row := r.db.QueryRow(query, id)

	turno := domain.Turno{}
	err := row.Scan(&turno.Id, &turno.FechaYHora, &turno.Descripcion, &turno.Dentista, &turno.Paciente)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, web.NewNotFoundApiError("Not found")
		}
		return nil, fmt.Errorf("could not get turno by id: %v", err)
	}

	return &turno, nil
}

func (r *SQLTurnoRepository) Update(turno *domain.TurnoPost) error {
	query := "UPDATE turnos SET dentista_id = ?, paciente_id = ?, fecha_y_hora = ?, descripcion = ? WHERE id = ?"
	_, err := r.db.Exec(query, turno.DentistaId, turno.PacienteId, turno.FechaYHora, turno.Descripcion, turno.Id)
	if err != nil {
		return fmt.Errorf("could not update turno: %v", err)
	}
	return nil
}

func (r *SQLTurnoRepository) UpdateField(id int, field string, value interface{}) error {
	query := fmt.Sprintf("UPDATE turnos SET %s = ? WHERE id = ?", field)
	_, err := r.db.Exec(query, value, id)
	if err != nil {
		return fmt.Errorf("could not update turno field: %v", err)
	}
	return nil
}

func (r *SQLTurnoRepository) Delete(id int) error {
	query := "DELETE FROM turnos WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("could not delete turno: %v", err)
	}
	return nil
}

func (repo *SQLTurnoRepository) CreateByDniAndMatricula(turno *domain.Turno) (*domain.Turno, error) {
	// Obtener paciente y dentista
	pacienteRepo := NewSQLPacienteRepository(repo.db)
	paciente, err := pacienteRepo.GetPacienteByDNI(turno.Paciente.Dni)
	if err != nil {
		return nil, err
	}

	dentistaRepo := NewSQLDentistRepository(repo.db)
	dentista, err := dentistaRepo.GetByMatricula(turno.Dentista.Matricula)
	if err != nil {
		return nil, err
	}

	// Insertar turno
	result, err := repo.db.Exec("INSERT INTO turnos (paciente_id, dentista_id, fecha_y_hora, descripcion) VALUES (?, ?, ?, ?)",
		paciente.Id, dentista.Id, turno.FechaYHora, turno.Descripcion)
	if err != nil {
		return nil, err
	}

	// Obtener ID del turno insertado
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Devolver el turno creado con el ID
	turno.Id = int(id)
	turno.Paciente = *paciente
	turno.Dentista = *dentista
	return turno, nil
}

func (repo *SQLTurnoRepository) GetTurnoByDniPaciente(dni string) ([]*domain.Turno, error) {
	rows, err := repo.db.Query(`SELECT t.id, t.fecha_y_hora, t.descripcion, p.dni as paciente_dni, p.nombre as paciente_nombre, d.matricula as dentista_matricula, d.nombre as dentista_nombre
								FROM turnos t 
								INNER JOIN pacientes p ON t.paciente_id = p.id 
								INNER JOIN dentistas d ON t.dentista_id = d.id 
								WHERE p.dni = ?`, dni)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	turnos := []*domain.Turno{}
	for rows.Next() {
		turno := &domain.Turno{}
		err = rows.Scan(&turno.Id, &turno.FechaYHora, &turno.Descripcion, &turno.Paciente.Dni, &turno.Paciente.Nombre, &turno.Dentista.Matricula, &turno.Dentista.Nombre)
		if err != nil {
			return nil, err
		}
		turnos = append(turnos, turno)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return turnos, nil
}

