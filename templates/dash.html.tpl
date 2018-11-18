<!DOCTYPE html>
<html>
    <head>
        <base href="/">
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=Edge"/>
        <!-- <meta http-equiv="Content-Type" content="text/html; charset=utf-8" /> -->
        <meta name="viewport" content="user-scalable=no, initial-scale=1.0, maximum-scale=1.0 minimal-ui"/>
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <meta name="apple-mobile-web-app-capable" content="yes"/>
        <meta name="apple-mobile-web-app-status-bar-style" content="black">

        <title>Geb's Pizza</title>
        <!-- <link rel="stylesheet" type="text/css" href="/styles/framework.css"> -->
        <link rel="stylesheet" type="text/css" href="/styles/font-awesome.css">
        <!-- <link rel="stylesheet" type="text/css" href="/styles/animate.css"> -->
        <!-- <link rel="stylesheet" type="text/css" href="/styles/grids.inuit.css"> -->
        <link href="https://unpkg.com/@angular/material/prebuilt-themes/indigo-pink.css" rel="stylesheet">
        <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">

        <script type="text/javascript">
            var global = global || window;
        </script>
    </head>
    <body>

    <app-home>
        <div id="preloader">
        	<div id="status">
            	<p class="center-text">
        			Loading the content...
                    <em>Loading depends on your connection speed!</em>
                </p>
            </div>
        </div>

    </app-home>

    <!-- <script src="/ui-app/dist/inline.bundle.js"></script>
    {{if eq .EnvironmentName "development"}}
    <script src='/ui-app/dist/vendor.bundle.js'></script>
    {{end}}
    <script src="/ui-app/dist/main.bundle.js?version={{.Version}}"> </script>-->
    </body>
    <script src="/ui/dist/ui/runtime.js?v={{.Version}}"></script>
    <script src='/ui/dist/ui/polyfills.js?v={{.Version}}'></script>
    <!-- <script src='/ui/dist/ui/scripts.js?v={{.Version}}'></script> -->
    {{if eq .EnvironmentName "development"}}
    <script src='/ui/dist/ui/vendor.js?v={{.Version}}'></script>
    <script src='/ui/dist/ui/styles.js?v={{.Version}}'></script>
    {{end}}

    <script src="/ui/dist/ui/main.js?version={{.Version}}"></script>
</html>
