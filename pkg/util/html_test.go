package util

import (
	"fmt"
	"testing"
)

var htmlData = `
<!DOCTYPE html>
<html>

  <head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">

  <title>数据库内核月报</title>
  <meta name="description" content="数据库内核月报，来自阿里云RDS-数据库内核组。
">

  <link rel="stylesheet" href="/monthly/css/typo.css">
  <link rel="stylesheet" href="/monthly/css/animate.css">
  <link rel="stylesheet" href="/monthly/css/main.css">
  <link rel="canonical" href="http://100.82.131.198:4000/monthly/">
  <link rel="alternate" type="application/rss+xml" title="数据库内核月报" href="http://100.82.131.198:4000/monthly/feed.xml" />

  <link rel="stylesheet" href="//cdn.staticfile.org/highlight.js/8.3/styles/tomorrow.min.css">
  <script src="/monthly/js/highlight.min.js"></script>
  <!-- <link rel="stylesheet" href="/monthly/themes/tomorrow.css">
  <script src="/monthly/highlight/highlight.pack.js"> -->
  <script>hljs.initHighlightingOnLoad();</script>

  <script src="http://cdn.staticfile.org/jquery/1.11.1/jquery.min.js"></script>
  <script src="http://cdn.staticfile.org/jquery/1.11.1/jquery.min.map"></script>

  <script src="/monthly/scripts/changeTarget.js"></script>
  
</head>


<!-- Google Analysis -->
<script>
  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');

  ga('create', 'UA-62056244-1', 'auto');
  ga('send', 'pageview');
</script>


  <body>

    <header>

  <a id="go-back-home" href="/monthly/">
    <h1>数据库内核月报</h1>
  </a>

</header>


        <div id = "container" class = "animated zoomIn">
  <div class="block">
    
  </div>

  <div class="content typo">

    <ul class="posts">
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2022/02">
              数据库内核月报 － 2022/02
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2022/01">
              数据库内核月报 － 2022/01
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2021/12">
              数据库内核月报 － 2021/12
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2021/11">
              数据库内核月报 － 2021/11
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2021/10">
              数据库内核月报 － 2021/10
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2021/09">
              数据库内核月报 － 2021/09
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2021/08">
              数据库内核月报 － 2021/08
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2021/07">
              数据库内核月报 － 2021/07
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2021/06">
              数据库内核月报 － 2021/06
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2021/05">
              数据库内核月报 － 2021/05
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2021/04">
              数据库内核月报 － 2021/04
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2021/03">
              数据库内核月报 － 2021/03
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2021/02">
              数据库内核月报 － 2021/02
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2021/01">
              数据库内核月报 － 2021/01
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2020/12">
              数据库内核月报 － 2020/12
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2020/11">
              数据库内核月报 － 2020/11
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2020/10">
              数据库内核月报 － 2020/10
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2020/09">
              数据库内核月报 － 2020/09
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2020/08">
              数据库内核月报 － 2020/08
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2020/07">
              数据库内核月报 － 2020/07
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2020/06">
              数据库内核月报 － 2020/06
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2020/05">
              数据库内核月报 － 2020/05
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2020/04">
              数据库内核月报 － 2020/04
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2020/03">
              数据库内核月报 － 2020/03
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2020/02">
              数据库内核月报 － 2020/02
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2020/01">
              数据库内核月报 － 2020/01
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2019/12">
              数据库内核月报 － 2019/12
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2019/11">
              数据库内核月报 － 2019/11
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2019/10">
              数据库内核月报 － 2019/10
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2019/09">
              数据库内核月报 － 2019/09
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2019/08">
              数据库内核月报 － 2019/08
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2019/07">
              数据库内核月报 － 2019/07
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2019/06">
              数据库内核月报 － 2019/06
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2019/05">
              数据库内核月报 － 2019/05
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2019/04">
              数据库内核月报 － 2019/04
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2019/03">
              数据库内核月报 － 2019/03
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2019/02">
              数据库内核月报 － 2019/02
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2019/01">
              数据库内核月报 － 2019/01
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2018/12">
              数据库内核月报 － 2018/12
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2018/11">
              数据库内核月报 － 2018/11
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2018/10">
              数据库内核月报 － 2018/10
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2018/09">
              数据库内核月报 － 2018/09
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2018/08">
              数据库内核月报 － 2018/08
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2018/07">
              数据库内核月报 － 2018/07
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2018/06">
              数据库内核月报 － 2018/06
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2018/05">
              数据库内核月报 － 2018/05
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2018/04">
              数据库内核月报 － 2018/04
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2018/03">
              数据库内核月报 － 2018/03
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2018/02">
              数据库内核月报 － 2018/02
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2018/01">
              数据库内核月报 － 2018/01
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2017/12">
              数据库内核月报 － 2017/12
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2017/11">
              数据库内核月报 － 2017/11
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2017/10">
              数据库内核月报 － 2017/10
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2017/09">
              数据库内核月报 － 2017/09
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2017/08">
              数据库内核月报 － 2017/08
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2017/07">
              数据库内核月报 － 2017/07
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2017/06">
              数据库内核月报 － 2017/06
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2017/05">
              数据库内核月报 － 2017/05
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2017/04">
              数据库内核月报 － 2017/04
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2017/03">
              数据库内核月报 － 2017/03
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2017/02">
              数据库内核月报 － 2017/02
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2017/01">
              数据库内核月报 － 2017/01
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2016/12">
              数据库内核月报 － 2016/12
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2016/11">
              数据库内核月报 － 2016/11
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2016/10">
              数据库内核月报 － 2016/10
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2016/09">
              数据库内核月报 － 2016/09
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2016/08">
              数据库内核月报 － 2016/08
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2016/07">
              数据库内核月报 － 2016/07
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2016/06">
              数据库内核月报 － 2016/06
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2016/05">
              数据库内核月报 － 2016/05
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2016/04">
              数据库内核月报 － 2016/04
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2016/03">
              数据库内核月报 － 2016/03
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2016/02">
              数据库内核月报 － 2016/02
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2016/01">
              数据库内核月报 － 2016/01
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2015/12">
              数据库内核月报 － 2015/12
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2015/11">
              数据库内核月报 － 2015/11
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2015/10">
              数据库内核月报 － 2015/10
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2015/09">
              数据库内核月报 － 2015/09
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2015/08">
              数据库内核月报 － 2015/08
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2015/07">
              数据库内核月报 － 2015/07
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2015/06">
              数据库内核月报 － 2015/06
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2015/05">
              数据库内核月报 － 2015/05
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2015/04">
              数据库内核月报 － 2015/04
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2015/03">
              数据库内核月报 － 2015/03
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2015/02">
              数据库内核月报 － 2015/02
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2015/01">
              数据库内核月报 － 2015/01
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2014/12">
              数据库内核月报 － 2014/12
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2014/11">
              数据库内核月报 － 2014/11
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2014/10">
              数据库内核月报 － 2014/10
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2014/09">
              数据库内核月报 － 2014/09
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
            <li><h3><a target="_top" class="main" href="/monthly/2014/08">
              数据库内核月报 － 2014/08
            </a></h3></li>
          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
          
          
          

          
      
    </ul>

  </div>

  <div class="block">
    
  </div>
</div>


    <footer>
  <script src="https://cdn1.lncld.net/static/js/av-mini-0.6.10.js"></script>
  <script src="http://jerry-cdn.b0.upaiyun.com/hit-kounter/hit-kounter-lc-0.2.0.js"></script>
  <a href="http://mysql.taobao.org/" target="_blank" class="muted">阿里云PolarDB-数据库内核组</a>
  <br>
  <a href="https://zhuanlan.zhihu.com/p/146555952" target="_blank" class="muted">社招/校招 欢迎投递简历 直达招聘信息</a>
  <br>
  <a href="https://www.zhihu.com/people/rong-rong-tong-xue-67" target="_blank" class="muted">直连月报小编</a>
  <br>
  <a href="https://github.com/alibaba/AliSQL" target="_blank" class="muted">欢迎在github上star AliSQL</a>
  <br>
  阅读：<span data-hk-page="current"> - </span><span class="pause">  </span>
</br>
<a rel="license" href="http://creativecommons.org/licenses/by-nc-sa/3.0/"><img alt="知识共享许可协议" style="border-width:0" src="https://i.creativecommons.org/l/by-nc-sa/3.0/88x31.png" /></a><br />本作品采用<a rel="license" href="http://creativecommons.org/licenses/by-nc-sa/3.0/">知识共享署名-非商业性使用-相同方式共享 3.0 未本地化版本许可协议</a>进行许可。
</footer>

<script type="text/javascript">
  jQuery(document).ready(function($){
    // browser window scroll (in pixels) after which the "back to top" link is shown
    var offset = 300,
      //browser window scroll (in pixels) after which the "back to top" link opacity is reduced
      offset_opacity = 1200,
      //duration of the top scrolling animation (in ms)
      scroll_top_duration = 700,
      //grab the "back to top" link
      $back_to_top = $('.cd-top');

    //hide or show the "back to top" link
    $(window).scroll(function(){
      ( $(this).scrollTop() > offset ) ? $back_to_top.addClass('cd-is-visible') : $back_to_top.removeClass('cd-is-visible cd-fade-out');
      if( $(this).scrollTop() > offset_opacity ) {
        $back_to_top.addClass('cd-fade-out');
      }
    });

    //smooth scroll to top
    $back_to_top.on('click', function(event){
      event.preventDefault();
      $('body,html').animate({
        scrollTop: 0 ,
        }, scroll_top_duration
      );
    });

  });
</script>



    <a href="#0" class="cd-top"><svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="10px"
   width="38px" height="60px" viewBox="0 0 16 16" enable-background="new 0 0 16 16" xml:space="preserve">
      <polygon fill="#FFFFFF" points="8,2.8 16,10.7 13.6,13.1 8.1,7.6 2.5,13.2 0,10.7 "/>
    </svg>
    </a>
  </body>

</html>

`

var htmlData2 = `
<!doctype html>
<html lang="en">
<head>
	<meta name="generator" content="Hugo 0.83.1" />
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width,initial-scale=1,maximum-scale=1,user-scalable=0">
  <title>CJ Ting&#39;s Blog - CJ Ting's Blog</title>
  <meta name="author" content="CJ Ting">
  <meta name="description" content="CJ Ting&#39;s Blog">
  <meta charset="utf-8" />
  <link rel="stylesheet" type="text/css" href="https://cdn.jsdelivr.net/npm/normalize.css@8.0.1/normalize.min.css">

  <link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png">
  <link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png">
  <link rel="stylesheet" type="text/css" href="//at.alicdn.com/t/font_1919785_agios4od1qk.css">

  <link rel="alternate" type="application/rss+xml" href="https://cjting.me/index.xml" title="CJ Ting's Blog" />
  <script src="https://cdn.jsdelivr.net/npm/gitalk@1/dist/gitalk.min.js"></script>
  <link rel="stylesheet" type="text/css" href="https://cdn.jsdelivr.net/npm/gitalk@1/dist/gitalk.css">

  <script src="https://cdn.jsdelivr.net/npm/clipboard@2.0.4/dist/clipboard.min.js"></script>

  <script src="https://cdn.jsdelivr.net/npm/notyf@3.1.0/notyf.min.js"></script>
  <link rel="stylesheet" type="text/css" href="https://cdn.jsdelivr.net/npm/notyf@3.1.0/notyf.min.css">

  <link rel="stylesheet" type="text/css" href="https://cdn.jsdelivr.net/npm/zoom-vanilla.js@2.0.6/css/zoom.css">

  <script src="https://polyfill.io/v3/polyfill.min.js?features=es6"></script>
  <script>
  MathJax = {
    tex: {
      inlineMath: [["$", "$"]],
    },
    chtml: {
      scale: 1.2,
    },
  }
  </script>
  <script id="MathJax-script" async src="https://cdn.jsdelivr.net/npm/mathjax@3.0.0/es5/tex-chtml.js"></script>
  <script type="text/javascript">
    window.PAGE = {
      title: "CJ Ting\u0027s Blog",
      tags:  null ,
    }
  </script>

  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/tocbot@4.11.2/dist/tocbot.css">
  <script src="https://cdn.jsdelivr.net/npm/tocbot@4.11.2/dist/tocbot.min.js"></script>

  <link rel="stylesheet" href="/asset/main.css">
</head>
<body>
  
    <header class="header">
  <a href="/">
    <img class="header__avatar" src="/image/avatar.jpeg">
  </a>

  <nav class="header__nav">
    <a href="/" class="header__nav__item  header__nav__item--active ">Posts</a>

    <a href="/archive" class="header__nav__item ">Archive</a>

    <a href="https://cjting.me/index.xml" class="header__nav__item" target="_blank">RSS</a>

    <a href="https://slides.cjting.me" target="_blank" class="header__nav__item">Slides</a>

    <a href="https://github.com/cj1128" class="header__nav__item" target="_blank">GitHub</a>
  </nav>
</header>

    <main class="main">
      
        <div class="container">
          
  <div class="index u-content-container">
    
    
      <div class="index__item">
        <a class="index__item__title hvr-underline-from-left" href="/2021/08/07/fourier-transform-and-audio-visualization/">
          音频可视化：采样、频率和傅里叶变换
        </a>

        <div class="index__item__content post__content">
          <p>印象中使用的第一个 PC 音乐播放器是「千千静听」，大概是 08 年左右。我还清楚地记得它自带了一首梁静茹的歌「Love is everything」，动听的旋律至今萦绕耳旁。</p>
<p><img src="qianqian.jpeg" alt=""></p>
<p>千千静听的左上角有一组随着音乐跳动的柱子，我想大家都习以为常了，很多播放器都有这功能。但是其实有没有想过，这是怎么实现的？</p>
<p>要理解这个问题，我们首先要理解声音是什么。</p>
        </div>

        <div class="index__item__meta">
          <div>
            2021.08.07
          </div>
        </div>
      </div>
    
      <div class="index__item">
        <a class="index__item__title hvr-underline-from-left" href="/2021/06/10/hot-reload-c/">
          热重载 C
        </a>

        <div class="index__item__content post__content">
          <p>热重载是一个非常好用的功能，可以在不重启的情况下更新应用，从而大大提高开发效率。</p>
<p>前端的 Wepback，后端的 Ruby/Python/Elixir，移动端的 Flutter 都有热重载，属于用过以后就回不去的 Killer Feature。</p>
<p>在我之前的认识中，一直认为只有脚本语言才可以支持热重载，因为虚拟机让热重载的实现变得非常简单，重新运行代码即可。</p>
<p>直到有一天，Casey 在 <a href="https://handmadehero.org/">HandmadeHero</a> 项目中用非常少的代码演示了怎样热重载 C，我才恍然大悟，编译语言一样可以热重载。</p>
        </div>

        <div class="index__item__meta">
          <div>
            2021.06.10
          </div>
        </div>
      </div>
    
      <div class="index__item">
        <a class="index__item__title hvr-underline-from-left" href="/2021/03/16/the-missing-div-instruction-part1/">
          消失的除法指令：Part1
        </a>

        <div class="index__item__content post__content">
          <p>之前学汇编的时候观察到一个现象，我在 C 语言中写了一个函数进行除法操作，但是编译得到的汇编代码中却没有除法指令，取而代之的是一条乘法指令。</p>
<p><img src="/image/FjCkN1q9ePhijDTEBzws-jI0B7w1.png" alt=""></p>
<p>图片对应的 GodBolt 地址在 <a href="https://gcc.godbolt.org/z/YrK4vnY1E">这里</a>，可以看到有一个 <code>imulq</code> 指令，这是一个乘法指令，乘了一个奇怪的数字 1431655766。</p>
<p>为什么编译器要这样操作？为什么能这样操作？1431655766 这个数字又是怎么来的？</p>
        </div>

        <div class="index__item__meta">
          <div>
            2021.03.16
          </div>
        </div>
      </div>
    
      <div class="index__item">
        <a class="index__item__title hvr-underline-from-left" href="/2021/03/02/how-to-validate-tls-certificate/">
          安全背后: 浏览器是如何校验证书的
        </a>

        <div class="index__item__content post__content">
          <p>现如今的 Web，HTTPS 早已经成为标配，公开的 HTTP 网站已经和 Flash 一样，慢慢在消亡了。</p>
<p>启用 HTTPS 的核心是一个叫做 <strong>证书</strong> 的东西。不知道大家是否有留意，前几年上 12306 的时候，浏览器都会提示「您的链接不是私密链接」，这其实就是因为 12306 的证书有问题。如果点击「继续前往」，打开 12306 网站，它会提示你下载安装它提供的“根证书”。</p>
<p><img src="/image/FrBOHRLcNkDuAUnDstdn4Sjk4xMr.png" alt=""></p>
<p>那么，证书是什么？里面含有什么内容？浏览器为什么会不信任 12306 的证书？为什么下载 12306 提供的根证书就可以解决这个问题？根证书又是什么？</p>
        </div>

        <div class="index__item__meta">
          <div>
            2021.03.02
          </div>
        </div>
      </div>
    
      <div class="index__item">
        <a class="index__item__title hvr-underline-from-left" href="/2020/12/10/tiny-x64-helloworld/">
          编写一个最小的 64 位 Hello World
        </a>

        <div class="index__item__content post__content">
          <p>Hello World 应该是每一位程序员的启蒙程序，出自于 <a href="https://en.wikipedia.org/wiki/Brian_Kernighan">Brian Kernighan</a> 和 <a href="https://en.wikipedia.org/wiki/Dennis_Ritchie">Dennis Ritchie</a> 的一代经典著作 <a href="https://en.wikipedia.org/wiki/The_C_Programming_Language">The C Programming Language</a>。</p>
<div class="highlight"><pre style="background-color:#fff;-moz-tab-size:2;-o-tab-size:2;tab-size:2"><code class="language-c" data-lang="c"><span style="color:#998;font-style:italic">// hello.c
</span><span style="color:#998;font-style:italic"></span><span style="color:#999;font-weight:bold;font-style:italic">#include</span> <span style="color:#999;font-weight:bold;font-style:italic">&lt;stdio.h&gt;</span><span style="color:#999;font-weight:bold;font-style:italic">
</span><span style="color:#999;font-weight:bold;font-style:italic"></span>
<span style="color:#458;font-weight:bold">int</span> <span style="color:#900;font-weight:bold">main</span>() {
  printf(<span style="color:#d14">&#34;hello, world</span><span style="color:#d14">\n</span><span style="color:#d14">&#34;</span>);
  <span style="color:#000;font-weight:bold">return</span> <span style="color:#099">0</span>;
}
</code></pre></div><p>这段代码我想大家应该都太熟悉了，熟悉到可以默写出来。虽然是非常简单的代码，但是如果细究起来，里面却隐含着很多细节：</p>
<ul>
<li><code>#include &lt;stdio.h&gt;</code> 和 <code>#include &quot;stdio.h&quot;</code> 有什么区别？</li>
<li><code>stdio.h</code> 文件在哪里？里面是什么内容？</li>
<li>为什么入口是 <code>main</code> 函数？可以写一个程序入口不是 <code>main</code> 吗？</li>
<li><code>main</code> 的 int 返回值有什么用？是谁在处理 <code>main</code> 的返回值？</li>
<li><code>printf</code> 是谁实现的？如果不用 <code>printf</code> 可以做到在终端中打印字符吗？</li>
</ul>
        </div>

        <div class="index__item__meta">
          <div>
            2020.12.10
          </div>
        </div>
      </div>
    
      <div class="index__item">
        <a class="index__item__title hvr-underline-from-left" href="/2020/10/31/tinytorrent-a-deno-bt-downloader/">
          tinyTorrent: 从头写一个 Deno 的 BitTorrent 下载器
        </a>

        <div class="index__item__content post__content">
          <p>BitTorrent 想必大家应该都不陌生，中文名叫做“种子”。</p>
<p>“种子”到底是什么？我一直不太清楚。在写这个项目之前，我对“种子”的认识停留在使用层面。</p>
<p>当我想找一个资源的时候，我会搜索 <code>xxx 种子</code>，一般会在一些非常不知名的小网站里面获取到以 <code>.torrent</code> 结尾的种子文件，然后使用迅雷或者 uTorrent 这样的下载器来下载。</p>
<p>如果迅雷有一点速度，哪怕几 kb，那么大概率我充个会员就搞定了。这个软件就是这么的恶心，不用有时候又没办法，像极了人生。</p>
<p>其他下载器比如 uTorrent 的话就一切随缘了，有些资源非常快，有些资源非常慢，有些一开始慢后来快。</p>
<p>这些问题是怎么回事？有没有改进的办法？在读 Jesse Li 的 <a href="https://blog.jse.li/posts/torrent/">Building a BitTorrent client from the ground up in Go</a> 之前，我从没想过。</p>
        </div>

        <div class="index__item__meta">
          <div>
            2020.10.31
          </div>
        </div>
      </div>
    
      <div class="index__item">
        <a class="index__item__title hvr-underline-from-left" href="/2020/09/11/cs107e-review/">
          CS107e: 树莓派，ARM 和操作系统
        </a>

        <div class="index__item__content post__content">
          <p>2020.10.09 更新：<strong>最近收到了一封官方的邮件，希望我把仓库隐藏起来，避免新同学直接 copy 我的代码。既然是来自官方的要求，自然毫不犹豫地配合。仓库和我归档的 Winter 2020 课程现在都不可见了，想要学习这门课程的同学可以直接去官方站点学习，祝大家学习愉快~</strong></p>
<p><a href="http://cs107e.github.io/">CS107e</a> 全称为 <em>CS107e: Computer Systems from the Ground Up</em>，是斯坦福大学的一门计算机课程，也是我目前为止发现的最好的关于硬件、底层和 C 的一门课程。</p>
<p>在课程学习过程中，我们会一步一步地从头开始使用 C 为树莓派开发一个操作系统核心。</p>
<p>根据官方的介绍，这门课程主要学习目标有两个：</p>
<ol>
<li>To understand how computers represent information, execute programs, and control peripherals. 理解计算机是怎样存储信息，执行程序以及控制外围设备。</li>
<li>To master command-line programming tools and the C programming language. 掌握命令行编程工具和 C 语言。</li>
</ol>
        </div>

        <div class="index__item__meta">
          <div>
            2020.09.11
          </div>
        </div>
      </div>
    
      <div class="index__item">
        <a class="index__item__title hvr-underline-from-left" href="/2020/08/16/shell-init-type/">
          Shell 启动类型探究 ── login &amp;&amp; interactive
        </a>

        <div class="index__item__content post__content">
          <p>Shell 对程序员来说是必不可少的生产力工具。</p>
<div class="highlight"><pre style="background-color:#fff;-moz-tab-size:2;-o-tab-size:2;tab-size:2"><code class="language-text" data-lang="text">$ figlet &lt;&lt;&lt; &#34;Hello Shell&#34;
 _   _      _ _         ____  _          _ _
| | | | ___| | | ___   / ___|| |__   ___| | |
| |_| |/ _ \ | |/ _ \  \___ \| &#39;_ \ / _ \ | |
|  _  |  __/ | | (_) |  ___) | | | |  __/ | |
|_| |_|\___|_|_|\___/  |____/|_| |_|\___|_|_|

</code></pre></div>
        </div>

        <div class="index__item__meta">
          <div>
            2020.08.16
          </div>
        </div>
      </div>
    
      <div class="index__item">
        <a class="index__item__title hvr-underline-from-left" href="/2020/07/01/douyu-crawler-and-font-anti-crawling/">
          斗鱼关注人数爬取 ── 字体反爬的攻与防
        </a>

        <div class="index__item__content post__content">
          <p>之前因为业务原因需要爬取一批斗鱼主播的相关数据，在这过程中我发现斗鱼使用了一种很有意思的反爬技术，字体反爬。</p>
<p>打开任何一个斗鱼主播的直播间，例如 <a href="https://www.douyu.com/7874579">这个主播</a>，他的关注人数数据显示在右上角：</p>
<p><img src="/image/Fp11DQiiqvE_vQddpXH1tgazywrA.png" alt=""></p>
<p>斗鱼在关注数据这里使用了字体反爬。什么是字体反爬？也就是通过自定义字体来自定义字符与渲染图形的映射。比如，字符 1 实际渲染的是 9，那么如果 HTML 中的数字是 111，实际显示就是 999。</p>
<p>在这种技术下，传统的通过解析 HTML 文档获取数据的方式就失效了，因为获取到的数据并不是真实数据。</p>
        </div>

        <div class="index__item__meta">
          <div>
            2020.07.01
          </div>
        </div>
      </div>
    
      <div class="index__item">
        <a class="index__item__title hvr-underline-from-left" href="/2020/06/07/chip8-emulator/">
          用 C 实现一个 CHIP-8 模拟器
        </a>

        <div class="index__item__content post__content">
          <p>很早之前我就想写一个 GBA 模拟器，因为小时候的 GBA 游戏给我留下了深刻的印象。</p>
<p>口袋妖怪、孤岛求生、牧场物语这些 GBA 的经典游戏，在那个时候还玩着小霸王的我眼中，无异于打开了新世界的大门，原来游戏可以这样的有趣。</p>
<p>因为对 GBA 的喜欢，所以有了编写一个 GBA 模拟器的想法。看过一些资料以后，我决定从最简单的 CHIP-8 开始练手。CHIP-8 是一个功能完整的平台，可以运行多个游戏，同时它的设计又非常简单，很适合新手入门。</p>
        </div>

        <div class="index__item__meta">
          <div>
            2020.06.07
          </div>
        </div>
      </div>
    

    <div class="index__paginator">
      
        <a href="javascript:void(0)" disabled="disabled">Prev</a>
      

      <div class="index__paginator__info">
        <span>1</span>
        /
        <span>4</span>
      </div>

      
        <a href="/page/2/">Next</a>
      
    </div>
  </div>

        </div>
      
    </main>
  

  <script src="https://cdn.jsdelivr.net/npm/zoom-vanilla.js@2.0.6/dist/zoom-vanilla.min.js"></script>
  <script src="/asset/main.js"></script>

  
  <script>
    (function() {
      var _hmt = _hmt || [];
      var hm = document.createElement("script");
      hm.src = "https://hm.baidu.com/hm.js?4f34ee3c85734c8235badd2b99b092a6";
      var s = document.getElementsByTagName("script")[0];
      s.parentNode.insertBefore(hm, s);
    })();
  </script>

  
  <div style="display: none">
    <script type="text/javascript" src="https://s23.cnzz.com/z_stat.php?id=1277776204&web_id=1277776204"></script>
  </div>
</body>
</html>

`

var htmlData3 = `
<a class="post-link" href="/2020/09/13/what-iam-doing-now.html" title="不写文章专注做视频，这里要荒废一段时间了1">第一个</a>
<a class="post-link" href="/2020/09/13/what-iam-doing-now.html" title="不写文章专注做视频，这里要荒废一段时间了2">第二个</a>
<a href="/2022/03/slg-server-mindmap/" class="post-title-link" itemprop="url">SLG游戏服务器随想</a>
`

// 整个网页，需要筛选的链接内容，是排列在一起，中间没有其他不需要的链接的
func TestHtmlMatchLink(t *testing.T) {
	res := HtmlMatchLink(htmlData)

	source := "<a target=\"_top\" class=\"main\" href=\"/monthly/2022/02\">"
	sourceIndex := -1
	for i, v := range res {
		if v.HtmlPre == source {
			sourceIndex = i
		}
	}

	hm10 := res[sourceIndex : sourceIndex+10]
	fmt.Println(hm10)

}

// 整个网页，需要筛选的链接内容，是排列在一起，中间没有其他不需要的链接的
func TestHtmlMatchLink11(t *testing.T) {
	res := HtmlMatchLink(htmlData2)

	source := "<a class=\"index__item__title hvr-underline-from-left\" href=\"/2021/08/07/fourier-transform-and-audio-visualization/\">"
	sourceIndex := -1
	for i, v := range res {
		if v.HtmlPre == source {
			sourceIndex = i
		}
	}

	hm10 := res[sourceIndex : sourceIndex+10]
	fmt.Println(hm10)

}

// 整个网页，需要筛选的链接内容，有其他不需要的链接的
//func TestHtmlMatchLink2(t *testing.T) {
//	res := HtmlMatchLink(htmlData2)
//
//	htmlIndexM := make(map[string]int)
//	htmlM := make(map[string]string)
//	linkM := make(map[string]string)
//	//source := "<a class=\"index__item__title hvr-underline-from-left\" href=\"/2021/08/07/fourier-transform-and-audio-visualization/\">"
//	sourcePre := "<a class=\"index__item__title hvr-underline-from-left\" href=\"/20"
//	//sourceIndex := -1
//	newRes := res[:0]
//	for i, v := range res {
//		if len(v) < 4 {
//			continue
//		}
//
//		htmlIndexM[v[0]] = i
//		htmlM[v[0]] = v[1]
//		linkM[v[2]] = v[3]
//
//		//if v[1] == source {
//		//	sourceIndex = i
//		//}
//
//		s := v[1]
//		if len(s) > len(sourcePre) {
//			s = s[:len(sourcePre)]
//			if CosineSimilarity(s, sourcePre, 0) == 1 {
//				newRes = append(newRes, v)
//			}
//		}
//
//	}
//
//	fmt.Println(newRes)
//
//	//
//	//cosineM := make(map[string]float64)
//	//for k, v := range htmlM {
//	//	cosineM[k] = CosineSimilarity(v, source, 0)
//	//	if cosineM[k] == 1 {
//	//		a := k
//	//		fmt.Println(a, v)
//	//		continue
//	//	}
//	//
//	//}
//	//
//	//consineS := make([]float64, 0)
//	//for _, v := range cosineM {
//	//	consineS = append(consineS, v)
//	//}
//	//
//	//ms := base.NewMapSorter(cosineM)
//	//sort.Sort(ms)
//	//
//	//ms10 := ms[len(ms)-10:]
//	//
//	//for _,v := range ms10 {
//	//
//	//	fmt.Println(htmlIndexM[v.Key])
//	//	fmt.Println(v)
//	//}
//
//}

func TestHtmlMatchLink3(t *testing.T) {
	res := HtmlMatchTitle(htmlData3)

	//sourcePre := "<a class=\"index__item__title hvr-underline-from-left\" href=\"/20"
	sourcePre := "<a class=\"post-link\" href=\"/20"
	newRes := res[:0]
	for _, v := range res {

		s := v.HtmlPre
		if len(s) > len(sourcePre) {
			s = s[:len(sourcePre)]
			if CosineSimilarity(s, sourcePre, 0) == 1 {
				newRes = append(newRes, v)
			}
		}

	}

	fmt.Println(newRes)
}
