package workflow

import "fmt"

func (m *Manager) RunWorkflow(contextName, workflowName, inputsFileUrl string, optionFileUrl string) (string, error) {
	m.readProjectSpec()
	m.setWorkflowSpec(workflowName)
	m.readConfig()
	m.setContext(contextName)
	m.setEngineForWorkflowType(contextName)
	m.validateContextIsDeployed(contextName)
	m.setOutputBucket()
	m.parseWorkflowLocation()
	m.readInput(inputsFileUrl)
	m.initializeTempDir()
	m.writeTempManifest()
	m.uploadInputsToS3()
	m.parseInputToArguments()
	if m.isUploadRequired() {
		m.setBaseObjectKey(contextName, workflowName)
		m.setWorkflowPath()
		m.packWorkflowPath() 
		m.uploadWorkflowToS3()
		m.cleanUpWorkflow()
	}
	m.calculateFinalLocation()
	m.readOptionFile(optionFileUrl)
	m.setContextStackInfo(contextName)
	m.setWesUrl()
	m.setWesClient()
	m.saveAttachments()
	m.setWorkflowParameters()
	m.setWorkflowEngineParameters()
	defer m.cleanUpAttachments()
	m.runWorkflow()
	m.recordWorkflowRun(workflowName, contextName)
	if m.err != nil {
		return "", fmt.Errorf("unable to run workflow: %w", m.err)
	}
	return m.runId, nil
}
