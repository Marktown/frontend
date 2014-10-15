<div class="container-fluid" ng-app="Main" ng-cloak>
  <div class="row">
    <div class="col-md-3">
      <div ng-controller="Home">
        <script type="text/ng-template" id="items_renderer.html">
          <div ui-tree-handle>
            ${item.name}
          </div>
          <ol ui-tree-nodes="" ng-model="item.items">
            <li ng-repeat="item in item.items" ui-tree-node ng-include="'items_renderer.html'">
            </li>
          </ol>
        </script>
        <div ui-tree>
          <ol ui-tree-nodes="" ng-model="items">
            <li ng-repeat="item in items" ui-tree-node ng-include="'items_renderer.html'"></li>
          </ol>
        </div>
      </div>
    </div>
    <div class="col-md-9 mt-editor" ng-controller="Editor">
      <div class="container-fluid">
        <div class="col-md-6" style="white-space: pre;">${file.data}</div>
        <div class="col-md-6" marked="file.data">
        </div>
      </div>
    </div>
  <div>
</div>
