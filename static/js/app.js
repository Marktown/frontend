var app = angular.module('Main', ['ui.tree']);
app.config(function($interpolateProvider) {
  $interpolateProvider.startSymbol('${');
  $interpolateProvider.endSymbol('}');
});
app.controller("Home", function($scope, $http) {
  $scope.items = [];
  $http.get('/files').success(function (data) {
    $scope.items = [data];
  });
});
