let success = document.getElementById("success"); //for toggling results
let error = document.getElementById("error"); //for toggling results
let drones = []; //store drone data
let markers = []; //array to hold markers for removing when needed
let map; //init map

const websocket = new WebSocket(`ws://${window.location.host}/drones`);

websocket.onopen = function(event) {
  console.log("Successfully connected to websocket server");
};

websocket.onerror = function(error) {
  console.log("Error connecting to websocket server");
  console.log(error);
};

websocket.onmessage = function(event) {

  drones = JSON.parse(event.data);

  let table = document.getElementById("tbody");

  //on every update clear all existing markers for full refresh
  for (let i = 0; i < markers.length; i++) {
    markers[i].setMap(null);
  }
  markers = [];

  drones.forEach(function(drone) {
    let exists = document.getElementById(drone.id);
    if (exists === null) {
      let row = table.insertRow(table.rows.length);
      row.id = drone.id;

      if (drone.stationary == true) {
        row.style.color='red';
      }

      let cell1 = row.insertCell(0);
      let cell2 = row.insertCell(1);
      let cell3 = row.insertCell(2);
      let cell4 = row.insertCell(3);
      let cell5 = row.insertCell(4);

      cell1.innerHTML = drone.id;
      cell2.innerHTML = drone.lat;
      cell3.innerHTML = drone.lon;
      cell4.innerHTML = drone.speed;
      cell5.innerHTML = `<button class="btn btn-danger" onclick="removeDrone(this)">Remove</button>`

    } else {
      let row = document.getElementById(drone.id)
      row.cells[0].innerHTML = drone.id;
      row.cells[1].innerHTML = drone.lat;
      row.cells[2].innerHTML = drone.lon;
      row.cells[3].innerHTML = drone.speed;

      if (drone.stationary == true) {
          row.style.color='red';
      } else {
          row.style.color='black';
      }
    }

      //add a marker for each drone
      let droneLatlng = new google.maps.LatLng(drone.lat, drone.lon);

      let marker = new google.maps.Marker({
          position: droneLatlng,
          title: drone.id,
      });

      //add marker to array
      markers.push(marker);
      //add the marker to the map
      marker.setMap(map);

  });

  success.style.display = "none"

  //clear rows of since deleted drones
  for (let i = 0, row; row = table.rows[i]; i++) {
    if (!drones.some(drone => drone.id === row.id)) {
      row.parentNode.removeChild(row);
    }

  }

};

//add drone
const droneForm = document.getElementById("droneForm");

droneForm.addEventListener("submit", function(e) {
  e.preventDefault();

  //reset error message if visible
  error.style.display="none";

  let formLat = document.getElementById("formLat");
  let formLon = document.getElementById("formLon");

  let lat = parseFloat(formLat.value);
  let lon = parseFloat(formLon.value);

  fetch("/add", {
    method: "POST",
    body: JSON.stringify({"lat": lat, "lon": lon})
  })
  .then(() => {
    this.formLat.value = "";
    this.formLon.value = "";
    success.style.display = "block";
  })
  .catch(() => {
    error.style.display="block";
  })

})

function removeDrone(e) {

  let p = e.parentNode.parentNode;
  let id = p.id;

  fetch(`/delete?id=${p.id}`, {
    method: "DELETE",
  })
  .then(() => {
    //remove row
    p.parentNode.removeChild(p);
    //remove marker
    let marker = markers.filter(m => m.title === id)[0]
    marker.setMap(null);
  })
  .catch(() => {
    error.style.display="block";
  })


}

function initMap() {
  map = new google.maps.Map(document.getElementById('map'), {
    center: {lat: 43.65, lng: -79.35},
    zoom: 8
  });
}



