package coverage

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"gotest.tools/assert"
)

func TestTerraformApplies(t *testing.T) {
	rootDir := "../../fixtures/examples/simple-resource"
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: rootDir,
		Logger:       logger.Discard,
		NoColor:      true,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
}

func TestTerraformNoApplyNoCoverage(t *testing.T) {
	// if we don't apply the root, none of the resources should be covered
	rootDir := "../../fixtures/examples/simple-resource"
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: rootDir,
		Logger:       logger.Discard,
		NoColor:      true,
	})
	r, err := NewReport(rootDir)
	assert.NilError(t, err)

	defer terraform.Destroy(t, terraformOptions)

	plan := terraform.InitAndPlanAndShowWithStructNoLogTempPlanFile(t, terraformOptions)
	r.AddCoverage(plan.RawPlan)

	// a Report based on prior state, when the root hasn't been applied, should
	// return 0% coverage
	r.Mode = PriorStateDeterminationMode
	percent, _, err := r.Coverage()
	assert.NilError(t, err)
	assert.Equal(t, 0.0, percent)

	// the same Report (of this root), based on planned values, should return 100%
	// coverage
	r.Mode = PlannedValuesDeterminationMode
	percent, _, err = r.Coverage()
	assert.NilError(t, err)
	assert.Equal(t, 1.0, percent)

	terraform.InitAndApply(t, terraformOptions)

	postApply := terraform.InitAndPlanAndShowWithStructNoLogTempPlanFile(t, terraformOptions)
	r.AddCoverage(postApply.RawPlan)
	// after applying this root, if the Report is based on prior state, it should
	// return 100% coverage
	r.Mode = PriorStateDeterminationMode
	percent, _, err = r.Coverage()
	assert.NilError(t, err)
	assert.Equal(t, 1.0, percent)
}
