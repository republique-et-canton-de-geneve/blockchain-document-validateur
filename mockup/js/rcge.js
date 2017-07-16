var endpoint = '/api';

angular.module('rc',[])
.controller('extract', function($scope, $http) {
  $scope.numLimit = 16;
  $http.get(endpoint + '/horodatage').
    success(function(data, status, headers, config) {
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
