function showSection(sectionId) {
    const sections = document.querySelectorAll('.section');
    sections.forEach(s => s.classList.add('hidden'));
  
    const selected = document.getElementById(sectionId);
    if (selected) {
      selected.classList.remove('hidden');
  
      if (sectionId === "get") {
        fetchIncidents(); 
      }
      else if(sectionId === "getById"){
        fetchIncidentID();
      }
    }
  }
  
  function fetchIncidents() {
    fetch("http://localhost:8080/incidents")
      .then(response => response.json())
      .then(data => renderIncidentList(data))
      .catch(error => {
        console.error("Error al cargar incidentes:", error);
        document.getElementById("incident-list").innerHTML = "<p style='color:red;'>Error al obtener los datos.</p>";
      });
  }
  function fetchIncidentID() {
    const id = document.getElementById("incident-id").value.trim();
    const resultDiv = document.getElementById("incident-found");
  
    if (!id) {
      resultDiv.innerHTML = "<h3 style='color: red;'> ERROR: Ingrese un ID válido</h3>";
      return;
    }
  
    fetch(`http://localhost:8080/incidents/${id}`)
      .then(response => {
        if (!response.ok) {
          throw new Error("404 Incidente no encontrado");
        }
        return response.json();
      })
      .then(incident => {
        resultDiv.innerHTML = `
          <div class="incident-card">
            <h3>#${incident.id} - ${incident.status}</h3>
            <p><strong>Reportado por:</strong> ${incident.reporter}</p>
            <p>${incident.description}</p>
            <p><small>${incident.created_at}</small></p>
          </div>
        `;
      })
      .catch(error => {
        resultDiv.innerHTML = `<h3 style="color:red;">${error.message}</h3>`;
      });
  }
  
  
  function renderIncidentList(incidents) {
    const container = document.getElementById("incident-list");
    container.innerHTML = "";
  
    if (incidents.length === 0) {
      container.innerHTML = "<p>No hay incidentes registrados.</p>";
      return;
    }
  
    incidents.forEach(incident => {
      const card = document.createElement("div");
      card.className = "incident-card";
  
      card.innerHTML = `
        <h3>#${incident.id} - ${incident.status}</h3>
        <p><strong>Reportado por:</strong> ${incident.reporter}</p>
        <p>${incident.description}</p>
        <p><small>${incident.created_at}</small></p>
      `;
  
      container.appendChild(card);
    });
  }
  function createIncident() {
    const reporter = document.getElementById("reporter").value.trim();
    const description = document.getElementById("description").value.trim();
    const status = document.getElementById("status").value;
    const resultDiv = document.getElementById("post-result");
  
    if (!reporter || !description || description.length < 10) {
      resultDiv.innerHTML = "<p style='color:red;'> Debes ingresar un reportero y una descripción válida (mínimo 10 caracteres)</p>";
      return;
    }
  
    const validStatuses = ["pendiente", "en proceso", "resuelto"];
    if (status && !validStatuses.includes(status)) {
      resultDiv.innerHTML = "<p style='color:red;'>Estado Inválido.</p>";
      return;
    }
  
    const payload = {
      reporter: reporter,
      description: description
    };
  
    if (status) {
      payload.status = status;
    }
  
    fetch("http://localhost:8080/incidents", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(payload)
    })
      .then(response => {
        if (!response.ok) {
          throw Error("ERROR al subir incidente");
        }
        return response.json();
      })
      .then(data => {
        resultDiv.innerHTML =
          `<p style="color:green;">Incidente creado con éxito:</p>
          <pre>${JSON.stringify(data, null, 2)}</pre>`;
      })
      .catch(error => {
        resultDiv.innerHTML = `<p style="color:red;">${error.message}</p>`;
      });
  }
  