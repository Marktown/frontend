<!DOCTYPE html>
<html>
    <head>
    <title>Marktown</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

    <meta name="_xsrf" content="{{.xsrf_token}}" />
    <!-- bower:css -->
    <link rel="stylesheet" href="../../static/bower_components/angular-ui-tree/dist/angular-ui-tree.min.css" />
    <!-- endbower -->
    <link rel="stylesheet" href="/static/css/app.css" />
    <link rel=icon href="/static/images/logo.ico" sizes="128x128" type="image/vnd.microsoft.icon">
    <link rel=icon href="/static/images/logo.png" sizes="16x16 32x32 48x48 64x64" type="image/png">
  </head>
  <body class="{{.context.ControllerName}} {{.context.ActionName}}">
    {{template "shared/navigation.html.tpl" .}}
    {{.LayoutContent}}
    <!-- bower:js -->
    <script src="../../static/bower_components/es5-shim/es5-shim.js"></script>
    <script src="../../static/bower_components/jquery/dist/jquery.js"></script>
    <script src="../../static/bower_components/bootstrap-sass/dist/js/bootstrap.js"></script>
    <script src="../../static/bower_components/angular/angular.js"></script>
    <script src="../../static/bower_components/angular-ui-tree/dist/angular-ui-tree.js"></script>
    <script src="../../static/bower_components/marked/lib/marked.js"></script>
    <script src="../../static/bower_components/angular-marked/angular-marked.js"></script>
    <!-- endbower -->
    <script src="../../static/js/app.js"></script>
    {{if eq .RunMode "dev"}}<script src="//localhost:35729/livereload.js"></script>{{end}}
  </body>
</html>
