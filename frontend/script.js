
// Funcion para mostrar una de las secciones 
function showSection(sectionId) {
    const sections = document.querySelectorAll('.section');
    sections.forEach(s => s.classList.add('hidden'));
  
    const selected = document.getElementById(sectionId);
    if (selected) {
      selected.classList.remove('hidden');
  
      if (sectionId === "get") {
        fetchIncidents(); 
      }
    }
  }
  
  // Traer todos los incidentes
  function fetchIncidents() {
    fetch("http://localhost:8080/incidents")
      .then(response => response.json())
      .then(data => renderIncidentList(data))
      .catch(error => {
        console.error("Error al cargar incidentes:", error);
        document.getElementById("incident-list").innerHTML = "<p style='color:red;'>Error al obtener los datos.</p>";
      });
  }

  // mostrar los incidentes 
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
        <h3>${incident.id} - ${incident.status}</h3>
        <p><strong>Reportado por:</strong> ${incident.reporter}</p>
        <p>${incident.description}</p>
        <p><small>${incident.created_at}</small></p>
      `;
  
      container.appendChild(card);
    });
  }

  // Traer un incidente por ID 
  function fetchIncidentID(inputID, outputID) {
    const id = document.getElementById(inputID).value.trim();
    const resultDiv = document.getElementById(outputID);
  
    if (!id || id <= 0) {
      resultDiv.innerHTML = "<p style='color:red;'> Ingresa un ID numérico válido mayor que 0</p>";
      return;
    }
  
    // traer el incidente especifico con su id 
    fetch(`http://localhost:8080/incidents/${id}`)
      .then(response => {
        if (!response.ok)
          throw new Error("404 Incidente no encontrado");
        {
          return response.json();
        }
      })
      .then(incident => {
        // mostrar el contenedor del incidente 
        resultDiv.innerHTML = `
          <div class="incident-card">
            <h3>${incident.id} - ${incident.status}</h3>
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
  

  // Crear un nuevo incidente y traer variables nuevas 
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
  

  function editIncidentID(){
    const id = document.getElementById("incident-id-put").value;
    const status = document.getElementById("status-put").value.trim();
    const resultDiv = document.getElementById("put-result");

    if(!id){
      resultDiv.innerHTML = "<p style= 'color:red'>Debe ingresar un ID valido</p>";
      return;
    }

    const validStatuses = ["pendiente", "en proceso", "resuelto"];
    if (!validStatuses.includes(status)) {
      resultDiv.innerHTML = "<p style='color:red;'>Estado Inválido.</p>";
      return;
    }
  
    const payload = {
      status: status
    }

    fetch(`http://localhost:8080/incidents/${id}`,{
      method:"PUT",
      headers: {
        "Content-Type": "application/json"
      }, body: JSON.stringify(payload)
    })
    .then(response => {
      if(!response.ok){
        throw Error("ERROR al actualizar incidente");
      }
      return response.json();
    })
    .then(()=>{
      resultDiv.innerHTML = `<p style="color:green;"> Incidente ${id} actualizado a estado: <strong>${status}</strong></p>`;
    })
    .catch(error => {
      resultDiv.innerHTML=  `<p style="color:red;"> ${error.message}</p>`;
    })
  }

  function deleteIncident(){
    const id = document.getElementById("incident-id-del").value;
    const resultDiv = document.getElementById("del-result");

    if(!id){
      resultDiv.innerHTML = "<p style= 'color:red'>Debe ingresar un ID valido</p>";
      return;
    }

    // Confirmación antes de eliminar 
    const confirmacion = confirm(`Deseas eliminar el incidente ${id}?`);
    if(!confirmacion){
      return;
    }

    fetch(`http://localhost:8080/incidents/${id}`,{
      method:"DELETE",
    })
    .then(response =>{
      if(response.status === 404){
        throw Error('404 Incidente no encontrado');
      }
      if(!response.ok){
        throw Error("Error al eliminar");
      }
      resultDiv.innerHTML = `<p style="color:green;">Incidente ${id} eliminado correctamente.</p>`;
    })
    .catch(error => {
      resultDiv.innerHTML = `<p style="color:red;"> ${error.message}.</p>`;
    })



  }