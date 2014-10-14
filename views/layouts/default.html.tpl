<!DOCTYPE html>
<html>
    <head>
    <title>Marktown</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

    <meta name="_xsrf" content="{{.xsrf_token}}" />
    <!-- bower:css -->
    <!-- endbower -->
    <link rel="stylesheet" href="/static/css/app.css" />
    <link rel=icon href="/static/images/logo.ico" sizes="128x128" type="image/vnd.microsoft.icon">
    <link rel=icon href="/static/images/logo.png" sizes="16x16 32x32 48x48 64x64" type="image/png">
  </head>
  <body class="{{.context.ControllerName}} {{.context.ActionName}}">
    <nav class="navbar navbar-default navbar-static-top" role="navigation">
      <div class="container">
        <div class="navbar-header">
          <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <a class="navbar-brand" href="#"></a>
        </div>
        <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
          <ul class="nav navbar-nav">
            <li {{if .context.Is "HomeController.Get" }}class="active"{{end}}><a href="{{urlfor "HomeController.Get"}}">Home</a></li>
          </ul>
        </div>
      </div>
    </nav>
		{{.LayoutContent}}
    <!-- bower:js -->
    <script src="../../static/bower_components/jquery/dist/jquery.js"></script>
    <script src="../../static/bower_components/bootstrap-sass/dist/js/bootstrap.js"></script>
    <!-- endbower -->
    <script src="../../static/js/app.js"></script>
    {{if eq .RunMode "dev"}}<script src="//localhost:35729/livereload.js"></script>{{end}}
  </body>
</html>
