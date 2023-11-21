$("#new_post").on("submit", newPost);
$("#update_post").on("click", updatePost);
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
      alert("There was an error creating the post!!");
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
    alert("Error when liking post")
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
    alert("Error when unliking post")
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
    alert("Editou com Sucesso")
  }).fail(function(){
    alert("Deu merda")
  }).always(function(){
    $('#update_post').prop('disabled', false)
  })
}