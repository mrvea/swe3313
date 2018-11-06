<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>Coke IMC Welcome</title>
  <meta name="description" content="Coke IMC Welcome">
   <link rel="stylesheet" href="/assets/outdated-browser/outdatedbrowser.min.css">
  <link rel="stylesheet" type="text/css" href="stylesheet.css">
  <link rel="stylesheet" type="text/css" href="juan.css">
  <link rel="stylesheet" type="text/css" href="welcome-portal-respo.css">
  <link rel="stylesheet" type="text/css" href="material.css">
  <link rel="icon" type="image/png" href="/favicon.ico?v=1" />
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <script
      src="//code.jquery.com/jquery-3.2.1.min.js"
      integrity="sha256-hwg4gsxgFZhOsEEamdOYGBf13FyQuiTwlAQgxVSNgt4="
      crossorigin="anonymous"></script>
      <script src="/vex/js/vex.combined.min.js"></script>
    <script>vex.defaultOptions.className = 'vex-theme-os eye-mod'</script>
    <link rel="stylesheet" href="/vex/css/vex.css" />
    <link rel="stylesheet" href="/vex/css/vex-theme-os.css" />
</head>
<body>

    <div id="menu" class="loginpage">
        <div class="menu-top">
            <div class="imc-logo"><img src="images/imc-logo-circle.png"></div>
        </div>      
        <ul class="main-menu">
            <li style="background-image:url(images/log-in-icon-light.png);" class="login-icon"><a href="#">Sign IN</a></li>
        </ul>
        <a href="mailto:support@coke-imc.com?subject=IMC Content Portal Support" class="need-help"><img src="images/question-mark.png"> Need Help?</a>
    </div><!--/close #menu-->

    <div id="everything-right" class="portal-updates">
    
        <div id="header">
            <div class="header-title"><h1>Sign In</h1></div>
        </div>

        <div id="body" class="login-wrap">          
            <div class="wrapper">
                {{if .flash_success}}
                    <div class="message">
                        {{.flash_success}}
                    </div>
                {{end}}
                {{if .error}}
                <div class="message">
                    {{.error}}
                </div>
                {{end}}
                <div class="login-text">
                    
                    <div class='actions'>
                        {{if .showRegister}}<a href="{{mountpathed "register"}}" class="button grey">Register Now<div class="ripples buttonRipples"><span class="ripplesCircle"></span></div></a>{{end}}
                    </div>
                </div>
                <div class="login-form">
                    <!-- <div class="login-form-head"><h2>Sign In</h2></div> -->
                    <form action="{{mountpathed "login"}}" method="POST">
                        <div class='group'>
                                <input type="text" name="{{.primaryID}}" value={{.primaryIDValue}} >
                                <span class='hightlight'></span><span class='bar'></span>
                            <label>{{title .primaryID}}:</label>
                        </div>
                        <div class='group'>
                                <input type="password" name="password" />
                            <span class='hightlight'></span><span class='bar'></span>
                            <label>Password:</label>
                        </div>
                        <div class="remember-box">
                            <div class="remember-me">       
                                {{if .showRemember}}<input type="checkbox" name="rm" value="true"> Remember Me{{end}}   
                            </div>
                        </div>
                        <div class="lateral-button">
                            <button type="submit" name="Sign In" class="button">Sign In
                                <div class="ripples buttonRipples"><span class="ripplesCircle"></span></div>
                            </button>
                        </div>
                        <div class="lateral-button">
                            {{if .showRecover}}<a href="{{mountpathed "recover"}}" class="recover">Reset Password
                            <div class="ripples buttonRipples"><span class="ripplesCircle"></span></div>
                            </a>{{end}}
                        </div>
                        <br />
                        <input type="hidden" name="{{.xsrfName}}" value="{{.xsrfToken}}" />
                        <button style="display: none;" id='downloadApp' class="button grey">Download Mobile App
                            <div class="ripples buttonRipples"><span class="ripplesCircle"></span></div>
                        </button>
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
        <script src="/assets/outdated-browser/outdatedbrowser.js"></script>
        <script>
        //event listener: DOM ready
        function addLoadEvent(func) {
            var oldonload = window.onload;
            if (typeof window.onload != 'function') {
                window.onload = func;
            } else {
                window.onload = function() {
                    if (oldonload) {
                        oldonload();
                    }
                    func();
                }
            }
        }
        //call plugin function after DOM ready
        addLoadEvent(function(){
            // if(document.cookie.indexOf('checked') === -1){
                outdatedBrowser({
                    bgColor: '#cc0000',
                    color: '#ffffff',
                    lowerThan: 'js:Promise',
                    languagePath: '/assets/outdated-browser/lang/en.html'
                });
                // document.cookie = "checked=1";
            // }
        });
    </script>
    <div id="outdated" style="display:none">

    </div>
    <script type="text/javascript" src='material.js'></script>
    <script type="text/javascript" src='mobile-app-button.js'></script>
</body>
</html>