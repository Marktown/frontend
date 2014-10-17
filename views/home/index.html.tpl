<div class="container-fluid" ng-app="Main" ng-cloak>
  <div class="row">
    <div class="col-md-3">
      <div ng-controller="Home">
        <script type="text/ng-template" id="items_renderer.html">
          <div ui-tree-handle>
            <a href="#/files/data?path=${item.name}">${item.name}</a>
          </div>
          <ol ui-tree-nodes="" ng-model="item.items">
            <li ng-repeat="item in item.items" ui-tree-node ng-include="'items_renderer.html'">
            </li>
          </ol>
        </script>
        <div ui-tree data-drag-enabled=false>
          <ol ui-tree-nodes="" ng-model="items">
            <li ng-repeat="item in items" ui-tree-node ng-include="'items_renderer.html'"></li>
          </ol>
        </div>
      </div>
    </div>
    <div class="col-md-9 mt-editor" ng-controller="Editor">
      <div class="container-fluid" ng-if="!error">
        <div class="col-md-6" style="white-space: pre;">${file.data}</div>
        <div class="col-md-6" marked="file.data"></div>
      </div>
      <div class="container-fluid" ng-if="error">
        <div class="col-md-12">
          <iframe class="editor-error" ng-init="initError()" style="width: 100%; height: 400px; border: 1px solid #eee; margin-bottom: 14px;"></iframe>
        </div>
      </div>
    </div>
  <div>
</div>
