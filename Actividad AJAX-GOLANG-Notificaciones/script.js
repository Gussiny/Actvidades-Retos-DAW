let xmlhttp;

        function mostrar() {

            if (window.XMLHttpRequest) {
                xmlhttp = new XMLHttpRequest();
            } else {
                xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
            }

            xmlhttp.onreadystatechange = function () {

                if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
                    document.getElementById("lugardetrabajo").innerHTML = "ARREGLO MODIFICADO";
                    // console.log(xmlhttp.responseText);
                }
            }

            // POST
            var cadena = "/reto"
            xmlhttp.open("POST", cadena, true);
            xmlhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
            xmlhttp.send("y=3&n=4");

            // alert("ARREGLO MODIFICADO");
        }

        function mostrar2() {

            if (window.XMLHttpRequest) {
                xmlhttp = new XMLHttpRequest();
            } else {
                xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
            }

            xmlhttp.onreadystatechange = function () {

                if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
                    document.getElementById("lugardetrabajo").innerHTML = xmlhttp.responseText;
                    console.log(xmlhttp.responseText);
                }
            }

            // GET
            xmlhttp.open("GET", "/reto", true);
            xmlhttp.send();
        }

        function mostrarJSON() {

            if (window.XMLHttpRequest) {
                xmlhttp = new XMLHttpRequest();
            } else {
                xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
            }

            xmlhttp.onreadystatechange = function () {

                if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
                    //document.getElementById("trabajo2").innerHTML = xmlhttp.responseText;
                    // console.log(xmlhttp.responseText);

                    var data = [xmlhttp.responseText]
                    var eljson = JSON.parse(xmlhttp.responseText)
                    // console.log(data);
                    console.log(eljson);

                    var tabla = document.getElementById("tabla")

                    $("#tabla").empty();

                    for (i = 0; i < eljson.length; i++) {
                        var row = `<tr>
                                <td>${eljson[i].Titulo}</td>
                                <td>${eljson[i].Director}</td>
                                <td>${eljson[i].Duracion}</td>
                                <td>${eljson[i].Calificacion}</td>
                            </tr>`

                        tabla.innerHTML += row
                    }
                }
            }

            // GET
            xmlhttp.open("GET", "/actividad2", true);
            xmlhttp.send();

        }

        function formulario() {
            var x = document.getElementById("formulario");
            var boton = document.getElementById("botonForm");

            if (x.style.display === "none") {
                x.style.display = "block";
                boton.style.background = 'red'
                boton.firstChild.data = "CANCELAR"
            } else {
                x.style.display = "none";
                boton.style.background = '#28a745'
                boton.firstChild.data = "NUEVO"
            }
        }

        function insertarDatos() {
            if (window.XMLHttpRequest) {
                xmlhttp = new XMLHttpRequest();
            } else {
                xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
            }

            xmlhttp.onreadystatechange = function () {
                if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
                    document.getElementById("lugardetrabajo").innerHTML = "ARREGLO MODIFICADO";
                }
            }

            // POST
            var cadena = "/insertar"
            var formulario = document.getElementById("formElement")

            var titulo = document.getElementById("titulo").value
            var director = document.getElementById("director").value
            var duracion = document.getElementById("duracion").value
            var calificacion = document.getElementById("calificacion").value

            console.log(titulo)
            // var datos = new FormData(formulario)
            var datos =  []
            datos.push("titulo=" + titulo)
            datos.push("director=" + director)
            datos.push("duracion=" + duracion)
            datos.push("calificacion=" + calificacion)

            console.log(datos)
            

            xmlhttp.open("POST", cadena, true);
            xmlhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
            xmlhttp.send(datos.join("&"));

            $("#formElement").trigger("reset");
        }