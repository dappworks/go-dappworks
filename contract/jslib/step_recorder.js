'use strict';

var StepRecorder = function () {

};

StepRecorder.prototype = {
    record: function (addr, steps) {
        var originalSteps = LocalStorage.get(addr);
        LocalStorage.set(addr, originalSteps + steps);
        return _native_reward.record(addr, steps);
    },
    dapp_schedule: function () {
    }
};

var stepRecorder = new StepRecorder;