/* 
 * 下記を参考に作成.
 * http://liginc.co.jp/web/tutorial/117900
 * 想像以上に快適な開発環境が手に入った！
 * 課題あり
 * ・watchではない任意のタイミングで全ファイルをコンパイルするタスクがない
 */

var gulp = require('gulp');
var less = require('gulp-less');
var autoprefixer = require('gulp-autoprefixer');
var uglifyjs = require('gulp-uglify');
var uglifycss = require('gulp-uglifycss');
var plumber = require("gulp-plumber");
var browserify = require('browserify');
var babelify = require('babelify');
var streamify = require('gulp-streamify');
var source = require('vinyl-source-stream');
var logger = require('gulp-logger');

gulp.task('watch-jsx', function() {
    gulp.watch(['src/jsx/*.jsx'], function(e) {
        var tokens = e.path.split('/');
        var fileName = tokens[tokens.length - 1].split('.')[0];
        browserify({
                entries: [e.path],
                transform: [babelify],
                debug: false // trueにするとsourcemapが生成される. 開発には有用だが、サイズが大きくなる
            })
            .bundle()
            .on("error", function (err) { console.log("Error : " + err.message); })
            .pipe(source(fileName + '.js'))
//            .pipe(streamify(uglifyjs()))
            .pipe(gulp.dest('assets/js'))
            .pipe(logger())
            ;
    });
});

gulp.task('watch-less', function() {
    gulp.watch(['src/less/**/*.less', '!src/less/**/_*.less'], function(e) {
        gulp.src(e.path)
            .pipe(plumber())
            .pipe(less())
            .pipe(autoprefixer())
            .pipe(uglifycss())
            .pipe(gulp.dest('assets/css'))
            .pipe(logger())
        ;
    });
});

gulp.task('watch-html', function() {
    gulp.watch('src/html/**/*.html', function(e) {
        gulp.src(e.path)
            .pipe(plumber())
            .pipe(gulp.dest('assets/html'))
            .pipe(logger())
        ;
    });
});

gulp.task('default', ['watch-jsx', 'watch-less', 'watch-html']);

