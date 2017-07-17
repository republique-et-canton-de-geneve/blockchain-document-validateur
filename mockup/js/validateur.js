Dropzone.autoDiscover = false;
var myDropzone;
function successmultiple(files, message, e) {
    from = message.from;
    target_hash = message.target_hash;
    t = new Date(message.time * 1000);
    from_link = '<a href="https://ropsten.etherscan.io/address/' + from + '">' + from + '</a>';
    display_str = "Horodaté le " + t + " par " + from_link + " (RCGE)<br/>Hash fichier: " + target_hash;
    $("#infobox").html(display_str);
    $("#infobox").attr("class", "alert alert-success");
    for (i = 0; i < files.length; i++)
        myDropzone.removeFile(files[i]);
}
function errormultiple(files, message, e) {
    $("#infobox").text(message);
    $("#infobox").attr("class", "alert alert-danger");
}
$(function() {
  myDropzone = new Dropzone("div#validatezone", {
    url : "/api/validate",
    uploadMultiple: true,
    paramName: "myfiles",
    dictDefaultMessage: "Cliquez ici ou déplacer l'extrait(.pdf) et son reçu (.json)",
    dictFallbackMessage: "Cliquez ici ou déplacer l'extrait(.pdf) et son reçu (.json)",
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
