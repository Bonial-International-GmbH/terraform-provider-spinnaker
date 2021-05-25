package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mitchellh/mapstructure"
	gate "github.com/spinnaker/spin/cmd/gateclient"
)

func CreatePipeline(client *gate.GatewayClient, pipeline interface{}) error {
	_, resp, err := retry(func() (map[string]interface{}, *http.Response, error) {
		resp, err := client.PipelineControllerApi.SavePipelineUsingPOST(client.Context, pipeline)

		return nil, resp, err
	})
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Encountered an error saving pipeline, status code: %d\n", resp.StatusCode)
	}

	return nil
}

func GetPipeline(client *gate.GatewayClient, applicationName, pipelineName string, dest interface{}) (map[string]interface{}, error) {
	jsonMap, resp, err := retry(func() (map[string]interface{}, *http.Response, error) {
		return client.ApplicationControllerApi.GetPipelineConfigUsingGET(
			client.Context,
			applicationName,
			pipelineName,
		)
	})
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return jsonMap, fmt.Errorf("%s", ErrCodeNoSuchEntityException)
		}
		return jsonMap, fmt.Errorf("Encountered an error getting pipeline %s. Error: %s\n",
			pipelineName,
			err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return jsonMap, fmt.Errorf("Encountered an error getting pipeline in pipeline %s with name %s, status code: %d\n",
			applicationName,
			pipelineName,
			resp.StatusCode)
	}

	if jsonMap == nil {
		return jsonMap, fmt.Errorf(ErrCodeNoSuchEntityException)
	}

	if err := mapstructure.Decode(jsonMap, dest); err != nil {
		return jsonMap, err
	}

	return jsonMap, nil
}

func UpdatePipeline(client *gate.GatewayClient, pipelineID string, pipeline interface{}) error {
	_, resp, err := retry(func() (map[string]interface{}, *http.Response, error) {
		return client.PipelineControllerApi.UpdatePipelineUsingPUT(client.Context, pipelineID, pipeline)
	})
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Encountered an error saving pipeline, status code: %d\n", resp.StatusCode)
	}

	return nil
}

func DeletePipeline(client *gate.GatewayClient, applicationName, pipelineName string) error {
	_, resp, err := retry(func() (map[string]interface{}, *http.Response, error) {
		resp, err := client.PipelineControllerApi.DeletePipelineUsingDELETE(
			client.Context,
			applicationName,
			pipelineName,
		)
		return nil, resp, err
	})
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Encountered an error deleting pipeline, status code: %d\n", resp.StatusCode)
	}
	log.Printf("deleted pipeline %v for application %v", pipelineName, applicationName)
	return nil
}

// RecreatePipeline is a convenience function for deleting and subsequently
// recreating a pipeline. It will return an error if either of the delete and
// create operations fails.
func RecreatePipeline(client *gate.GatewayClient, applicationName, pipelineName string, pipeline interface{}) error {
	err := DeletePipeline(client, applicationName, pipelineName)
	if err != nil {
		return err
	}

	return CreatePipeline(client, pipeline)
}
