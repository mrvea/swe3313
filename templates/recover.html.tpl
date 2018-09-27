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
        <link rel="stylesheet" type="text/css" href="/styles/owl.theme.css">
        <link rel="stylesheet" type="text/css" href="/styles/swipebox.css">
        <link rel="stylesheet" type="text/css" href="/styles/font-awesome.css">
        <link rel="stylesheet" type="text/css" href="/styles/animate.css">

        <script type="text/javascript" src="/js/jquery.js"></script>
        <script type="text/javascript" src="/js/jqueryui.js"></script>
        <script type="text/javascript" src="/js/framework.plugins.js"></script>
        <script type="text/javascript" src="/js/custom.js"></script>
        <script type="text/javascript" src="/js/jquery-ui-tabs-and-date"></script>
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
						<div class="panel-heading">Account Recover</div>
						<div class="panel-body">
							<form method="POST">
								{{$pid := .primaryID}}
								<div class="form-group {{with .errs}}{{with $errlist := index . $pid}}has-error{{end}}{{end}}">
									<input type="text" class="form-control contactField" name="{{.primaryID}}" placeholder="{{title .primaryID}}" value="{{.primaryIDValue}}" />
									{{with .errs}}{{with $errlist := index . $pid}}{{range $errlist}}<span class="help-block">{{.}}</span>{{end}}{{end}}{{end}}
								</div>
								{{$cpid := .primaryID | printf "confirm_%s"}}
								<div class="form-group {{with .errs}}{{with $errlist := index . $cpid}}has-error{{end}}{{end}}">
									<input type="text" class="form-control contactField" name="confirm_{{.primaryID}}" placeholder="Confirm {{title .primaryID}}" value="{{.confirmPrimaryIDValue}}" />
									{{with .errs}}{{with $errlist := index . $cpid}}{{range $errlist}}<span class="help-block">{{.}}</span>{{end}}{{end}}{{end}}
								</div>
								<input type="hidden" name="{{.xsrfName}}" value="{{.xsrfToken}}" />
								<div class="row">
									<div class="col-md-offset-1 col-md-10">
										<button class="button" type="submit">Recover</button>
									</div>
								</div>
								<div class="row">
									<div class="col-md-offset-1 col-md-10">
										<a class="button" href="{{mountpathed "login"}}">Cancel</a>
									</div>
								</div>
							</form>
						</div>
					</div>
				</div>
			</div>
        </div>
    </div>
    </body>
</html>