"use strict";

function breakRings(rings){
    return 1
}

var assert = require('assert');

if (!global.is_checking) {
    assert.equal(breakRings([[1, 2], [2, 3], [3, 4], [4, 5], [5, 6], [4, 6]]), 3, "Example");
    assert.equal(breakRings([[1, 2], [1, 3], [1, 4], [2, 3], [2, 4], [3, 4]]), 3, "All to all");
    assert.equal(breakRings([[5, 6], [4, 5], [3, 4], [3, 2], [2, 1], [1, 6]]), 3, "Chain");
    assert.equal(breakRings([[8, 9], [1, 9], [1, 2], [2, 3], [3, 4], [4, 5], [5, 6], [6, 7], [8, 7]]), 5, "Long chain");
    console.log("Coding complete? Click 'Check' to review your tests and earn cool rewards!");
}
