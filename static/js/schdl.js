document.getElementById("schdl_table").hidden = true;
function loadScheduleTable() {

	document.getElementById("schdl_table").hidden = false;
    for (let i = document.getElementById('schdl_table').getElementsByTagName('tr').length -1; i; i--) {
            document.getElementById('schdl_table').deleteRow(i);
        }

    let x = document.getElementById("search_input").value;

    console.log(x);
    let xhttp = new XMLHttpRequest();

    xhttp.onreadystatechange = function() {
        if (this.readyState === 4 && this.status === 200) {
            let resp = JSON.parse(xhttp.response);
            console.log(JSON.parse(xhttp.response))
            console.log(resp[0].working_shift)
            

            let table = document.getElementById("schdl_table");
            for (let i = 0; i < 365; i++) {
                let tr = document.createElement('tr')
                for (let j = 0; j < 2; j++) {
                    let td = document.createElement('td')
                    if (j === 0) td.innerHTML = resp[i].date;
                    else if (j === 1) td.innerHTML = resp[i].working_shift;
                    tr.appendChild(td)
                }
                table.appendChild(tr)
            }
        }
    };
    xhttp.open("GET", "http://127.0.0.1:8080/api/employees/"+x, true);
    xhttp.send();
}

