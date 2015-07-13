var gulp = require('gulp');
var browserify = require('browserify');
var reactify = require('reactify');
var watchify = require('watchify');
var source = require('vinyl-source-stream');
var concat = require('gulp-concat');

gulp.task('browserify', function(){
    browserify('./src/app/app.jsx')
        .transform(reactify)
        .bundle()
        .pipe(source('app.js'))
        .pipe(gulp.dest('server/static/'));
});

gulp.task('default', ['browserify']);
