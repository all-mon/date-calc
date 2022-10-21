function loadDoc() {
    var x = document.getElementById("myInput").value;
    console.log(x);
    var xhttp = new XMLHttpRequest();

    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var resp = JSON.parse(xhttp.response);
            //console.log(JSON.parse(xhttp.response))
            //console.log(resp[0].sch)

            var table = document.getElementById("test");
            for (var i = 0; i < 30; i++) {
                var tr = document.createElement('tr')
                for (var j = 0; j < 2; j++) {
                    var td = document.createElement('td')
                    if (j === 0) td.innerHTML = resp[i].date;
                    else if (j === 1) td.innerHTML = resp[i].sch;
                    tr.appendChild(td)
                }
                table.appendChild(tr)
            }
        }
    };
    xhttp.open("GET", "http://127.0.0.1:8080/api/"+x, true);
    xhttp.send();
}

