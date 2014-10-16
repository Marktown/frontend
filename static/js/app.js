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
app.controller("Editor", function($scope, $http, $rootScope, $location, $sce) {
  $rootScope.$on('$locationChangeSuccess', function(event, next, current) {
    var path = $location.$$search.path;
    if (typeof(path) === 'undefined') {
      path = '/';
    }
    $http.get('/files/data?path='+path).success(function (data, status, headers, config) {
      $scope.error = null;
      $scope.file = data;
    }).error(function (data, status, headers, config) {
      $scope.error = $sce.trustAsHtml(data);
      $scope.file = null;
    });
  });
});
