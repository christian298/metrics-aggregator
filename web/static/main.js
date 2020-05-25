import {getCLS, getFID, getLCP} from 'https://unpkg.com/web-vitals@0.2.2/dist/web-vitals.es5.min.js?module';

const sendData = async ({value, id, name}) => {
    if("sendBeacon" in navigator) {
        navigator.sendBeacon('http://localhost:4000/metrics', JSON.stringify({value, id, name}));
    } else {
        await fetch("http://localhost:4000/metrics", {method: "POST", body:JSON.stringify({value, id, name})});
    }
};

getCLS(sendData);
getFID(sendData);
getLCP(sendData);