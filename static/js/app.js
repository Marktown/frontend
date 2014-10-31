var app = angular.module('Main', ['ui.tree', 'hc.marked']);
app.config(function($interpolateProvider) {
  $interpolateProvider.startSymbol('${');
  $interpolateProvider.endSymbol('}');
});
app.controller("Home", function($scope, $http) {
  function loadFiles(){
    $http.get('/files').success(function (data) {
      $scope.items = [data];
    });
  }
  loadFiles();
  $scope.addFile = function(){
    var filename = prompt('Please choose a filename', 'Untitled.md');
    if (filename) {
      $http.post('/files', {filename: filename, content: "Hello World!"}).success(function(data){
        loadFiles();
      });
    }
  }
});
app.controller("Editor", function($scope, $http, $rootScope, $location, $sce) {
  $rootScope.changePath = function(path) {
    $scope.path = path;
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
  $scope.updateFile= function(){
    $http.post('/files', {filename: $scope.path, content: $scope.file.data}).success(function(data){
    });
  }
});
