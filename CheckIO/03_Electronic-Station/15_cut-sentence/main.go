"use strict";

function cutSentence(line, length) {
    // cut a given sentence in a way so it become shorter than a given length

    // your code here

    return ''
}

var assert = require('assert');

if (!global.is_checking) {
    assert.equal(cutSentence("Hi my name is Alex", 8), "Hi my...", "First")
    assert.equal(cutSentence("Hi my name is Alex", 4), "Hi...", "Second")
    assert.equal(cutSentence("Hi my name is Alex", 20), "Hi my name is Alex", "Third")
    console.log("Done! Do you like it? Go Check it!");
}
