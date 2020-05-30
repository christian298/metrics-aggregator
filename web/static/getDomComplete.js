export function getDomComplete(onReport) {
    if (PerformanceObserver.supportedEntryTypes.includes('navigation')) {
        try {
            // Create the performance observer.
            const po = new PerformanceObserver((list) => {
                for (const entry of list.getEntries()) {
                    const { domComplete } = entry;

                    if (domComplete > 0) {
                        onReport({
                            value: domComplete,
                            name: 'DOM_COMPLETE',
                        });
                        po.disconnect();
                    }
                }
            });
            // Start listening for `navigation` entries to be dispatched.
            po.observe({ type: 'navigation', buffered: true });
        } catch (e) {
            // Do nothing if the browser doesn't support this API.
        }
    } else {
        const perf = window.performance;

        if (perf && perf.timing) {
            window.onload = function () {
                onReport({
                    value:
                        perf.timing.domComplete - perf.timing.navigationStart,
                    name: 'DOM_COMPLETE',
                });
            };
        }
    }
}
