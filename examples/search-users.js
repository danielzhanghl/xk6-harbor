// test the performance for the search users API
import { Rate } from 'k6/metrics'
import harbor from 'k6/x/harbor'

import { Settings } from '../config.js'
import { getUsernames, randomItem } from '../helpers.js'
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

    return {
        usernames: getUsernames(settings)
    }
}

export default function ({ usernames }) {
    const username = randomItem(usernames)

    try {
        harbor_instance.searchUsers({ username })
        successRate.add(true)
    } catch (e) {
        successRate.add(false)
        
    }
}

export function handleSummary(data) {
    return generateSummary('search-users')(data)
}
