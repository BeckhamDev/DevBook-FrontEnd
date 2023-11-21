$("#new_post").on("submit", newPost);
$("#update_post").on("click", updatePost);
$(".delete-post").on("click", deletePost);
$(document).on("click", '.like-post', likePost);
$(document).on("click", '.unlike-post', unlikePost);


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
      Swal.fire("Error", "Error creating the post", "error")
    });
}

function likePost(event) {
  event.preventDefault();

  const targetElement = $(event.target);
  const postId = targetElement.closest("div").data("post-id")

  targetElement.prop('disabled', true);
  $.ajax({
    url: `/posts/${postId}/like`,
    method: "POST",
  }).done(function () {
    const likesCounter = targetElement.next("span")
    const likes = parseInt(likesCounter.text())

    likesCounter.text(likes + 1)

    targetElement.addClass("unlike-post")
    targetElement.addClass("text-danger")
    targetElement.removeClass("like-post")

  }).fail(function(){
    Swal.fire("Error", "Error when liking post", "error")
  }).always(function(){
    targetElement.prop('disabled', false);
  });
}

function unlikePost(event){
  event.preventDefault();

  const targetElement = $(event.target);
  const postId = targetElement.closest("div").data("post-id")

  targetElement.prop('disabled', true);
  $.ajax({
    url: `/posts/${postId}/unlike`,
    method: "POST",
  }).done(function () {
    const likesCounter = targetElement.next("span")
    const likes = parseInt(likesCounter.text())

    likesCounter.text(likes - 1)

    targetElement.removeClass("unlike-post")
    targetElement.removeClass("text-danger")
    targetElement.addClass("like-post")

  }).fail(function(){
    Swal.fire("Error", "Error when unliking post", "error")
  }).always(function(){
    targetElement.prop('disabled', false);
  });
}

function updatePost(event) {
  $(this).prop('disabled', true);
  
  const postId = $(this).data("post-id");
  
  $.ajax({
    url: `/posts/${postId}`,
    method: "PUT",
    data: {
      title: $('#title').val(),
      content: $('#content').val(), 
    }
  }).done(function(){
    Swal.fire(
      "Success",
      "Post created successfully!",
      "success",
    ).then(function(){
      window.location = "/home"
    })
  }).fail(function(){
    Swal.fire("Error", "Error creating the post", "error")
  }).always(function(){
    $('#update_post').prop('disabled', false)
  })
}

function deletePost(event) {
  event.preventDefault();

  Swal.fire({
    title: "Warning",
    text: "Are you sure you want to delete this post?",
    showCancelButton: true,
    cancelButtonText: "Cancel",
    icon: "warning",
  }).then(function(confirm){
    if (!confirm.value) return;

    const targetElement = $(event.target);
    const post =targetElement.closest("div")
    const postId = post.data("post-id")

    targetElement.prop('disabled', true);

    $.ajax({
      url: `/posts/${postId}`,
      method: "DELETE"
    }).done(function(){
      post.fadeOut("slow", function(){
        $(this).remove()
      })
    }).fail(function(){
      Swal.fire("Error", "Error deleting the post", "error")
    })
  })

}