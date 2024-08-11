document.addEventListener("DOMContentLoaded", function() {
    // Form Validation
    const form = document.querySelector("form");
    if (form) {
        form.addEventListener("submit", function(event) {
            const name = document.getElementById("name").value.trim();
            const price = document.getElementById("price").value.trim();
            const image = document.getElementById("image").files[0];
            
            if (!name || !price || !image) {
                alert("Please fill out all fields and select an image.");
                event.preventDefault();
                return false;
            }

            if (isNaN(price) || Number(price) <= 0) {
                alert("Please enter a valid price.");
                event.preventDefault();
                return false;
            }
        });
    }

    // Image Preview
    const imageInput = document.getElementById("image");
    const imagePreview = document.createElement("img");
    if (imageInput) {
        imageInput.parentNode.appendChild(imagePreview);
        imageInput.addEventListener("change", function() {
            const file = this.files[0];
            if (file) {
                const reader = new FileReader();
                reader.onload = function(e) {
                    imagePreview.src = e.target.result;
                    imagePreview.style.maxWidth = "200px";
                    imagePreview.style.marginTop = "10px";
                }
                reader.readAsDataURL(file);
            }
        });
    }

    // Smooth Scroll
    const navLinks = document.querySelectorAll("nav ul li a");
    navLinks.forEach(link => {
        link.addEventListener("click", function(event) {
            if (this.hash !== "") {
                event.preventDefault();
                const hash = this.hash;

                document.querySelector(hash).scrollIntoView({
                    behavior: "smooth"
                });
            }
        });
    });
});
