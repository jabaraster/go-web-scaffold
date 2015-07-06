'use strict';

var request = require('superagent');
var AppAjax = function() {
    function get(url) {
        return request.get('/resource' + url);
    }
    function post(url) {
        return request.post('/resource' + url);
    }
}
module.exports = AppAjax;
