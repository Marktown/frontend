<div class="container">
  <h1>home-index</h1>
  <div ng-app="Main" ng-controller="Home">
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
