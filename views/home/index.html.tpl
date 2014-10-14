<div class="container">
  <h1>home-index</h1>
  <div ng-app="Main" ng-controller="Home">
    <div ui-tree>
      <ol ui-tree-nodes="" ng-model="list">
        <li ng-repeat="item in list" ui-tree-node>
          <div ui-tree-handle>
            ${item.title}
          </div>
          <ol ui-tree-nodes="" ng-model="item.items">
            <li ng-repeat="subItem in item.items" ui-tree-node>
              <div ui-tree-handle>
                ${subItem.title}
              </div>
            </li>
          </ol>
        </li>
      </ol>
    </div>
  </div>
</div>
