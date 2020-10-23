import http from 'k6/http';
import { check, sleep, group } from 'k6';
import uuid from "uuid"

export let options = {
    thresholds: {
        'failed requests': ['rate<0.1'],
        'http_req_duration{name:Insert patients}': ['avg<60'],
        'http_req_duration{name:Get patient list}': ['avg<60', 'max<120'],
    },
    scenarios: {
        search_patients: {
            executor: 'constant-vus',
            exec: 'search_patients',
            vus: 10,
            startTime: '5s',
            duration: '25s',
        },
        insert_patients: {
            executor: 'per-vu-iterations',
            exec: 'insert_patients',
            vus: 20,
            iterations: 50,
            maxDuration: '30s',
        },
    },
};

function getRandomNumber() {
    return Math.floor(Math.random() * 1000000);
}

const SLEEP_DURATION = 0.1;

export function setup() {
    // Login
    let body = JSON.stringify({
        email: 'admin@gmail.com',
        password: 'admin',
    });
    let params = {
        headers: {
            'Content-Type': 'application/json',
        },
        tags: {
            name: 'login',
        },
    };

    let login_response = http.post('http://localhost:8080/v1/login', body, params);

    check(login_response, {
        'is status 200': (r) => r.status === 200,
        'is token present': (r) => login_response.json('token') !== '',
    });

    let authToken = login_response.json('token');

    return authToken;
}

export function search_patients(authToken) {
    let params = {
        headers: { 'Content-Type': 'application/json', authorization: `Bearer ${authToken}` },
        tags: { name: 'Get patient list' },
    };

    group("Search patients", function () {
        let get_patients_response = http.get(`http://localhost:8080/v1/patient?name=Name${getRandomNumber()}&email=mail${getRandomNumber()}&address=Address${getRandomNumber()}`, params);

        check(get_patients_response, {
            'is status 200': (r) => r.status === 200
        });

        sleep(SLEEP_DURATION);
    });
}

export function insert_patients(authToken) {
    let params = {
        headers: { 'Content-Type': 'application/json', authorization: `Bearer ${authToken}` },
        tags: { name: 'Insert patients' },
    };

    group("Insert patients", function () {
        let bodyInsert = JSON.stringify({
            "id": uuid.v1(),
            "name": `Name${getRandomNumber()}`,
            "address": `Address${getRandomNumber()}`,
            "email": `mail${getRandomNumber()}@gmail.com`,
            "phone": `968${getRandomNumber()}`
        });

        let post_patients_response = http.post('http://localhost:8080/v1/patient', bodyInsert, params);

        check(post_patients_response, {
            'is status 201': (r) => r.status === 201
        });

        sleep(SLEEP_DURATION);
    });
}

