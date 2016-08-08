/*
    Copyright (c) 2016 VMware, Inc. All Rights Reserved.
    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
*/
(function() {

  'use strict';

  angular
    .module('harbor.repository')
    .directive('addLabel', addLabel);

  AddLabelController.$inject = ['AddLabelService', '$scope'];

  function AddLabelController(AddLabelService, $scope) {
    var vm = this;

    $scope.p = {};
    var vm0 = $scope.p;
    vm0.labelName = '';

    vm.cancel = cancel;

    vm.reset = reset;

    vm.hasError = false;
    vm.errorMessage = '';
    vm.addLabel = addLabel;

    // 添加label
    function addLabel(repoName, p) {
      if(p && angular.isDefined(p.labelName)) {
        AddLabelService(repoName, p.labelName)
          .success(addLabelSuccess)
          .error(addLabelFailed);
      }
    }

    function addLabelSuccess(data, status) {
      $scope.$emit('addedSuccess', true);
      vm.hasError = false;
      vm.errorMessage = '';
      vm.isOpen = false;
      //成功之后更新页面
      vm.refresh();
    }

    function addLabelFailed(data, status) {
      vm.hasError = true;
      vm.errorMessage = status;
      console.log('Failed to add project:' + status);
      if(status === 409 && vm0.projectName !== '') {
        vm.errorMessage = 'label_already_exist';
      }
    }

    function cancel(form){
      if(form) {
        form.$setPristine();
        form.$setUntouched();
      }
      vm.isOpen = false;
      vm0.labelName = '';

      vm.hasError = false; vm.close = close;
      vm.errorMessage = '';
    }

    function reset() {
      vm.hasError = false;
      vm.errorMessage = '';
    }
  }

  function addLabel() {
    var directive = {
      'restrict': 'E',
      'templateUrl': '/static/resources/js/components/repository/add-label.directive.html',
      'controller': AddLabelController,
      'scope' : {
        'isOpen': '=',
        'repoName' : '=',
        'associateId' : '=',
        'refresh': '&'
      },
      'link': link,
      'controllerAs': 'vm',
      'bindToController': true
    };
    return directive;

    function link(scope, element, attrs, ctrl) {

      scope.form.$setPristine();
      scope.form.$setUntouched();
    }
  }

})();
