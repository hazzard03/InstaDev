$("#formulario-cadastro").on("submit", criarUsuario);

function criarUsuario(evento) {
  evento.preventDefault();

  $.ajax({
    url: "/usuarios",
    method: "POST",
    data: {
      nome: $("#nome").val(),
      email: $("#email").val(),
      nick: $("#nick").val(),
      senha: $("#senha").val(),
    },
  });
}

let inputPass = document.getElementById("senha");
let inputConfirmPass = document.getElementById("confirmacao");

inputConfirmPass.addEventListener("focusout", () => {
  if (inputPass.value !== inputConfirmPass.value) {
    alert("As senhas não estão iguais");
  }
});
