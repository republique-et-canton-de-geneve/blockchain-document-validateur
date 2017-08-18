var endpoint = 'api';
Dropzone.autoDiscover = false;
var myDropzone;
var messageValide = "Documents valides";
var messageInvalide = "Documents non valides";

function successmultiple(files, message, e) {
    $("#infobox").attr("class", "alert alert-success");
    $("#iconbox").attr("class", "fa fa-check fa-stack-1x fa-inverse");
    $("#msgbox").html(messageValide);
    for (i = 0; i < files.length; i++)
        myDropzone.removeFile(files[i]);
}
function errormultiple(files, message, e) {
    $("#infobox").attr("class", "alert alert-danger");
    $("#iconbox").attr("class", "fa fa-exclamation fa-stack-1x fa-inverse");
    $("#msgbox").html(messageInvalide);
}
$(function() {
  myDropzone = new Dropzone("div#validatezone", {
    url : endpoint + "/validate",
    uploadMultiple: true,
    paramName: "myfiles",
    dictDefaultMessage: "Téléverser votre fichier ainsi que son reçu (fichier avec une extension .json).",
    dictFallbackMessage:"Téléverser votre fichier ainsi que son reçu (fichier avec une extension .json).",
    maxFile: 2,
    autoProcessQueue: false,
    successmultiple: successmultiple,
      errormultiple: errormultiple,
      addRemoveLinks: true,
  });
});

function processValidate() {
    myDropzone.processQueue();
}
