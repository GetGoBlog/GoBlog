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
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bootswatch/3.3.6/paper/bootstrap.min.css">
		
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
        <li {{if not .User}}class="hidden"{{end}}{{if eq .PageName "admin" }}class="active"{{end}}><a href="/admin">Admin</a></li>
      </ul>
      <ul class="nav navbar-nav navbar-right">
		<li {{if not .User}}class="hidden"{{end}}> <a href="/admin">{{if .User}}{{.User}}{{end}}</a> </li>
		<li {{if not .User}}class="hidden"{{end}}> <a href="/logout">Log Out</a> </li>
		<li {{if .User}}class="hidden"{{end}} {{if eq .PageName "signup" }}class="active"{{end}}><a href="/signup">Sign Up</a></li>
	<li {{if .User}}class="hidden"{{end}} class="dropdown">
		<a href="/login">Sign in</a>
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
		{{if .Success}}
			<div class="alert alert-success">
				<h1>Success: Your blog, {{.Success}}, has been created!</h1>
			</div>
		{{end}}
	
		<div class="list-group">
			<div class="list-group-item active">Your Blogs</div>
			{{ range .Blogs }}
				<li class="list-group-item"> 
					{{.Blogname}} 
					<div class="btn-group pull-right"> 
						<a class="btn btn-success" href="http://{{.Website}}" target="_blank">View Blog</a> 
						<a class="btn btn-info" href="http://{{.Website}}/admin" target="_blank">Blog Admin</a> 
					</div>
				</li>
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
  <div class="col-sm-6">
  <input id="blogname" name="blogname" type="text" placeholder="example" class="form-control input-md" required="">
</div>
<div class="col-sm-6">
<p>.goblog.pw</p>
</div>
</div>

<!-- Text input-->
<!--
<div class="form-group">
  <label class="col-md-4 control-label" for="website">Blog Website</label>  
  <div class="col-md-6">
  <input id="blogname" name="website" type="text" placeholder="example.com" class="form-control input-md" required="">
    
  </div>
</div>
-->

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

var errorPage = `
{{define "content"}}

<div class="col-md-6 col-md-offset-3">
	<div class="alert alert-danger" role="alert">
	  <span class="glyphicon glyphicon-exclamation-sign" aria-hidden="true"></span>
	  <span class="sr-only">Error:</span>
	  {{.Error}}
	</div>
</div>

{{end}}
`

var mainPage = `
{{define "content"}}

<div class="row">
	<div class="col-md-5 col-md-offset-1">
		<h1>GoBlog</h1>
		<h2>Simple Free Blog Hosting</h2>
		<hr>
		<h3>Deploy your own blog in as little as 15 seconds</h3>
	</div>
	<div class="col-md-6">
		<img src="https://placeholdit.imgix.net/~text?txtsize=100&txt=Animated%20GIF%20goes%20here&w=600&h=500" style="width:100%">
	</div>
</div>

<hr>

<div class="row text-center">
	<h2>Simple to Set Up</h2>
	<div class="col-md-4">
		<h3>Easy to Deploy</h3>
		<i class="glyphicon glyphicon-ok" style="font-size: 20vw; color: #2196f3"></i>
	</div>
	<div class="col-md-4">
		<h3>Upload a Custom Theme</h3>
		<i class="glyphicon glyphicon-cloud-upload" style="font-size: 20vw; color: #2196f3"></i>
	</div>
	<div class="col-md-4">
		<h3>Use Your Own Domain Name</h3>
		<i class="glyphicon glyphicon-globe" style="font-size: 20vw; color: #2196f3"></i>
	</div>

</div>

{{end}}
`

var newMainPage = `

<!DOCTYPE html>
<html lang="en">

<head>
	<meta charset="UTF-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title>GoBlog</title>
	<meta name="description" content="GoBlog is a platform for users to easily deploy their blog." />
	<meta name="keywords" content="free, ghost, journey, theme, blog" />
	<!-- Favicons (created with http://realfavicongenerator.net/)-->
	<link rel="apple-touch-icon" sizes="57x57" href="img/favicons/apple-touch-icon-57x57.png">
	<link rel="apple-touch-icon" sizes="60x60" href="img/favicons/apple-touch-icon-60x60.png">
	<link rel="icon" type="image/png" href="img/favicons/favicon-32x32.png" sizes="32x32">
	<link rel="icon" type="image/png" href="img/favicons/favicon-16x16.png" sizes="16x16">
	<link rel="manifest" href="img/favicons/manifest.json">
	<link rel="shortcut icon" href="img/favicons/favicon.ico">
	<meta name="msapplication-TileColor" content="#00a8ff">
	<meta name="msapplication-config" content="img/favicons/browserconfig.xml">
	<meta name="theme-color" content="#ffffff">
	<!-- Normalize -->
	<link rel="stylesheet" type="text/css" href="css/normalize.css">
	<!-- Bootstrap -->
	<link rel="stylesheet" type="text/css" href="css/bootstrap.css">
	<!-- Owl -->
	<link rel="stylesheet" type="text/css" href="css/owl.css">
	<!-- Animate.css -->
	<link rel="stylesheet" type="text/css" href="css/animate.css">
	<!-- Font Awesome -->
	<link rel="stylesheet" type="text/css" href="fonts/font-awesome-4.1.0/css/font-awesome.min.css">
	<!-- Elegant Icons -->
	<link rel="stylesheet" type="text/css" href="fonts/eleganticons/et-icons.css">
	<!-- Main style -->
	<link rel="stylesheet" type="text/css" href="css/cardio.css">
</head>

<body>
	<div class="preloader">
		<img src="img/loader.gif" alt="Preloader image">
	</div>
	<nav class="navbar">
		<div class="container">
			<!-- Brand and toggle get grouped for better mobile display -->
			<div class="navbar-header">
				<button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1">
					<span class="sr-only">Toggle navigation</span>
					<span class="icon-bar"></span>
					<span class="icon-bar"></span>
					<span class="icon-bar"></span>
				</button>
				<a class="navbar-brand" href="#"><img src="img/logo.png" data-active-url="img/logo-active.png" alt=""></a>
			</div>
			<!-- Collect the nav links, forms, and other content for toggling -->
			<div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
				<ul class="nav navbar-nav navbar-right main-nav">
					<li><a href="#intro">Intro</a></li>
					<li><a href="#services">Features</a></li>
					<li><a href="#team">Team</a></li>
					<li><a href="#pricing">Pricing</a></li>
					<li><a href="/login/">Login</a></li>
					<li><a href="#" data-toggle="modal" data-target="#modal1" class="btn btn-blue">Sign Up</a></li>
				</ul>
			</div>
			<!-- /.navbar-collapse -->
		</div>
		<!-- /.container-fluid -->
	</nav>
	<header id="intro">
		<div class="container">
			<div class="table">
				<div class="header-text">
					<div class="row">
						<div class="col-md-12 text-center">
							<h3 class="light white">Welcome to GoBlog</h3>
							<h1 class="white typed">Deploy your own blog in seconds.</h1>
							<span class="typed-cursor">|</span>
						</div>
					</div>
				</div>
			</div>
		</div>
	</header>
	
	<section id="services" class="section section-padded">
		<div class="container">
			<div class="row text-center title">
				<h2>Features</h2>
				<h4 class="light muted">Here is why we are different.</h4>
			</div>
			<div class="row services">
				<div class="col-md-4">
					<div class="service">
						<div class="icon-holder">
							<img src="img/icons/heart-blue.png" alt="" class="icon">
						</div>
						<h4 class="heading">Open Source</h4>
						<p class="description">All of our source code is available on GitHub.</p>
					</div>
				</div>
				<div class="col-md-4">
					<div class="service">
						<div class="icon-holder">
							<img src="img/icons/guru-blue.png" alt="" class="icon">
						</div>
						<h4 class="heading">Easy Setup</h4>
						<p class="description">Deploying your own blog takes seconds.</p>
					</div>
				</div>
				<div class="col-md-4">
					<div class="service">
						<div class="icon-holder">
							<img src="img/icons/weight-blue.png" alt="" class="icon">
						</div>
						<h4 class="heading">Modern Themes</h4>
						<p class="description">With the ability to upload any Ghost compatible theme, you have thousands of choices!</p>
					</div>
				</div>
			</div>
		</div>
		<div class="cut cut-bottom"></div>
	</section>
	<section id="team" class="section gray-bg">
		<div class="container">
			<div class="row title text-center">
				<h2 class="margin-top">Team</h2>
				<h4 class="light muted">We worked together to build GoBlog</h4>
			</div>
			<div class="row">
				<div class="col-md-4">
					<div class="team text-center">
						<img src="img/team/faraz.png" alt="Team Image" class="avatar">
						<div class="title">
							<h4>Faraz Fazli</h4>
							<h5 class="muted regular">Backend Golang Developer</h5>
						</div>
					</div>
				</div>
				<div class="col-md-4">
					<div class="team text-center">
						<img src="img/team/zain.jpg" alt="Team Image" class="avatar">
						<div class="title">
							<h4>Zain Hoda</h4>
							<h5 class="muted regular">Frontend Developer</h5>
						</div>
					</div>
				</div>
				<div class="col-md-4">
					<div class="team text-center">
						<img src="img/team/anne.jpg" alt="Team Image" class="avatar">
						<div class="title">
							<h4>Anne Moroney</h4>
							<h5 class="muted regular">UX Developer</h5>
						</div>
					</div>
				</div>
			</div>
		</div>
	</section>
	<section id="pricing" class="section">
		<div class="container">
			<div class="row title text-center">
				<h2 class="margin-top white">Pricing</h2>
				<h4 class="light white">Choose your favorite pricing plan and sign up today!</h4>
			</div>
			<div class="row no-margin">
				<div class="col-md-7 no-padding col-md-offset-5 pricings text-center">
					<div class="pricing">
						<div class="box-main active" data-img="img/cover1.jpg">
							<h4 class="white">GoBlog Beta</h4>
							<h4 class="white regular light">$0.00 <span class="small-font">/ life</span></h4>
							<a href="#" data-toggle="modal" data-target="#modal1" class="btn btn-white-fill">Sign Up Now</a>
							<i class="info-icon icon_question"></i>
						</div>
						<div class="box-second active">
							<ul class="white-list text-left">
								<li>Unlimited Blogs</li>
								<li>Unlimited Views</li>
								<li>Unlimited Transfer/Storage</li>
								<li>Premium DDOS Protection</li>
								<li>Worldwide CDN</li>
							</ul>
						</div>
					</div>
				</div>
			</div>
		</div>
	</section>
	<section class="section section-padded blue-bg">
		<div class="container">
			<div class="row">
				<div class="col-md-8 col-md-offset-2">
					<div class="owl-twitter owl-carousel">
						<div class="item text-center">
							<i class="icon fa fa-twitter"></i>
							<h4 class="white light">To enjoy the glow of good health, you must exercise.</h4>
							<h4 class="light-white light">#health #training #exercise</h4>
						</div>
						<div class="item text-center">
							<i class="icon fa fa-twitter"></i>
							<h4 class="white light">To enjoy the glow of good health, you must exercise.</h4>
							<h4 class="light-white light">#health #training #exercise</h4>
						</div>
						<div class="item text-center">
							<i class="icon fa fa-twitter"></i>
							<h4 class="white light">To enjoy the glow of good health, you must exercise.</h4>
							<h4 class="light-white light">#health #training #exercise</h4>
						</div>
					</div>
				</div>
			</div>
		</div>
	</section>
	<div class="modal fade" id="modal1" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
		<div class="modal-dialog">
			<div class="modal-content modal-popup">
				<a href="#" class="close-link"><i class="icon_close_alt2"></i></a>
				<h3 class="white">Sign Up</h3>
				<form action="/signup/" method="POST" class="popup-form">
					<input name="email" type="text" class="form-control form-white" placeholder="Email Address">
					<input name="password" type="password" class="form-control form-white" placeholder="Password">
					<div class="checkbox-holder text-left">
						<div class="checkbox">
							<input type="checkbox" value="None" id="squaredOne" name="check" />
							<label for="squaredOne"><span>I Agree to the <strong>Terms &amp; Conditions</strong></span></label>
						</div>
					</div>
					<button type="submit" class="btn btn-submit">Submit</button>
				</form>
			</div>
		</div>
	</div>
	<footer>
		<div class="container">
			<div class="row bottom-footer text-center-mobile">
				<div class="col-sm-8">
					<p>&copy; 2015 GoBlog All Rights Reserved. Powered by <a href="http://www.phir.co/">PHIr</a> exclusively for <a href="http://tympanus.net/codrops/">Codrops</a></p>
				</div>
				<div class="col-sm-4 text-right text-center-mobile">
					<ul class="social-footer">
						<li><a href="#"><i class="fa fa-facebook"></i></a></li>
						<li><a href="#"><i class="fa fa-twitter"></i></a></li>
						<li><a href="#"><i class="fa fa-google-plus"></i></a></li>
					</ul>
				</div>
			</div>
		</div>
	</footer>
	<!-- Holder for mobile navigation -->
	<div class="mobile-nav">
		<ul>
		</ul>
		<a href="#" class="close-link"><i class="arrow_up"></i></a>
	</div>
	<!-- Scripts -->
	<script src="js/jquery-1.11.1.min.js"></script>
	<script src="js/owl.carousel.min.js"></script>
	<script src="js/bootstrap.min.js"></script>
	<script src="js/wow.min.js"></script>
	<script src="js/typewriter.js"></script>
	<script src="js/jquery.onepagenav.js"></script>
	<script src="js/main.js"></script>
</body>

</html>


`
