$("#form-register").on("submit", createUser);

function createUser(event) {
  event.preventDefault();

  if ($("#password").val() != $("#confirm_password").val()) {
    alert("The passwords must match!");
    return;
  }

  $.ajax({
    url: "/users",
    method: "POST",
    data: {
      name: $("#name").val(),
      email: $("#email").val(),
      nick: $("#nick").val(),
      password: $("#password").val(),
    },
  })
    .done(function () {
      alert("User created successfuly!");
    })
    .fail(function () {
      alert("Error creating user!");
    });
}
