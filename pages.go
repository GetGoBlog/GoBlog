package main

var base = `
 
<!DOCTYPE html>
<html lang="en">
	<head>
    	<meta charset="utf-8">
    	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
    	<title>GoBlog</title>
		<!-- Latest compiled and minified CSS -->
		<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css" integrity="sha384-1q8mTJOASx8j1Au+a5WDVnPi2lkFfwwEAa8hDDdjZlpLegxhjVME1fgjWPGmkzs7" crossorigin="anonymous">
		
		<!-- Optional theme -->
		<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap-theme.min.css" integrity="sha384-fLW2N01lMqjakBkx3l/M9EahuwpSfeNvV63J5ezn3uZzapT0u7EYsXMjQV+0En5r" crossorigin="anonymous">
		
		<!-- Latest compiled and minified JavaScript -->
		<script src="https://code.jquery.com/jquery-2.2.0.min.js"></script>
		<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/js/bootstrap.min.js" integrity="sha384-0mSbJDEHialfmuBBQP6A4Qrprq5OVfW37PRR3j5ELqxss1yVqOtnepnHVP9aJ7xS" crossorigin="anonymous"></script>
	</head>
	<body>
  <body>

<nav class="navbar navbar-default">
  <div class="container-fluid">
    <!-- Brand and toggle get grouped for better mobile display -->
    <div class="navbar-header">
      <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false">
        <span class="sr-only">Toggle navigation</span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
      </button>
      <a class="navbar-brand" href="/">GoBlog</a>
    </div>

    <!-- Collect the nav links, forms, and other content for toggling -->
    <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
      <ul class="nav navbar-nav">
        <li {{if eq .PageName "admin" }}class="active"{{end}}><a href="/admin">Admin</a></li>
      </ul>
      <form class="navbar-form navbar-left" role="search">
        <div class="form-group">
          <input type="text" class="form-control" placeholder="Search">
        </div>
        <button type="submit" class="btn btn-default">Submit</button>
      </form>
      <ul class="nav navbar-nav navbar-right">
        <li {{if eq .PageName "login" }}class="active"{{end}}><a href="/login">Login</a></li>
        <li {{if eq .PageName "signup" }}class="active"{{end}}><a href="/signup">Sign Up</a></li>
      </ul>
    </div><!-- /.navbar-collapse -->
  </div><!-- /.container-fluid -->
</nav>

    {{ template "content" }}
  </body>
</html>
 
`

var login = `
{{ define "content" }}
	  	<div class="container-fluid">
	  		<div class="page-header">
				<h1>Login</h1>
			</div>
			<form class="form-horizontal" action="/login/" method="POST">
			    <div class="form-group">
			        <label for="name" class="col-sm-2 control-label">User name</label>
			        <div class="col-sm-4">
			            <input type="text" class="form-control" id="name" name="name" required>
			        </div>
			    </div>
			    <div class="form-group">
			        <label for="password" class="col-sm-2 control-label">Password</label>
			        <div class="col-sm-4">
			            <input type="password" class="form-control" id="password" name="password" required>
			        </div>
			    </div>
			    <div class="col-sm-6">
			        <button type="submit" class="btn btn-primary pull-right">Login</button>
			    </div>
			</form>
		</div>
{{ end }}

`

var admin = `
{{define "content"}}
Admin
{{end}}
`

var signup = `
{{define "content"}}
Signup
{{end}}
`

var mainPage = `
{{define "content"}}
Main
{{end}}
`
