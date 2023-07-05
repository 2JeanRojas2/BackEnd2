package storage

import (
	"FinalBack/internal/domain"
	"database/sql"

	"fmt"
)

type SQLDentistRepository struct {
    db *sql.DB
}

func NewSQLDentistRepository(db *sql.DB) *SQLDentistRepository {
    return &SQLDentistRepository{db}
}

func (repo *SQLDentistRepository) Create(dentist *domain.Dentista) error {
    query := "INSERT INTO dentistas (nombre, apellido, matricula) VALUES (?, ?, ?)"
    result, err := repo.db.Exec(query, dentist.Nombre, dentist.Apellido, dentist.Matricula)
    if err != nil {
        return fmt.Errorf("could not create dentist: %v", err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return fmt.Errorf("could not get last insert id: %v", err)
    }
    dentist.Id = int(id)

    return nil
}

func (repo *SQLDentistRepository) GetById(id int) (*domain.Dentista, error) {
    query := "SELECT id, nombre, apellido, matricula FROM dentistas WHERE id = ?"
    row := repo.db.QueryRow(query, id)

    dentist := domain.Dentista{}
    err := row.Scan(&dentist.Id, &dentist.Nombre, &dentist.Apellido, &dentist.Matricula)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("dentist not found")
        }
        return nil, fmt.Errorf("could not get dentist: %v", err)
    }

    return &dentist, nil
}

func (repo *SQLDentistRepository) GetByMatricula(matricula string) (*domain.Dentista, error) {
    query := "SELECT id, nombre, apellido, matricula FROM dentistas WHERE matricula = ?"
    row := repo.db.QueryRow(query, matricula)

    dentist := domain.Dentista{}
    err := row.Scan(&dentist.Id, &dentist.Nombre, &dentist.Apellido, &dentist.Matricula)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("dentist not found")
        }
        return nil, fmt.Errorf("could not get dentist: %v", err)
    }

    return &dentist, nil
}

func (repo *SQLDentistRepository) Update(dentist *domain.Dentista) error {
    query := "UPDATE dentistas SET nombre = ?, apellido = ?, matricula = ? WHERE id = ?"
    _, err := repo.db.Exec(query, dentist.Nombre, dentist.Apellido, dentist.Matricula, dentist.Id)
    if err != nil {
        return fmt.Errorf("could not update dentist: %v", err)
    }

    return nil
}

// UpdateDentistField actualiza un campo de un dentista específico
func (repo *SQLDentistRepository) UpdateDentistField(id int, field string, value interface{}) error {
	// Verificar si el campo a actualizar es válido
	switch field {
	case "nombre", "apellido", "matricula":
	default:
		return fmt.Errorf("el campo a actualizar no es válido")
	}

	// Crear sentencia SQL para actualizar el campo
	query := fmt.Sprintf("UPDATE dentistas SET %s = ? WHERE id = ?", field)

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
		return fmt.Errorf("no se encontró el dentista especificado")
	}

	return nil
}

func (s *SQLDentistRepository) Exist(codeValue string) bool{
	var exist bool
	var id int

	query := "SELECT id FROM dentistas WHERE code_value = ?;"
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

func (repo *SQLDentistRepository) Delete(id int) error {
    query := "DELETE FROM dentistas WHERE id = ?"
    result, err := repo.db.Exec(query, id)
    if err != nil {
        return fmt.Errorf("could not delete dentist: %v", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("could not get rows affected: %v", err)
    }
    if rowsAffected == 0 {
        return fmt.Errorf("no dentist with id %d found", id)
    }

    return nil
}