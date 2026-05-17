package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"servidorHTTP/servidorHTTP/app/utils"
)

func PacientesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		listarPacientes(w)
	case http.MethodPost:
		inserirPaciente(w, r)
	case http.MethodPut:
		atualizarPaciente(w, r)
	case http.MethodDelete:
		excluirPaciente(w, r)
	default:
		http.Error(w, `{"erro":"metodo nao permitido"}`, http.StatusMethodNotAllowed)
	}
}

func listarPacientes(w http.ResponseWriter) {
	pacientes, err := utils.ListarPacientes()
	if err != nil {
		http.Error(w, `{"erro":"erro ao buscar pacientes"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(pacientes)
}

func inserirPaciente(w http.ResponseWriter, r *http.Request) {
	var paciente utils.Paciente

	err := json.NewDecoder(r.Body).Decode(&paciente)
	if err != nil {
		http.Error(w, `{"erro":"json invalido"}`, http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(paciente.Name) == "" ||
		paciente.BirthDate == "" ||
		strings.TrimSpace(paciente.Diagnosis) == "" ||
		paciente.DischargeDate == "" {
		http.Error(w, `{"erro":"preencha todos os campos"}`, http.StatusBadRequest)
		return
	}

	paciente, err = utils.InsertPaciente(paciente)
	if err != nil {
		http.Error(w, `{"erro":"erro ao inserir paciente"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(paciente)
}

func atualizarPaciente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id <= 0 {
		http.Error(w, `{"erro":"id invalido"}`, http.StatusBadRequest)
		return
	}

	var paciente utils.Paciente
	err = json.NewDecoder(r.Body).Decode(&paciente)
	if err != nil {
		http.Error(w, `{"erro":"json invalido"}`, http.StatusBadRequest)
		return
	}

	paciente.ID = id
	if strings.TrimSpace(paciente.Name) == "" ||
		paciente.BirthDate == "" ||
		strings.TrimSpace(paciente.Diagnosis) == "" ||
		paciente.DischargeDate == "" {
		http.Error(w, `{"erro":"preencha todos os campos"}`, http.StatusBadRequest)
		return
	}

	err = utils.AtualizarPaciente(paciente)
	if err != nil {
		http.Error(w, `{"erro":"erro ao atualizar paciente"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(paciente)
}

func excluirPaciente(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id <= 0 {
		http.Error(w, `{"erro":"id invalido"}`, http.StatusBadRequest)
		return
	}

	err = utils.DeletarPaciente(id)
	if err != nil {
		http.Error(w, `{"erro":"erro ao excluir paciente"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
