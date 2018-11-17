<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>Geb's Pizza</title>
  <meta name="description" content="Coke IMC Welcome">
  <link rel="stylesheet" type="text/css" href="stylesheet.css">
  <link rel="stylesheet" type="text/css" href="material.css">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <script
      src="//code.jquery.com/jquery-3.2.1.min.js"
      integrity="sha256-hwg4gsxgFZhOsEEamdOYGBf13FyQuiTwlAQgxVSNgt4="
      crossorigin="anonymous"></script>
</head>
<body>

    <div id="menu" class='loginpage'>
        <div class="menu-top">
            <div class="imc-logo"><img src="images/imc-logo-circle.png"></div>
        </div>      
        <ul class="main-menu">
            <li style="background-image:url(images/log-in-icon-light.png);" class="login-icon"><a href="/login">Sign In</a></li>
        </ul>
        <a href="mailto:support@coke-imc.com?subject=IMC Content Portal Support" class="need-help"><img src="images/question-mark.png"> Need Help?</a>
    </div><!--/close #menu-->

    <div id="everything-right" class="portal-updates">
    
        <div id="header">
            <div class="header-title"><h1>Welcome</h1></div>
        </div>

        <div id="body" class="login-wrap">          
            <div class="wrapper">
                <div class="login-form">
                    <form id='register-form' action="{{mountpathed "register"}}" method="post">
                        <div class='group'>                                         
                            <input name="{{.primaryID}}" type="text" value="{{with .primaryIDValue}}{{.}}{{end}}" />
                            
                            <span class='hightlight'></span><span class='bar'></span>

                            <label for="{{.primaryID}}">{{title .primaryID}}:</label>
                            {{$pid := .primaryID}}{{with .errs}}{{with $errlist := index . $pid}}{{range $errlist}}<span class="error">{{.}}</span><br />{{end}}{{end}}{{end}}
                        </div>
                        <div class='group'>
                            <input type="password" name="password">
                            <span class='hightlight'></span><span class='bar'></span>

                            <label>Create Password:</label>
                            {{with .errs}}{{range .password}}<span class="error">{{.}}</span><br />{{end}}{{end}}
                        </div>
                        <div class='group'>
                            <input type="password" name="confirm_password">
                            <span class='hightlight'></span><span class='bar'></span>
                            <label>Confirm Password:<br /></label>
                            {{with .errs}}{{range .confirm_password}}<span class="error">{{.}}</span><br />{{end}}{{end}}
                        </div>
                        <button id='submit-button' type="submit" class="button">

                            Register
                            <div class="ripples buttonRipples"><span class="ripplesCircle"></span></div>
                        </button>
                        <input type="hidden" name="{{.xsrfName}}" value="{{.xsrfToken}}" />
                    </form>
                </div>
            </div>
        </div>  

        <div id="footer">
            <hr class="red-faded-border" />
            <div class="footer-content">
                <p>&copy; 2017 All Rights Reserved.</p>
            </div>
        </div>
    </div><!--/close #everything-right-->
    <script type="text/javascript" src='material.js'></script>
    <script type="text/javascript" src='main.js'></script>
</body>

</html>