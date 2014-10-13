<!DOCTYPE html>
<html>
    <head>
      <title>UserStack</title>
      <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

      <meta name="_xsrf" content="{{.xsrf_token}}" />
      <!-- bower:css -->
      <!-- endbower -->
      <link rel="stylesheet" href="/static/css/app.css" />
      <link rel=icon href="/static/images/logo.ico" sizes="128x128" type="image/vnd.microsoft.icon">
      <link rel=icon href="/static/images/logo.png" sizes="128x128" type="image/png">
  </head>
    <body class="{{.context.ControllerName}} {{.context.ActionName}}">
      {{if .loggedIn}}
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
              <li {{if .context.Is "MainController.Get" }}class="active"{{end}}><a href="{{urlfor "MainController.Get"}}">Home</a></li>
            </ul>
          </div>
        </div>
      </nav>
      <div class="container">
        {{end}}
        {{if .flash.error}}
          <div class="alert alert-danger">{{.flash.error}}</div>
        {{end}}
        {{if .flash.notice}}
          <div class="alert alert-info">{{.flash.notice}}</div>
        {{end}}
      </div>
      {{.LayoutContent}}
      <!-- bower:js -->
      <script src="../../static/bower_components/jquery/dist/jquery.js"></script>
      <script src="../../static/bower_components/bootstrap-sass/dist/js/bootstrap.js"></script>
      <!-- endbower -->
      <script src="../../static/scripts/app.js"></script>
    </body>
</html>
