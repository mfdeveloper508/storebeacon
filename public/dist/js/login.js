$(window).ready(function() {
    $('#verifysend').click(function(){
        var email = $("#email").val();
        var password = $("#password").val();
        window.location.href = '/verify/store/?email='+email+'&password='+password;
    });
});