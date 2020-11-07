import http from 'k6/http';
import { check, sleep, group } from 'k6';

export let options = {
    thresholds: {
        'failed requests': ['rate<0.1'], // threshold on a custom metric
        'http_req_duration': ['p(95)<500']  // threshold on a standard metric
    },
    vus: 10,
    duration: '10s',
};

const SLEEP_DURATION = 0.1;
const API_URL = 'http://localhost:3001/api'

export default function () {
    group("Get patient list", function () {
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

        let login_response = http.post(`${API_URL}/v1/login`, body, params);
        check(login_response, {
            'is status 200': (r) => r.status === 200,
            'is token present': (r) => login_response.json('token') !== '',
        });

        params.headers['authorization'] = `Bearer ${login_response.json('token')}`
        params.tags['name'] = 'Get user list'

        sleep(SLEEP_DURATION);

        // Get patients
        params.tags.name = 'find-patients';

        let get_patients_response = http.get(`${API_URL}/v1/patient?name=pepe&email=x&address=avda`, params);
        check(get_patients_response, {
            'is status 200': (r) => r.status === 200
        });

        sleep(SLEEP_DURATION);
    });
}
