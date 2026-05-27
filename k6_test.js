import http from 'k6/http';
import { check } from 'k6';

export const options = {
    scenarios: {
        constant_load: {
            executor: 'constant-arrival-rate',
            rate: 8000, // 8k requests/sec total
            timeUnit: '1s',
            duration: '10s',
            preAllocatedVUs: 50,
            maxVUs: 200,
        },
    },
};

const endpoints = [
    { url: 'http://localhost:8085/api/hello', name: 'hello', checkFn: (r) => r.body.includes('Kitwork') },
    { url: 'http://localhost:8085/api/cached', name: 'cached', checkFn: (r) => r.body.includes('Static cached') },
    { url: 'http://localhost:8085/background', name: 'background', checkFn: (r) => r.body.includes('Background task started!') },
    { url: 'http://localhost:8085/', name: 'root', checkFn: (r) => r.body.includes('tabs-container') },
    { url: 'http://localhost:8085/test-query', name: 'query', checkFn: (r) => r.body.includes('count') }
];

export default function () {
    const rand = Math.random();
    let endpoint;
    
    if (rand < 0.35) {
        endpoint = endpoints[0]; // /api/hello (35%)
    } else if (rand < 0.70) {
        endpoint = endpoints[1]; // /api/cached (35%)
    } else if (rand < 0.85) {
        endpoint = endpoints[2]; // /background (15%)
    } else if (rand < 0.95) {
        endpoint = endpoints[3]; // / (10%)
    } else {
        endpoint = endpoints[4]; // /test-query (5% - remote database query)
    }

    const res = http.get(endpoint.url);

    const checkObj = {};
    checkObj[`${endpoint.name} status is 200`] = (r) => r.status === 200;
    checkObj[`${endpoint.name} body is valid`] = (r) => endpoint.checkFn(r);

    check(res, checkObj);
}


