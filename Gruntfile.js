module.exports = function(grunt) {
  require("load-grunt-tasks")(grunt);

  grunt.initConfig({
    babel: {
      options: {
        sourceMap: true
      },
      dist: {
        files: [{
          expand: true,
          cwd: 'src/',
          src: ['**/*.js'],
          dest: 'dist/',
          ext: '.js'
        }]
      }
    },

    clean: {
      dist: [
        'dist'
      ]
    }

  });

  grunt.registerTask("build", ["clean", "babel"]);
  grunt.registerTask("default", ["build"]);
};