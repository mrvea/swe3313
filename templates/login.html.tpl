<!DOCTYPE HTML>
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
        <meta name="viewport" content="user-scalable=no, initial-scale=1.0, maximum-scale=1.0 minimal-ui"/>
        <meta name="apple-mobile-web-app-capable" content="yes"/>
        <meta name="apple-mobile-web-app-status-bar-style" content="black">

        <title>Pizza</title>
        <link rel="stylesheet" type="text/css" href="/styles/style.css">
        <link rel="stylesheet" type="text/css" href="/styles/framework.css">
        <link rel="stylesheet" type="text/css" href="/styles/font-awesome.css">
        <link rel="stylesheet" type="text/css" href="/styles/animate.css">

        <script type="text/javascript" src="/js/jquery.js"></script>
        <!-- <script type="text/javascript" src="/js/framework.plugins.js"></script> -->
        <!-- <script type="text/javascript" src="/js/custom.js"></script> -->
    </head>
    <body>

    <div id="preloader">
        <div id="status">
            <p class="center-text">
                Loading the content...
                <em>Loading depends on your connectiocd n speed!</em>
            </p>
        </div>
    </div>

    <div class="all-elements">
        <!-- Page Content-->
        <div id="content" class="snap-content">
            <div class="header">
                <a href="#" class="sidebar-deploy"></a>
                <h3><span>Pizza</span><img src="/images/logo-topbar.png" /></h3>
            </div>
            <div class="decoration"></div>
            <div class="row login-page bgwhite">
				<div class="login-page-wrapper">
					<div>
						<div>Please sign in to continue!</div>
						<div>
							{{if .error}}
							<div class="alert alert-danger">{{.error}}</div>
							{{end}}
							<form method="POST">
									<input type="text" id="email" class="form-control login-username" name="{{.primaryID}}" placeholder="{{title .primaryID}}" value="{{.primaryIDValue}}">
									<input id="password" type="password" class="form-control login-password" name="password" placeholder="Password">
								<input type="hidden" name="{{.xsrfName}}" value="{{.xsrfToken}}" />
								<div class="row">
									<div class="col-md-offset-1 col-md-10">
										<button class="button" type="submit">Login</button>
									</div>
								</div>
								{{if .showRecover}}
										<a class="btn btn-link btn-block" href="{{mountpathed "recover"}}">Recover Account</a>
								{{end}}
							</form>
						</div>
					</div>
				</div>
			</div>
        </div>
    </div>
    <script>
        (function ($) {
            $(window).load(function() {
                $("#status").fadeOut(); // will first fade out the loading animation
                $("#preloader").delay(400).fadeOut("medium"); // will fade out the white DIV that covers the website.
            });

            // $(document).ready(function() {

            //     //Remove 300ms lag set by -webkit-browsers
            //     window.addEventListener('load', function() {
            //         FastClick.attach(document.body);
            //     }, false);
            // });
        })(jQuery);
    </script>
    </body>
</html>
