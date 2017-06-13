var endpoint = '/api';

angular.module('rc',[])
.controller('Transaction', function($scope, $http) {
  var tx_id = getParameterByName('tx_id');
  $http.get(endpoint + '/tx?tx_id=' + tx_id).
    success(function(data, status, headers, config) {
      $scope.transaction_list = data;
    }).
    error(function(data, status, headers, config) {
      console.log(data);
      console.log(status);
      console.log(headers);
    });
  $scope.action = function (i) {
    var input = $scope.transaction_list[i];
    return set_state(input, !input.state, $http);
  };
  $http.get(endpoint + '/stat').
    success(function(data, status, headers, config) {
      console.log(data);
      $scope.info = data;
    });
})
.controller('EBilletList', function($scope, $http) {
  var user_id = getParameterByName('user_id');
  $scope.action = function (i) {
    var input = $scope.ebillet_list[i];
    if (input.is_mine && input.is_sellable)
      return sell(input, $http);
    
    if (input.is_mine && input.is_transferable)
      return stop_sell(input, $http);
    
    if (!input.is_mine && input.is_transferable)
      return transfer(input, $http);

    console.log("Should never go here");
    return 'ERROR';
  };
  $http.get(endpoint + '/search?user_id=' + user_id).
    success(function(data, status, headers, config) {
      $scope.ebillet_list = data;
      console.log(data);
    }).
    error(function(data, status, headers, config) {
      console.log(data);
      console.log(status);
      console.log(headers);
    });
  $http.get(endpoint + '/stat').
    success(function(data, status, headers, config) {
      console.log(data);
      $scope.info = data;
    });
})
.controller('extract', function($scope, $http) {
  $scope.extract_list = [];
  $scope.submit = function() {
    if (add_to_extract_list($scope)) {
      $scope.Ch1 = '';
      $scope.Ch2 = '';
      $scope.Ch3 = '';
      $scope.Name = '';
      $scope.Mail = '';
    }
  };
  $scope.action = function(index) {
      $scope.extract_list.splice(index, 1);
  };
});

function add_to_extract_list(scope) {
    if (scope.Name == '' ||
        scope.Mail == '' ||
        scope.Ch1  == '' ||
        scope.Ch2  == '' ||
        scope.Ch3  == '')
      return false;

    var to_add = {'name': scope.Name,
                  'mail': scope.Mail,
                  'ch1': scope.Ch1,
                  'ch2': scope.Ch2,
                  'ch3': scope.Ch3};
    scope.extract_list.push(to_add);

    return true;
}

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
