package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVideoMime(t *testing.T) {
	//key - source url
	//val - expexted result
	testData := map[string]string{
		"":    "",
		"   ": "",
		"http://test.com/image.mp4":  MimeVideoMp4,
		"http://test.com/image.m4v":  MimeVideoMp4,
		"http://test.com/image.mp4v": MimeVideoMp4,
		"http://test.com/image.3gpp": MimeVideo3gpp,
		"http://test.com/image.3gp":  MimeVideo3gpp,

		"http://test.com/image.webm": MimeVideoWebm,
		"http://test.com/image.flv":  MimeVideoXflv,

		"http://test.com/image.m1v":  MimeVideoMpeg,
		"http://test.com/image.m2v":  MimeVideoMpeg,
		"http://test.com/image.mod":  MimeVideoMpeg,
		"http://test.com/image.mp2v": MimeVideoMpeg,
		"http://test.com/image.mpa":  MimeVideoMpeg,
		"http://test.com/image.mpe":  MimeVideoMpeg,
		"http://test.com/image.mpeg": MimeVideoMpeg,
		"http://test.com/image.mpg":  MimeVideoMpeg,
		"http://test.com/image.mpv2": MimeVideoMpeg,
	}

	for sourceURL, result := range testData {
		assert.Equal(t, result, GetVideoMime(sourceURL))
	}
}

func TestIsValidVideoURL(t *testing.T) {
	//key - source url
	//val - expexted result
	testData := map[string]bool{
		"":    false,
		"   ": false,
		"http://test.com/index.html":  false,
		"http://test.com/index.xml":   false,
		"http://test.com/image.mp4":   true,
		"http://test./image.mp4":      false,
		"http://test.com/image.mpeg":  true,
		"http://te!st.com/image.mpeg": false,
		"http://test.com/image.3gpp":  true,
		"http://test.t/image.3gpp":    false,
		"http://test.com/image.webm":  true,
		"ftp://test.com/image.webm":   false,
		"http://test.com/image.flv":   true,
		"http://te_st.com/image.flv":  false,
	}

	for sourceURL, result := range testData {
		assert.Equal(t, result, IsValidVideoURL(sourceURL))
	}
}

func TestIsValidImageURL(t *testing.T) {
	//key - source url
	//val - expexted result
	testData := map[string]bool{
		"":    false,
		"   ": false,
		"http://test.com/index.html":  false,
		"http://test.com/index.xml":   false,
		"http://test.com/image.png":   true,
		"http://test./image.mp4":      false,
		"http://test.com/image.jpeg":  true,
		"http://te!st.com/image.mpeg": false,
		"http://test.com/image.jpg":   true,
		"http://test.t/image.3gpp":    false,
		"http://test.com/image.gif":   true,
		"ftp://test.com/image.webm":   false,
	}

	for sourceURL, result := range testData {
		assert.Equal(t, result, IsValidImageURL(sourceURL))
	}
}

func TestIsValidHTMLURL(t *testing.T) {
	//key - source url
	//val - expexted result
	testData := map[string]bool{
		"":    false,
		"   ": false,
		"http://test.com/index.html": true,
		"http://test.com/index.xml":  false,
		"http://test.com/image.htm":  true,
		"http://test./image.mp4":     false,
		"http://test.com/image.jpeg": false,
	}

	for sourceURL, result := range testData {
		assert.Equal(t, result, IsValidHTMLURL(sourceURL))
	}
}
