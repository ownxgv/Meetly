document.addEventListener('DOMContentLoaded', function () {
    const submitButton = document.querySelector('#signup-submit');
    
    // Перевірка на існування елемента перед додаванням обробника події
    if (submitButton) {
        submitButton.onclick = function (event) {
            event.preventDefault();
            
            // Отримуємо значення полів форми
            let name = document.querySelector('#signup-name').value;
            let pass = document.querySelector('#signup-pass').value;
            let email = document.querySelector('#signup-email').value;
            let birthday = document.querySelector('#signup-birthday').value;
            let PIB = document.querySelector('#signup-PIB').value; // Використовуємо querySelector замість querySelectorAll

            // Створюємо об'єкт з даними
            let data = {
                "name": name,
                "pass": pass,
                "email": email,
                "birthday": birthday,
                "PIB": PIB
            };
            
            // Викликаємо функцію для відправки даних
            reg('registr/signup.php', 'POST', login, data);
        };
    } else {
        console.error('Element with id "signup-submit" not found');
    }

    // Функція для відправки даних
    function reg(url, method, functionName, dataArray) {
        let xhttp = new XMLHttpRequest();
        xhttp.open(method, url, true);
        xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
        xhttp.send(requestData(dataArray));

        xhttp.onreadystatechange = function () {
            if (this.readyState == 4 && this.status == 200) {
                functionName(this.responseText);
            }
        };
    }

    // Функція для перетворення даних в рядок запиту
    function requestData(dataArr) {
        let out = '';
        for (let key in dataArr) {
            out += `${key}=${encodeURIComponent(dataArr[key])}&`;  // Використовуємо encodeURIComponent для безпеки
        }
        return out;
    }

    // Функція обробки результату
    function login(result) {
        console.log(result);
        if (result == 2) {
            alert('Заповніть всі поля');
        } else if (result == 1) {
            alert('Успіх! Тепер можна увійти!');
        } else {
            alert('Помилка! Спробуйте пізніше.');
        }
    }
});
