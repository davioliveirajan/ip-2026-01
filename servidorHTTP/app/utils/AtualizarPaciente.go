package utils

func AtualizarPaciente(paciente Paciente) error {
	query := `
		UPDATE pacientes
		SET nome = $1, data_de_nascimento = $2, diagnostico = $3, alta = $4
		WHERE id = $5
	`

	_, err := DB.Exec(
		query,
		paciente.Name,
		paciente.BirthDate,
		paciente.Diagnosis,
		paciente.DischargeDate,
		paciente.ID,
	)

	return err
}
