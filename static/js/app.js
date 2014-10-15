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
app.controller("Editor", function($scope, $http) {
  $http.get('/files/1').success(function (data) {
    $scope.file = data;
  });
});
