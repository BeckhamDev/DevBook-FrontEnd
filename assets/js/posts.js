$("#new_post").on("submit", newPost);

function newPost(event) {
  event.preventDefault();

  $.ajax({
    url: "/posts",
    method: "POST",
    data: {
      title: $("#title").val(),
      content: $("#content").val(),
    },
  })
    .done(function () {
      window.location.reload();
    })
    .fail(function () {
      alert("There was an error creating the post!!");
    });
}
