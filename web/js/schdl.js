function loadDoc() {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var resp = JSON.parse(xhttp.response);
            console.log(JSON.parse(xhttp.response))

            console.log(resp[0].sch)
            var table = document.getElementById("test");
            for (var i = 0; i < 5; i++) {
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
    xhttp.open("GET", "http://127.0.0.1:8080/api/%D0%9C%D0%BE%D0%BD%D0%B0%D1%85%D0%BE%D0%B2", true);
    xhttp.send();
}

