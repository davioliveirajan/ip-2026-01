package utils

func DeletarPaciente(id int) error {
	_, err := DB.Exec("DELETE FROM pacientes WHERE id = $1", id)
	return err
}
