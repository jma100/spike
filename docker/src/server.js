var express = require('express');
var os = require("os");

var app = express();
var hostname = os.hostname();

var request_count = 0;

app.get('/', function (req, res) {
    ++request_count;
    ret = '<html><body>'
        + 'Hello from Node.js container ' + hostname + '.'
        + ' I have processed ' + request_count + ' requests since startup.'
        + '</body></html>';
    res.send(ret);
});

app.listen(80);
console.log('Running on http://localhost');
