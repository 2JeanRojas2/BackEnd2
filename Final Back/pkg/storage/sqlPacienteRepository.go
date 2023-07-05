package storage

import (
	"FinalBack/internal/domain"
	"database/sql"

	"fmt"
)

type SQLPacienteRepository struct {
    db *sql.DB
}

func NewSQLPacienteRepository(db *sql.DB) *SQLPacienteRepository {
    return &SQLPacienteRepository{db}
}

func (repo *SQLPacienteRepository) Create(paciente *domain.Paciente) error {
    query := "INSERT INTO pacientes (nombre, apellido, domicilio, dni, fechadealta) VALUES (?, ?, ?, ?, ?)"
    result, err := repo.db.Exec(query, paciente.Nombre, paciente.Apellido, paciente.Domicilio, paciente.Dni, paciente.FechaDeAlta)
    if err != nil {
        return fmt.Errorf("could not create paciente: %v", err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return fmt.Errorf("could not get last insert id: %v", err)
    }
    paciente.Id = int(id)

    return nil
}

func (repo *SQLPacienteRepository) GetById(id int) (*domain.Paciente, error) {
    query := "SELECT id, nombre, apellido, domicilio, dni, fechadealta FROM pacientes WHERE id = ?"
    row := repo.db.QueryRow(query, id)

    paciente := domain.Paciente{}
    err := row.Scan(&paciente.Id, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio, &paciente.Dni, &paciente.FechaDeAlta)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("paciente not found")
        }
        return nil, fmt.Errorf("could not get paciente: %v", err)
    }

    return &paciente, nil
}

func (repo *SQLPacienteRepository) GetPacienteByDNI(dni string) (*domain.Paciente, error) {
    query := "SELECT id, nombre, apellido, domicilio, dni, fechadealta FROM pacientes WHERE dni = ?"
    row := repo.db.QueryRow(query, dni)

    paciente := domain.Paciente{}
    err := row.Scan(&paciente.Id, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio, &paciente.Dni, &paciente.FechaDeAlta)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("paciente not found")
        }
        return nil, fmt.Errorf("could not get paciente: %v", err)
    }

    return &paciente, nil
}

func (repo *SQLPacienteRepository) Update(paciente *domain.Paciente) error {
    query := "UPDATE pacientes SET nombre = ?, apellido = ?, domicilio = ?, dni = ?, fechadealta = ? WHERE id = ?"
    _, err := repo.db.Exec(query, paciente.Nombre, paciente.Apellido, paciente.Domicilio, paciente.Dni, paciente.FechaDeAlta, paciente.Id)
    if err != nil {
        return fmt.Errorf("could not update paciente: %v", err)
    }

    return nil
}

func (repo *SQLPacienteRepository) UpdatePacienteField(id int, field string, value interface{}) error {
	// Verificar si el campo a actualizar es válido
	switch field {
	case "nombre", "apellido", "domicilio", "dni", "fechaDeAlta":
	default:
		return fmt.Errorf("el campo a actualizar no es válido")
	}

	// Crear sentencia SQL para actualizar el campo
	query := fmt.Sprintf("UPDATE pacientes SET %s = ? WHERE id = ?", field)

	// Ejecutar sentencia SQL
	result, err := repo.db.Exec(query, value, id)
	if err != nil {
		return err
	}

	// Verificar si se actualizó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró el paciente especificado")
	}

	return nil
}

func (s *SQLPacienteRepository) Exist(codeValue string) bool{
	var exist bool
	var id int

	query := "SELECT id FROM pacientes WHERE code_value = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&id)
	if err != nil {
		return exist
	}

	if id > 0 {
		exist = true
	}

	return exist

}

func (repo *SQLPacienteRepository) Delete(id int) error {
    query := "DELETE FROM pacientes WHERE id = ?"
    result, err := repo.db.Exec(query, id)
    if err != nil {
        return fmt.Errorf("could not delete paciente: %v", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("could not get rows affected: %v", err)
    }
    if rowsAffected == 0 {
        return fmt.Errorf("no paciente with id %d found", id)
    }

    return nil
}