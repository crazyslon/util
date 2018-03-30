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

//GetVideoMime return video mime type by source url
//If file in url has not supported extension return empty string
func GetVideoMime(sourceURL string) string {
	ext := filepath.Ext(sourceURL)
	mime, _ := videoMimes[ext]
	return mime
}

//IsValidVideoURL check is string is valid video path or not
func IsValidVideoURL(rawurl string) bool {
	return IsValidURL(rawurl) && len(GetVideoMime(rawurl)) > 0
}
