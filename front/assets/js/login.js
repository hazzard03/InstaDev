$('#login').on("submit", function(event) {
    event.preventDefault();

    var senha = $('#senha').val();

    if (senha.length < 8 || senha.length > 72) {
        Swal.fire("Erro", "A senha deve ter entre 8 e 72 caracteres.", "error");
        return; 
    }

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#email').val(),
            senha: senha,
        }
    }).done(function() {
        alert("Login com Sucesso")
        window.location = "/home";
    }).fail(function() {
        alert("Erro");
        Swal.fire("Ops...", "Usuário ou senha incorretos!", "error");
    });
});
