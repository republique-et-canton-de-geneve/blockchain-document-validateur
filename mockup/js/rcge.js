var endpoint = '/api';

angular.module('rc',[])
.controller('extract', function($scope, $http) {
  $scope.numLimit = 16;
  $http.get(endpoint + '/horodatage').
    success(function(data, status, headers, config) {
      data.reverse();
      $scope.extract_list = data;
      console.log(data);
    }).
    error(function(data, status, headers, config) {
      console.log(data);
      console.log(status);
      console.log(headers);
    });
});

function getParameterByName(name, url) {
    if (!url) {
      url = window.location.href;
    }
    name = name.replace(/[\[\]]/g, "\\$&");
    var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
        results = regex.exec(url);
    if (!results) return null;
    if (!results[2]) return '';
    return decodeURIComponent(results[2].replace(/\+/g, " "));
}

Dropzone.autoDiscover = false;
var uploadextraitzone;
function successmultiple(files, message, e) {
    from = message.from;
    target_hash = message.target_hash;
    t = new Date(message.time * 1000);
    display_str = "En attente d'ancrage dans un block Ethereum";
    $("#infobox").html(display_str);
    $("#infobox").attr("class", "alert alert-success");
    setTimeout(function () {
        location.reload();
    }, 25000);
}
function errormultiple(files, message, e) {
    $("#infobox").text(message);
    $("#infobox").attr("class", "alert alert-danger");
    console.log(message);
}
$(function() {
  uploadextraitzone = new Dropzone("div#uploadextraitzone", {
    url : "/api/upload",
    uploadMultiple: true,
    paramName: "myfiles",
    dictDefaultMessage: "Cliquez ici ou déplacer les extraits au format PDF",
    dictFallbackMessage: "Cliquez ici ou déplacer les extraits au format PDF",
    parallelUploads: 256,
    autoProcessQueue: false,
    successmultiple: successmultiple,
      errormultiple: errormultiple,
      addRemoveLinks: true,
  });
});

function processValidate() {
    uploadextraitzone.processQueue();
}
