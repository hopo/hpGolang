"use strict";

function powerSupply(network, power_plants) {
    return [];
}

var assert = require('assert');

if (!global.is_checking) {
    assert.deepEqual(
        powerSupply([['p1', 'c1'], ['c1', 'c2']], {'p1': 1}).sort(),
        ['c2'],
        "one blackout"
    )
    assert.deepEqual(
        powerSupply([['c0', 'c1'], ['c1', 'p1'], ['c1', 'c3'], ['p1', 'c4']], {'p1': 1}).sort(),
        ['c0', 'c3'].sort(),
        "two blackout"
    )
    assert.deepEqual(
        powerSupply([['p1', 'c1'], ['c1', 'c2'], ['c2', 'c3']], {'p1': 3}).sort(),
        [],
        "no blackout"
    )
    assert.deepEqual(
        powerSupply([['c0', 'p1'], ['p1', 'c2']], {'p1': 0}).sort(),
        ['c0', 'c2'].sort(),
        "weak power-plant"
    )
    assert.deepEqual(
        powerSupply([['p0', 'c1'], ['p0', 'c2'], ['c2', 'c3'], ['c3', 'p4'], ['p4', 'c5']],
          {'p0': 1, 'p4': 1}).sort(),
        [],
        "cooperation"
    )
    assert.deepEqual(
        powerSupply([['c0', 'p1'], ['p1', 'c2'], ['c2', 'c3'], ['c2', 'c4'], ['c4', 'c5'],
                     ['c5', 'c6'], ['c5', 'p7']],
                     {'p1': 1, 'p7': 1}).sort(),
        ['c3', 'c4', 'c6'].sort(),
        "complex cities 1"
    )
    assert.deepEqual(
        powerSupply([['p0', 'c1'], ['p0', 'c2'], ['p0', 'c3'],
                       ['p0', 'c4'], ['c4', 'c9'], ['c4', 'c10'],
                       ['c10', 'c11'], ['c11', 'p12'], ['c2', 'c5'],
                       ['c2', 'c6'], ['c5', 'c7'], ['c5', 'p8']],
                      {'p0': 1, 'p12': 4, 'p8': 1}).sort(),
        ['c6', 'c7'].sort(),
        "complex cities 2"
    )
    assert.deepEqual(
        powerSupply([['c1', 'c2'], ['c2', 'c3']], {}).sort(),
        ['c1', 'c2', 'c3'].sort(),
        "no power plants"
    )
    assert.deepEqual(
        powerSupply([['p1', 'c2'], ['p1', 'c4'], ['c4', 'c3'], ['c2', 'c3']],
                    {'p1': 1}).sort(),
        ['c3'],
        "circle"
    )
    assert.deepEqual(
        powerSupply([['p1', 'c2'], ['p1', 'c4'], ['c2', 'c3']], {'p1': 4}).sort(),
        [],
        "more than enough"
    )
    console.log("Push!! Push Check NOW!!");
}
