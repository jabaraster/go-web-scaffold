'use strict';

const request = require('superagent');

module.exports = {
    get: function(url) {
        return request.get('/resource' + url);
    },
    post: function(url) {
        return request.post('/resource' + url);
    }
};
