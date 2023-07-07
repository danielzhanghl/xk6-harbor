// test the performance for the list repositories API
import { SharedArray } from 'k6/data'
import { Rate } from 'k6/metrics'
import harbor from 'k6/x/harbor'

import { Settings } from '../config.js'
import { getProjectName, randomItem } from '../helpers.js'
import { generateSummary } from '../report.js'

const settings = Settings()

const projectNames = new SharedArray('projectNames', function () {
    const results = []

    for (let i = 0; i < settings.ProjectsCount; i++) {
        results.push(getProjectName(settings, i))
    }

    return results
});

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
}

export default function () {
    try {
        harbor_instance.listRepositories(randomItem(projectNames))
        successRate.add(true)
    } catch (e) {
        successRate.add(false)
        
    }
}

export function handleSummary(data) {
    return generateSummary('list-repositories')(data)
}
