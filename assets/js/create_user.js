$("#form-register").on("submit", createUser);

function createUser(event) {
  event.preventDefault();

  if ($("#password").val() != $("#confirm_password").val()) {
    Swal.fire("We have a problem", "The passwords do not match", "error")
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
      Swal.fire("All Done!", "User created successfuly", "success")
        .then(function(){
          $.ajax({
            url:"/login",
            method:"POST",
            data:{
              email: $("#email").val(),
              password: $("password").val(),
            }
          }).done(function () {
            window.location = "/home"
          }).fail(function () {
            Swal.fire("Error", "Error authenticating the user", "error")
          });
        })
    })
    .fail(function () {
      Swal.fire("Error", "Error creating the user", "error")
    });
}
