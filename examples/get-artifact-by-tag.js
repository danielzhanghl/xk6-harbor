// test the performance for the get artifact API
import { Rate } from 'k6/metrics'
import harbor from 'k6/x/harbor'

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
harbor_instance.initialize( settings.Harbor)

export function setup() {
    // find an artifact
    const { projects } = harbor_instance.listProjects({ page: 1, pageSize: 10 })
    for (let i = 0; i < projects.length; i++) {
        const projectName = projects[i].name
        const { repositories } = harbor_instance.listRepositories(projectName)
        for (let j = 0; j < repositories.length; j++) {
            const repositoryName = repositories[j].name.replace(`${projectName}/`, '')
            const { artifacts } = harbor_instance.listArtifacts(projectName, repositoryName)
            if (artifacts.length > 0) {
                for (let k = 0; k < artifacts.length; k++) {
                    if (artifacts[k].tags != null && artifacts[k].tags.length > 0) {
                        const reference = randomItem(artifacts[k].tags).name
                        return {
                            projectName,
                            repositoryName,
                            reference
                        }
                    }
                }
            }
        }
    }

    return {}
}

export default function ({ projectName, repositoryName, reference }) {
    try {
        harbor_instance.getArtifact(projectName, repositoryName, reference)
        successRate.add(true)
    } catch (e) {
        successRate.add(false)
        
    }
}

export function handleSummary(data) {
    return generateSummary('get-artifact-by-tag')(data)
}
