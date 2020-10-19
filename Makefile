build:
	go build wallpaper.go

prod:
	go build -ldflags "-s -w" wallpaper.go

hidden:
	go build -ldflags "-s -w -H windowsgui" wallpaper.go