/*********************************************************************************************
 * Purpose of this Script: Limit home battery discharge during EV charge with evcc
 * To avoid battery2battery charging of e-vehicles with high c-rates and deep discharge
 * cycles for the home battery, this script monitors evcc charging mode and soc and
 * limits the discharge rate to 10% (400W Grundlast) on Sunny5 inverters during immediate
 * ev charge sessions.
 *********************************************************************************************/

var isCharging = getState("mqtt.0.evcc.loadpoints.1.charging").val;
var mode = getState("mqtt.0.evcc.loadpoints.1.mode").val;
var soc = getState("mqtt.0.sunny5.soc").val;

function checkState() {
    var discpp = getState("mqtt.0.sunny5.dischg_power_percent").val;
    if (isCharging && mode == 'now' && soc <= 60) {
        log('limit discharge power to 10%'); 
        if (discpp != '10') setState("mqtt.0.sunny5.dischg_power_percent.set", '10');
    } else {
        log('remove discharge limit');
        if (discpp != '100') setState("mqtt.0.sunny5.dischg_power_percent.set", '100');
    }
}

on({ id: 'mqtt.0.evcc.loadpoints.1.charging', change:'any'}, (obj) => {
    isCharging = obj.state.val;
    log('ev charging = ' + isCharging);
    
    checkState();
})

on({ id: 'mqtt.0.evcc.loadpoints.1.mode', change:'any'}, (obj) => {
    log('setting evcc charge mode to ' + obj.state.val);
    mode = obj.state.val;

    checkState();
})

on({ id: 'mqtt.0.evcc.loadpoints.1.mode.set', change:'any'}, (obj) => {
    log('set mode was confirmed ' + obj.state.val);
    //need this ugly workaround to work in ioBroker
    if (obj.state.val != 'ok') setState("mqtt.0.evcc.loadpoints.1.mode.set", 'ok');
})

on({ id: 'mqtt.0.sunny5.soc', change:'any'}, (obj) => {
    soc = obj.state.val;
})
