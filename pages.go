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
			</div>
			<form class="form-horizontal" action="/login" method="POST">
			    <div class="form-group">
			        <label for="email" class="col-sm-2 control-label">Email</label>
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
			<div class="alert alert-success">
			<!-- This is needed. We are gaining users but they are confused on what to do next. -->
				Your blog account is separate from your main account. Click "Blog Admin" to get started on your blog!
			</div>
	<div class="col-md-6 col-md-offset-3">
		{{if .Success}}
			<div class="alert alert-success">
				<h1>Success: Your blog, {{.Success}}, has been created!</h1>
			</div>
		{{end}}
		</div>
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

<form class="form-horizontal" action="/admin" method="POST">
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

<form class="form-horizontal" action="/signup" method="POST">
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
	<form class="form-horizontal" action="/" method="GET">
			    <div class="col-sm-6">
			        <button type="submit" class="btn btn-primary pull-right">Go Back</button>
			    </div>
			</form>
</div>

{{end}}
`

var newMainPage = `
<!DOCTYPE html>
<html class="gr__codeply_com"><head>
<meta http-equiv="content-type" content="text/html; charset=UTF-8">
  <meta charset="utf-8">
  <title>GoBlog</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta name="google" value="notranslate">



  <link rel="stylesheet" href="css/bootstrap.css">
  <link href="css/animate.css" rel="stylesheet">
  <link rel="stylesheet" href="css/main.css">

  
</head>
<body class="bg-faded">
  <header id="top">
    <div class="header-content">
        <div class="inner">
            <h2 class="text-primary wow fadeIn">Welcome to GoBlog</h2>
            <h5 style="visibility: visible; animation-name: fadeIn;" class="wow fadeIn text-normal">Endless writing, completely free.</h5>
            <hr>
            <a style="visibility: visible; animation-name: fadeInUp;" href="/signup" class="btn btn-primary-outline btn-xl page-scroll wow fadeInUp">Get Started</a>
        </div>
    </div>
</header>
<nav id="topNav" class="navbar navbar-default navbar-fixed-top">
    <div class="container">
        <button class="navbar-toggler hidden-md-up pull-right" type="button" data-toggle="collapse" data-target="#collapsingNavbar">
            ☰
        </button>
        <a class="navbar-brand page-scroll" href="#top">GoBlog</a>
        <div class="collapse navbar-toggleable-sm" id="collapsingNavbar">
            <ul class="nav navbar-nav">
                <li class="nav-item">
                    <a class="nav-link page-scroll" href="#team">Team</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link page-scroll" href="#contact">Contact</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link page-scroll" href="/login">Login</a>
                </li>
				<li class="nav-item">
                    <a class="nav-link page-scroll" href="/signup">Sign Up</a>
                </li>
            </ul>
        </div>
    </div>
</nav>
<section id="team">
       <div style="visibility: hidden; animation-name: none;" class="container wow fadeInUp">
        <h2 class="text-primary wow fadeIn">Our Team</h2>
        <p class="lead">
            Together, we build GoBlog, a powerful free open-source blogging platform written in Go.
        </p>
        <div class="card-group">
            <!-- card -->
            <div class="card">
               <center><img class="card-img-top img-fluid" style="padding: 30px 0px 0px 0px;" src="img/faraz.jpg" alt="Faraz"></center>
                <div class="card-block">
				<center>
                    <h4 class="card-title">Faraz Fazli</h4>
                    <p class="card-text">Backend Go Developer</p>
				</center>
                </div>
            </div>
            <!-- card -->
            <div class="card">
                <center><img class="card-img-top img-fluid" style="padding: 30px 0px 0px 0px;" src="img/zain.jpg" alt="Zain"></center>
                <div class="card-block">
				<center>
                    <h4 class="card-title">Zain Hoda</h4>
                    <p class="card-text">Frontend Developer</p>
				</center>
                </div>
            </div>
            <!-- card -->
            <div class="card">
                <center><img class="card-img-top img-fluid" style="padding: 30px 0px 0px 0px;" src="img/anne.jpg" alt="Anne"></center>
                <div class="card-block">
				<center>
                    <h4 class="card-title">Anne Moroney</h4>
                    <p class="card-text">DevOps Engineer</p>
                </center> 
                </div>
            </div>
        </div>
    </div>
</section>
<section id="contact">
<aside class="bg-alt">
    <div class="container text-xs-center">
        <div class="call-to-action">
            <h2 class="text-primary wow fadeIn">What are you waiting for?</h2>
            <a style="visibility: hidden; animation-name: none;" href="/signup" class="btn btn-primary-outline btn-lg wow flipInX text-uppercase">Signup today!</a>
        </div>
</aside>
    <div class="container">
        <div class="row">
            <div class="col-lg-8 col-lg-offset-2 text-xs-center">
			<hr>
                <h2 class="text-primary wow fadeIn">Contact Us</h2>
                <hr class="primary">
                <p>Want to get in touch with us? Send us an email using the form below and we'll get back to you.</p>
            </div>
            <div class="col-lg-10 col-lg-offset-1 text-xs-center">
                <form class="contact-form row" action="http://formspree.io/farazfazli@gmail.com" method="POST">
                    <div class="col-md-4">
                        <input type="text" name="name" class="form-control" placeholder="Name">
                    </div>
                    <div class="col-md-4">
                        <input type="email" name="_replyto" class="form-control" placeholder="Email">
                    </div>
                    <div class="col-md-4">
                        <input type="text" name="_subject" class="form-control" placeholder="Subject">
                    </div>
                    <div class="col-md-12">
                        <label></label>
                        <textarea name="message" class="form-control" rows="9" placeholder="Your message here.."></textarea>
						<input type="text" name="_gotcha" style="display:none"/>
                    </div>
                    <div class="col-md-4 col-md-offset-4">
                        <br />
                        <button type="submit" value="Send" class="btn btn-primary-outline btn-lg wow flipInX text-uppercase">Send</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</section>
<footer id="footer">
    <div class="container">
        <center><span class="pull-right text-muted small">Bootstrap 4 + Faraz Fazli</span></center>
    </div>
</footer>
  
  <script>
    // sandbox disable popups
    if (window.self !== window.top && window.name!="view1") {;
      window.alert = function(){/*disable alert*/};
      window.confirm = function(){/*disable confirm*/};
      window.prompt = function(){/*disable prompt*/};
      window.open = function(){/*disable open*/};
    }
    
    // prevent href=# click jump
    document.addEventListener("DOMContentLoaded", function() {
      var links = document.getElementsByTagName("A");
      for(var i=0; i < links.length; i++) {
        if(links[i].href.indexOf('#')!=-1) {
          links[i].addEventListener("click", function(e) {
          console.debug("prevent href=# click");
              if (this.hash) {
                if (this.hash=="#") {
                  e.preventDefault();
                  return false;
                }
                else {
                  /*
                  var el = document.getElementById(this.hash.replace(/#/, ""));
                  if (el) {
                    el.scrollIntoView(true);
                  }
                  */
                }
              }
              return false;
          })
        }
      }
    }, false);
    
  </script>
  
  <!--scripts loaded here-->
  
  <script src="js/jquery.js"></script>
  <script src="js/bootstrap.js"></script>
  
  
  <script src="js/jquery_002.js"></script>
  <script src="js/wow.js"></script>
  
  <script>
  (function($) {
    "use strict";

    $('body').scrollspy({
        target: '.navbar-fixed-top',
        offset: 60
    });

    new WOW().init();
    
    $('a.page-scroll').bind('click', function(event) {
        var $ele = $(this);
        $('html, body').stop().animate({
            scrollTop: ($($ele.attr('href')).offset().top - 60)
        }, 1450, 'easeInOutExpo');
        event.preventDefault();
    });
    
    $('#collapsingNavbar li a').click(function() {
        /* always close responsive nav after click */
        $('.navbar-toggler:visible').click();
    });

    $('#galleryModal').on('show.bs.modal', function (e) {
       $('#galleryImage').attr("src",$(e.relatedTarget).data("src"));
    });

})(jQuery);
  </script>



</body>
</html>
`
