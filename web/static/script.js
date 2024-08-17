document.addEventListener('DOMContentLoaded', function () {
    document.getElementById('loginForm').addEventListener('submit', function (event) {
        event.preventDefault();
        var formData = new FormData(this);

        fetch('/login', {
            method: 'POST',
            body: formData
        })
            .then(response => response.json())
            .then(data => {
                if (data.error) {
                    // Exibe a mensagem de erro no HTMLdocument.getElementById('errorMessage').innerText = data.error;
                    document.getElementById('errorMessage').innerText = data.error;
                } else {
                    // Redireciona para a página de alunoswindow.location.href = data.redirect;
                    window.location.href = data.redirect;
                }
            })
            .catch(error => console.error('Error:', error));
    });
});
