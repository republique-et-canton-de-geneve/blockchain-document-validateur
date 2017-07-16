Dropzone.autoDiscover = false;
$(function() {
    $("div#validatezone").dropzone({
        url : "/api/validate",
        uploadMultiple: true,
        paramName: "myfiles"
    });
});
