// Получение всех пользователей
async function getSchenduleByName() {
    // отправляет запрос и получаем ответ
    const response = await fetch("http://localhost:8080/api/Перехода", {
        method: "GET",
        headers: { "Accept": "application/json" }
    });
    // если запрос прошел нормально
    if (response.ok === true) {
        // получаем данные
        const datesList = await response.json();
        let rows = document.querySelector("tbody"); 
        datesList.forEach(date => {
            // добавляем полученные элементы в таблицу
            rows.append(row(date));
        });
    }
}



// создание строки для таблицы
function row(date) {

    const tr = document.createElement("tr");
    tr.setAttribute("data-row", date.date);

    const idTd = document.createElement("td");
    idTd.append(date.date);
    tr.append(idTd);

    const nameTd = document.createElement("td");
    nameTd.append(date.sch);
    tr.append(nameTd);

    return tr;
}      
getSchenduleByName();