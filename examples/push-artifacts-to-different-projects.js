// test the performance for pushing artifacts to different projects
import { Rate } from 'k6/metrics'
import counter from 'k6/x/counter'
import harbor from 'k6/x/harbor'
import { ContentStore } from 'k6/x/harbor'

import { Settings } from '../config.js'
import { getProjectNames, randomItem, numberToPadString } from '../helpers.js'
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

    const projectNames = getProjectNames(settings)

    const blobsArr = []
    const refs = []
    for (let i = 0; i < options.iterations; i++) {
        blobsArr.push(store.generateMany(settings.BlobSize, settings.BlobsCountPerArtifact))
        refs.push(`${randomItem(projectNames)}/repository-${Date.now()}:tag-${numberToPadString(i, options.iterations)}`)
    }

    return {
        blobsArr,
        refs
    }
}

export default function ({ blobsArr, refs }) {
    const i = counter.up() - 1

    try {
        harbor_instance.push({ ref: refs[i], store, blobs: blobsArr[i] })
        successRate.add(true)
    } catch (e) {
        successRate.add(false)
        
    }
}

export function teardown({ refs }) {
    store.free()

    for (const ref of refs) {
        const r = /([^/]+)\/([^:]+):(.*)/.exec(ref)
        harbor.deleteArtifact(r[1], r[2], r[3])
    }
}

export function handleSummary(data) {
    return generateSummary('push-artifacts-to-different-projects')(data)
}
