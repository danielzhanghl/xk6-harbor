// test the performance for pulling artifacts from same project
import { Rate } from 'k6/metrics'
import harbor from 'k6/x/harbor'
import { ContentStore } from 'k6/x/harbor'

import { Settings } from '../config.js'
import { getProjectNames, randomItem } from '../helpers.js'
import { generateSummary } from '../report.js'

const settings = Settings()

const store = new ContentStore('data')

export let successRate = new Rate('success')

export let options = {
    setupTimeout: '6h',
    duration: '24h',
    teardownTimeout: '6h',
    vus: 500,
    iterations: 1000,
    thresholds: {
        'iteration_duration{scenario:default}': [
            `max>=0`,
        ],
        'iteration_duration{group:::setup}': [`max>=0`],
        'iteration_duration{group:::teardown}': [`max>=0`],
    }
};

export let harbor_instance = harbor
harbor_instance.initialize( settings.Harbor)

export function setup() {

    const projectName = randomItem(getProjectNames(settings))
    const repositoryName = `repository-${Date.now()}`
    const reference = `tag-${Date.now()}`

    harbor_instance.push({
        ref: `${projectName}/${repositoryName}:${reference}`,
        store,
        blobs: store.generateMany(settings.BlobSize, settings.BlobsCountPerArtifact)
    })

    return {
        projectName,
        repositoryName,
        reference
    }
}

export default function ({ projectName, repositoryName, reference }) {
    try {
        harbor_instance.pull(`${projectName}/${repositoryName}:${reference}`)
        successRate.add(true)
    } catch (e) {
        successRate.add(false)
        
    }
}

export function teardown({ projectName, repositoryName, reference }) {
    harbor_instance.deleteArtifact(projectName, repositoryName, reference)

    harbor.free()
}

export function handleSummary(data) {
    return generateSummary('pull-artifacts-from-same-project')(data)
}
