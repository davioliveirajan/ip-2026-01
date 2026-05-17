package utils

type Paciente struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	BirthDate     string `json:"birthDate"`
	Diagnosis     string `json:"diagnosis"`
	DischargeDate string `json:"dischargeDate"`
}

func EnsurePacientesTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS pacientes (
			id SERIAL PRIMARY KEY,
			nome TEXT NOT NULL,
			data_de_nascimento DATE NOT NULL,
			diagnostico TEXT NOT NULL,
			alta DATE NOT NULL
		)
	`

	_, err := DB.Exec(query)
	return err
}

func InsertPaciente(paciente Paciente) (Paciente, error) {
	query := `
		INSERT INTO pacientes (nome, data_de_nascimento, diagnostico, alta)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := DB.QueryRow(
		query,
		paciente.Name,
		paciente.BirthDate,
		paciente.Diagnosis,
		paciente.DischargeDate,
	).Scan(&paciente.ID)

	return paciente, err
}
