$('#follow').on('click', follow);
$('#unfollow').on('click', unfollow);

function follow(event) {
    const userID = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/user/${userID}/follow`,
        method: 'POST'
    }).done(function(){
        window.location = `/user/${userID}`
    }).fail(function(){
        Swal.fire("Ops...", "Error when following the user","error")
        $('#follow').prop('disabled', false);

    });
}

function unfollow(event) {
    const userID = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/user/${userID}/unfollow`,
        method: 'POST'
    }).done(function(){
        window.location = `/user/${userID}`
    }).fail(function(){
        Swal.fire("Ops...", "Error when unfollowing user","error")
        $('#unfollow').prop('disabled', false);
    });

}