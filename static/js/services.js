var clAppServices = angular.module('clAppServices', ['ngResource']);

clAppServices.factory('DeliverableRes', ['$resource', function($resource){
  return $resource('/admin/api/module/:moduleName/deliverable/:title',{moduleName:'@moduleName', title:'@title'},{
    'query':{method:'GET', isArray:false}
  } );
}]);

clAppServices.factory('DeliverablesForMod', ['$resource', function($resource){
  return $resource('/admin/api/module/:moduleName',{moduleName:'@moduleName'},{
    'query':{method:'GET', isArray:true}
  } );
}]);


clAppServices.factory('UploadUrlForDeliverable', ['$resource', function($resource){
  return $resource('/admin/api/module/:moduleName/deliverable/:title/getUploadURL',{moduleName:'@moduleName', title:'@title'},{
    'query':{method:'GET', isArray:false}
  } );
}]);
