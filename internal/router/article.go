package router

import (
	"log"
	"strconv"
	"time"

	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	"github.com/megrez/internal/dao"
	"github.com/megrez/internal/entity/po"
	"github.com/megrez/internal/entity/vo"
)

var DAO *dao.DAO

func routeArticle(g *gin.Engine, dao *dao.DAO) {
	DAO = dao
	g.GET("/", index)
	g.GET("/index/:pageNum", index)
	g.GET("/article/:id", articleDetail)
	g.GET("/article/:id/comment-page/:pageNum", articleDetail)
	g.POST("/admin/article", createArticle)
}

func index(c *gin.Context) {
	var pageNum, pageSize int
	var err error
	if c.Param("pageNum") == "" {
		pageNum = 1
	} else {
		pageNum, err = strconv.Atoi(c.Param("pageNum"))
		if err != nil {
			log.Println("incorrect param pageNum, err:", err)
			// TODO: 应该是 4XX?
			c.Redirect(500, "/error")
		}
	}
	if c.Param("pageSize") == "" {
		pageSize = 10
	} else {
		pageSize, err = strconv.Atoi(c.Param("pageSize"))
		if err != nil {
			log.Println("incorrect param pageSize, err:", err)
			c.Redirect(500, "/error")
		}
	}

	articlePOs, err := DAO.ListAllArticles(pageNum, pageSize)
	if err != nil {
		log.Println("get articles from db failed, err:", err)
		c.Redirect(500, "/error")
	}
	articleVOs := []vo.IndexArticle{}
	for _, articlePO := range articlePOs {
		articleVO, err := vo.GetIndexArticleFromPO(&articlePO)
		if err != nil {
			c.Redirect(500, "/error")
		}
		articleVOs = append(articleVOs, articleVO)
	}
	globalOption, err := vo.GetGlobalOption()
	if err != nil {
		c.Redirect(500, "/error")
	}
	// pageInfo
	count, err := DAO.CountAllArticles()
	if err != nil {
		c.Redirect(500, "/error")
	}
	page := vo.CaculatePage(pageNum, pageSize, int(count))
	// TODO: 过滤器格式化时间
	c.HTML(200, "index.html", pongo2.Context{"articles": articleVOs, "global": globalOption, "page": page})
}

func articleDetail(c *gin.Context) {
	var pageNum, pageSize int
	var err error
	if c.Param("pageNum") == "" {
		pageNum = 1
	} else {
		pageNum, err = strconv.Atoi(c.Param("pageNum"))
		if err != nil {
			log.Println("incorrect param pageNum, err:", err)
			// TODO: 应该是 4XX?
			c.Redirect(500, "/error")
		}
	}

	if c.Param("pageSize") == "" {
		pageSize = 8
	} else {
		pageSize, err = strconv.Atoi(c.Param("pageSize"))
		if err != nil {
			log.Println("incorrect param pageSize, err:", err)
			c.Redirect(500, "/error")
		}
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("incorrect param id, err:", err)
		c.Redirect(500, "/error")
	}
	articlePO, err := DAO.GetArticleByID(uint(id))
	if err != nil {
		log.Println("query article from db failed, err: ", err)
	}
	articleDetial, err := vo.GetArticleDetailFromPO(&articlePO, pageNum, pageSize)
	if err != nil {
		log.Println("get article detail failed, err:", err)
		c.Redirect(500, "/error")
	}
	globalOption, err := vo.GetGlobalOption()
	if err != nil {
		c.Redirect(500, "/error")
	}
	c.HTML(200, "article.html", pongo2.Context{"article": articleDetial, "global": globalOption})
}

func createArticle(c *gin.Context) {
	for i := 2; i <= 200; i++ {
		article := &po.Article{
			Title:           "测试标题",
			OriginalContent: "测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文测试正文",
			Summary:         "测试摘要",
			Thumb:           "https://cdn.rawchen.com/2020/11/student-cloud-server/1.png",
			FormatContent: `<blockquote>
		<p>购买地址：<br>腾讯云：<a
				href="https://cloud.tencent.com/act/cps/redirect?redirect=1575&cps_key=b58209dca42a9decdae2f276d61acd77&from=console"
				target="_blank">https://cloud.tencent.com</a><br>阿里云：<a
				href="https://www.aliyun.com/activity/1111?userCode=kui4i1ow"
				target="_blank">https://www.aliyun.com</a></p>
	</blockquote>
	<p><img src="https://cdn.rawchen.com/2020/11/student-cloud-server/1.png" alt="" title=""></p>
	<p><img src="https://cdn.rawchen.com/2020/11/student-cloud-server/0.png" alt="" title=""></p>
	<h1>简介</h1>
	<p>云服务器(Elastic Compute Service,
		ECS)是一种简单高效、安全可靠、处理能力可弹性伸缩的计算服务。其管理方式比物理服务器更简单高效。用户无需提前购买硬件，即可迅速创建或释放任意多台云服务器。</p>
	<p>什么意思呢？通俗来说就是一台能通过外部网络连接的全天开机的托管服务器，你能用它来干什么呢？</p>
	<h1>用处</h1>
	<ol>
		<li><strong>建站（个人网站、博客、电商网站、论坛等）</strong>，不管是静态的html还是动态的php，jsp，.net编写的网页都可以放到web服务器并以域名形式发布供外网浏览。</li>
		<li><strong>云存储（云文件下载库）</strong>，现如今各种知名网盘，都不好用，要不就是限流限的厉害，要不就是上传大小限制等。自己的服务器多爽啊，传个资料分分钟的事。而且还安全，什么？有人攻击，你来攻击我阿里云服务器试试?
		</li>
		<li><strong>办公系统应用（专属电子邮箱、OA、会员管理系统等）</strong></li>
		<li><strong>搭建微信公众号后台 5. 渲染和视频转码 6. 用作下载机 7. 内网穿透 8. 跑些Python、Java小程序等</strong></li>
	</ol>
	<h1>建站</h1>
	<p>好像听起来自己也就只会建站，也只想建个自己的博客？那就接着看。<br>步骤大概分这几个：<br>注册阿里云帐号，通过学生实名认证，通过<a
			href="https://www.aliyun.com/activity/1111?userCode=kui4i1ow"
			target="_blank">https://www.aliyun.com</a>购买学生云服务器，更改服务器实例的登录密码，通过远程命令一键安装宝塔管理面板，一键安装LNMP环境，创建网站并外网访问。
	</p>
	<p>步骤好像也挺简单，那具体有教程吗，其实网站挺多的了，就是可能不太详细。要不我写个详细教程？</p>
	<h1>详细建站教程</h1>`,
			TopPriority: 0,
			Status:      0,
			Private:     false,
			Visits:      5,
			Slug:        strconv.Itoa(i),
		}
		err := DAO.CreateArticle(article)
		if err != nil {
			log.Println("create article failed, err: ", err)
			return
		}
	}
	c.JSON(200, "success")
}

func FormatTime(t time.Time, str string) string {
	return t.Format(str)
}
