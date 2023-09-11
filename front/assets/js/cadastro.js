$('#formulario-cadastro').on("submit", function(event) {
    event.preventDefault(); 
    $form = $(this); 
    var senha = $('#senha').val();
    if (senha.length < 8 || senha.length > 72) {
        Swal.fire("Erro", "A senha deve ter entre 8 e 72 caracteres.", "error");
       
    }else{
        $.ajax({
            url: "/usuarios",
            method: "POST",
            data: {
                nome: $('#nome').val(), 
                email: $('#email').val(),
                nick: $('#nick').val(),
                senha: senha 
            }
        }).done(function() {
            Swal.fire("Sucesso!", "Usuário cadastrado com sucesso!", "success")
                .then(function() {
                    $.ajax({
                        url: "/login",
                        method: "POST",
                        data: {
                            email: $('#email').val(),
                            senha: senha 
                        }
                    }).done(function() {
                        window.location = "/home";
                    }).fail(function() {
                        Swal.fire("Ops...", "Erro ao autenticar o usuário!", "error");
                    })
                })
        }).fail(function() {
            Swal.fire("Ops...", "Erro ao cadastrar o usuário!", "error");
        });
    }  
});
