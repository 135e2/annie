package extractors

import (
	"net/url"
	"strings"

	"github.com/135e2/annie/extractors/acfun"
	"github.com/135e2/annie/extractors/bcy"
	"github.com/135e2/annie/extractors/bilibili"
	"github.com/135e2/annie/extractors/douyin"
	"github.com/135e2/annie/extractors/douyu"
	"github.com/135e2/annie/extractors/eporner"
	"github.com/135e2/annie/extractors/facebook"
	"github.com/135e2/annie/extractors/geekbang"
	"github.com/135e2/annie/extractors/haokan"
	"github.com/135e2/annie/extractors/hupu"
	"github.com/135e2/annie/extractors/huya"
	"github.com/135e2/annie/extractors/instagram"
	"github.com/135e2/annie/extractors/iqiyi"
	"github.com/135e2/annie/extractors/mgtv"
	"github.com/135e2/annie/extractors/miaopai"
	"github.com/135e2/annie/extractors/netease"
	"github.com/135e2/annie/extractors/pixivision"
	"github.com/135e2/annie/extractors/pornhub"
	"github.com/135e2/annie/extractors/qq"
	"github.com/135e2/annie/extractors/streamtape"
	"github.com/135e2/annie/extractors/tangdou"
	"github.com/135e2/annie/extractors/tiktok"
	"github.com/135e2/annie/extractors/tumblr"
	"github.com/135e2/annie/extractors/twitter"
	"github.com/135e2/annie/extractors/types"
	"github.com/135e2/annie/extractors/udn"
	"github.com/135e2/annie/extractors/universal"
	"github.com/135e2/annie/extractors/vimeo"
	"github.com/135e2/annie/extractors/weibo"
	"github.com/135e2/annie/extractors/ximalaya"
	"github.com/135e2/annie/extractors/xvideos"
	"github.com/135e2/annie/extractors/yinyuetai"
	"github.com/135e2/annie/extractors/youku"
	"github.com/135e2/annie/extractors/youtube"
	"github.com/135e2/annie/utils"
)

var extractorMap map[string]types.Extractor

func init() {
	douyinExtractor := douyin.New()
	youtubeExtractor := youtube.New()
	stExtractor := streamtape.New()

	extractorMap = map[string]types.Extractor{
		"": universal.New(), // universal extractor

		"douyin":     douyinExtractor,
		"iesdouyin":  douyinExtractor,
		"bilibili":   bilibili.New(),
		"bcy":        bcy.New(),
		"pixivision": pixivision.New(),
		"youku":      youku.New(),
		"youtube":    youtubeExtractor,
		"youtu":      youtubeExtractor, // youtu.be
		"iqiyi":      iqiyi.New(iqiyi.SiteTypeIqiyi),
		"iq":         iqiyi.New(iqiyi.SiteTypeIQ),
		"mgtv":       mgtv.New(),
		"tangdou":    tangdou.New(),
		"tumblr":     tumblr.New(),
		"vimeo":      vimeo.New(),
		"facebook":   facebook.New(),
		"douyu":      douyu.New(),
		"miaopai":    miaopai.New(),
		"163":        netease.New(),
		"weibo":      weibo.New(),
		"ximalaya":   ximalaya.New(),
		"instagram":  instagram.New(),
		"twitter":    twitter.New(),
		"qq":         qq.New(),
		"yinyuetai":  yinyuetai.New(),
		"geekbang":   geekbang.New(),
		"pornhub":    pornhub.New(),
		"xvideos":    xvideos.New(),
		"udn":        udn.New(),
		"tiktok":     tiktok.New(),
		"haokan":     haokan.New(),
		"acfun":      acfun.New(),
		"eporner":    eporner.New(),
		"streamtape": stExtractor,
		"streamta":   stExtractor, // streamta.pe
		"hupu":       hupu.New(),
		"huya":       huya.New(),
	}
}

// Extract is the main function to extract the data.
func Extract(u string, option types.Options) ([]*types.Data, error) {
	u = strings.TrimSpace(u)
	var domain string

	bilibiliShortLink := utils.MatchOneOf(u, `^(av|BV|ep)\w+`)
	if len(bilibiliShortLink) > 1 {
		bilibiliURL := map[string]string{
			"av": "https://www.bilibili.com/video/",
			"BV": "https://www.bilibili.com/video/",
			"ep": "https://www.bilibili.com/bangumi/play/",
		}
		domain = "bilibili"
		u = bilibiliURL[bilibiliShortLink[1]] + u
	} else {
		u, err := url.ParseRequestURI(u)
		if err != nil {
			return nil, err
		}
		if u.Host == "haokan.baidu.com" {
			domain = "haokan"
		} else {
			domain = utils.Domain(u.Host)
		}
	}
	extractor := extractorMap[domain]
	if extractor == nil {
		extractor = extractorMap[""]
	}
	videos, err := extractor.Extract(u, option)
	if err != nil {
		return nil, err
	}
	for _, v := range videos {
		v.FillUpStreamsData()
	}
	return videos, nil
}
