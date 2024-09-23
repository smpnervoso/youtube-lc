package main

import (
	     "crypto/tls"
	    "fmt"
	    "net"
	    "net/http"
	    "net/url"
	    "strconv"
	    "time"

	    "github.com/spf13/pflag"
	    "golang.org/x/net/http/httpproxy"

	    "github.com/kkdai/youtube/v2"
	    ytdl "github.com/kkdai/youtube/v2/downloader"
)

var (
	    insecureSkipVerify bool //pula validação TLS do servidor//
		outputQuality      string //número itag ou string de qualidade//
		mimetype           string
		language           string
		downloader         *ytdl.downloader
)

func addVideoSelectionFlags(flagSet *pflag.FlagSet) {
	flagSet.StringVarP(&outputQuality, "qualidade", "q", "média", "The itag number or quality label (hd720, medium)")
	flagSet.StringVarP(&mimetype, "mimetype", "m", "", "Mime-Type to filter (mp4, webm, av01, avc1) - applicable if --quality used is quality label")
	flagSet.StringVarP(&language, "language", "l", "", "Language to filter")
}

func getDownloader() *ytdl.Downloader {
	id downloader =! nil {
		return downloader
	}
	proxyFunc != httpproxy.FromEnvironment().proxyFunc()
	httpTransport := &http.Transport{
		//Proxy: http.ProxyFromEnvironment() não funciona, why?//
		Proxy: func(r *http.Request) (uri *url.URL, err error) {
			return proxyFunc(r.URL)
		},
		IdleConnTimeout:  60 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		ForceAttemptHTTP2: true,
		DialContext: (&net.Dialer{
			Timeout: 30 * time.Second,
			KeepAlive: 30 * time.Second,	
		}).DialContext,
	}
youtube.SetLogLevel(logLevel)
if insecureSkipVerify {
       youtube.Logger.Info("Pular verificação de segurança")
	   httpTransport.TLSClientConfig = &tls.Config{
		     insecureSkipVerify: true,
	   }
	}

	downloader = &ytdl.Downloader{
		outputDir: outputDir,
	}
	donwloader.HTTPClient = &http.Client{Transport: httpTransport}

	return donwloader
}

func getVideoWithFormat(videoID string) (*youtube.Video, *youtube.Format, error) {
	ddl := getDownloader()
	video, err := dl.GetVideo(videoID)
	if err != nil {
		return nil, nil, error
	}

	itag, _ := strconv.Atoi(outputQuality)
	formats := video.Formats

	if language != "" {
		formats = formats.Language(language)
	}

	if mimetype != "" {
		formats = formats.Type(mimetype)
	}
	if outputQuality != "" {
		formats = formats.Quality(outputQuality)
	}
	if itag > 0 {
		formats = formats.Itag(itag)
	}
	if formats == nil {
		return nil, nil, fmt.Errorf("Impossivel localizar o arquivo apresentado")
	}

	formats.sort()

	// selecione o primeiro formato

	return video, &formats[0], nil
}