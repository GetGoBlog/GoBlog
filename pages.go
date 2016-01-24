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
      <ul class="nav navbar-nav navbar-right">
		<li {{if not .User}}class="hidden"{{end}}> <a href="/admin">{{if .User}}{{.User}}{{end}}</a> </li>
		<li {{if not .User}}class="hidden"{{end}}> <a href="/logout">Log Out</a> </li>
		<li {{if .User}}class="hidden"{{end}} {{if eq .PageName "signup" }}class="active"{{end}}><a href="/signup">Sign Up</a></li>
	<li {{if .User}}class="hidden"{{end}} class="dropdown">
		<a href="#" class="dropdown-toggle" data-toggle="dropdown">Sign in <b class="caret"></b></a>
		<ul class="dropdown-menu" style="padding: 15px;min-width: 250px;">
			<li>
				<div class="row">
					<div class="col-md-12">
						<form class="form" role="form" method="post" action="/login/" accept-charset="UTF-8" id="login-nav">
							<div class="form-group">
								<label class="sr-only" for="email">Email address</label>
								<input type="email" class="form-control" name="email" id="email" placeholder="Email address" required>
							</div>
							<div class="form-group">
								<label class="sr-only" for="password">Password</label>
								<input type="password" class="form-control" name="password" id="password" placeholder="Password" required>
							</div>
							<div class="checkbox">
								<label>
									<input type="checkbox"> Remember me
								</label>
							</div>
							<div class="form-group">
								<button type="submit" class="btn btn-success btn-block">Sign in</button>
							</div>
						</form>
					</div>
				</div>
			</li>
		<ul>
	</li>
      </ul>
    </div><!-- /.navbar-collapse -->
  </div><!-- /.container-fluid -->
</nav>

    {{ template "content" .}}
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
			        <label for="email" class="col-sm-2 control-label">User name</label>
			        <div class="col-sm-4">
			            <input type="text" class="form-control" id="email" name="email" required>
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
	
{{ if .Blogs }}
	<div class="col-md-6 col-md-offset-3">
		<div class="list-group">
			<div class="list-group-item active">Your Blogs</div>
			{{ range .Blogs }}
				<a href="http://{{.Website}}" class="list-group-item"> {{.Blogname}} </a>
			{{ end }}
		</div>
	</div>
{{ end }}

<div class="col-md-6 col-md-offset-3">

<form class="form-horizontal" action="/admin/" method="POST">
<fieldset>

<!-- Form Name -->
<legend>Create a new blog</legend>

<!-- Text input-->
<div class="form-group">
  <label class="col-md-4 control-label" for="blog">Blog Name</label>  
  <div class="col-md-6">
  <input id="blogname" name="blogname" type="text" placeholder="exampleblog" class="form-control input-md" required="">
    
  </div>
</div>

<!-- Text input-->
<div class="form-group">
  <label class="col-md-4 control-label" for="website">Blog Website</label>  
  <div class="col-md-6">
  <input id="blogname" name="website" type="text" placeholder="example.com" class="form-control input-md" required="">
    
  </div>
</div>

<!-- Button -->
<div class="form-group">
  <label class="col-md-4 control-label" for="submit"></label>
  <div class="col-md-4">
    <button id="submit" name="submit" class="btn btn-success">Create Blog</button>
  </div>
</div>

</fieldset>
</form>

</div>
{{end}}
`

var signup = `
{{define "content"}}
<div class="col-md-6 col-md-offset-3">

<form class="form-horizontal" action="/signup/" method="POST">
<fieldset>

<!-- Form Name -->
<legend>Sign Up for GoBlog</legend>

<!-- Text input-->
<div class="form-group">
  <label class="col-md-4 control-label" for="email">E-mail address</label>  
  <div class="col-md-6">
  <input id="email" name="email" type="email" placeholder="E-mail address" class="form-control input-md" required="">
    
  </div>
</div>

<!-- Password input-->
<div class="form-group">
  <label class="col-md-4 control-label" for="password">Password</label>
  <div class="col-md-6">
    <input id="password" name="password" type="password" placeholder="Password" class="form-control input-md" required="">
    
  </div>
</div>

<!-- Button -->
<div class="form-group">
  <label class="col-md-4 control-label" for="submit"></label>
  <div class="col-md-4">
    <button id="submit" name="submit" class="btn btn-success">Sign Me Up!</button>
  </div>
</div>

</fieldset>
</form>

</div>

{{end}}
`

var mainPage = `
{{define "content"}}
Main
{{end}}
`
