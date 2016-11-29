var app = angular.module('game', ['ngRoute']);
var engine;

app.config(function($routeProvider){
  $routeProvider.when('/', {
    controller: 'GameController',
    templateUrl: 'templates/home.html'
  });
});

app.controller('GameController', ['$scope', '$timeout', function($scope, $timeout){
  $timeout(function(){
    engine = new Engine('game_screen');
    engine.client = initClient(new WebSocket("ws://" + location.host + "/connect")); 
  }, 0);

}]);

app.directive('menubar', function(){
  return {
    templateUrl: 'templates/menubar.html',
    controller: 'MenuController'
  }
});

app.controller('MenuController', ['$scope', function($scope){
  $scope.loggedIn = true;
}]);