const imageBoxes = document.querySelectorAll('.image-box');
        let currentIndex = 0;

        function selectImage(index) {
            currentIndex = index;
            document.getElementById('file-input').click();
        }

        function uploadImage(event) {
            const file = event.target.files[0];
            if (file) {
                const reader = new FileReader();
                reader.onload = function(e) {
                    imageBoxes[currentIndex].innerHTML = `<img src="${e.target.result}" alt="Uploaded Image">`;
                };
                reader.readAsDataURL(file);
            }
            event.target.value = ''; // Reset input to allow re-upload
        }
        const cities = ["Київ", "Львів", "Одеса", "Харків", "Дніпро", "Запоріжжя", "Вінниця", "Івано-Франківськ", "Полтава", "Чернігів"];
        const input = document.getElementById("cityInput");
        const dropdown = document.getElementById("dropdown");
        
        input.addEventListener("input", function() {
            const value = this.value.toLowerCase();
            dropdown.innerHTML = "";
            if (value) {
                const filteredCities = cities.filter(city => city.toLowerCase().includes(value));
                if (filteredCities.length > 0) {
                    dropdown.style.display = "block";
                    filteredCities.forEach(city => {
                        const div = document.createElement("div");
                        div.textContent = city;
                        div.addEventListener("click", function() {
                            input.value = city;
                            dropdown.style.display = "none";
                        });
                        dropdown.appendChild(div);
                    });
                } else {
                    dropdown.style.display = "none";
                }
            } else {
                dropdown.style.display = "none";
            }
        });
        
        document.addEventListener("click", function(event) {
            if (!event.target.closest(".search-container")) {
                dropdown.style.display = "none";
            }
        });               