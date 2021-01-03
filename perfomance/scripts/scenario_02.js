import http from 'k6/http'
import { check, sleep, group } from 'k6'
import uuid from "uuid"
import { name, internet, address } from 'faker/locale/es'
const API_URL = 'http://localhost:3001/api'

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
}

function getRandomArbitraryBetween(max) {
    const min = 0
    return Math.floor(Math.random() * (max - min) + min)
}

function getAvatarPath() {
    const id = getRandomArbitraryBetween(220)
    return `https://marmelab.com/posters/avatar-${id}.jpeg`
}

const SLEEP_DURATION = 0.1

export function setup() {
    // Login
    let body = JSON.stringify({
        email: 'admin@gmail.com',
        password: 'admin',
    })
    let params = {
        headers: {
            'Content-Type': 'application/json',
        },
        tags: {
            name: 'login',
        },
    }

    console.log('Init connection')
    let login_response = http.post(`${API_URL}/v1/login`, body, params)
    console.log('login_response', login_response)

    check(login_response, {
        'is status 200': (r) => r.status === 200,
        'is token present': (r) => login_response.json('token') !== '',
    })

    let authToken = login_response.json('token')

    return authToken
}

export function search_patients(authToken) {
    let params = {
        headers: { 'Content-Type': 'application/json', authorization: `Bearer ${authToken}` },
        tags: { name: 'Get patient list' },
    }

    group("Search patients", function () {
        const first_name = name.firstName()
        const last_name = name.lastName()
        const email = internet.email(first_name, last_name)
        const street =  address.streetAddress(false)

        let get_patients_response = http.get(`${API_URL}/v1/patient?name=${first_name}&email=${email}&address=${street}`, params)

        check(get_patients_response, {
            'is status 200': (r) => r.status === 200
        })

        sleep(SLEEP_DURATION)
    })
}

export function insert_patients(authToken) {
    let params = {
        headers: { 'Content-Type': 'application/json', authorization: `Bearer ${authToken}` },
        tags: { name: 'Insert patients' },
    }

    group("Insert patients", function () {
        const first_name = name.firstName()
        const last_name = name.lastName()
        const email = internet.email(first_name, last_name)
        const street =  address.streetAddress(false)

        let bodyInsert = JSON.stringify({
            "id": uuid.v1(),
            "name": `${first_name} ${last_name}`,
            "address": `${street}`,
            "email": `${email}`,
            "phone": `968${getRandomArbitraryBetween(10000)}`,
            "avatar_path": getAvatarPath()
        })

        let post_patients_response = http.post(`${API_URL}/v1/patient`, bodyInsert, params)

        check(post_patients_response, {
            'is status 201': (r) => r.status === 201
        })

        sleep(SLEEP_DURATION)
    })
}

