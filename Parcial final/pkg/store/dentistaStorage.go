package store

import ( "database/sql"
"parcialfinal/internal/domain")

type dentistaStorage struct {
	DB *sql.DB
}

func NewdentistaStorage(db *sql.DB) *dentistaStorage {
	return &dentistaStorage{DB: db}
}

func (s *dentistaStorage) GetOne(id int) (*domain.Dentista, error) {
	var dentistReturn domain.Dentista

	query := "SELECT * FROM dentistas WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&dentistReturn.Id, 
					&dentistReturn.Nombre, 
					&dentistReturn.Apellido,
					&dentistReturn.Matricula)
	if err != nil {
		return nil, err
	}
	return &dentistReturn, nil
}


func (s *dentistaStorage) Create(dentista domain.Dentista) (*domain.Dentista, error){
	query := "INSERT INTO dentistas (nombre, apellido, matricula) VALUES (?, ?, ?);"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(dentista.Nombre, dentista.Apellido, dentista.Matricula)
	if err != nil {
		return nil, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}

	lid, _ := res.LastInsertId()
	dentista.Id = int(lid)
	return &dentista, nil
}

func (s *dentistaStorage) Exist(codeValue string) bool{
	var exist bool
	var id int

	query := "SELECT id FROM dentistas WHERE code_value = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&id)
	if err != nil {
		return exist
	}

	if id > 0 {
		exist = true
	}

	return exist

}

func(s *dentistaStorage) UpdateOne(dentista domain.Dentista) error {
    stmt, err := s.DB.Prepare("UPDATE dentistas SET nombre = ?, apellido = ?, matricula = ? WHERE id = ?")
    if err != nil {
        return err
    }
    _, err = stmt.Exec(
        dentista.Id,
		dentista.Nombre,
		dentista.Apellido,
		dentista.Matricula,
    )
    if err != nil {
        return err
    }

    return nil
}

func(s *dentistaStorage) DeleteOne(id int) error {
	stmt, err := s.DB.Prepare("delete from dentistas where id = ?")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	if _, err := res.RowsAffected(); err != nil {
		return err
	}

	return nil
}