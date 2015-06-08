package debug

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"path"
	"path/filepath"
)

// bindata_read reads the given file from disk. It returns an error on failure.
func bindata_read(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// assets_css_common_app_thema_css reads file data from disk. It returns an error on failure.
func assets_css_common_app_thema_css() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/css/common/app-thema.css"
	name := "assets/css/common/app-thema.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_common_app_thema_css_min reads file data from disk. It returns an error on failure.
func assets_css_common_app_thema_css_min() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/css/common/app-thema.css.min"
	name := "assets/css/common/app-thema.css.min"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_common_app_css reads file data from disk. It returns an error on failure.
func assets_css_common_app_css() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/css/common/app.css"
	name := "assets/css/common/app.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_common_app_css_min reads file data from disk. It returns an error on failure.
func assets_css_common_app_css_min() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/css/common/app.css.min"
	name := "assets/css/common/app.css.min"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_common_bootstrap_min_css reads file data from disk. It returns an error on failure.
func assets_css_common_bootstrap_min_css() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/css/common/bootstrap.min.css"
	name := "assets/css/common/bootstrap.min.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_common_bootstrap_min_css_min reads file data from disk. It returns an error on failure.
func assets_css_common_bootstrap_min_css_min() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/css/common/bootstrap.min.css.min"
	name := "assets/css/common/bootstrap.min.css.min"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_common_images_ui_bg_glass_100_f6f6f6_1x400_png reads file data from disk. It returns an error on failure.
func assets_css_common_images_ui_bg_glass_100_f6f6f6_1x400_png() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/css/common/images/ui-bg_glass_100_f6f6f6_1x400.png"
	name := "assets/css/common/images/ui-bg_glass_100_f6f6f6_1x400.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_common_images_ui_bg_gloss_wave_35_f6a828_500x100_png reads file data from disk. It returns an error on failure.
func assets_css_common_images_ui_bg_gloss_wave_35_f6a828_500x100_png() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/css/common/images/ui-bg_gloss-wave_35_f6a828_500x100.png"
	name := "assets/css/common/images/ui-bg_gloss-wave_35_f6a828_500x100.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_common_images_ui_bg_highlight_soft_100_eeeeee_1x100_png reads file data from disk. It returns an error on failure.
func assets_css_common_images_ui_bg_highlight_soft_100_eeeeee_1x100_png() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/css/common/images/ui-bg_highlight-soft_100_eeeeee_1x100.png"
	name := "assets/css/common/images/ui-bg_highlight-soft_100_eeeeee_1x100.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_common_images_ui_bg_highlight_soft_75_ffe45c_1x100_png reads file data from disk. It returns an error on failure.
func assets_css_common_images_ui_bg_highlight_soft_75_ffe45c_1x100_png() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/css/common/images/ui-bg_highlight-soft_75_ffe45c_1x100.png"
	name := "assets/css/common/images/ui-bg_highlight-soft_75_ffe45c_1x100.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_common_images_ui_icons_ffffff_256x240_png reads file data from disk. It returns an error on failure.
func assets_css_common_images_ui_icons_ffffff_256x240_png() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/css/common/images/ui-icons_ffffff_256x240.png"
	name := "assets/css/common/images/ui-icons_ffffff_256x240.png"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_common_jquery_ui_min_css reads file data from disk. It returns an error on failure.
func assets_css_common_jquery_ui_min_css() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/css/common/jquery-ui.min.css"
	name := "assets/css/common/jquery-ui.min.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_common_jquery_datetimepicker_css reads file data from disk. It returns an error on failure.
func assets_css_common_jquery_datetimepicker_css() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/css/common/jquery.datetimepicker.css"
	name := "assets/css/common/jquery.datetimepicker.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_fonts_glyphicons_halflings_regular_eot reads file data from disk. It returns an error on failure.
func assets_css_fonts_glyphicons_halflings_regular_eot() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/css/fonts/glyphicons-halflings-regular.eot"
	name := "assets/css/fonts/glyphicons-halflings-regular.eot"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_fonts_glyphicons_halflings_regular_svg reads file data from disk. It returns an error on failure.
func assets_css_fonts_glyphicons_halflings_regular_svg() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/css/fonts/glyphicons-halflings-regular.svg"
	name := "assets/css/fonts/glyphicons-halflings-regular.svg"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_fonts_glyphicons_halflings_regular_ttf reads file data from disk. It returns an error on failure.
func assets_css_fonts_glyphicons_halflings_regular_ttf() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/css/fonts/glyphicons-halflings-regular.ttf"
	name := "assets/css/fonts/glyphicons-halflings-regular.ttf"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_fonts_glyphicons_halflings_regular_woff reads file data from disk. It returns an error on failure.
func assets_css_fonts_glyphicons_halflings_regular_woff() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/css/fonts/glyphicons-halflings-regular.woff"
	name := "assets/css/fonts/glyphicons-halflings-regular.woff"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_fonts_glyphicons_halflings_regular_woff2 reads file data from disk. It returns an error on failure.
func assets_css_fonts_glyphicons_halflings_regular_woff2() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/css/fonts/glyphicons-halflings-regular.woff2"
	name := "assets/css/fonts/glyphicons-halflings-regular.woff2"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_index_css reads file data from disk. It returns an error on failure.
func assets_css_index_css() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/css/index.css"
	name := "assets/css/index.css"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_css_index_css_min reads file data from disk. It returns an error on failure.
func assets_css_index_css_min() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/css/index.css.min"
	name := "assets/css/index.css.min"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_html_common_basic_layout_html reads file data from disk. It returns an error on failure.
func assets_html_common_basic_layout_html() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/html/common/basic-layout.html"
	name := "assets/html/common/basic-layout.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_html_index_html reads file data from disk. It returns an error on failure.
func assets_html_index_html() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/html/index.html"
	name := "assets/html/index.html"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_jsxtransformer_js reads file data from disk. It returns an error on failure.
func assets_js_common_jsxtransformer_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/js/common/JSXTransformer.js"
	name := "assets/js/common/JSXTransformer.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_bootstrap_min_js reads file data from disk. It returns an error on failure.
func assets_js_common_bootstrap_min_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/js/common/bootstrap.min.js"
	name := "assets/js/common/bootstrap.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_gmaps_min_js reads file data from disk. It returns an error on failure.
func assets_js_common_gmaps_min_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/js/common/gmaps.min.js"
	name := "assets/js/common/gmaps.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_indicator_js reads file data from disk. It returns an error on failure.
func assets_js_common_indicator_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/js/common/indicator.js"
	name := "assets/js/common/indicator.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_indicator_js_min reads file data from disk. It returns an error on failure.
func assets_js_common_indicator_js_min() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/js/common/indicator.js.min"
	name := "assets/js/common/indicator.js.min"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_inputs_js reads file data from disk. It returns an error on failure.
func assets_js_common_inputs_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/js/common/inputs.js"
	name := "assets/js/common/inputs.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_jquery_ui_min_js reads file data from disk. It returns an error on failure.
func assets_js_common_jquery_ui_min_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/js/common/jquery-ui.min.js"
	name := "assets/js/common/jquery-ui.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_jquery_activity_indicator_1_0_0_min_js reads file data from disk. It returns an error on failure.
func assets_js_common_jquery_activity_indicator_1_0_0_min_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/js/common/jquery.activity-indicator-1.0.0.min.js"
	name := "assets/js/common/jquery.activity-indicator-1.0.0.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_jquery_datetimepicker_js reads file data from disk. It returns an error on failure.
func assets_js_common_jquery_datetimepicker_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/js/common/jquery.datetimepicker.js"
	name := "assets/js/common/jquery.datetimepicker.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_jquery_min_js reads file data from disk. It returns an error on failure.
func assets_js_common_jquery_min_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/js/common/jquery.min.js"
	name := "assets/js/common/jquery.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_jquery_ui_datepicker_ja_min_js reads file data from disk. It returns an error on failure.
func assets_js_common_jquery_ui_datepicker_ja_min_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/js/common/jquery.ui.datepicker-ja.min.js"
	name := "assets/js/common/jquery.ui.datepicker-ja.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_message_js reads file data from disk. It returns an error on failure.
func assets_js_common_message_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/js/common/message.js"
	name := "assets/js/common/message.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_message_js_min reads file data from disk. It returns an error on failure.
func assets_js_common_message_js_min() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/js/common/message.js.min"
	name := "assets/js/common/message.js.min"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_modal_dialog_js reads file data from disk. It returns an error on failure.
func assets_js_common_modal_dialog_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/js/common/modal-dialog.js"
	name := "assets/js/common/modal-dialog.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_react_with_addons_js reads file data from disk. It returns an error on failure.
func assets_js_common_react_with_addons_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/js/common/react-with-addons.js"
	name := "assets/js/common/react-with-addons.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_react_with_addons_min_js reads file data from disk. It returns an error on failure.
func assets_js_common_react_with_addons_min_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/js/common/react-with-addons.min.js"
	name := "assets/js/common/react-with-addons.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_react_js reads file data from disk. It returns an error on failure.
func assets_js_common_react_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/js/common/react.js"
	name := "assets/js/common/react.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_common_react_min_js reads file data from disk. It returns an error on failure.
func assets_js_common_react_min_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/js/common/react.min.js"
	name := "assets/js/common/react.min.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_index_js reads file data from disk. It returns an error on failure.
func assets_js_index_js() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/js/index.js"
	name := "assets/js/index.js"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assets_js_index_js_min reads file data from disk. It returns an error on failure.
func assets_js_index_js_min() (*asset, error) {
	path := "/Users/jabaraster/Documents/Develop/Go/myproject/src/github.com/jabaraster/go-scaffold/assets/js/index.js.min"
	name := "assets/js/index.js.min"
	bytes, err := bindata_read(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if (err != nil) {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"assets/css/common/app-thema.css": assets_css_common_app_thema_css,
	"assets/css/common/app-thema.css.min": assets_css_common_app_thema_css_min,
	"assets/css/common/app.css": assets_css_common_app_css,
	"assets/css/common/app.css.min": assets_css_common_app_css_min,
	"assets/css/common/bootstrap.min.css": assets_css_common_bootstrap_min_css,
	"assets/css/common/bootstrap.min.css.min": assets_css_common_bootstrap_min_css_min,
	"assets/css/common/images/ui-bg_glass_100_f6f6f6_1x400.png": assets_css_common_images_ui_bg_glass_100_f6f6f6_1x400_png,
	"assets/css/common/images/ui-bg_gloss-wave_35_f6a828_500x100.png": assets_css_common_images_ui_bg_gloss_wave_35_f6a828_500x100_png,
	"assets/css/common/images/ui-bg_highlight-soft_100_eeeeee_1x100.png": assets_css_common_images_ui_bg_highlight_soft_100_eeeeee_1x100_png,
	"assets/css/common/images/ui-bg_highlight-soft_75_ffe45c_1x100.png": assets_css_common_images_ui_bg_highlight_soft_75_ffe45c_1x100_png,
	"assets/css/common/images/ui-icons_ffffff_256x240.png": assets_css_common_images_ui_icons_ffffff_256x240_png,
	"assets/css/common/jquery-ui.min.css": assets_css_common_jquery_ui_min_css,
	"assets/css/common/jquery.datetimepicker.css": assets_css_common_jquery_datetimepicker_css,
	"assets/css/fonts/glyphicons-halflings-regular.eot": assets_css_fonts_glyphicons_halflings_regular_eot,
	"assets/css/fonts/glyphicons-halflings-regular.svg": assets_css_fonts_glyphicons_halflings_regular_svg,
	"assets/css/fonts/glyphicons-halflings-regular.ttf": assets_css_fonts_glyphicons_halflings_regular_ttf,
	"assets/css/fonts/glyphicons-halflings-regular.woff": assets_css_fonts_glyphicons_halflings_regular_woff,
	"assets/css/fonts/glyphicons-halflings-regular.woff2": assets_css_fonts_glyphicons_halflings_regular_woff2,
	"assets/css/index.css": assets_css_index_css,
	"assets/css/index.css.min": assets_css_index_css_min,
	"assets/html/common/basic-layout.html": assets_html_common_basic_layout_html,
	"assets/html/index.html": assets_html_index_html,
	"assets/js/common/JSXTransformer.js": assets_js_common_jsxtransformer_js,
	"assets/js/common/bootstrap.min.js": assets_js_common_bootstrap_min_js,
	"assets/js/common/gmaps.min.js": assets_js_common_gmaps_min_js,
	"assets/js/common/indicator.js": assets_js_common_indicator_js,
	"assets/js/common/indicator.js.min": assets_js_common_indicator_js_min,
	"assets/js/common/inputs.js": assets_js_common_inputs_js,
	"assets/js/common/jquery-ui.min.js": assets_js_common_jquery_ui_min_js,
	"assets/js/common/jquery.activity-indicator-1.0.0.min.js": assets_js_common_jquery_activity_indicator_1_0_0_min_js,
	"assets/js/common/jquery.datetimepicker.js": assets_js_common_jquery_datetimepicker_js,
	"assets/js/common/jquery.min.js": assets_js_common_jquery_min_js,
	"assets/js/common/jquery.ui.datepicker-ja.min.js": assets_js_common_jquery_ui_datepicker_ja_min_js,
	"assets/js/common/message.js": assets_js_common_message_js,
	"assets/js/common/message.js.min": assets_js_common_message_js_min,
	"assets/js/common/modal-dialog.js": assets_js_common_modal_dialog_js,
	"assets/js/common/react-with-addons.js": assets_js_common_react_with_addons_js,
	"assets/js/common/react-with-addons.min.js": assets_js_common_react_with_addons_min_js,
	"assets/js/common/react.js": assets_js_common_react_js,
	"assets/js/common/react.min.js": assets_js_common_react_min_js,
	"assets/js/index.js": assets_js_index_js,
	"assets/js/index.js.min": assets_js_index_js_min,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"assets": &_bintree_t{nil, map[string]*_bintree_t{
		"css": &_bintree_t{nil, map[string]*_bintree_t{
			"common": &_bintree_t{nil, map[string]*_bintree_t{
				"app-thema.css": &_bintree_t{assets_css_common_app_thema_css, map[string]*_bintree_t{
				}},
				"app-thema.css.min": &_bintree_t{assets_css_common_app_thema_css_min, map[string]*_bintree_t{
				}},
				"app.css": &_bintree_t{assets_css_common_app_css, map[string]*_bintree_t{
				}},
				"app.css.min": &_bintree_t{assets_css_common_app_css_min, map[string]*_bintree_t{
				}},
				"bootstrap.min.css": &_bintree_t{assets_css_common_bootstrap_min_css, map[string]*_bintree_t{
				}},
				"bootstrap.min.css.min": &_bintree_t{assets_css_common_bootstrap_min_css_min, map[string]*_bintree_t{
				}},
				"images": &_bintree_t{nil, map[string]*_bintree_t{
					"ui-bg_glass_100_f6f6f6_1x400.png": &_bintree_t{assets_css_common_images_ui_bg_glass_100_f6f6f6_1x400_png, map[string]*_bintree_t{
					}},
					"ui-bg_gloss-wave_35_f6a828_500x100.png": &_bintree_t{assets_css_common_images_ui_bg_gloss_wave_35_f6a828_500x100_png, map[string]*_bintree_t{
					}},
					"ui-bg_highlight-soft_100_eeeeee_1x100.png": &_bintree_t{assets_css_common_images_ui_bg_highlight_soft_100_eeeeee_1x100_png, map[string]*_bintree_t{
					}},
					"ui-bg_highlight-soft_75_ffe45c_1x100.png": &_bintree_t{assets_css_common_images_ui_bg_highlight_soft_75_ffe45c_1x100_png, map[string]*_bintree_t{
					}},
					"ui-icons_ffffff_256x240.png": &_bintree_t{assets_css_common_images_ui_icons_ffffff_256x240_png, map[string]*_bintree_t{
					}},
				}},
				"jquery-ui.min.css": &_bintree_t{assets_css_common_jquery_ui_min_css, map[string]*_bintree_t{
				}},
				"jquery.datetimepicker.css": &_bintree_t{assets_css_common_jquery_datetimepicker_css, map[string]*_bintree_t{
				}},
			}},
			"fonts": &_bintree_t{nil, map[string]*_bintree_t{
				"glyphicons-halflings-regular.eot": &_bintree_t{assets_css_fonts_glyphicons_halflings_regular_eot, map[string]*_bintree_t{
				}},
				"glyphicons-halflings-regular.svg": &_bintree_t{assets_css_fonts_glyphicons_halflings_regular_svg, map[string]*_bintree_t{
				}},
				"glyphicons-halflings-regular.ttf": &_bintree_t{assets_css_fonts_glyphicons_halflings_regular_ttf, map[string]*_bintree_t{
				}},
				"glyphicons-halflings-regular.woff": &_bintree_t{assets_css_fonts_glyphicons_halflings_regular_woff, map[string]*_bintree_t{
				}},
				"glyphicons-halflings-regular.woff2": &_bintree_t{assets_css_fonts_glyphicons_halflings_regular_woff2, map[string]*_bintree_t{
				}},
			}},
			"index.css": &_bintree_t{assets_css_index_css, map[string]*_bintree_t{
			}},
			"index.css.min": &_bintree_t{assets_css_index_css_min, map[string]*_bintree_t{
			}},
		}},
		"html": &_bintree_t{nil, map[string]*_bintree_t{
			"common": &_bintree_t{nil, map[string]*_bintree_t{
				"basic-layout.html": &_bintree_t{assets_html_common_basic_layout_html, map[string]*_bintree_t{
				}},
			}},
			"index.html": &_bintree_t{assets_html_index_html, map[string]*_bintree_t{
			}},
		}},
		"js": &_bintree_t{nil, map[string]*_bintree_t{
			"common": &_bintree_t{nil, map[string]*_bintree_t{
				"JSXTransformer.js": &_bintree_t{assets_js_common_jsxtransformer_js, map[string]*_bintree_t{
				}},
				"bootstrap.min.js": &_bintree_t{assets_js_common_bootstrap_min_js, map[string]*_bintree_t{
				}},
				"gmaps.min.js": &_bintree_t{assets_js_common_gmaps_min_js, map[string]*_bintree_t{
				}},
				"indicator.js": &_bintree_t{assets_js_common_indicator_js, map[string]*_bintree_t{
				}},
				"indicator.js.min": &_bintree_t{assets_js_common_indicator_js_min, map[string]*_bintree_t{
				}},
				"inputs.js": &_bintree_t{assets_js_common_inputs_js, map[string]*_bintree_t{
				}},
				"jquery-ui.min.js": &_bintree_t{assets_js_common_jquery_ui_min_js, map[string]*_bintree_t{
				}},
				"jquery.activity-indicator-1.0.0.min.js": &_bintree_t{assets_js_common_jquery_activity_indicator_1_0_0_min_js, map[string]*_bintree_t{
				}},
				"jquery.datetimepicker.js": &_bintree_t{assets_js_common_jquery_datetimepicker_js, map[string]*_bintree_t{
				}},
				"jquery.min.js": &_bintree_t{assets_js_common_jquery_min_js, map[string]*_bintree_t{
				}},
				"jquery.ui.datepicker-ja.min.js": &_bintree_t{assets_js_common_jquery_ui_datepicker_ja_min_js, map[string]*_bintree_t{
				}},
				"message.js": &_bintree_t{assets_js_common_message_js, map[string]*_bintree_t{
				}},
				"message.js.min": &_bintree_t{assets_js_common_message_js_min, map[string]*_bintree_t{
				}},
				"modal-dialog.js": &_bintree_t{assets_js_common_modal_dialog_js, map[string]*_bintree_t{
				}},
				"react-with-addons.js": &_bintree_t{assets_js_common_react_with_addons_js, map[string]*_bintree_t{
				}},
				"react-with-addons.min.js": &_bintree_t{assets_js_common_react_with_addons_min_js, map[string]*_bintree_t{
				}},
				"react.js": &_bintree_t{assets_js_common_react_js, map[string]*_bintree_t{
				}},
				"react.min.js": &_bintree_t{assets_js_common_react_min_js, map[string]*_bintree_t{
				}},
			}},
			"index.js": &_bintree_t{assets_js_index_js, map[string]*_bintree_t{
			}},
			"index.js.min": &_bintree_t{assets_js_index_js_min, map[string]*_bintree_t{
			}},
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

