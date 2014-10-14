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
      <link rel=icon href="/static/images/logo.png" sizes="128x128" type="image/png">
  </head>
    <body class="{{.context.ControllerName}} {{.context.ActionName}}">
		<div class="navbar">
		</div>
		<div class="file-browser">
		</div>
		<div class="content">
			{{.LayoutContent}}
		</div>
		<footer>Copyright &copy; 2014 {{.Authors}}</footer>
      <!-- bower:js -->
      <script src="../../static/bower_components/jquery/dist/jquery.js"></script>
      <script src="../../static/bower_components/bootstrap-sass/dist/js/bootstrap.js"></script>
      <!-- endbower -->
      <script src="../../static/js/app.js"></script>
      {{if eq .RunMode "dev"}}<script src="//localhost:35729/livereload.js"></script>{{end}}
    </body>
</html>
