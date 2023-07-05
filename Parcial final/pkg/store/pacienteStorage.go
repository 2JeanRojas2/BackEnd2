package store

import (
	"database/sql"
	"parcialfinal/internal/domain"
	"parcialfinal/pkg/web"
)

type PacienteStorage struct {
	db *sql.DB
}

func NewPacienteStorage(db *sql.DB) *PacienteStorage {
	return &PacienteStorage{
		db: db,
	}
}

func (s *PacienteStorage) GetOne(id int) (*domain.Paciente, error) {
	var paciente domain.Paciente
	row := s.db.QueryRow("SELECT id, nombre, apellido, domicilio, dni, fecha_de_alta FROM pacientes WHERE id = ?", id)
	err := row.Scan(&paciente.Id, 
					&paciente.Nombre, 
					&paciente.Apellido, 
					&paciente.Domicilio, 
					&paciente.Dni, 
					&paciente.FechaDeAlta)
	if err != nil {
		return nil, err
	}
	return &paciente, nil
}

func (s *PacienteStorage) Create(paciente domain.Paciente) (*domain.Paciente, error) {
	result, err := s.db.Exec("INSERT INTO pacientes(nombre, apellido, domicilio, dni, fecha_de_alta) VALUES (?, ?, ?, ?, ?)",
		paciente.Nombre, paciente.Apellido, paciente.Domicilio, paciente.Dni, paciente.FechaDeAlta)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	paciente.Id = int(id)

	return &paciente, nil
}

func (s *PacienteStorage) Exist(dni string) bool {
	var count int
	s.db.QueryRow("SELECT COUNT(*) FROM pacientes WHERE dni = ?", dni).Scan(&count)
	return count > 0
}

func (s *PacienteStorage) UpdateOne(paciente domain.Paciente) error {
	_, err := s.db.Exec("UPDATE pacientes SET nombre = ?, apellido = ?, domicilio = ?, dni = ?, fecha_de_alta = ? WHERE id = ?",
		paciente.Nombre, paciente.Apellido, paciente.Domicilio, paciente.Dni, paciente.FechaDeAlta, paciente.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *PacienteStorage) DeleteOne(id int) error {
	result, err := s.db.Exec("DELETE FROM pacientes WHERE id = ?", id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return web.NewBadRequestApiError("failed to delete")
	}
	return nil
}


