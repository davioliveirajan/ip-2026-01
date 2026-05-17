const openPatientFormButton = document.querySelector("#openPatientForm");
const closePatientFormButton = document.querySelector("#closePatientForm");
const patientModal = document.querySelector("#patientModal");
const patientForm = document.querySelector("#patientForm");
const patientsGrid = document.querySelector("#patientsGrid");
const emptyState = document.querySelector("#emptyState");
const patientCount = document.querySelector("#patientCount");
const patientSearch = document.querySelector("#patientSearch");
const formTitle = document.querySelector("#formTitle");
const savePatientButton = document.querySelector("#savePatientButton");
const deletePatientButton = document.querySelector("#deletePatientButton");

let patients = [];
let editingPatientId = null;

function formatDate(dateValue) {
  if (!dateValue) {
    return "";
  }

  return new Date(`${dateValue}T00:00:00`).toLocaleDateString("pt-BR");
}

function toDatabaseDate(dateValue) {
  const value = dateValue.trim();

  if (/^\d{4}-\d{2}-\d{2}$/.test(value)) {
    return value;
  }

  const match = value.match(/^(\d{2})\/(\d{2})\/(\d{4})$/);
  if (!match) {
    return "";
  }

  const [, day, month, year] = match;
  const date = new Date(Number(year), Number(month) - 1, Number(day));

  if (
    date.getFullYear() !== Number(year) ||
    date.getMonth() !== Number(month) - 1 ||
    date.getDate() !== Number(day)
  ) {
    return "";
  }

  return `${year}-${month}-${day}`;
}

function toInputDate(dateValue) {
  if (!dateValue) {
    return "";
  }

  const match = dateValue.match(/^(\d{4})-(\d{2})-(\d{2})$/);
  if (!match) {
    return dateValue;
  }

  const [, year, month, day] = match;
  return `${day}/${month}/${year}`;
}

function applyDateMask(event) {
  const numbers = event.target.value.replace(/\D/g, "").slice(0, 8);
  const parts = [];

  if (numbers.length > 0) {
    parts.push(numbers.slice(0, 2));
  }

  if (numbers.length > 2) {
    parts.push(numbers.slice(2, 4));
  }

  if (numbers.length > 4) {
    parts.push(numbers.slice(4, 8));
  }

  event.target.value = parts.join("/");
}

function openModal() {
  formTitle.textContent = editingPatientId ? "Editar paciente" : "Adicionar paciente";
  savePatientButton.textContent = editingPatientId ? "Salvar alteracoes" : "Salvar paciente";
  deletePatientButton.classList.toggle("is-hidden", !editingPatientId);
  patientModal.classList.add("is-open");
  patientModal.setAttribute("aria-hidden", "false");
  document.querySelector("#patientName").focus();
}

function closeModal() {
  patientModal.classList.remove("is-open");
  patientModal.setAttribute("aria-hidden", "true");
  patientForm.reset();
  editingPatientId = null;
  formTitle.textContent = "Adicionar paciente";
  savePatientButton.textContent = "Salvar paciente";
  deletePatientButton.classList.add("is-hidden");
}

function renderPatients() {
  patientsGrid.innerHTML = "";
  patientCount.textContent = patients.length;

  const searchTerm = patientSearch.value.trim().toLowerCase();
  const visiblePatients = patients.filter((patient) =>
    patient.name.toLowerCase().includes(searchTerm)
  );

  emptyState.classList.toggle("is-visible", visiblePatients.length === 0);

  visiblePatients.forEach((patient) => {
    const card = document.createElement("article");
    const cardHeader = document.createElement("div");
    const title = document.createElement("h3");
    const editButton = document.createElement("button");
    const details = document.createElement("dl");

    card.className = "patient-card";
    cardHeader.className = "patient-card-header";
    title.textContent = patient.name;
    editButton.className = "edit-button";
    editButton.type = "button";
    editButton.textContent = "Editar";
    editButton.addEventListener("click", () => startEditingPatient(patient));

    [
      ["Data de nascimento", formatDate(patient.birthDate)],
      ["Diagnostico", patient.diagnosis],
      ["Data de alta medica", formatDate(patient.dischargeDate)],
    ].forEach(([label, value]) => {
      const row = document.createElement("div");
      const term = document.createElement("dt");
      const description = document.createElement("dd");

      term.textContent = label;
      description.textContent = value;
      row.append(term, description);
      details.appendChild(row);
    });

    cardHeader.append(title, editButton);
    card.append(cardHeader, details);
    patientsGrid.appendChild(card);
  });
}

function startEditingPatient(patient) {
  editingPatientId = patient.id;
  patientForm.elements.name.value = patient.name;
  patientForm.elements.birthDate.value = toInputDate(patient.birthDate);
  patientForm.elements.diagnosis.value = patient.diagnosis;
  patientForm.elements.dischargeDate.value = toInputDate(patient.dischargeDate);
  openModal();
}

async function loadPatients() {
  const response = await fetch("/pacientes");

  if (!response.ok) {
    throw new Error("Nao foi possivel carregar os pacientes.");
  }

  patients = await response.json();
  renderPatients();
}

openPatientFormButton.addEventListener("click", openModal);
closePatientFormButton.addEventListener("click", closeModal);
patientSearch.addEventListener("input", renderPatients);
patientForm.elements.birthDate.addEventListener("input", applyDateMask);
patientForm.elements.dischargeDate.addEventListener("input", applyDateMask);

deletePatientButton.addEventListener("click", async () => {
  if (!editingPatientId) {
    return;
  }

  const confirmed = confirm("Tem certeza que deseja excluir este paciente?");
  if (!confirmed) {
    return;
  }

  const response = await fetch(`/pacientes?id=${editingPatientId}`, {
    method: "DELETE",
  });

  if (!response.ok) {
    alert("Erro ao excluir paciente.");
    return;
  }

  await loadPatients();
  closeModal();
});

patientModal.addEventListener("click", (event) => {
  if (event.target.matches("[data-close-modal]")) {
    closeModal();
  }
});

document.addEventListener("keydown", (event) => {
  if (event.key === "Escape" && patientModal.classList.contains("is-open")) {
    closeModal();
  }
});

patientForm.addEventListener("submit", async (event) => {
  event.preventDefault();

  const formData = new FormData(patientForm);
  const patient = {
    name: formData.get("name").trim(),
    birthDate: toDatabaseDate(formData.get("birthDate")),
    diagnosis: formData.get("diagnosis").trim(),
    dischargeDate: toDatabaseDate(formData.get("dischargeDate")),
  };

  if (!patient.birthDate || !patient.dischargeDate) {
    alert("Digite as datas no formato dd/mm/aaaa.");
    return;
  }

  const url = editingPatientId ? `/pacientes?id=${editingPatientId}` : "/pacientes";
  const method = editingPatientId ? "PUT" : "POST";

  const response = await fetch(url, {
    method,
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(patient),
  });

  if (!response.ok) {
    alert("Erro ao salvar paciente.");
    return;
  }

  await loadPatients();
  closeModal();
});

loadPatients().catch(() => {
  patients = [];
  renderPatients();
  alert("Nao foi possivel carregar os pacientes do banco.");
});
