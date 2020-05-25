(async function () {
    if ("performance" in window) {
        if ("PerformanceObserver" in window) {
            const observer = new PerformanceObserver((list) => {
                for (const entry of list.getEntriesByName('first-contentful-paint')) {

                    fcp = entry.startTime;
                    console.log('FCP:', entry.startTime);
                    send();
                    observer.disconnect();
                }
            });

            const observer2 = new PerformanceObserver((list) => {
                for (const entry of list.getEntries()) {
                    const ttfb = entry.responseStart - entry.requestStart;
                    const transferSize = entry.transferSize;
                    const compressionRatio = entry.decodedBodySize / entry.encodedBodySize;

                    console.log({compressionRatio, transferSize});
                }
            });

            observer.observe({
                type: 'paint',
                buffered: true,
            });

            observer2.observe({type: 'navigation', buffered: true})
        } else {

        }
    }

    let fcp;

    function finishDocumentLoadTime() {
        if (window.PerformanceNavigationTiming && performance.timeOrigin) {
            const ntEntry = performance.getEntriesByType('navigation')[0];
            return (ntEntry.domContentLoadedEventEnd + performance.timeOrigin) / 1000;
        } else {
            return performance.timing.domContentLoadedEventEnd / 1000;
        }
    }

    function ttfb() {
        return window.performance.timing.responseStart - window.performance.timing.requestStart;
    }

    async function send() {
        const payload = {
            finishDocumentTime: finishDocumentLoadTime(),
            ttfb: ttfb(),
            fcp
        };

        return await fetch("http://localhost:4000/metrics", {method: "POST", body:JSON.stringify(payload)});
    }

    // await send();
})();

export const sendData = async (payload) => {
    if("sendBeacon" in navigator) {
        navigator.sendBeacon('http://localhost:4000/metrics', payload);
    } else {
        await fetch("http://localhost:4000/metrics", {method: "POST", body:JSON.stringify(payload)});
    }
};