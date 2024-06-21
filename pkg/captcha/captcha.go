package captcha

import (
	"errors"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

type Config struct {
	Type       string
	Height     int
	Width      int
	NoiseCount int
	Length     int
	ExpireTime int
}

var source = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func NewCaptcha(conf Config, store base64Captcha.Store) (*base64Captcha.Captcha, error) {
	col := &color.RGBA{255, 255, 255, 255}
	fontsStorage := base64Captcha.DefaultEmbeddedFonts
	fonts := make([]string, 0)
	showLineOption := base64Captcha.OptionShowHollowLine

	var driver base64Captcha.Driver
	switch conf.Type {
	case "math":
		// 图形化算数验证码
		driver = base64Captcha.NewDriverMath(conf.Height, conf.Width, conf.NoiseCount, showLineOption, col, fontsStorage, fonts)
	case "string":
		//showLineOption = showLineOption | base64Captcha.OptionShowSineLine
		driver = base64Captcha.NewDriverString(conf.Height, conf.Width, conf.NoiseCount, showLineOption, conf.Length, source, col, fontsStorage, fonts)
	case "audio":
		driver = base64Captcha.NewDriverAudio(conf.Length, "zh")
	//case "lang":
	//	driver = base64Captcha.NewDriverLanguage(conf.Height, conf.Width, conf.NoiseCount, showLineOption, conf.Length, col, fontsStorage, nil, source)
	//case "digit":
	//	driver = base64Captcha.NewDriverDigit(conf.Height, conf.Width, conf.Length, 0.45, 80)
	//case "chinese":
	//	driver = base64Captcha.NewDriverChinese(conf.Height, conf.Width, conf.NoiseCount, showLineOption, conf.Length, source, col, fontsStorage, fonts)
	default:
		return nil, errors.New("验证码类型不正确，请配置正确的验证码类型")
	}

	return base64Captcha.NewCaptcha(driver, store), nil
}
