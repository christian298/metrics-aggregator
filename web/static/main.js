import {
    getCLS,
    getFID,
    getLCP,
    getTTFB,
    getFCP,
} from 'https://unpkg.com/web-vitals@0.2.2/dist/web-vitals.es5.min.js?module';
import { getDomComplete } from './getDomComplete.js';

const sendData = async ({ value, id, name, delta, entries }) => {
    console.log({ name, entries, value, delta });

    if ('sendBeacon' in navigator) {
        navigator.sendBeacon(
            'http://localhost:4000/metrics',
            JSON.stringify({ value, id, name })
        );
    } else {
        await fetch('http://localhost:4000/metrics', {
            method: 'POST',
            body: JSON.stringify({ value, id, name }),
        });
    }
};

getCLS(sendData);
getFID(sendData);
getLCP(sendData);
getTTFB(sendData);
getFCP(sendData);
getDomComplete(sendData);

// console.log(PerformanceObserver.supportedEntryTypes);
