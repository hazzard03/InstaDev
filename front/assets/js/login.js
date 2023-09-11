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
        Swal.fire("Sucesso!", "Usuário logado com sucesso!", "success")
        window.location = "/home";
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao autenticar o usuário!", "error");
    });
});
