var app = angular.module('Main', ['ui.tree']);
app.config(function($interpolateProvider) {
  $interpolateProvider.startSymbol('${');
  $interpolateProvider.endSymbol('}');
});
app.controller("Home", function($scope, $http) {
  $scope.list = [];
  $http.get('/files').success(function (data) {
    console.log(data);
    $scope.list = [data];
  });
});
