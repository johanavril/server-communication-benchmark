import http from "k6/http";
import { check } from "k6";

import { Counter } from "k6/metrics";
import { Rate } from "k6/metrics";
import { Trend } from "k6/metrics";


const BASE_URL = "http://localhost:8080";
const COMMUNICATION_TYPE_HEADER = "X-Communication-Type"
let counterError = new Counter("Healthcheck Counter HTTP non 2xx Response");

export function ping(type) {
	const PATH = "/ping";
	let params = {
	    headers: { [COMMUNICATION_TYPE_HEADER]: type },
	};

	let response = http.get(`${BASE_URL}${PATH}`, params)
	check(response, {
		"status is 200": r => r.status == 200
	})
}
