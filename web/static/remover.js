function removerAluno(ra) {
    if (confirm("Tem certeza que deseja remover este aluno?")) {
        fetch(`/alunos/${ra}`, {
            method: 'DELETE'
        })
            .then(response => {
                if (response.ok) {
                    alert("Aluno removido com sucesso!");
                    window.location.reload();
                } else {
                    alert("Erro ao remover aluno.");
                }
            })
            // Nao vai acontecer se o usuario agir normalmente.
            .catch(error => {
                console.error('Erro:', error);
                alert("Erro ao remover aluno.");
            });
    }
}