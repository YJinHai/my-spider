package regexp

import (
	"testing"
)

var body = `
<!DOCTYPE html>
<html lang="zh-Hans">
<head>
  <meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=2">
<meta name="theme-color" content="#222">
<meta name="generator" content="Hexo 4.0.0">
  <link rel="apple-touch-icon" sizes="180x180" href="/images/apple-touch-icon-next.png">
  <link rel="icon" type="image/png" sizes="32x32" href="/images/favicon-32x32-next.png">
  <link rel="icon" type="image/png" sizes="16x16" href="/images/favicon-16x16-next.png">
  <link rel="mask-icon" href="/images/logo.svg" color="#222">
  <link rel="alternate" href="/atom.xml" title="Legendtkl" type="application/atom+xml">

<link rel="stylesheet" href="/css/main.css">


<link rel="stylesheet" href="/lib/font-awesome/css/font-awesome.min.css">


<script id="hexo-configurations">
  var NexT = window.NexT || {};
  var CONFIG = {
    root: '/',
    scheme: 'Mist',
    version: '7.5.0',
    exturl: false,
    sidebar: {"position":"right","display":"always","offset":12,"onmobile":false},
    copycode: {"enable":false,"show_result":false,"style":null},
    back2top: {"enable":true,"sidebar":false,"scrollpercent":false},
    bookmark: {"enable":false,"color":"#222","save":"auto"},
    fancybox: false,
    mediumzoom: false,
    lazyload: false,
    pangu: false,
    algolia: {
      appID: '',
      apiKey: '',
      indexName: '',
      hits: {"per_page":10},
      labels: {"input_placeholder":"Search for Posts","hits_empty":"We didn't find any results for the search: ${query}","hits_stats":"${hits} results found in ${time} ms"}
    },
    localsearch: {"enable":false,"trigger":"auto","top_n_per_article":1,"unescape":false,"preload":false},
    path: '',
    motion: {"enable":true,"async":false,"transition":{"post_block":"fadeIn","post_header":"slideDownIn","post_body":"slideDownIn","coll_header":"slideLeftIn","sidebar":"slideUpIn"}},
    translation: {
      copy_button: 'Copy',
      copy_success: 'Copied',
      copy_failure: 'Copy failed'
    },
    sidebarPadding: 40
  };
</script>

  <meta name="description" content="Do not go gentle into that good night.">
<meta property="og:type" content="website">
<meta property="og:title" content="Legendtkl">
<meta property="og:url" content="http:&#x2F;&#x2F;yoursite.com&#x2F;index.html">
<meta property="og:site_name" content="Legendtkl">
<meta property="og:description" content="Do not go gentle into that good night.">
<meta property="og:locale" content="zh-Hans">
<meta name="twitter:card" content="summary">

<link rel="canonical" href="http://yoursite.com/">


<script id="page-configurations">
  // https://hexo.io/docs/variables.html
  CONFIG.page = {
    sidebar: "",
    isHome: true,
    isPost: false,
    isPage: false,
    isArchive: false
  };
</script>

  <title>Legendtkl</title>
  






  <noscript>
  <style>
  .use-motion .brand,
  .use-motion .menu-item,
  .sidebar-inner,
  .use-motion .post-block,
  .use-motion .pagination,
  .use-motion .comments,
  .use-motion .post-header,
  .use-motion .post-body,
  .use-motion .collection-header { opacity: initial; }

  .use-motion .site-title,
  .use-motion .site-subtitle {
    opacity: initial;
    top: initial;
  }

  .use-motion .logo-line-before i { left: initial; }
  .use-motion .logo-line-after i { right: initial; }
  </style>
</noscript>

</head>

<body itemscope itemtype="http://schema.org/WebPage">
  <div class="container use-motion">
    <div class="headband"></div>

    <header class="header" itemscope itemtype="http://schema.org/WPHeader">
      <div class="header-inner"><div class="site-brand-container">
  <div class="site-meta">

    <div>
      <a href="/" class="brand" rel="start">
        <span class="logo-line-before"><i></i></span>
        <span class="site-title">Legendtkl</span>
        <span class="logo-line-after"><i></i></span>
      </a>
    </div>
        <p class="site-subtitle">abc</p>
  </div>

  <div class="site-nav-toggle">
    <div class="toggle" aria-label="Toggle navigation bar">
      <span class="toggle-line toggle-line-first"></span>
      <span class="toggle-line toggle-line-middle"></span>
      <span class="toggle-line toggle-line-last"></span>
    </div>
  </div>
</div>


<nav class="site-nav">
  
  <ul id="menu" class="menu">
        <li class="menu-item menu-item-archives">

    <a href="/archives/" rel="section"><i class="fa fa-fw fa-archive"></i>Archives</a>

  </li>
        <li class="menu-item menu-item-categories">

    <a href="/categories/" rel="section"><i class="fa fa-fw fa-th"></i>Categories</a>

  </li>
        <li class="menu-item menu-item-booklist">

    <a href="/booklist/" rel="section"><i class="fa fa-fw fa-book"></i>booklist</a>

  </li>
        <li class="menu-item menu-item-about">

    <a href="/about/" rel="section"><i class="fa fa-fw fa-user"></i>About</a>

  </li>
  </ul>

</nav>
</div>
    </header>

    
  <div class="back-to-top">
    <i class="fa fa-arrow-up"></i>
    <span>0%</span>
  </div>


    <main class="main">
      <div class="main-inner">
        <div class="content-wrap">
          

          <div class="content">
            

  <div class="posts-expand">
        
  
  
  <article itemscope itemtype="http://schema.org/Article" class="post-block home" lang="zh-Hans">
    <link itemprop="mainEntityOfPage" href="http://yoursite.com/2021/03/06/no-rules/">

    <span hidden itemprop="author" itemscope itemtype="http://schema.org/Person">
      <meta itemprop="image" content="/img/avatar.jpg">
      <meta itemprop="name" content="legendtkl">
      <meta itemprop="description" content="Do not go gentle into that good night.">
    </span>

    <span hidden itemprop="publisher" itemscope itemtype="http://schema.org/Organization">
      <meta itemprop="name" content="Legendtkl">
    </span>
      <header class="post-header">
        <h1 class="post-title" itemprop="name headline">
          
            <a href="/2021/03/06/no-rules/" class="post-title-link" itemprop="url">读《不拘一格：网飞的自由与责任工作法》，从网飞我们可以学到什么</a>
        </h1>

        <div class="post-meta">
            <span class="post-meta-item">
              <span class="post-meta-item-icon">
                <i class="fa fa-calendar-o"></i>
              </span>
              <span class="post-meta-item-text">Posted on</span>
              

              <time title="Created: 2021-03-06 14:02:42 / Modified: 18:11:20" itemprop="dateCreated datePublished" datetime="2021-03-06T14:02:42+08:00">2021-03-06</time>
            </span>
            <span class="post-meta-item">
              <span class="post-meta-item-icon">
                <i class="fa fa-folder-o"></i>
              </span>
              <span class="post-meta-item-text">In</span>
                <span itemprop="about" itemscope itemtype="http://schema.org/Thing">
                  <a href="/categories/reading/" itemprop="url" rel="index">
                    <span itemprop="name">reading</span>
                  </a>
                </span>
            </span>

          

        </div>
      </header>

    
    
    
    <div class="post-body" itemprop="articleBody">

      
          <p>这是一篇关于书籍的《不拘一格：网飞的自由与责任工作法》的笔记，只记录一些感兴趣的点。书籍的英文原名为 <strong><em>No Rules Rules: Netflix and the Culture of Reinvention</em></strong>，直白一点说就是网飞的企业文化，而中文译名直接道出了网飞的企业文化的核心：自由与责任。</p>
          <!--noindex-->
            <div class="post-button">
              <a class="btn" href="/2021/03/06/no-rules/#more" rel="contents">
                Read more &raquo;
              </a>
            </div>
          <!--/noindex-->
        
      
    </div>

    
    
    
      <footer class="post-footer">
        <div class="post-eof"></div>
      </footer>
  </article>
  
  
  

        
  
  
  <article itemscope itemtype="http://schema.org/Article" class="post-block home" lang="zh-Hans">
    <link itemprop="mainEntityOfPage" href="http://yoursite.com/2020/11/27/flink-ha-kubernetes/">

    <span hidden itemprop="author" itemscope itemtype="http://schema.org/Person">
      <meta itemprop="image" content="/img/avatar.jpg">
      <meta itemprop="name" content="legendtkl">
      <meta itemprop="description" content="Do not go gentle into that good night.">
    </span>

    <span hidden itemprop="publisher" itemscope itemtype="http://schema.org/Organization">
      <meta itemprop="name" content="Legendtkl">
    </span>
      <header class="post-header">
        <h1 class="post-title" itemprop="name headline">
          
            <a href="/2020/11/27/flink-ha-kubernetes/" class="post-title-link" itemprop="url">Flink JM HA 在 Kubernetes 上的实现</a>
        </h1>

        <div class="post-meta">
            <span class="post-meta-item">
              <span class="post-meta-item-icon">
                <i class="fa fa-calendar-o"></i>
              </span>
              <span class="post-meta-item-text">Posted on</span>

              <time title="Created: 2020-11-27 23:18:57" itemprop="dateCreated datePublished" datetime="2020-11-27T23:18:57+08:00">2020-11-27</time>
            </span>
              <span class="post-meta-item">
                <span class="post-meta-item-icon">
                  <i class="fa fa-calendar-check-o"></i>
                </span>
                <span class="post-meta-item-text">Edited on</span>
                <time title="Modified: 2020-12-06 14:21:25" itemprop="dateModified" datetime="2020-12-06T14:21:25+08:00">2020-12-06</time>
              </span>
            <span class="post-meta-item">
              <span class="post-meta-item-icon">
                <i class="fa fa-folder-o"></i>
              </span>
              <span class="post-meta-item-text">In</span>
                <span itemprop="about" itemscope itemtype="http://schema.org/Thing">
                  <a href="/categories/flink/" itemprop="url" rel="index">
                    <span itemprop="name">flink</span>
                  </a>
                </span>
                  , 
                <span itemprop="about" itemscope itemtype="http://schema.org/Thing">
                  <a href="/categories/flink/kubernetes/" itemprop="url" rel="index">
                    <span itemprop="name">kubernetes</span>
                  </a>
                </span>
            </span>

          

        </div>
      </header>

    
    
    
    <div class="post-body" itemprop="articleBody">

      
          <p>这篇文章介绍一下 Flink 的 JobManager HA 在 Kubernetes 上面的实现思路。Flink 1.12 还没有 release，但是在开发计划中已经看到了这块内容。但是这篇文章主要介绍我们内部的实现。下一篇在 Flink 1.12 正式 release 之后再进行介绍官方的实现。</p>
          <!--noindex-->
            <div class="post-button">
              <a class="btn" href="/2020/11/27/flink-ha-kubernetes/#more" rel="contents">
                Read more &raquo;
              </a>
            </div>
          <!--/noindex-->
        
      
    </div>

    
    
    
      <footer class="post-footer">
        <div class="post-eof"></div>
      </footer>
  </article>
  
  
  

        
  
  
  <article itemscope itemtype="http://schema.org/Article" class="post-block home" lang="zh-Hans">
    <link itemprop="mainEntityOfPage" href="http://yoursite.com/2020/09/15/secure-container-overview/">

    <span hidden itemprop="author" itemscope itemtype="http://schema.org/Person">
      <meta itemprop="image" content="/img/avatar.jpg">
      <meta itemprop="name" content="legendtkl">
      <meta itemprop="description" content="Do not go gentle into that good night.">
    </span>

    <span hidden itemprop="publisher" itemscope itemtype="http://schema.org/Organization">
      <meta itemprop="name" content="Legendtkl">
    </span>
      <header class="post-header">
        <h1 class="post-title" itemprop="name headline">
          
            <a href="/2020/09/15/secure-container-overview/" class="post-title-link" itemprop="url">安全容器综述</a>
        </h1>

        <div class="post-meta">
            <span class="post-meta-item">
              <span class="post-meta-item-icon">
                <i class="fa fa-calendar-o"></i>
              </span>
              <span class="post-meta-item-text">Posted on</span>
              

              <time title="Created: 2020-09-15 22:53:17 / Modified: 15:26:57" itemprop="dateCreated datePublished" datetime="2020-09-15T22:53:17+08:00">2020-09-15</time>
            </span>
            <span class="post-meta-item">
              <span class="post-meta-item-icon">
                <i class="fa fa-folder-o"></i>
              </span>
              <span class="post-meta-item-text">In</span>
                <span itemprop="about" itemscope itemtype="http://schema.org/Thing">
                  <a href="/categories/docker/" itemprop="url" rel="index">
                    <span itemprop="name">docker</span>
                  </a>
                </span>
                  , 
                <span itemprop="about" itemscope itemtype="http://schema.org/Thing">
                  <a href="/categories/docker/%E5%AE%B9%E5%99%A8/" itemprop="url" rel="index">
                    <span itemprop="name">容器</span>
                  </a>
                </span>
            </span>

          

        </div>
      </header>

    
    
    
    <div class="post-body" itemprop="articleBody">

      
          <p>最近两年安全容器非常的火，这篇文章就带大家来看一下何谓安全容器技术，以及目前主流的安全容器都有哪些，最后还会附上很多有价值的参考链接。本文将通过如下的方式进行展开：</p>
<ol>
<li>何谓安全容器</li>
<li>安全容器技术的思路是什么样的</li>
<li>目前的主流安全容器有哪些</li>
</ol>
          <!--noindex-->
            <div class="post-button">
              <a class="btn" href="/2020/09/15/secure-container-overview/#more" rel="contents">
                Read more &raquo;
              </a>
            </div>
          <!--/noindex-->
        
      
    </div>

    
    
    
      <footer class="post-footer">
        <div class="post-eof"></div>
      </footer>
  </article>
  
  
  

        
  
  
  <article itemscope itemtype="http://schema.org/Article" class="post-block home" lang="zh-Hans">
    <link itemprop="mainEntityOfPage" href="http://yoursite.com/2020/09/14/imooc-md/">

    <span hidden itemprop="author" itemscope itemtype="http://schema.org/Person">
      <meta itemprop="image" content="/img/avatar.jpg">
      <meta itemprop="name" content="legendtkl">
      <meta itemprop="description" content="Do not go gentle into that good night.">
    </span>

    <span hidden itemprop="publisher" itemscope itemtype="http://schema.org/Organization">
      <meta itemprop="name" content="Legendtkl">
    </span>
      <header class="post-header">
        <h1 class="post-title" itemprop="name headline">
          
            <a href="/2020/09/14/imooc-md/" class="post-title-link" itemprop="url">推荐一门自己做的 Docker 相关的付费课程</a>
        </h1>

        <div class="post-meta">
            <span class="post-meta-item">
              <span class="post-meta-item-icon">
                <i class="fa fa-calendar-o"></i>
              </span>
              <span class="post-meta-item-text">Posted on</span>
              

              <time title="Created: 2020-09-14 22:36:19 / Modified: 22:46:49" itemprop="dateCreated datePublished" datetime="2020-09-14T22:36:19+08:00">2020-09-14</time>
            </span>
            <span class="post-meta-item">
              <span class="post-meta-item-icon">
                <i class="fa fa-folder-o"></i>
              </span>
              <span class="post-meta-item-text">In</span>
                <span itemprop="about" itemscope itemtype="http://schema.org/Thing">
                  <a href="/categories/docker/" itemprop="url" rel="index">
                    <span itemprop="name">docker</span>
                  </a>
                </span>
                  , 
                <span itemprop="about" itemscope itemtype="http://schema.org/Thing">
                  <a href="/categories/docker/kubernetes/" itemprop="url" rel="index">
                    <span itemprop="name">kubernetes</span>
                  </a>
                </span>
            </span>

          

        </div>
      </header>

    
    
    
    <div class="post-body" itemprop="articleBody">

      
          <p>过去半年基本很少在博客或者知乎发表文章，主要原因是一直在做一个知识付费的课程，慕课网的专栏，这个是课程链接：<a href="https://www.imooc.com/read/84" target="_blank" rel="noopener">跟 BAT 技术专家学 Docker + K8S</a> 。下面摘抄一下网站的课程介绍如下，感兴趣的可以报名学习一下。这么多年一直在写免费的文章，第一次尝试付费内容。说来惭愧，目前的收入相比写文章付出的时间实在是惨不忍睹，所以博客上面做一个广告，希望多多包涵。下面是课程介绍。</p>
          <!--noindex-->
            <div class="post-button">
              <a class="btn" href="/2020/09/14/imooc-md/#more" rel="contents">
                Read more &raquo;
              </a>
            </div>
          <!--/noindex-->
        
      
    </div>

    
    
    
      <footer class="post-footer">
        <div class="post-eof"></div>
      </footer>
  </article>
  
  
  

        
  
  
  <article itemscope itemtype="http://schema.org/Article" class="post-block home" lang="zh-Hans">
    <link itemprop="mainEntityOfPage" href="http://yoursite.com/2020/08/01/spark-catalog-plugin/">

    <span hidden itemprop="author" itemscope itemtype="http://schema.org/Person">
      <meta itemprop="image" content="/img/avatar.jpg">
      <meta itemprop="name" content="legendtkl">
      <meta itemprop="description" content="Do not go gentle into that good night.">
    </span>

    <span hidden itemprop="publisher" itemscope itemtype="http://schema.org/Organization">
      <meta itemprop="name" content="Legendtkl">
    </span>
      <header class="post-header">
        <h1 class="post-title" itemprop="name headline">
          
            <a href="/2020/08/01/spark-catalog-plugin/" class="post-title-link" itemprop="url">Spark Catalog Plugin 机制介绍</a>
        </h1>

        <div class="post-meta">
            <span class="post-meta-item">
              <span class="post-meta-item-icon">
                <i class="fa fa-calendar-o"></i>
              </span>
              <span class="post-meta-item-text">Posted on</span>
              

              <time title="Created: 2020-08-01 17:04:54 / Modified: 17:10:29" itemprop="dateCreated datePublished" datetime="2020-08-01T17:04:54+08:00">2020-08-01</time>
            </span>
            <span class="post-meta-item">
              <span class="post-meta-item-icon">
                <i class="fa fa-folder-o"></i>
              </span>
              <span class="post-meta-item-text">In</span>
                <span itemprop="about" itemscope itemtype="http://schema.org/Thing">
                  <a href="/categories/Spark/" itemprop="url" rel="index">
                    <span itemprop="name">Spark</span>
                  </a>
                </span>
            </span>

          

        </div>
      </header>

    
    
    
    <div class="post-body" itemprop="articleBody">

      
          <p>Spark 3.0 推出了 Catalog Plugin 特性。在 Release Note 里面位于 <strong><em>Highlight</em></strong> 部分。我们这篇文章就来介绍一下 Catalog Plugin 机制。</p>
          <!--noindex-->
            <div class="post-button">
              <a class="btn" href="/2020/08/01/spark-catalog-plugin/#more" rel="contents">
                Read more &raquo;
              </a>
            </div>
          <!--/noindex-->
        
      
    </div>

    
    
    
      <footer class="post-footer">
        <div class="post-eof"></div>
      </footer>
  </article>
  
  
  

        
  
  
  <article itemscope itemtype="http://schema.org/Article" class="post-block home" lang="zh-Hans">
    <link itemprop="mainEntityOfPage" href="http://yoursite.com/2020/07/26/flink-catalog/">

    <span hidden itemprop="author" itemscope itemtype="http://schema.org/Person">
      <meta itemprop="image" content="/img/avatar.jpg">
      <meta itemprop="name" content="legendtkl">
      <meta itemprop="description" content="Do not go gentle into that good night.">
    </span>

    <span hidden itemprop="publisher" itemscope itemtype="http://schema.org/Organization">
      <meta itemprop="name" content="Legendtkl">
    </span>
      <header class="post-header">
        <h1 class="post-title" itemprop="name headline">
          
            <a href="/2020/07/26/flink-catalog/" class="post-title-link" itemprop="url">Flink Catalog 介绍</a>
        </h1>

        <div class="post-meta">
            <span class="post-meta-item">
              <span class="post-meta-item-icon">
                <i class="fa fa-calendar-o"></i>
              </span>
              <span class="post-meta-item-text">Posted on</span>
              

              <time title="Created: 2020-07-26 11:49:54 / Modified: 13:34:47" itemprop="dateCreated datePublished" datetime="2020-07-26T11:49:54+08:00">2020-07-26</time>
            </span>
            <span class="post-meta-item">
              <span class="post-meta-item-icon">
                <i class="fa fa-folder-o"></i>
              </span>
              <span class="post-meta-item-text">In</span>
                <span itemprop="about" itemscope itemtype="http://schema.org/Thing">
                  <a href="/categories/Flink/" itemprop="url" rel="index">
                    <span itemprop="name">Flink</span>
                  </a>
                </span>
            </span>

          

        </div>
      </header>

    
    
    
    <div class="post-body" itemprop="articleBody">

      
          <h2 id="0-引言"><a href="#0-引言" class="headerlink" title="0. 引言"></a>0. 引言</h2><p>这篇文章我们介绍了一下 Flink 的 Catalog，基于 Flink 1.11，熟悉 Flink 或者 Spark 等大数据引擎的同学应该都知道这两个计算引擎都有一个共同的组件叫 Catalog。下面是 Flink 的 Catalog 的官方定义。</p>
<blockquote>
<p>Catalog 提供了元数据信息，例如数据库、表、分区、视图以及数据库或其他外部系统中存储的函数和信息。</p>
<p>数据处理最关键的方面之一是管理元数据。 元数据可以是临时的，例如临时表、或者通过 TableEnvironment 注册的 UDF。 元数据也可以是持久化的，例如 Hive Metastore 中的元数据。Catalog 提供了一个统一的API，用于管理元数据，并使其可以从 Table API 和 SQL 查询语句中来访问。</p>
</blockquote>
<p>简单来说，Catalog 就是元数据管理中心，其中元数据包括数据库、表、表结构等信息。</p>
          <!--noindex-->
            <div class="post-button">
              <a class="btn" href="/2020/07/26/flink-catalog/#more" rel="contents">
                Read more &raquo;
              </a>
            </div>
          <!--/noindex-->
        
      
    </div>

    
    
    
      <footer class="post-footer">
        <div class="post-eof"></div>
      </footer>
  </article>
  
  
  

        
  
  
  <article itemscope itemtype="http://schema.org/Article" class="post-block home" lang="zh-Hans">
    <link itemprop="mainEntityOfPage" href="http://yoursite.com/2020/06/06/kubernetes-scheduler-sourcecode/">

    <span hidden itemprop="author" itemscope itemtype="http://schema.org/Person">
      <meta itemprop="image" content="/img/avatar.jpg">
      <meta itemprop="name" content="legendtkl">
      <meta itemprop="description" content="Do not go gentle into that good night.">
    </span>

    <span hidden itemprop="publisher" itemscope itemtype="http://schema.org/Organization">
      <meta itemprop="name" content="Legendtkl">
    </span>
      <header class="post-header">
        <h1 class="post-title" itemprop="name headline">
          
            <a href="/2020/06/06/kubernetes-scheduler-sourcecode/" class="post-title-link" itemprop="url">源码面前，了无密码：Kuberentes Scheduler 源码剖析</a>
        </h1>

        <div class="post-meta">
            <span class="post-meta-item">
              <span class="post-meta-item-icon">
                <i class="fa fa-calendar-o"></i>
              </span>
              <span class="post-meta-item-text">Posted on</span>

              <time title="Created: 2020-06-06 11:29:00" itemprop="dateCreated datePublished" datetime="2020-06-06T11:29:00+08:00">2020-06-06</time>
            </span>
              <span class="post-meta-item">
                <span class="post-meta-item-icon">
                  <i class="fa fa-calendar-check-o"></i>
                </span>
                <span class="post-meta-item-text">Edited on</span>
                <time title="Modified: 2020-07-26 13:29:36" itemprop="dateModified" datetime="2020-07-26T13:29:36+08:00">2020-07-26</time>
              </span>

          

        </div>
      </header>

    
    
    
    <div class="post-body" itemprop="articleBody">

      
          <p>本篇文章介绍一下 Kubernetes 的默认调度器 kube-scheduler 的源码实现。kubernetes 代码版本：v1.18.4-rc.0。</p>
          <!--noindex-->
            <div class="post-button">
              <a class="btn" href="/2020/06/06/kubernetes-scheduler-sourcecode/#more" rel="contents">
                Read more &raquo;
              </a>
            </div>
          <!--/noindex-->
        
      
    </div>

    
    
    
      <footer class="post-footer">
        <div class="post-eof"></div>
      </footer>
  </article>
  
  
  

        
  
  
  <article itemscope itemtype="http://schema.org/Article" class="post-block home" lang="zh-Hans">
    <link itemprop="mainEntityOfPage" href="http://yoursite.com/2019/11/11/terminal-10x-bash/">

    <span hidden itemprop="author" itemscope itemtype="http://schema.org/Person">
      <meta itemprop="image" content="/img/avatar.jpg">
      <meta itemprop="name" content="legendtkl">
      <meta itemprop="description" content="Do not go gentle into that good night.">
    </span>

    <span hidden itemprop="publisher" itemscope itemtype="http://schema.org/Organization">
      <meta itemprop="name" content="Legendtkl">
    </span>
      <header class="post-header">
        <h1 class="post-title" itemprop="name headline">
          
            <a href="/2019/11/11/terminal-10x-bash/" class="post-title-link" itemprop="url">终端 10X 工作法</a>
        </h1>

        <div class="post-meta">
            <span class="post-meta-item">
              <span class="post-meta-item-icon">
                <i class="fa fa-calendar-o"></i>
              </span>
              <span class="post-meta-item-text">Posted on</span>
              

              <time title="Created: 2019-11-11 10:36:01 / Modified: 10:37:52" itemprop="dateCreated datePublished" datetime="2019-11-11T10:36:01+08:00">2019-11-11</time>
            </span>
            <span class="post-meta-item">
              <span class="post-meta-item-icon">
                <i class="fa fa-folder-o"></i>
              </span>
              <span class="post-meta-item-text">In</span>
                <span itemprop="about" itemscope itemtype="http://schema.org/Thing">
                  <a href="/categories/linux/" itemprop="url" rel="index">
                    <span itemprop="name">linux</span>
                  </a>
                </span>
            </span>

          

        </div>
      </header>

    
    
    
    <div class="post-body" itemprop="articleBody">

      
          <p>在 github 上面有一个 700 多人 star 的 repo 叫做 Bash-Oneliner，介绍了很多实用并且可以有效提高工作效率的命令，我们来了解一下。原文直达：<a href="https://github.com/onceupon/Bash-Oneliner" target="_blank" rel="noopener">Bash-Oneliner</a> 。注：去除了部分看上去没啥用的命令，可以原文查看所有内容。</p>
          <!--noindex-->
            <div class="post-button">
              <a class="btn" href="/2019/11/11/terminal-10x-bash/#more" rel="contents">
                Read more &raquo;
              </a>
            </div>
          <!--/noindex-->
        
      
    </div>

    
    
    
      <footer class="post-footer">
        <div class="post-eof"></div>
      </footer>
  </article>
  
  
  

        
  
  
  <article itemscope itemtype="http://schema.org/Article" class="post-block home" lang="zh-Hans">
    <link itemprop="mainEntityOfPage" href="http://yoursite.com/2019/10/13/uber-go-style/">

    <span hidden itemprop="author" itemscope itemtype="http://schema.org/Person">
      <meta itemprop="image" content="/img/avatar.jpg">
      <meta itemprop="name" content="legendtkl">
      <meta itemprop="description" content="Do not go gentle into that good night.">
    </span>

    <span hidden itemprop="publisher" itemscope itemtype="http://schema.org/Organization">
      <meta itemprop="name" content="Legendtkl">
    </span>
      <header class="post-header">
        <h1 class="post-title" itemprop="name headline">
          
            <a href="/2019/10/13/uber-go-style/" class="post-title-link" itemprop="url">Uber Go 语言编程规范</a>
        </h1>

        <div class="post-meta">
            <span class="post-meta-item">
              <span class="post-meta-item-icon">
                <i class="fa fa-calendar-o"></i>
              </span>
              <span class="post-meta-item-text">Posted on</span>

              <time title="Created: 2019-10-13 12:07:36" itemprop="dateCreated datePublished" datetime="2019-10-13T12:07:36+08:00">2019-10-13</time>
            </span>
              <span class="post-meta-item">
                <span class="post-meta-item-icon">
                  <i class="fa fa-calendar-check-o"></i>
                </span>
                <span class="post-meta-item-text">Edited on</span>
                <time title="Modified: 2019-11-10 12:11:19" itemprop="dateModified" datetime="2019-11-10T12:11:19+08:00">2019-11-10</time>
              </span>
            <span class="post-meta-item">
              <span class="post-meta-item-icon">
                <i class="fa fa-folder-o"></i>
              </span>
              <span class="post-meta-item-text">In</span>
                <span itemprop="about" itemscope itemtype="http://schema.org/Thing">
                  <a href="/categories/golang/" itemprop="url" rel="index">
                    <span itemprop="name">golang</span>
                  </a>
                </span>
            </span>

          

        </div>
      </header>

    
    
    
    <div class="post-body" itemprop="articleBody">

      
          <p>相信很多人前两天都看到 Uber 在 github 上面开源的 Go 语言编程规范了，原文在这里：<a href="https://github.com/uber-go/guide/blob/master/style.md" target="_blank" rel="noopener">https://github.com/uber-go/guide/blob/master/style.md</a> 。我们今天就来简单了解一下国外大厂都是如何来写代码的。行文仓促，错误之处，多多指正。另外如果觉得还不错，也欢迎分享给更多的人。</p>
          <!--noindex-->
            <div class="post-button">
              <a class="btn" href="/2019/10/13/uber-go-style/#more" rel="contents">
                Read more &raquo;
              </a>
            </div>
          <!--/noindex-->
        
      
    </div>

    
    
    
      <footer class="post-footer">
        <div class="post-eof"></div>
      </footer>
  </article>
  
  
  

        
  
  
  <article itemscope itemtype="http://schema.org/Article" class="post-block home" lang="zh-Hans">
    <link itemprop="mainEntityOfPage" href="http://yoursite.com/2019/04/18/about-wx/">

    <span hidden itemprop="author" itemscope itemtype="http://schema.org/Person">
      <meta itemprop="image" content="/img/avatar.jpg">
      <meta itemprop="name" content="legendtkl">
      <meta itemprop="description" content="Do not go gentle into that good night.">
    </span>

    <span hidden itemprop="publisher" itemscope itemtype="http://schema.org/Organization">
      <meta itemprop="name" content="Legendtkl">
    </span>
      <header class="post-header">
        <h1 class="post-title" itemprop="name headline">
          
            <a href="/2019/04/18/about-wx/" class="post-title-link" itemprop="url">合抱之木，生于毫末</a>
        </h1>

        <div class="post-meta">
            <span class="post-meta-item">
              <span class="post-meta-item-icon">
                <i class="fa fa-calendar-o"></i>
              </span>
              <span class="post-meta-item-text">Posted on</span>

              <time title="Created: 2019-04-18 21:55:17" itemprop="dateCreated datePublished" datetime="2019-04-18T21:55:17+08:00">2019-04-18</time>
            </span>
              <span class="post-meta-item">
                <span class="post-meta-item-icon">
                  <i class="fa fa-calendar-check-o"></i>
                </span>
                <span class="post-meta-item-text">Edited on</span>
                <time title="Modified: 2019-11-08 11:09:25" itemprop="dateModified" datetime="2019-11-08T11:09:25+08:00">2019-11-08</time>
              </span>
            <span class="post-meta-item">
              <span class="post-meta-item-icon">
                <i class="fa fa-folder-o"></i>
              </span>
              <span class="post-meta-item-text">In</span>
                <span itemprop="about" itemscope itemtype="http://schema.org/Thing">
                  <a href="/categories/%E4%B8%AA%E4%BA%BA%E5%85%B4%E8%B6%A3/" itemprop="url" rel="index">
                    <span itemprop="name">个人兴趣</span>
                  </a>
                </span>
            </span>

          

        </div>
      </header>

    
    
    
    <div class="post-body" itemprop="articleBody">

      
          <blockquote>
<p>合抱之木，生于毫末；九层之台，起于垒土；千里之行，始于足下。</p>
</blockquote>
<p>“合抱之木，生于毫末”是我非常喜欢的一句话，强调万事积于忽微，也就是积累的重要性。</p>
<p>要说到积累，我积累的最多的应该就是写代码和写作。由于接触计算机比较晚，本科阶段一直想做学术，所以本科阶段的编码训练并不是很多。在本科毕业之后，本来是去中科院硕博连读，经过一个学期，终于决定还是要去工业界。确定了之后，就开始高强度的代码编程训练。光《算法导论》我就刷了三遍，课后习题做了一遍。《深入理解计算机系统》等书籍也是刷了三遍以上。研究生阶段基本由个人负责的大型项目的代码行数也是在十万行以上。找工作的时候参加 Google Code Jam 校招比赛，打到全球前两百名，最后通过 Google 的校招算法面试。</p>
          <!--noindex-->
            <div class="post-button">
              <a class="btn" href="/2019/04/18/about-wx/#more" rel="contents">
                Read more &raquo;
              </a>
            </div>
          <!--/noindex-->
        
      
    </div>

    
    
    
      <footer class="post-footer">
        <div class="post-eof"></div>
      </footer>
  </article>
  
  
  

  </div>

  
  <nav class="pagination">
    <span class="page-number current">1</span><a class="page-number" href="/page/2/">2</a><span class="space">&hellip;</span><a class="page-number" href="/page/12/">12</a><a class="extend next" rel="next" href="/page/2/"><i class="fa fa-angle-right" aria-label="Next page"></i></a>
  </nav>



          </div>
          

        </div>
          
  
  <div class="toggle sidebar-toggle">
    <span class="toggle-line toggle-line-first"></span>
    <span class="toggle-line toggle-line-middle"></span>
    <span class="toggle-line toggle-line-last"></span>
  </div>

  <aside class="sidebar">
    <div class="sidebar-inner">

      <ul class="sidebar-nav motion-element">
        <li class="sidebar-nav-toc">
          Table of Contents
        </li>
        <li class="sidebar-nav-overview">
          Overview
        </li>
      </ul>

      <!--noindex-->
      <div class="post-toc-wrap sidebar-panel">
      </div>
      <!--/noindex-->

      <div class="site-overview-wrap sidebar-panel">
        <div class="site-author motion-element" itemprop="author" itemscope itemtype="http://schema.org/Person">
  <img class="site-author-image" itemprop="image" alt="legendtkl"
    src="/img/avatar.jpg">
  <p class="site-author-name" itemprop="name">legendtkl</p>
  <div class="site-description" itemprop="description">Do not go gentle into that good night.</div>
</div>
<div class="site-state-wrap motion-element">
  <nav class="site-state">
      <div class="site-state-item site-state-posts">
          <a href="/archives/">
        
          <span class="site-state-item-count">114</span>
          <span class="site-state-item-name">posts</span>
        </a>
      </div>
      <div class="site-state-item site-state-categories">
            <a href="/categories/">
          
        <span class="site-state-item-count">33</span>
        <span class="site-state-item-name">categories</span></a>
      </div>
      <div class="site-state-item site-state-tags">
            <a href="/tags/">
        <span class="site-state-item-count">87</span>
        <span class="site-state-item-name">tags</span></a>
      </div>
  </nav>
</div>
  <div class="feed-link motion-element">
    <a href="/atom.xml" rel="alternate">
      <i class="fa fa-rss"></i>RSS
    </a>
  </div>
  <div class="links-of-author motion-element">
      <span class="links-of-author-item">
        <a href="https://github.com/legendtkl" title="Github &amp;rarr; https:&#x2F;&#x2F;github.com&#x2F;legendtkl" rel="noopener" target="_blank"><i class="fa fa-fw fa-github"></i>Github</a>
      </span>
      <span class="links-of-author-item">
        <a href="http://weibo.com/HIT_Achilles" title="Weibo &amp;rarr; http:&#x2F;&#x2F;weibo.com&#x2F;HIT_Achilles" rel="noopener" target="_blank"><i class="fa fa-fw fa-weibo"></i>Weibo</a>
      </span>
      <span class="links-of-author-item">
        <a href="http://www.zhihu.com/people/legendtkl" title="Zhihu &amp;rarr; http:&#x2F;&#x2F;www.zhihu.com&#x2F;people&#x2F;legendtkl" rel="noopener" target="_blank"><i class="fa fa-fw fa-zhihu"></i>Zhihu</a>
      </span>
  </div>



      </div>

    </div>
  </aside>
  <div id="sidebar-dimmer"></div>


      </div>
    </main>

    <footer class="footer">
      <div class="footer-inner">
        

<div class="copyright">
  
  &copy; 
  <span itemprop="copyrightYear">2021</span>
  <span class="with-love">
    <i class="fa fa-user"></i>
  </span>
  <span class="author" itemprop="copyrightHolder">legendtkl</span>
</div>
  <div class="powered-by">Powered by <a href="https://hexo.io/" class="theme-link" rel="noopener" target="_blank">Hexo</a> v4.0.0
  </div>
  <span class="post-meta-divider">|</span>
  <div class="theme-info">Theme – <a href="https://mist.theme-next.org/" class="theme-link" rel="noopener" target="_blank">NexT.Mist</a> v7.5.0
  </div>

        












        
      </div>
    </footer>
  </div>

  
  <script src="/lib/anime.min.js"></script>
  <script src="/lib/velocity/velocity.min.js"></script>
  <script src="/lib/velocity/velocity.ui.min.js"></script>
<script src="/js/utils.js"></script><script src="/js/motion.js"></script>
<script src="/js/schemes/muse.js"></script>
<script src="/js/next-boot.js"></script>



  
















  

  

</body>
</html>

`

func Test_getArticleBody(t *testing.T) {
	//type args struct {
	//	str string
	//}
	//tests := []struct {
	//	name string
	//	args args
	//	want string
	//}{
	//	// TODO: Add test cases.
	//	{
	//		name: "test",
	//		args: args{str: body},
	//	},
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		getArticleBody(tt.args.str)
	//	})
	//}
	getArticleBodyPreContent(body)
}
