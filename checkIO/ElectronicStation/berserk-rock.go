"use strict";
function berserkRook(marbles, step) {
    return 0.5
}

var assert = require('assert');

if (!global.is_checking) {
    assert.equal(berserkRook('d3', ['d6', 'b6', 'c8', 'g4', 'b8', 'g6']), 5, "First");
    assert.equal(berserkRook('a2', ['f6', 'f2', 'a6', 'f8', 'h8', 'h6']), 6, "Second");
    assert.equal(berserkRook('a2', ['f6', 'f8', 'f2', 'a6', 'h6']), 4, "Third");
    console.log("Coding complete? Click 'Check' to review your tests and earn cool rewards!");
}
