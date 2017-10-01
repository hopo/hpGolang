"use strict";

function wallKeeper(on_panels){
    return [];
}


var assert = require('assert');

if (!global.is_checking) {

    function checker(solution, on_panels) {
        var answer = solution(on_panels);
        var p = [];
        for (var i=0; i < 5; i += 1) {
            var row = [];
            for (var j=0; j < 5; j += 1) {
                row.push(on_panels.indexOf(i*5+j+1) > -1 ? 1: 0);
            }
            p.push(row);
        }

        answer.forEach(a=>{
            var r = Math.floor((a-1) / p.length);
            var c = (a-1) % p[0].length;

            p[r][c] = 1 - p[r][c];

            if (r+1 < 5)
                p[r+1][c] = 1 - p[r+1][c];
            if (r-1 > -1)
                p[r-1][c] = 1 - p[r-1][c];
            if (c+1 < 5)
                p[r][c+1] = 1 - p[r][c+1];
            if (c-1 > -1)
                p[r][c-1] = 1 - p[r][c-1];
        });
        var total = 0;
        for (var k=0; k < 5; k += 1) {
            p[k].forEach(function(v){
                total += v
            });
        }
        return total === 0;
    }
    assert.equal(checker(wallKeeper, [5, 7, 13, 14, 18]),
        true, 'basic');

    assert.equal(checker(wallKeeper, [1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
        11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25]),
        true, 'all_lights');

    console.log("Coding complete? Click 'Check' to review your tests and earn cool rewards!");
}
