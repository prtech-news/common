package common

import (
	"bytes"
	"html/template"
	"log"
)

const htmlTemplate = `
<html op="news">
<head>
    <meta name="referrer" content="origin">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" type="text/css" href="">
    <link rel="shortcut icon" href="">
    <!--<link rel="alternate" type="application/rss+xml" title="RSS" href="rss">-->
    <title>Puerto Rico - Tech News</title></head>

<body>
<center>
    <table id="hnmain" border="0" cellpadding="0" cellspacing="0" width="85%" bgcolor="#F4F5F7">
        <tr>
            <td bgcolor="#3a91ca">
                <table border="0" cellpadding="0" cellspacing="0" width="100%" style="padding:2px">
                    <tr>
                        <td style="width:18px;padding-right:4px"><a href="/prtech-news"><img
                                src="" width="18" height="18"
                                style="border:1px white solid;"></a></td>
                        <td style="line-height:12pt; height:10px;"><span class="pagetop"><b class="hnname">
                            Puerto Rico Tech News</b></span>
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
        <tr id="pagespace" title="" style="height:10px"></tr>
        <tr>
            <td>
                <table border="0" cellpadding="0" cellspacing="0" class="itemlist">
                    {{range .}}
                    <tr class='athing' id=''>
                        <td align="right" valign="top" class="title"><span class="rank"></span>
                        </td>
                        <td valign="top" class="votelinks">
                            <center><a id=''>
                                <div class='votearrow' title='upvote'></div>
                            </a></center>
                        </td>
                        <td class="title"><a href="{{ .Link }}" target="_blank" class="storylink">{{ .Title }}
                        </a><span class="sitebit comhead"> (<a href="/prtech-news/from?site={{ .Source }}"
                                                               target="_blank"><span
                                class="sitestr">{{ .Source }}</span></a>)</span></td>
                    </tr>
                    <tr>
                        <td colspan="2"></td>
                        <td class="subtext">
                            Added by<a href="/" target="_blank" class="hnuser"> prtech.news</a> |
                            <span class="age">Published on {{ .PubDate }}</span>
                        </td>
                    </tr>
                    <tr class="spacer" style="height:5px"></tr>
                    {{end}}

                    <tr class="morespace" style="height:10px"></tr>
                    <tr>
                        <td colspan="2"></td>
                        <td class="title"><a href="" class="morelink"
                                             rel="next">Top</a></td>
                    </tr>
                </table>
            </td>
        </tr>
        <tr>
            <td><img src="s.gif" height="10" width="0">
                <table width="100%" cellspacing="0" cellpadding="1">
                    <tr>
                        <td bgcolor="#3a91ca"></td>
                    </tr>
                </table>
                <br>
                <center><span class="yclinks">
                </span><br><br>
                    <!--<form method="get" action="">Search:-->
                    <!--<input type="text" name="q" value="" size="17" autocorrect="off" spellcheck="false"-->
                    <!--autocapitalize="off" autocomplete="false"></form>-->
                    <!--<p>This site is not affiliated or linked with <a href="https://news.ycombinator.com/" target="_blank" style="color: blue">Hacker News [Y Combinator]</a>.</p>-->
                    <br><br>
                </center>
            </td>
        </tr>
    </table>
</center>
</body>
</html>
`

func CreateHtmlFromArticles(articles []*Article) ([]byte, error) {
	// template.ParseFiles("views/layout.html")
	tmpl, err := template.New("index").Parse(htmlTemplate)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	tmpl.Execute(&buf, articles) // TODO verify anonymous struct syntax
	bytes := buf.Bytes()
	log.Printf("Created HTML page with %d number of bytes\n", len(bytes))
	return bytes, nil
}
