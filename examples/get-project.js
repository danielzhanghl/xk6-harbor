// test the performance for the get project API
import { Rate } from 'k6/metrics'
import harbor from 'k6/x/harbor'
// import { vu } from 'k6/execution';
// import exec from 'k6/execution';

import { Settings } from '../config.js'
import { randomItem } from '../helpers.js'
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
harbor_instance.initialize(settings.Harbor)

export function setup() {
    // find a project
    const { projects } = harbor_instance.listProjects({ page: 1, pageSize: 10 })
    const projectName = randomItem(projects).name
    return {
        projectName,
    }
}

export default function ({ projectName}) {
    try {
        harbor_instance.getProject(projectName)
        successRate.add(true)
    } catch (e) {
        successRate.add(false)
        
    }
}

export function handleSummary(data) {
    return generateSummary('get-project')(data)
}
