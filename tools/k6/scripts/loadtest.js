import http from 'k6/http';
import * as k6 from 'k6';
import * as data from 'k6/data';
import faker from "k6/x/faker";

// 1. init code
const shared = new data.SharedArray('shared', function () {
  return JSON.parse(open('../shared/bundles.json'));
});

export const options = {
  scenarios: {
    "getBundle": {
      executor: 'shared-iterations',
      vus: 10,
      iterations: 2000,
      startTime: '0s',
      exec: "getBundle"
    },
    "createBundle": {
      executor: 'shared-iterations',
      vus: 10,
      iterations: 2000,
      startTime: '0s',
      exec: "createBundle"
    },
    "updateBundle": {
      executor: 'shared-iterations',
      vus: 10,
      iterations: 2000,
      startTime: '0s',
      exec: "updateBundle"
    }
  },
  // http errors should be less than 5%
  // 95% of requests should be below 200ms
  thresholds: {
    http_req_failed: ['rate<0.05'],
    http_req_duration: ['p(95)<250'],
  },
};

// 2. setup code
export function setup() {
  let host = `${__ENV.HOST}`
  let port = `${__ENV.PORT}`
  let baseURL = host.concat(":", port)
  return { baseURL }
};

export function getBundle(data) {
  const url = data.baseURL.concat("/bundle/", faker.strings.uuid());
  const payload = null
  const params = {
    headers: { 'Content-Type': 'application/json' },
  };

  let res = http.request('GET', url, payload, params);
  if (!k6.check(res, { 'is status 200': (r) => r.status === 200 })) {
    console.error(res);
    k6.fail('unexpected status');
  };
  k6.sleep(1);
};

export function createBundle(data) {
  let rand = shared[Math.floor(Math.random() * shared.length)];
  const url = data.baseURL.concat("/bundle");
  const payload = JSON.stringify(rand);
  const params = {
    headers: { 'Content-Type': 'application/json' },
  };

  let res = http.request('POST', url, payload, params);
  if (!k6.check(res, { 'is status 202': (r) => r.status === 202 })) {
    console.error(res);
    k6.fail('unexpected status');
  };
  k6.sleep(1);
};


export function updateBundle(data) {
  let rand = shared[Math.floor(Math.random() * shared.length)];
  const url = data.baseURL.concat("/bundle/", rand.external_reference_id);
  const payload = JSON.stringify(rand);
  const params = {
    headers: { 'Content-Type': 'application/json' },
  };

  let res = http.request('PATCH', url, payload, params);
  if (!k6.check(res, { 'is status 204': (r) => r.status === 204 })) {
    console.error(res);
    k6.fail('unexpected status');
  };
  k6.sleep(1);
};
