// test the performance for the list users API
import { Rate } from 'k6/metrics'
import harbor from 'k6/x/harbor'

import { Settings } from '../config.js'
import { generateSummary } from '../report.js'

const settings = Settings()

export let successRate = new Rate('success')

export let options = {
    setupTimeout: '6h',
    duration: '24h',
    vus: 500,
    iterations: 1000,
    thresholds: {
        'iteration_duration{scenario:default}': [
            `max>=0`,
        ],
        'iteration_duration{group:::setup}': [`max>=0`],
    }
};

export let harbor_instance = harbor
harbor_instance.initialize( settings.Harbor)

export function setup() {

    const { total } = harbor_instance.listUsers({ page: 1, pageSize: 1 })

    console.log(`total users: ${total}`)

    return {
        usersCount: total
    }
}

export default function ({ usersCount }) {
    const pageSize = 15
    const pages = Math.ceil(usersCount / pageSize)
    const page = Math.floor(Math.random() * pages) + 1

    try {
        harbor_instance.listUsers({ page, pageSize })
        successRate.add(true)
    } catch (e) {
        successRate.add(false)
        
    }
}

export function handleSummary(data) {
    return generateSummary('list-users')(data)
}
