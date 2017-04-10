"use strict";

function radiationSearch(data) {
    return [0, 0]
}

var assert = require('assert');

if (!global.is_checking) {
    assert.deepEqual(radiationSearch([
        [1, 2, 3, 4, 5],
        [1, 1, 1, 2, 3],
        [1, 1, 1, 2, 2],
        [1, 2, 2, 2, 1],
        [1, 1, 1, 1, 1]
    ]), [14, 1], "14 of 1");
    assert.deepEqual(radiationSearch([
        [2, 1, 2, 2, 2, 4],
        [2, 5, 2, 2, 2, 2],
        [2, 5, 4, 2, 2, 2],
        [2, 5, 2, 2, 4, 2],
        [2, 4, 2, 2, 2, 2],
        [2, 2, 4, 4, 2, 2]
    ]), [19, 2], "19 of 2");
    console.log("Coding complete? Click 'Check' to review your tests and earn cool rewards!");
}
