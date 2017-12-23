package bot

type Bot interface {
	Greeting() string
	Reply(string) string
	Name() string
	Title() string
	ThumbnailPath() string
	SetName(string)
	SetTitle(string)
	SetThumbnailPath(string)
}
