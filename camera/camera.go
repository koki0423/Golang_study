package main

import (
	"gocv.io/x/gocv"
)

func main() {
	// カメラデバイスを開く
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		panic(err)
	}
	defer webcam.Close()

	// ウィンドウを作成
	window := gocv.NewWindow("Hello")
	defer window.Close()

	// 画像フレームを格納するMatを作成
	img := gocv.NewMat()
	defer img.Close()

	for {
		// 新しいフレームを読み込む
		if ok := webcam.Read(&img); !ok {
			println("カメラから画像を読み取れませんでした。")
			return
		}
		if img.Empty() {
			continue
		}

		// ウィンドウに画像を表示
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
