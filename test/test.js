import { group } from 'k6';
import { ping } from './api.js';
export let options = {
  scenarios: {
        vus_50: {
          executor: "constant-vus",
          vus: 50,
          duration: "30s",
          exec: "testPing",
          gracefulStop: "30s",
        },
  },

  thresholds: {
    http_req_duration: ["p(95)<500"],
    errorRate: [
      { threshold: "rate < 0.01", abortOnFail: true },
    ],
  },
};

export function testPing() {
    group('REST Test', function() {
        ping("REST")
    })
//    group('GRPC Test', function() {
//        ping("GRPC")
//    })
}