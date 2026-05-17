package utils

import "time"

func ListarPacientes() ([]Paciente, error) {
	rows, err := DB.Query(`
		SELECT id, nome, data_de_nascimento, diagnostico, alta
		FROM pacientes
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pacientes := []Paciente{}

	for rows.Next() {
		var paciente Paciente
		var birthDate time.Time
		var dischargeDate time.Time

		err := rows.Scan(
			&paciente.ID,
			&paciente.Name,
			&birthDate,
			&paciente.Diagnosis,
			&dischargeDate,
		)
		if err != nil {
			return nil, err
		}

		paciente.BirthDate = birthDate.Format("2006-01-02")
		paciente.DischargeDate = dischargeDate.Format("2006-01-02")
		pacientes = append(pacientes, paciente)
	}

	return pacientes, rows.Err()
}
