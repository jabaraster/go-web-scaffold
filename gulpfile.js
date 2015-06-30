/* 
 * 下記を参考に作成.
 * http://liginc.co.jp/web/tutorial/117900
 * 想像以上に快適な開発環境が手に入った！
 */

var gulp = require('gulp');
var less = require('gulp-less');
var autoprefixer = require('gulp-autoprefixer');
var reactify = require('gulp-react');
var uglifyjs = require('gulp-uglify');
var uglifycss = require('gulp-uglifycss');
var browser = require('browser-sync');
var plumber = require("gulp-plumber");
var browserify = require('gulp-browserify');

gulp.task('default', function() {
    gulp.watch(['src/jsx/**/*.jsx']  , ['js']);
    gulp.watch(['src/less/**/*.less'], ['less']);
    gulp.watch(['src/html/**/*.html'], ['html']);
});

gulp.task("server", function() {
    browser({
        server: {
            baseDir: "./assets"
        }
    });
});

gulp.task('less', function() {
    gulp.src(['src/less/**/*.less', '!src/less/**/_*.less'])
        .pipe(plumber())
        .pipe(less())
        .pipe(autoprefixer())
        .pipe(uglifycss())
        .pipe(gulp.dest('assets/css'))
        .pipe(browser.reload({stream: true}))
    ;
});

gulp.task('js', function() {
    gulp.src('src/jsx/**/*.jsx')
        .pipe(plumber())
        .pipe(reactify({harmony: true}))
        .pipe(browserify()) // require()を解決する
        .pipe(uglifyjs())
        .pipe(gulp.dest('assets/js'))
        .pipe(browser.reload({stream: true}))
    ;
});

gulp.task('html', function() {
    gulp.src('src/html/**/*.html')
        .pipe(plumber())
        .pipe(gulp.dest('assets'))
        .pipe(browser.reload({stream: true}))
    ;
});

