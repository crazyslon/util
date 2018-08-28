package util

import "path/filepath"

const (
	//MimeVideoMp4 video/mp4
	MimeVideoMp4 = "video/mp4"

	//MimeVideoWebm  "video/webm"
	MimeVideoWebm = "video/webm"

	//MimeVideo3gpp "video/3gpp"
	MimeVideo3gpp = "video/3gpp"

	//MimeVideoXflv "video/x-flv"
	MimeVideoXflv = "video/x-flv"

	//MimeVideoMpeg "video/mpeg"
	MimeVideoMpeg = "video/mpeg"

	//MimeImageJpeg "image/jpeg"
	MimeImageJpeg = "image/jpeg"

	//MimeImagePng "image/png"
	MimeImagePng = "image/png"

	//MimeImageGif "image/gif"
	MimeImageGif = "image/gif"

	//MimeTextHTML "text/html"
	MimeTextHTML = "text/html"
)

var videoMimes = map[string]string{
	".mp4":  MimeVideoMp4,
	".m4v":  MimeVideoMp4,
	".mp4v": MimeVideoMp4,
	".webm": MimeVideoWebm,
	".3gp":  MimeVideo3gpp,
	".3gpp": MimeVideo3gpp,
	".flv":  MimeVideoXflv,
	".m1v":  MimeVideoMpeg,
	".m2v":  MimeVideoMpeg,
	".mod":  MimeVideoMpeg,
	".mp2v": MimeVideoMpeg,
	".mpa":  MimeVideoMpeg,
	".mpe":  MimeVideoMpeg,
	".mpeg": MimeVideoMpeg,
	".mpg":  MimeVideoMpeg,
	".mpv2": MimeVideoMpeg,
}

var imageMimes = map[string]string{
	".gif":  MimeImageGif,
	".png":  MimeImagePng,
	".jpeg": MimeImageJpeg,
	".jpg":  MimeImageJpeg,
}

var htmlMimes = map[string]string{
	".html": MimeTextHTML,
	".htm":  MimeTextHTML,
}

//GetVideoMime return video mime type by source url
//If file in url has not supported extension return empty string
func GetVideoMime(sourceURL string) string {
	return getMime(sourceURL, videoMimes)
}

//GetImageMime return image mime type by source url
//If file in url has not supported extension return empty string
func GetImageMime(sourceURL string) string {
	return getMime(sourceURL, imageMimes)
}

//GetHTMLMime return html mime type by source url
//If file in url has not supported extension return empty string
func GetHTMLMime(sourceURL string) string {
	return getMime(sourceURL, htmlMimes)
}

func getMime(sourceURL string, mimes map[string]string) string {
	ext := filepath.Ext(sourceURL)
	mime, _ := mimes[ext]
	return mime
}

//IsValidVideoURL check is string is valid video path or not
func IsValidVideoURL(rawurl string) bool {
	return IsValidURL(rawurl) && len(GetVideoMime(rawurl)) > 0
}

//IsValidImageURL check is string is valid image path or not
func IsValidImageURL(rawurl string) bool {
	return IsValidURL(rawurl) && len(GetImageMime(rawurl)) > 0
}

//IsValidHTMLURL check is string is valid html path or not
func IsValidHTMLURL(rawurl string) bool {
	return IsValidURL(rawurl) && len(GetHTMLMime(rawurl)) > 0
}
