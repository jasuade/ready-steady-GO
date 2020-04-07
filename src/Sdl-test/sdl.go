package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	w, r, err := initSdl()
	if err != nil {
		fmt.Printf("Ërror %v", err)
	}
	defer sdl.Quit()
	defer w.Destroy()
	defer r.Destroy()
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
	}
}
func initSdl() (*sdl.Window, *sdl.Renderer, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, nil, err
	}
	// func CreateWindowAndRenderer(w, h int32, flags uint32) (*Window, *Renderer, error)
	w, r, err := sdl.CreateWindowAndRenderer(1000, 500, sdl.WINDOW_OPENGL)
	if err != nil {
		return nil, nil, err
	}
	texture, err := loadPNG("hundo.png", r)
	if err != nil {
		return nil, nil, err
	}
	//Draw
	r.SetDrawColor(45, 205, 139, 255)
	r.Clear()
	//Create src and dst rect
	srcRect := &sdl.Rect{X: 0, Y: 0, W: 370, H: 300}
	dstRect := &sdl.Rect{(1000 - 370) / 2, (500 - 300) / 2, 370, 300}
	if err := r.Copy(texture, srcRect, dstRect); err != nil {
		return nil, nil, err
	}
	r.Present()
	return w, r, nil
}
func loadPNG(filePath string, r *sdl.Renderer) (*sdl.Texture, error) {
	file := sdl.RWFromFile(filePath, "rb")
	defer file.Close()
	img, err := img.LoadPNGRW(file)
	defer img.Free()
	if err != nil {
		return nil, err
	}
	sdlTexture, err := r.CreateTextureFromSurface(img)
	if err != nil {
		return nil, err
	}
	return sdlTexture, nil
}
