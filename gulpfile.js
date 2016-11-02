var gulp = require('gulp');
var path = require('path');
var shell = require('gulp-shell');

var goPath = 'src/github.com/rdeepk/temperature-monitoring/mypackage/**/*.go';

gulp.task('compilepkg', function() {
    return gulp.src(goPath, {read: false})
    .pipe(shell(['go install <%= stripPath(filr.path) %>'],
        {
            tempateData: {
                stripPath: function(filePath) {
                    var subPath = filePath.substring(process.cwd().length + 5);
                    var pkg = subPath.substring(0, subPath.lastIndexOf(path.sep));
                    return pkg;
                }
            }
        })
     );
});

gulp.task('watch', function() {
    gulp.watch(goPath, ['compilepkg']);
});