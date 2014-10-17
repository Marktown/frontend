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
  $rootScope.changePath = function(path) {
    $http.get('/files/data?path='+path).success(function (data, status, headers, config) {
      $scope.error = null;
      $scope.file = data;
    }).error(function (data, status, headers, config) {
      $scope.file = null;
      $scope.error = data;
    });
  };
  $scope.initError = function () {
    $('.editor-error').get(0).contentWindow.document.write($scope.error);
  }
});
