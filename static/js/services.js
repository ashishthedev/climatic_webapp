var clAppServices = angular.module('clAppServices', ['ngResource']);

clAppServices.factory('DELIVERABLES', ['$resource', function($resource){
  return $resource('/admin/api/deliverables',{},{
  } );
}]);


clAppServices.factory('UPLOADURL', ['$resource', function($resource){
  return $resource('/admin/api/uploadurl',{},{
  } );
}]);

clAppServices.factory('REMOVETASKFILE', ['$resource', function($resource){
  return $resource('/admin/api/removetaskfile',{},{
  } );
}]);

