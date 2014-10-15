var app = angular.module('Main', ['ui.tree', 'hc.marked']);
app.config(function($interpolateProvider) {
  $interpolateProvider.startSymbol('${');
  $interpolateProvider.endSymbol('}');
});
app.controller("Home", function($scope, $http) {
  $http.get('/files').success(function (data) {
    $scope.items = [data];
  });
});
app.controller("Editor", function($scope, $http, $rootScope, $location) {
  $rootScope.$on('$locationChangeSuccess', function(event, next, current) {
    var path = $location.$$search.path;
    $http.get('/files/data?path='+path).success(function (data) {
      $scope.file = data;
    });
  });
});
